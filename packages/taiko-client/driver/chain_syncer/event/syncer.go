package event

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/metadata"
	pacayaBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/pacaya"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/chain_syncer/beaconsync"
	blocksInserter "github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/chain_syncer/event/blocks_inserter"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/state"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/metrics"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"

	minimalBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
	anchorTxConstructor "github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/anchor_tx_constructor"
	txListDecompressor "github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/txlist_decompressor"
	txlistFetcher "github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/txlist_fetcher"
	eventIterator "github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/chain_iterator/event_iterator"
)

// Syncer responsible for letting the L2 execution engine catching up with protocol's latest
// pending block through deriving L1 calldata.
type Syncer struct {
	ctx                context.Context
	rpc                *rpc.Client
	state              *state.State
	progressTracker    *beaconsync.SyncProgressTracker        // Sync progress tracker
	txListDecompressor *txListDecompressor.TxListDecompressor // Transactions list decompressor

	// Blocks inserters
	blocksInserterPacaya blocksInserter.Inserter // Pacaya blocks inserter

	lastInsertedBlockID *big.Int
	reorgDetectedFlag   bool
}

// NewSyncer creates a new syncer instance.
func NewSyncer(
	ctx context.Context,
	client *rpc.Client,
	state *state.State,
	progressTracker *beaconsync.SyncProgressTracker,
	blobServerEndpoint *url.URL,
) (*Syncer, error) {
	constructor, err := anchorTxConstructor.New(client)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize anchor constructor: %w", err)
	}

	protocolConfigs, err := client.GetProtocolConfigs(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	blobDataSource := rpc.NewBlobDataSource(
		ctx,
		client,
		blobServerEndpoint,
	)

	txListDecompressor := txListDecompressor.NewTxListDecompressor(
		uint64(protocolConfigs.BlockMaxGasLimit()),
		rpc.BlockMaxTxListBytes,
		client.L2.ChainID,
	)

	var (
		txListFetcherBlob = txlistFetcher.NewBlobTxListFetcher(client, blobDataSource)
	)
	return &Syncer{
		ctx:                ctx,
		rpc:                client,
		state:              state,
		progressTracker:    progressTracker,
		txListDecompressor: txListDecompressor,
		blocksInserterPacaya: blocksInserter.NewBlocksInserterPacaya(
			client,
			progressTracker,
			blobDataSource,
			txListDecompressor,
			constructor,
			txListFetcherBlob,
		),
	}, nil
}

// ProcessL1Blocks fetches all `PublicationFeed.Published` events between given
// L1 block heights, and then tries inserting them into L2 execution engine's blockchain.
func (s *Syncer) ProcessL1Blocks(ctx context.Context) error {
	for {
		if err := s.processL1Blocks(ctx); err != nil {
			return err
		}

		// If the L1 chain has been reorged, we process the new L1 blocks again with
		// the new L1Current cursor.
		if s.reorgDetectedFlag {
			s.reorgDetectedFlag = false
			continue
		}

		return nil
	}
}

