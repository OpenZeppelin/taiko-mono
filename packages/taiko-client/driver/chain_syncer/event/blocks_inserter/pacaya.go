package blocksinserter

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/holiman/uint256"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/metadata"
	pacayaBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/pacaya"
	anchorTxConstructor "github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/anchor_tx_constructor"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/chain_syncer/beaconsync"
	txListDecompressor "github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/txlist_decompressor"
	txlistFetcher "github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/txlist_fetcher"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/internal/metrics"
	eventIterator "github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/chain_iterator/event_iterator"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/utils"
)

// BlocksInserterPacaya is responsible for inserting Pacaya blocks to the L2 execution engine.
type BlocksInserterPacaya struct {
	rpc                *rpc.Client
	progressTracker    *beaconsync.SyncProgressTracker
	blobDatasource     *rpc.BlobDataSource
	txListDecompressor *txListDecompressor.TxListDecompressor   // Transactions list decompressor
	anchorConstructor  *anchorTxConstructor.AnchorTxConstructor // TaikoAnchor.anchorV3 transactions constructor
	calldataFetcher    txlistFetcher.TxListFetcher
	blobFetcher        txlistFetcher.TxListFetcher
	mutex              sync.Mutex
}

// NewBlocksInserterPacaya creates a new BlocksInserterPacaya instance.
func NewBlocksInserterPacaya(
	rpc *rpc.Client,
	progressTracker *beaconsync.SyncProgressTracker,
	blobDatasource *rpc.BlobDataSource,
	txListDecompressor *txListDecompressor.TxListDecompressor,
	anchorConstructor *anchorTxConstructor.AnchorTxConstructor,
	blobFetcher txlistFetcher.TxListFetcher,
) *BlocksInserterPacaya {
	return &BlocksInserterPacaya{
		rpc:                rpc,
		progressTracker:    progressTracker,
		blobDatasource:     blobDatasource,
		txListDecompressor: txListDecompressor,
		anchorConstructor:  anchorConstructor,
		blobFetcher:        blobFetcher,
	}
}

