package metadata

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
	minimalBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"

	"github.com/ethereum/go-ethereum/log"
)

// Ensure TaikoDataPublicationAlethia implements TaikoBlockMetaData.
var _ TaikoPublicationData = (*TaikoDataPublicationAlethia)(nil)

type TaikoDataPublicationAlethia struct {
	header          minimal.IInboxPublicationHeader
	attributes      TaikoPublicationAttributesAlethia
	publicationHash common.Hash
	attributesHash  common.Hash
	types.Log
}

// NewTaikoDataBlockMetadataPacaya creates a new instance of TaikoDataBlockMetadataAlethia
// from the TaikoInbox.Published event.
func NewTaikoDataBlockMetadataAlethia(e *minimalBindings.IInboxPublished) *TaikoDataPublicationAlethia {

	// var metadata minimal.IInboxPublicationMetadata
	// var blobRef minimal.IBlobRegistryBlobRef

	// if len(e.Attributes) >= 2 {
	// 	if err := abi.UnpackIntoInterface(&metadata, "Metadata", e.Attributes[0]); err != nil {
	// 		panic(fmt.Sprintf("failed to decode Metadata: %v", err))
	// 	}
	// 	if err := abi.UnpackIntoInterface(&blobRef, "BlobRef", e.Attributes[1]); err != nil {
	// 		panic(fmt.Sprintf("failed to decode BlobRef: %v", err))
	// 	}
	// }
	//
	// attr := &TaikoPublicationAttributesAlethia{
	// 	metadata: metadata,
	// 	blobRef:  blobRef,
	// }
	// abi.UnpackRevert(attr, "Metadata")
	//
	// e.Attributes[0]
	ab, err := abi.UnpackRevert(e.Attributes[0])
	if err != nil {
		panic(err)
	}
	log.Info("unpacked", "attr", ab)

	return &TaikoDataPublicationAlethia{
		header: e.Header,
		// attributes:      e.Attributes,
		publicationHash: e.PubHash,
		attributesHash:  e.Header.AttributesHash,
		Log:             e.Raw,
	}
}

func (t *TaikoDataPublicationAlethia) Header() minimal.IInboxPublicationHeader {
	return t.header
}

func (t *TaikoDataPublicationAlethia) PublicationHash() common.Hash {
	return t.publicationHash
}

func (t *TaikoDataPublicationAlethia) Attributes() TaikoPublicationAttributesAlethia {
	return t.attributes
}

func (t *TaikoDataPublicationAlethia) AttributesHash() common.Hash {
	return t.attributesHash
}

// GetRawBlockHeight returns the raw L1 block height.
func (m *TaikoDataPublicationAlethia) GetRawBlockHeight() *big.Int {
	return new(big.Int).SetUint64(m.BlockNumber)
}

// GetRawBlockHash returns the raw L1 block hash.
func (m *TaikoDataPublicationAlethia) GetRawBlockHash() common.Hash {
	return m.BlockHash
}

// GetTxIndex returns the transaction index.
func (m *TaikoDataPublicationAlethia) GetTxIndex() uint {
	return m.Log.TxIndex
}

// GetTxHash returns the transaction hash.
func (m *TaikoDataPublicationAlethia) GetTxHash() common.Hash {
	return m.Log.TxHash
}