// processL1Blocks is the inner method which responsible for processing
// all new L1 blocks.
func (s *Syncer) processL1Blocks(ctx context.Context) error {
	var (
		l1End          = s.state.GetL1Head()
		startL1Current = s.state.GetL1Current()
	)
	log.Info("Processing L1 blocks", "startL1Current", startL1Current.Number.Uint64(), "l1End", l1End.Number.Uint64())
	// If there is a L1 reorg, sometimes this will happen.
	if startL1Current.Number.Uint64() >= l1End.Number.Uint64() && startL1Current.Hash() != l1End.Hash() {
		newL1Current, err := s.rpc.L1.HeaderByNumber(ctx, new(big.Int).Sub(l1End.Number, common.Big1))
		if err != nil {
			return err
		}

		log.Info(
			"Reorg detected",
			"oldL1CurrentHeight", startL1Current.Number,
			"oldL1CurrentHash", startL1Current.Hash(),
			"newL1CurrentHeight", newL1Current.Number,
			"newL1CurrentHash", newL1Current.Hash(),
			"l1Head", l1End.Number,
		)

		s.state.SetL1Current(newL1Current)
		s.lastInsertedBlockID = nil
	}

	log.Info("Creating PublishedIterator")
	iter, err := eventIterator.NewPublishedIterator(ctx, &eventIterator.PublishedIteratorConfig{
		Client:           s.rpc.L1,
		TaikoInbox:       s.rpc.MinimalRollupClients.Inbox,
		StartHeight:      s.state.GetL1Current().Number,
		EndHeight:        l1End.Number,
		OnPublishedEvent: s.onPublished,
	})
	if err != nil {
		return err
	}

	if err := iter.Iter(); err != nil {
		return err
	}

	// If there is a L1 reorg, we don't update the L1Current cursor.
	if !s.reorgDetectedFlag {
		s.state.SetL1Current(l1End)
		metrics.DriverL1CurrentHeightGauge.Set(float64(s.state.GetL1Current().Number.Uint64()))
	}

	return nil
}

// onBatchProposed is a `BatchProposed` event callback which responsible for
// inserting the proposed block one by one to the L2 execution engine.
func (s *Syncer) onBatchProposed(
	ctx context.Context,
	meta metadata.TaikoProposalMetaData,
	endIter eventIterator.EndBatchProposedEventIterFunc,
) error {
	var (
		// firstBlockID = new(big.Int).SetUint64(meta.Pacaya().GetLastBlockID() - uint64(len(meta.Pacaya().GetBlocks())) + 1)
		lastBlockID = new(big.Int).SetUint64(meta.Pacaya().GetLastBlockID())
		timestamp   = meta.Pacaya().GetLastBlockTimestamp()
	)

	// We simply ignore the genesis block's `BatchesProposed` event.
	if lastBlockID.Cmp(common.Big0) == 0 {
		return nil
	}

	// If we are not inserting a block whose parent block is the latest verified block in protocol,
	// and the node hasn't just finished the P2P sync, we check if the L1 chain has been reorged.
	if !s.progressTracker.Triggered() {
		// reorgCheckResult, err := s.checkReorg(ctx, firstBlockID)
		// if err != nil {
		// 	return err
		// }
		//
		// if reorgCheckResult.IsReorged {
		// 	log.Info(
		// 		"Reset L1Current cursor due to L1 reorg",
		// 		"l1CurrentHeightOld", s.state.GetL1Current().Number,
		// 		"l1CurrentHashOld", s.state.GetL1Current().Hash(),
		// 		"l1CurrentHeightNew", reorgCheckResult.L1CurrentToReset.Number,
		// 		"l1CurrentHashNew", reorgCheckResult.L1CurrentToReset.Hash(),
		// 		"lastInsertedBlockIDOld", s.lastInsertedBlockID,
		// 		"lastInsertedBlockIDNew", reorgCheckResult.LastHandledBlockIDToReset,
		// 	)
		// 	s.state.SetL1Current(reorgCheckResult.L1CurrentToReset)
		// 	s.lastInsertedBlockID = reorgCheckResult.LastHandledBlockIDToReset
		// 	s.reorgDetectedFlag = true
		// 	endIter()
		//
		// 	return nil
		// }
	}

	// Ignore those already inserted blocks.
	if s.lastInsertedBlockID != nil && lastBlockID.Cmp(s.lastInsertedBlockID) <= 0 {
		log.Debug(
			"Skip already inserted block",
			"blockID", lastBlockID,
			"lastInsertedBlockID", s.lastInsertedBlockID,
		)
		return nil
	}

	// If the event's timestamp is in the future, we wait until the timestamp is reached, should
	// only happen when testing.
	if timestamp > uint64(time.Now().Unix()) {
		log.Warn(
			"Future L2 block, waiting",
			"L2BlockTimestamp", timestamp,
			"now", time.Now().Unix(),
		)
		time.Sleep(time.Until(time.Unix(int64(timestamp), 0)))
	}

	// Insert new blocks to L2 EE's chain.
	log.Info(
		"New BatchProposed event",
		"l1Height", meta.GetRawBlockHeight(),
		"l1Hash", meta.GetRawBlockHash(),
		"batchID", meta.Pacaya().GetBatchID(),
		"lastBlockID", lastBlockID,
		"lastTimestamp", meta.Pacaya().GetLastBlockTimestamp(),
		"blocks", len(meta.Pacaya().GetBlocks()),
	)
	if err := s.blocksInserterPacaya.InsertBlocks(ctx, meta, endIter); err != nil {
		return err
	}

	metrics.DriverL1CurrentHeightGauge.Set(float64(meta.GetRawBlockHeight().Uint64()))
	s.lastInsertedBlockID = lastBlockID

	if s.progressTracker.Triggered() {
		s.progressTracker.ClearMeta()
	}

	return nil
}

