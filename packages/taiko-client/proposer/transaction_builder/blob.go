package builder

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	// "github.com/ethereum/go-ethereum/crypto"

	// "github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	// pacayaBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/pacaya"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/config"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/utils"
)

// BlobTransactionBuilder is responsible for building a TaikoInbox.proposeBatch transaction with txList
// bytes saved in blob.
type BlobTransactionBuilder struct {
	rpc                     *rpc.Client
	proposerPrivateKey      *ecdsa.PrivateKey
	taikoInboxAddress       common.Address
	proverSetAddress        common.Address
	l2SuggestedFeeRecipient common.Address
	gasLimit                uint64
	chainConfig             *config.ChainConfig
	revertProtectionEnabled bool
}

// NewBlobTransactionBuilder creates a new BlobTransactionBuilder instance based on giving configurations.
func NewBlobTransactionBuilder(
	rpc *rpc.Client,
	proposerPrivateKey *ecdsa.PrivateKey,
	taikoInboxAddress common.Address,
	proverSetAddress common.Address,
	l2SuggestedFeeRecipient common.Address,
	gasLimit uint64,
	chainConfig *config.ChainConfig,
	revertProtectionEnabled bool,
) *BlobTransactionBuilder {
	return &BlobTransactionBuilder{
		rpc,
		proposerPrivateKey,
		taikoInboxAddress,
		proverSetAddress,
		l2SuggestedFeeRecipient,
		gasLimit,
		chainConfig,
		revertProtectionEnabled,
	}
}

// TODO: Add forced inclusion

// BuildPacaya implements the ProposeBlocksTransactionBuilder interface.
func (b *BlobTransactionBuilder) BuildPacaya(
	ctx context.Context,
	txBatch []types.Transactions,
	anchorBlockId uint64,
	// minTxsPerForcedInclusion *big.Int,
) (*txmgr.TxCandidate, error) {
	var (
		to = &b.taikoInboxAddress
		// proposer      = crypto.PubkeyToAddress(b.proposerPrivateKey.PublicKey)
		data  []byte
		blobs []*eth.Blob
		// encodedParams []byte
		allTxs types.Transactions
	)

	// if b.proverSetAddress != rpc.ZeroAddress {
	// 	to = &b.proverSetAddress
	// 	proposer = b.proverSetAddress
	// }

	// if forcedInclusion != nil {
	// 	blobParams, blockParams := buildParamsForForcedInclusion(forcedInclusion, minTxsPerForcedInclusion)
	// 	forcedInclusionParams = &encoding.BatchParams{
	// 		Proposer:                 proposer,
	// 		Coinbase:                 b.l2SuggestedFeeRecipient,
	// 		RevertIfNotFirstProposal: b.revertProtectionEnabled,
	// 		BlobParams:               *blobParams,
	// 		Blocks:                   blockParams,
	// 	}
	// }

	for _, txs := range txBatch {
		allTxs = append(allTxs, txs...)
	}

	txListsBytes, err := utils.EncodeAndCompressTxList(allTxs)
	if err != nil {
		return nil, err
	}

	if blobs, err = b.splitToBlobsWithMetadata(txListsBytes); err != nil {
		return nil, err
	}

	// params := &encoding.BatchParams{
	// 	Proposer:                 proposer,
	// 	Coinbase:                 b.l2SuggestedFeeRecipient,
	// 	RevertIfNotFirstProposal: b.revertProtectionEnabled,
	// 	BlobParams: encoding.BlobParams{
	// 		BlobHashes:     [][32]byte{},
	// 		FirstBlobIndex: 0,
	// 		NumBlobs:       uint8(len(blobs)),
	// 		ByteOffset:     0,
	// 		ByteSize:       uint32(len(txListsBytes)),
	// 	},
	// 	Blocks: blockParams,
	// }

	// if b.revertProtectionEnabled {
	// if forcedInclusionParams != nil {
	// 	forcedInclusionParams.ParentMetaHash = parentMetahash
	// } else {
	// params.ParentMetaHash = parentMetahash
	// }
	// }

	if data, err = encoding.TaikoInboxAlethiaABI.Pack("publish", new(big.Int).SetUint64(uint64(len(blobs))), anchorBlockId); err != nil {
		return nil, err
	}

	return &txmgr.TxCandidate{
		TxData:   data,
		Blobs:    blobs,
		To:       to,
		GasLimit: b.gasLimit,
	}, nil
}

func (b *BlobTransactionBuilder) splitToBlobsWithMetadata(txListBytes []byte) ([]*eth.Blob, error) {
	// 4 bytes for uint32 (for now)
	const metadataSize = 4
	maxDataPerBlob := eth.MaxBlobDataSize - metadataSize

	var blobs []*eth.Blob
	totalBlobs := (len(txListBytes) + maxDataPerBlob - 1) / maxDataPerBlob

	for i := 0; i < totalBlobs; i++ {
		start := i * maxDataPerBlob
		end := start + maxDataPerBlob
		if end > len(txListBytes) {
			end = len(txListBytes)
		}

		metadata := make([]byte, metadataSize)

		// Two random values for now just to test it works
		// Field 1
		binary.BigEndian.PutUint16(metadata[0:2], uint16(999))

		// Field 2
		binary.BigEndian.PutUint16(metadata[2:4], uint16(111))

		blobData := append(metadata, txListBytes[start:end]...)

		var blob = &eth.Blob{}
		if err := blob.FromData(blobData); err != nil {
			return nil, err
		}
		blobs = append(blobs, blob)
	}
	return blobs, nil
}

// splitToBlobs splits the txListBytes into multiple blobs.
func (b *BlobTransactionBuilder) splitToBlobs(txListBytes []byte) ([]*eth.Blob, error) {
	var blobs []*eth.Blob
	for start := 0; start < len(txListBytes); start += eth.MaxBlobDataSize {
		end := start + eth.MaxBlobDataSize
		if end > len(txListBytes) {
			end = len(txListBytes)
		}

		var blob = &eth.Blob{}
		if err := blob.FromData(txListBytes[start:end]); err != nil {
			return nil, err
		}

		blobs = append(blobs, blob)
	}

	return blobs, nil
}