// InsertBlocks inserts new Pacaya blocks to the L2 execution engine.
func (i *BlocksInserterPacaya) InsertBlocks(
	ctx context.Context,
	meta metadata.TaikoPublicationData,
	endIter eventIterator.EndPublishedEventIterFunc,
) (err error) {
	// if !metadata.IsPacaya() {
	// 	return fmt.Errorf("metadata is not for Pacaya fork")
	// }
	// i.mutex.Lock()
	// defer i.mutex.Unlock()
	//
	var (
		// 	// meta        = metadata.Pacaya()
		txListBytes []byte
	)
	//
	// var publicationId = meta.Header().Id.Uint64()
	//
	// Fetch transactions list.
	if txListBytes, err = i.blobFetcher.FetchPacaya(ctx, meta); err != nil {
		return fmt.Errorf("failed to fetch tx list from blob: %w", err)
	}
	// NOTE: We dont use calldata fetch for Alethia
	// else {
	// 	if txListBytes, err = i.calldataFetcher.FetchPacaya(ctx, meta); err != nil {
	// 		return fmt.Errorf("failed to fetch tx list from calldata: %w", err)
	// 	}
	// }

	var (
		txsWithMetadata = i.txListDecompressor.TryDecompressWithMetadata(i.rpc.L2.ChainID, txListBytes, true)
		parent          *types.Header
		lastPayloadData *engine.ExecutableData
	)
	log.Info("allTxs", "Field 1", txsWithMetadata.Field1)
	log.Info("allTxs", "Field 2", txsWithMetadata.Field2)

	// TODO: this should loop through 'blocks'
	for j := range 2 {
		// Fetch the L2 parent block, if the node is just finished a P2P sync, we simply use the tracker's
		// last synced verified block as the parent, otherwise, we fetch the parent block from L2 EE.
		// if i.progressTracker.Triggered() {
		// Already synced through beacon sync, just skip this event.
		// if new(big.Int).SetUint64(meta.GetLastBlockID()).Cmp(i.progressTracker.LastSyncedBlockID()) <= 0 {
		// 	return nil
		// }

		// parent, err = i.rpc.L2.HeaderByHash(ctx, i.progressTracker.LastSyncedBlockHash())
		// } else {
		// 	var parentNumber *big.Int
		// 	if lastPayloadData == nil {
		// 		if meta.GetBatchID().Uint64() == i.rpc.MinimalRollupClients.ForkHeights.Pacaya {
		// 			parentNumber = new(big.Int).SetUint64(meta.GetBatchID().Uint64() - 1)
		// 		} else {
		// 			lastBatch, err := i.rpc.GetBatchByID(ctx, new(big.Int).SetUint64(meta.GetBatchID().Uint64()-1))
		// 			if err != nil {
		// 				return fmt.Errorf("failed to fetch last batch (%d): %w", meta.GetBatchID().Uint64()-1, err)
		// 			}
		// 			parentNumber = new(big.Int).SetUint64(lastBatch.LastBlockId)
		// 		}
		// 	} else {
		// 		parentNumber = new(big.Int).SetUint64(lastPayloadData.Number)
		// 	}
		//
		// TODO: Good enough for now -  we eventually need to get parent some other way
		parent, err = i.rpc.L2ParentByCurrentBlockID(ctx, new(big.Int).Add(big.NewInt(0), common.Big1))
		// }
		if err != nil {
			return fmt.Errorf("failed to fetch L2 parent block: %w", err)
		}

		log.Debug(
			"Parent block",
			"blockID", parent.Number,
			"hash", parent.Hash(),
			"beaconSyncTriggered", i.progressTracker.Triggered(),
		)

		// Otherwise, we need to create a new execution payload and set it as the head block in L2 EE.
		createExecutionPayloadsMetaData, anchorTx, err := assembleCreateExecutionPayloadMetaPacaya(
			ctx,
			i.rpc,
			i.anchorConstructor,
			meta,
			txsWithMetadata.Transactions,
			meta.Header().Id,
			parent,
			j,
		)
		if err != nil {
			return fmt.Errorf("failed to assemble execution payload creation metadata: %w", err)
		}

		// TODO: Thisi is just a placeholder for now
		config := &pacayaBindings.LibSharedDataBaseFeeConfig{
			AdjustmentQuotient:     10,
			SharingPctg:            80,
			GasIssuancePerSecond:   100000,
			MinGasExcess:           500000000,
			MaxGasIssuancePerBlock: 200000,
		}

		// Decompress the transactions list and try to insert a new head block to L2 EE.
		if lastPayloadData, err = createPayloadAndSetHead(
			ctx,
			i.rpc,
			&createPayloadAndSetHeadMetaData{
				createExecutionPayloadsMetaData: createExecutionPayloadsMetaData,
				AnchorBlockID:                   new(big.Int).SetUint64(meta.Attributes().Metadata.AnchorBlockId.Uint64()),
				AnchorBlockHash:                 meta.Attributes().Metadata.AnchorBlockHash,
				BaseFeeConfig:                   config,
				Parent:                          parent,
			},
			anchorTx,
		); err != nil {
			return fmt.Errorf("failed to insert new head to L2 execution engine: %w", err)
		}

		log.Debug("Payload data", "hash", lastPayloadData.BlockHash, "txs", len(lastPayloadData.Transactions))

		// Wait till the corresponding L2 header to be existed in the L2 EE.
		if _, err := i.rpc.WaitL2Header(ctx, new(big.Int).SetUint64(lastPayloadData.Number)); err != nil {
			return fmt.Errorf("failed to wait for L2 header (%d): %w", lastPayloadData.Number, err)
		}

		log.Info(
			"🔗 New L2 block inserted",
			"blockID", lastPayloadData.Number,
			"hash", lastPayloadData.BlockHash,
			"coinbase", lastPayloadData.FeeRecipient.Hex(),
			"transactions", len(lastPayloadData.Transactions),
			"timestamp", lastPayloadData.Timestamp,
			"baseFee", utils.WeiToGWei(lastPayloadData.BaseFeePerGas),
			"withdrawals", len(lastPayloadData.Withdrawals),
			"publication ID", meta.Header().Id,
			"gasLimit", lastPayloadData.GasLimit,
			"gasUsed", lastPayloadData.GasUsed,
			"parentHash", lastPayloadData.ParentHash,
			"indexInBatch", j,
		)

		metrics.DriverL2HeadHeightGauge.Set(float64(lastPayloadData.Number))
	}

	return nil
}