// onPublished is a `Published` event callback which responsible for
func (s *Syncer) onPublished(
	ctx context.Context,
	event *minimalBindings.IInboxPublished,
	endIter eventIterator.EndPublishedEventIterFunc,
) error {
	log.Info("onPublished called - processing Published event")
	// Extract block information from the published event
	blockID := event.Header.Id
	timestamp := event.Header.Timestamp.Uint64()

	// We simply ignore the genesis block's event
	if blockID.Cmp(common.Big0) == 0 {
		return nil
	}

	// If we are not inserting a block whose parent block is the latest verified block in protocol,
	// and the node hasn't just finished the P2P sync, we check if the L1 chain has been reorged.
	if !s.progressTracker.Triggered() {
		// reorgCheckResult, err := s.checkReorg(ctx, blockID)
		// if err != nil {
		// 	return err
		// }
		//
		// if reorgCheckResult.IsReorged {
		// 	log.Info(
		// 		"Reset L1Current cursor due to L1 reorg",
		// 		"l1CurrentHeightOld", s.state.GetL1Current().Number,
		// 		"l1CurrentHashOld", s.state.GetL1Current().Hash(),
		// 		"l1CurrentHeightNew", reorgCheckResult.L1CurrentToReset.Number,
		// 		"l1CurrentHashNew", reorgCheckResult.L1CurrentToReset.Hash(),
		// 		"lastInsertedBlockIDOld", s.lastInsertedBlockID,
		// 		"lastInsertedBlockIDNew", reorgCheckResult.LastHandledBlockIDToReset,
		// 	)
		// 	s.state.SetL1Current(reorgCheckResult.L1CurrentToReset)
		// 	s.lastInsertedBlockID = reorgCheckResult.LastHandledBlockIDToReset
		// 	s.reorgDetectedFlag = true
		// 	endIter()
		//
		// 	return nil
		// }
	}

	// Ignore those already inserted blocks.
	if s.lastInsertedBlockID != nil && blockID.Cmp(s.lastInsertedBlockID) <= 0 {
		log.Debug(
			"Skip already inserted block",
			"blockID", blockID,
			"lastInsertedBlockID", s.lastInsertedBlockID,
		)
		return nil
	}

	// If the event's timestamp is in the future, we wait until the timestamp is reached, should
	// only happen when testing.
	if timestamp > uint64(time.Now().Unix()) {
		log.Warn(
			"Future L2 block, waiting",
			"L2BlockTimestamp", timestamp,
			"now", time.Now().Unix(),
		)
		time.Sleep(time.Until(time.Unix(int64(timestamp), 0)))
	}

	// Insert new blocks to L2 EE's chain.
	log.Info(
		"New Publication event",
		"l1Height", event.Raw.BlockNumber,
		"l1Hash", event.Raw.BlockHash,
		"pubId", event.Header.Id,
		"timestamp", timestamp,
		"attributes", len(event.Attributes),
	)

	// NOTE: This block of code is creating a dummy metadata object to pass to the blocks inserter
	// This is a temporary solution until we have a better way to handle this
	// What we need to do is to create a new metadata object that implements the TaikoBatchMetaDataPacaya interface
	// I suspect some things might fail down the line if we don't do this

	//*************************************************************************

	// Fetch the transaction to get the blob hash
	tx, isPending, err := s.rpc.L1.TransactionByHash(ctx, event.Raw.TxHash)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction by hash: %w", err)
	}

	// Extract blob hashes from the transaction
	var blobHashes []common.Hash
	if tx.Type() == 3 { // EIP-4844 blob transaction
		blobHashes = tx.BlobHashes()
		log.Info("Found blob transaction", "txHash", tx.Hash(), "blobHashes", blobHashes)
	} else {
		log.Info("Transaction is not a blob transaction", "txHash", tx.Hash(), "txType", tx.Type())
	}

	// Create a placeholder metadata object from the Published event
	meta := &placeholderPacayaMetadata{
		event:      event,
		blockID:    blockID,
		endIter:    endIter,
		blobHashes: blobHashes,
	}

	// Create a wrapper that implements TaikoProposalMetaData
	wrappedMeta := &placeholderProposalMetadata{
		pacayaMeta: meta,
	}

	// Create an adapter function to convert EndPublishedEventIterFunc to EndBatchProposedEventIterFunc
	endIterAdapter := func() {
		endIter()
	}

	log.Info("Inserting blocks with metadata",
		"blockID", blockID,
		"txHash", event.Raw.TxHash,
		"isPending", isPending,
		"blobHashes", blobHashes)

	// ************************************************************************

	if err := s.blocksInserterPacaya.InsertBlocks(ctx, wrappedMeta, endIterAdapter); err != nil {
		return err
	}

	metrics.DriverL1CurrentHeightGauge.Set(float64(event.Raw.BlockNumber))
	s.lastInsertedBlockID = blockID

	if s.progressTracker.Triggered() {
		s.progressTracker.ClearMeta()
	}

	return nil
}

