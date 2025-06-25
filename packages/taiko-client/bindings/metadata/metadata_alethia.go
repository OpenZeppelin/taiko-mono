package metadata

import (
	// "encoding/hex"
	// "encoding/hex"
	"fmt"
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

// NewTaikoDataBlockMetadataPacaya creates a new instance of TaikoDataBlockMetadataAlethia
// from the TaikoInbox.Published event.
func NewTaikoDataBlockMetadataAlethia(e *minimalBindings.IInboxPublished) *TaikoDataPublicationAlethia {

	if len(e.Attributes) >= 2 {
		MetadataComponents := []abi.ArgumentMarshaling{
			{Name: "anchorBlockId", Type: "uint64"},
			{Name: "anchorBlockHash", Type: "bytes32"},
			{Name: "isDelayedInclusion", Type: "bool"},
		}
		PublicationMetadataType, err := abi.NewType("tuple(uint256,bytes32,bool)", "IInbox.PublicationMetadata", MetadataComponents)
		if err != nil {
			log.Info("NewType", "err", err)
		}

		PublicationMetadataArg := abi.Arguments{
			{Name: "IInbox.PublicationMetadata", Type: PublicationMetadataType},
		}

		decodedPublicationMetadataStruct, err := PublicationMetadataArg.Unpack(e.Attributes[0])
		if err != nil {
			log.Info("Unpack", "err", err)
		}
		log.Info("decodedPublicationMetadataStruct", "decodedPublicationMetadataStruct", decodedPublicationMetadataStruct)

		BlobRefComponents := []abi.ArgumentMarshaling{
			{Name: "blockNumber", Type: "uint256"},
			{Name: "blobhashes", Type: "bytes32[]"},
		}
		BlobRefType, err := abi.NewType("tuple", "IInbox.BlobRef", BlobRefComponents)
		if err != nil {
			log.Info("NewType", "err", err)
		}

		BlobRefArg := abi.Arguments{
			{Name: "IInbox.BlobRef", Type: BlobRefType},
		}

		blob, err := BlobRefArg.Unpack(e.Attributes[1])
		if err != nil {
			log.Info("Unpack", "err", err)
		}
		fmt.Println("blob", blob[0])

		// metadata := TaikoPublicationAttributesAlethia {
		// 	metadata: decodedPublicationMetadataStruct.(minimal.IInboxPublicationMetadata),
		// 	blobRef: TaikoBlobRef {
		// 		BlockNumber: decodedBlobRefStruct.(struct {
		// 			BlockNumber *big.Int
		// 		}).BlockNumber,
		// 		BlobHashes: decodedBlobRefStruct.(struct {
		// 			BlobHashes []common.Hash
		// 		}).BlobHashes,
		// 	},

		//
		// }

	}

	return &TaikoDataPublicationAlethia{
		header: e.Header,
		// attributes:      e.Attributes,
		publicationHash: e.PubHash,
		attributesHash:  e.Header.AttributesHash,
		Log:             e.Raw,
	}
}

// helper to must-parse ABI types
func mustType(str string) abi.Type {
	t, err := abi.NewType(str, "", nil)
	if err != nil {
		panic(err)
	}
	return t
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