// InsertPreconfBlocksFromExecutionPayloads inserts preconf blocks from the given execution payloads.
func (i *BlocksInserterPacaya) InsertPreconfBlocksFromExecutionPayloads(
	ctx context.Context,
	executionPayloads []*eth.ExecutionPayload,
	fromCache bool,
) ([]*types.Header, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	headers := make([]*types.Header, len(executionPayloads))
	for j, executableData := range executionPayloads {
		header, err := i.insertPreconfBlockFromExecutionPayload(ctx, executableData)
		if err != nil {
			return nil, fmt.Errorf("failed to insert preconf block: %w", err)
		}
		log.Info(
			"⏰ New preconfirmation L2 block inserted",
			"blockID", header.Number,
			"hash", header.Hash(),
			"coinbase", header.Coinbase.Hex(),
			"timestamp", header.Time,
			"baseFee", utils.WeiToGWei(header.BaseFee),
			"withdrawalsHash", header.WithdrawalsHash,
			"gasLimit", header.GasLimit,
			"gasUsed", header.GasUsed,
			"parentHash", header.ParentHash,
			"fromCache", fromCache,
		)
		headers[j] = header
	}

	return headers, nil
}

// insertPreconfBlockFromExecutionPayload the inner method to insert a preconf block from the given execution payload.
func (i *BlocksInserterPacaya) insertPreconfBlockFromExecutionPayload(
	ctx context.Context,
	executableData *eth.ExecutionPayload,
) (*types.Header, error) {
	// Ensure the preconfirmation block number is greater than the current head L1 origin block ID.
	headL1Origin, err := i.rpc.L2.HeadL1Origin(ctx)
	if err != nil && err.Error() != ethereum.NotFound.Error() {
		return nil, fmt.Errorf("failed to fetch head L1 origin: %w", err)
	}

	// When the chain only has the genesis block, we shall skip this check.
	if headL1Origin != nil {
		if uint64(executableData.BlockNumber) <= headL1Origin.BlockID.Uint64() {
			return nil, fmt.Errorf(
				"preconfirmation block ID (%d) is less than or equal to the current head L1 origin block ID (%d)",
				executableData.BlockNumber,
				headL1Origin.BlockID,
			)
		}
	}

	if len(executableData.Transactions) == 0 {
		return nil, fmt.Errorf("no transactions data in the payload")
	}

	// Decompress the transactions list.
	if executableData.Transactions[0], err = utils.DecompressPacaya(executableData.Transactions[0]); err != nil {
		return nil, fmt.Errorf("failed to decompress transactions list bytes: %w", err)
	}

	var u256BaseFee = uint256.Int(executableData.BaseFeePerGas)
	payload, err := createExecutionPayloadsAndSetHead(
		ctx,
		i.rpc,
		&createExecutionPayloadsMetaData{
			BlockID:               new(big.Int).SetUint64(uint64(executableData.BlockNumber)),
			ExtraData:             executableData.ExtraData,
			SuggestedFeeRecipient: executableData.FeeRecipient,
			GasLimit:              uint64(executableData.GasLimit),
			Difficulty:            common.Hash(executableData.PrevRandao),
			Timestamp:             uint64(executableData.Timestamp),
			ParentHash:            executableData.ParentHash,
			L1Origin: &rawdb.L1Origin{
				BlockID:       new(big.Int).SetUint64(uint64(executableData.BlockNumber)),
				L2BlockHash:   common.Hash{}, // Will be set by taiko-geth.
				L1BlockHeight: nil,
				L1BlockHash:   common.Hash{},
			},
			BaseFee:     u256BaseFee.ToBig(),
			Withdrawals: make([]*types.Withdrawal, 0),
		},
		executableData.Transactions[0],
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create execution data: %w", err)
	}

	metrics.DriverL2PreconfHeadHeightGauge.Set(float64(executableData.BlockNumber))

	return i.rpc.L2.HeaderByHash(ctx, payload.BlockHash)
}

// RemovePreconfBlocks removes preconf blocks from the L2 execution engine.
func (i *BlocksInserterPacaya) RemovePreconfBlocks(ctx context.Context, newLastBlockID uint64) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	newHead, err := i.rpc.L2.HeaderByNumber(ctx, new(big.Int).SetUint64(newLastBlockID))
	if err != nil {
		return err
	}

	fc := &engine.ForkchoiceStateV1{HeadBlockHash: newHead.Hash()}
	fcRes, err := i.rpc.L2Engine.ForkchoiceUpdate(ctx, fc, nil)
	if err != nil {
		return err
	}
	if fcRes.PayloadStatus.Status != engine.VALID {
		return fmt.Errorf("unexpected ForkchoiceUpdate response status: %s", fcRes.PayloadStatus.Status)
	}

	return nil
}