// checkLastVerifiedBlockMismatch checks if there is a mismatch between protocol's last verified block hash and
// the corresponding L2 EE block hash.
// func (s *Syncer) checkLastVerifiedBlockMismatch(ctx context.Context) (*rpc.ReorgCheckResult, error) {
// Fetch the latest verified block hash.
// ts, err := s.rpc.GetLastVerifiedTransitionPacaya(ctx)
// if err != nil {
// 	return nil, err
// }
//
// var (
// 	reorgCheckResult    = new(rpc.ReorgCheckResult)
// 	lastVerifiedBatchID = ts.BatchId
// )
//
// // If the current L2 chain is behind of the last verified block, we skip the check.
// if s.state.GetL2Head().Number.Uint64() < ts.BlockId ||
// 	(s.lastInsertedBlockID != nil && s.lastInsertedBlockID.Uint64() < ts.BlockId) {
// 	return reorgCheckResult, nil
// }
//
// header, err := s.rpc.L2.HeaderByNumber(ctx, new(big.Int).SetUint64(ts.BlockId))
// if err != nil {
// 	return nil, fmt.Errorf("failed to fetch L2 header by number: %w", err)
// }
//
// // If the last verified block hash matches the L2 EE block hash, we skip the check.
// if header.Hash() == ts.Ts.BlockHash {
// 	return reorgCheckResult, nil
// }
//
// for {
// 	batch, err := s.rpc.GetBatchByID(ctx, new(big.Int).SetUint64(lastVerifiedBatchID))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch batch by ID: %w", err)
// 	}
// 	previousBatch, err := s.rpc.GetBatchByID(ctx, new(big.Int).SetUint64(lastVerifiedBatchID-1))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch previous batch by ID: %w", err)
// 	}
//
// 	if batch.VerifiedTransitionId.Cmp(common.Big0) == 0 {
// 		lastVerifiedBatchID = previousBatch.BatchId
// 		continue
// 	}
// 	// ts, err := s.rpc.PacayaClients.TaikoInbox.GetBatchVerifyingTransition(&bind.CallOpts{Context: ctx}, batch.BatchId)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("failed to fetch Pacaya transition: %w", err)
// 	// }
// 	if header, err = s.rpc.L2.HeaderByNumber(ctx, new(big.Int).SetUint64(batch.LastBlockId)); err != nil {
// 		return nil, fmt.Errorf("failed to fetch L2 header by number: %w", err)
// 	}
//
// 	if header.Hash() == ts.BlockHash {
// 		log.Info(
// 			"Verified block matched, start reorging",
// 			"currentHeightToCheck", batch.LastBlockId,
// 			"chainBlockHash", header.Hash(),
// 			"transitionBlockHash", common.BytesToHash(ts.BlockHash[:]),
// 		)
// 		reorgCheckResult.IsReorged = true
// 		if reorgCheckResult.L1CurrentToReset, err = s.rpc.L1.HeaderByNumber(
// 			ctx,
// 			new(big.Int).SetUint64(batch.AnchorBlockId),
// 		); err != nil {
// 			return nil, fmt.Errorf("failed to fetch L1 header by number: %w", err)
// 		}
// 		reorgCheckResult.LastHandledBlockIDToReset = header.Number
// 		return reorgCheckResult, nil
// 	}
//
// 	log.Info(
// 		"Verified block mismatch",
// 		"currentHeightToCheck", batch.LastBlockId,
// 		"chainBlockHash", header.Hash(),
// 		"transitionBlockHash", common.BytesToHash(ts.BlockHash[:]),
// 	)
//
// 	lastVerifiedBatchID = previousBatch.BatchId
// }
// }

