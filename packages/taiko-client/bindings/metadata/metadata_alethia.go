package metadata

// import (
// 	"math/big"
//
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
//
// 	minimalBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
// )
//
// // Ensure TaikoDataBlockMetadataAlethia implements TaikoBlockMetaData.
// var _ TaikoProposalMetaData = (*TaikoDataBlockMetadataAlethia)(nil)
//
// // TaikoDataBlockMetadataAlethia is the metadata of an Alethia Taiko blocks batch.
// type TaikoDataBlockMetadataAlethia struct {
// 	minimalBindings.IInboxPublicationHeader
// 	pubHash common.Hash
// 	types.Log
// }
//
// // NewTaikoDataBlockMetadataPacaya creates a new instance of TaikoDataBlockMetadataAlethia
// // from the TaikoInbox.Published event.
// func NewTaikoDataBlockMetadataAlethia(e *minimalBindings.IInboxPublished) *TaikoDataBlockMetadataAlethia {
// 	return &TaikoDataBlockMetadataAlethia{
// 		IInboxPublicationHeader: e.Header,
// 		pubHash:                 e.PubHash,
// 		Log:                     e.Raw,
// 	}
// }
//
// // Alethia implements TaikoProposalMetaData interface.
// func (m *TaikoDataBlockMetadataAlethia) Alethia() TaikoPublicationMetaDataAlethia {
// 	return m
// }
//
// // IsAlethia implements TaikoProposalMetaData interface.
// func (m *TaikoDataBlockMetadataAlethia) IsAlethia() bool {
// 	return true
// }
//
// // GetTxListHash returns the hash of calldata txlist.
// func (m *TaikoDataBlockMetadataAlethia) GetTxListHash() common.Hash {
// 	return m.TxsHash
// }
//
// // GetTxListHash returns block extradata.
// func (m *TaikoDataBlockMetadataAlethia) GetExtraData() []byte {
// 	return m.ExtraData[:]
// }
//
// // GetCoinbase returns block coinbase.
// func (m *TaikoDataBlockMetadataAlethia) GetCoinbase() common.Address {
// 	return m.Coinbase
// }
//
// // GetTxListHash returns batch ID.
// func (m *TaikoDataBlockMetadataAlethia) GetPublicationID() *big.Int {
// 	return new(big.Int).SetUint64(m.PublicationID)
// }
//
// // GetGasLimit returns gas limit of each L2 block.
// func (m *TaikoDataBlockMetadataAlethia) GetGasLimit() uint32 {
// 	return m.GasLimit
// }
//
// // GetLastBlockTimestamp returns last block's timestamp in this batch.
// func (m *TaikoDataBlockMetadataAlethia) GetLastBlockTimestamp() uint64 {
// 	return m.LastBlockTimestamp
// }
//
// // GetBlobHashes returns blob hashes in this batch.
// func (m *TaikoDataBlockMetadataAlethia) GetBlobHashes() []common.Hash {
// 	var blobHashes []common.Hash
// 	for _, hash := range m.BlobHashes {
// 		blobHashes = append(blobHashes, hash)
// 	}
// 	return blobHashes
// }
//
// // GetLastBlockID returns last block's ID in this batch.
// func (m *TaikoDataBlockMetadataAlethia) GetLastBlockID() uint64 {
// 	return m.LastBlockId
// }
//
// // GetProposer returns the proposer of this batch.
// func (m *TaikoDataBlockMetadataAlethia) GetProposer() common.Address {
// 	return m.Proposer
// }
//
// // GetProposedAt returns the proposing timestamp of this batch.
// func (m *TaikoDataBlockMetadataAlethia) GetProposedAt() uint64 {
// 	return m.ProposedAt
// }
//
// // ProposedIn returns the proposing L1 block number of this batch.
// func (m *TaikoDataBlockMetadataAlethia) GetProposedIn() uint64 {
// 	return m.ProposedIn
// }
//
// // GetBlobCreatedIn returns the L1 block number when the blob created.
// func (m *TaikoDataBlockMetadataAlethia) GetBlobCreatedIn() *big.Int {
// 	return new(big.Int).SetUint64(m.BlobCreatedIn)
// }
//
// // GetTxListOffset returns calldata tx list offset.
// func (m *TaikoDataBlockMetadataAlethia) GetTxListOffset() uint32 {
// 	return m.BlobByteOffset
// }
//
// // GetTxListSize returns calldata tx list size.
// func (m *TaikoDataBlockMetadataAlethia) GetTxListSize() uint32 {
// 	return m.BlobByteSize
// }
//
// // GetAnchorBlockID returns the anchor block ID.
// func (m *TaikoDataBlockMetadataAlethia) GetAnchorBlockID() uint64 {
// 	return m.AnchorBlockId
// }
//
// // GetAnchorBlockHash returns the anchor block hash.
// func (m *TaikoDataBlockMetadataAlethia) GetAnchorBlockHash() common.Hash {
// 	return m.AnchorBlockHash
// }
//
// // GetBlocks returns block params of this batch.
// func (m *TaikoDataBlockMetadataAlethia) GetBlocks() []pacayaBindings.ITaikoInboxBlockParams {
// 	return m.Blocks
// }
//
// // GetBaseFeeConfig returns the L2 block basefee configs.
// func (m *TaikoDataBlockMetadataAlethia) GetBaseFeeConfig() *pacayaBindings.LibSharedDataBaseFeeConfig {
// 	return &m.BaseFeeConfig
// }
//
// // GetRawBlockHeight returns the raw L1 block height.
// func (m *TaikoDataBlockMetadataAlethia) GetRawBlockHeight() *big.Int {
// 	return new(big.Int).SetUint64(m.BlockNumber)
// }
//
// // GetRawBlockHash returns the raw L1 block hash.
// func (m *TaikoDataBlockMetadataAlethia) GetRawBlockHash() common.Hash {
// 	return m.BlockHash
// }
//
// // GetTxIndex returns the transaction index.
// func (m *TaikoDataBlockMetadataAlethia) GetTxIndex() uint {
// 	return m.Log.TxIndex
// }
//
// // GetTxHash returns the transaction hash.
// func (m *TaikoDataBlockMetadataAlethia) GetTxHash() common.Hash {
// 	return m.Log.TxHash
// }
//
// // // IsOntakeBlock returns whether the block is an ontake block.
// // func (m *TaikoDataBlockMetadataAlethia) IsOntakeBlock() bool {
// // 	return true
// // }
