package metadata

import (
	// "encoding/hex"
	// "encoding/hex"
	"math/big"
	"reflect"

	// "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
	minimalBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"

	"github.com/ethereum/go-ethereum/accounts/abi"

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

func decodeBlobRef(data []byte) (*BlobRef, error) {
	components := []abi.ArgumentMarshaling{
		{Name: "blockNumber", Type: "uint256"},
		{Name: "blobhashes", Type: "bytes32[]"},
	}
	tupleTy, err := abi.NewType("tuple", "IInbox.BlobRef", components)
	if err != nil {
		return nil, err
	}

	args := abi.Arguments{{Name: "blobRef", Type: tupleTy}}

	out, err := args.Unpack(data)
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(out[0])
	blockNumber := v.FieldByName("BlockNumber").Interface().(*big.Int)
	blobhashes := v.FieldByName("Blobhashes").Interface().([][32]uint8)

	result := &BlobRef{
		BlockNumber: blockNumber,
		Blobhashes:  blobhashes,
	}

	return result, nil
}

func decodePublicationMetadata(data []byte) (*minimal.IInboxPublicationMetadata, error) {
	components := []abi.ArgumentMarshaling{
		{Name: "anchorBlockId", Type: "uint64"},
		{Name: "anchorBlockHash", Type: "bytes32"},
		{Name: "isDelayedInclusion", Type: "bool"},
	}
	PublicationMetadataType, err := abi.NewType("tuple", "IInbox.PublicationMetadata", components)
	if err != nil {
		return nil, err
	}
	args := abi.Arguments{{Name: "publicationMetadata", Type: PublicationMetadataType}}

	out, err := args.Unpack(data)
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(out[0])
	anchorBlockId := v.FieldByName("AnchorBlockId").Interface().(uint64)
	anchorBlockHash := v.FieldByName("AnchorBlockHash").Interface().([32]uint8)
	isDelayedInclusion := v.FieldByName("IsDelayedInclusion").Interface().(bool)

	result := &minimal.IInboxPublicationMetadata{
		AnchorBlockId:      big.NewInt(int64(anchorBlockId)),
		AnchorBlockHash:    anchorBlockHash,
		IsDelayedInclusion: isDelayedInclusion,
	}

	return result, nil
}

// NewTaikoDataBlockMetadataPacaya creates a new instance of TaikoDataBlockMetadataAlethia
// from the TaikoInbox.Published event.
func NewTaikoDataBlockMetadataAlethia(e *minimalBindings.IInboxPublished) *TaikoDataPublicationAlethia {

	var taikoPubAttr TaikoPublicationAttributesAlethia
	if len(e.Attributes) >= 2 {
		publicationMetadata, err := decodePublicationMetadata(e.Attributes[0])
		if err != nil {
			log.Info("Unpack", "err", err)
		}
		publicationBlobRef, err := decodeBlobRef(e.Attributes[1])
		if err != nil {
			log.Info("Unpack", "err", err)
		}
		taikoPubAttr = TaikoPublicationAttributesAlethia{
			Metadata: *publicationMetadata,
			BlobRef:  *publicationBlobRef,
		}

	}

	return &TaikoDataPublicationAlethia{
		header:          e.Header,
		attributes:      taikoPubAttr,
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