// // checkReorg checks whether the L1 chain has been reorged, and resets the L1Current cursor if necessary.
// func (s *Syncer) checkReorg(ctx context.Context, blockID *big.Int) (*rpc.ReorgCheckResult, error) {
// 	// If the L2 chain is at genesis, we don't need to check L1 reorg.
// 	if s.state.GetL1Current().Number == s.state.GenesisL1Height {
// 		return new(rpc.ReorgCheckResult), nil
// 	}
//
// 	// 1. Check if the verified blocks in L2 EE have been reorged.
// 	reorgCheckResult, err := s.checkLastVerifiedBlockMismatch(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to check if the verified blocks in L2 EE have been reorged: %w", err)
// 	}
//
// 	// 2. If the verified blocks check is passed, we check the unverified blocks.
// 	if reorgCheckResult == nil || !reorgCheckResult.IsReorged {
// 		if reorgCheckResult, err = s.rpc.CheckL1Reorg(ctx, new(big.Int).Sub(blockID, common.Big1)); err != nil {
// 			return nil, fmt.Errorf("failed to check whether L1 chain has been reorged: %w", err)
// 		}
// 	}
//
// 	return reorgCheckResult, nil
// }

// BlocksInserterPacaya returns the Pacaya blocks inserter.
func (s *Syncer) BlocksInserterPacaya() *blocksInserter.BlocksInserterPacaya {
	return s.blocksInserterPacaya.(*blocksInserter.BlocksInserterPacaya)
}

// placeholderPacayaMetadata is a placeholder implementation of TaikoBatchMetaDataPacaya
type placeholderPacayaMetadata struct {
	event      *minimalBindings.IInboxPublished
	blockID    *big.Int
	endIter    eventIterator.EndPublishedEventIterFunc
	blobHashes []common.Hash
	blocks     []pacayaBindings.ITaikoInboxBlockParams
}

// Implement the TaikoBatchMetaDataPacaya interface methods
func (m *placeholderPacayaMetadata) GetTxListHash() common.Hash {
	return m.event.Raw.TxHash
}

func (m *placeholderPacayaMetadata) GetExtraData() []byte {
	return []byte{}
}

func (m *placeholderPacayaMetadata) GetCoinbase() common.Address {
	return common.Address{}
}

func (m *placeholderPacayaMetadata) GetBatchID() *big.Int {
	return m.blockID
}

func (m *placeholderPacayaMetadata) GetGasLimit() uint32 {
	return 30000000 // Default gas limit
}

