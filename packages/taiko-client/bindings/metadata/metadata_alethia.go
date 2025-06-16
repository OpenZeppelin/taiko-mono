package metadata

//
// import (
// 	"math/big"
//
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
//
// 	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
// 	minimalBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
// )
//
// // Ensure TaikoDataBlockMetadataAlethia implements TaikoBlockMetaData.
// var _ TaikoPublicationData = (*TaikoPublicationDataAlethia)(nil)
//
// // TaikoDataBlockMetadataAlethia is the metadata of an Alethia Taiko blocks batch.
// type TaikoPublicationDataAlethia struct {
// 	minimalBindings.IInboxPublicationHeader
// 	pubHash common.Hash
// 	types.Log
// }
//
// // NewTaikoDataBlockMetadataPacaya creates a new instance of TaikoDataBlockMetadataAlethia
// // from the TaikoInbox.Published event.
// func NewTaikoDataBlockMetadataAlethia(e *minimalBindings.IInboxPublished) *TaikoPublicationDataAlethia {
// 	return &TaikoPublicationDataAlethia{
// 		IInboxPublicationHeader: e.Header,
// 		pubHash:                 e.PubHash,
// 		Log:                     e.Raw,
// 	}
// }
//
// // IsAlethia implements TaikoProposalMetaData interface.
// func (m *TaikoPublicationDataAlethia) Header() minimal.IInboxPublicationHeader {
// 	return m.IInboxPublicationHeader
// }
//
// func (m *TaikoPublicationDataAlethia) PublicationHash() common.Hash {
// 	return m.pubHash
// }
//
// func (m *TaikoPublicationDataAlethia) Attributes() minimal.IInboxPublicationMetadata {
// 	return m.Attributes()
// }
//
// func (m *TaikoPublicationDataAlethia) AttributesHash() common.Hash {
// 	return m.AttributesHash
// }