func (m *placeholderPacayaMetadata) GetLastBlockTimestamp() uint64 {
	return m.event.Header.Timestamp.Uint64()
}

func (m *placeholderPacayaMetadata) GetProposer() common.Address {
	return common.Address{}
}

func (m *placeholderPacayaMetadata) GetProposedAt() uint64 {
	return m.event.Raw.BlockNumber
}

func (m *placeholderPacayaMetadata) GetProposedIn() uint64 {
	return m.event.Raw.BlockNumber
}

func (m *placeholderPacayaMetadata) GetBlobCreatedIn() *big.Int {
	return new(big.Int).SetUint64(m.event.Raw.BlockNumber)
}

func (m *placeholderPacayaMetadata) GetTxListOffset() uint32 {
	return 0
}

func (m *placeholderPacayaMetadata) GetTxListSize() uint32 {
	return 0
}

func (m *placeholderPacayaMetadata) GetLastBlockID() uint64 {
	return m.blockID.Uint64()
}

func (m *placeholderPacayaMetadata) GetBlobHashes() []common.Hash {
	return m.blobHashes
}

func (m *placeholderPacayaMetadata) GetAnchorBlockID() uint64 {
	return m.event.Raw.BlockNumber
}

func (m *placeholderPacayaMetadata) GetAnchorBlockHash() common.Hash {
	return m.event.Raw.BlockHash
}

func (m *placeholderPacayaMetadata) GetBlocks() []pacayaBindings.ITaikoInboxBlockParams {
	return m.blocks
}

func (m *placeholderPacayaMetadata) GetBaseFeeConfig() *pacayaBindings.LibSharedDataBaseFeeConfig {
	return &pacayaBindings.LibSharedDataBaseFeeConfig{}
}

func (m *placeholderPacayaMetadata) GetRawBlockHeight() *big.Int {
	return new(big.Int).SetUint64(m.event.Raw.BlockNumber)
}

func (m *placeholderPacayaMetadata) GetRawBlockHash() common.Hash {
	return m.event.Raw.BlockHash
}

func (m *placeholderPacayaMetadata) GetTxIndex() uint {
	return m.event.Raw.TxIndex
}

func (m *placeholderPacayaMetadata) GetTxHash() common.Hash {
	return m.event.Raw.TxHash
}

func (m *placeholderPacayaMetadata) InnerMetadata() *pacayaBindings.ITaikoInboxBatchMetadata {
	return &pacayaBindings.ITaikoInboxBatchMetadata{}
}

// placeholderProposalMetadata is a placeholder implementation of TaikoProposalMetaData
type placeholderProposalMetadata struct {
	pacayaMeta *placeholderPacayaMetadata
}

// Implement the TaikoProposalMetaData interface methods
func (m *placeholderProposalMetadata) Pacaya() metadata.TaikoBatchMetaDataPacaya {
	return m.pacayaMeta
}

func (m *placeholderProposalMetadata) IsPacaya() bool {
	return true
}

func (m *placeholderProposalMetadata) GetRawBlockHeight() *big.Int {
	return m.pacayaMeta.GetRawBlockHeight()
}

func (m *placeholderProposalMetadata) GetRawBlockHash() common.Hash {
	return m.pacayaMeta.GetRawBlockHash()
}

func (m *placeholderProposalMetadata) GetTxIndex() uint {
	return m.pacayaMeta.GetTxIndex()
}

func (m *placeholderProposalMetadata) GetTxHash() common.Hash {
	return m.pacayaMeta.GetTxHash()
}

func (m *placeholderProposalMetadata) GetProposer() common.Address {
	return m.pacayaMeta.GetProposer()
}

func (m *placeholderProposalMetadata) GetCoinbase() common.Address {
	return m.pacayaMeta.GetCoinbase()
}

func (m *placeholderProposalMetadata) GetBlobCreatedIn() *big.Int {
	return m.pacayaMeta.GetBlobCreatedIn()
}
