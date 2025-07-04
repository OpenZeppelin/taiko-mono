package anchortxconstructor

import (
	"context"
	"fmt"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	consensus "github.com/ethereum/go-ethereum/consensus/taiko"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/encoding"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/driver/signer"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/utils"
)

// AnchorTxConstructor is responsible for assembling the anchor transaction (TaikoAnchor.anchorV3) in
// each L2 block, which must be the first transaction, and its sender must be the golden touch account.
type AnchorTxConstructor struct {
	rpc    *rpc.Client
	signer *signer.FixedKSigner
}

// New creates a new AnchorConstructor instance.
func New(rpc *rpc.Client) (*AnchorTxConstructor, error) {
	signer, err := signer.NewFixedKSigner("0x" + encoding.GoldenTouchPrivKey)
	if err != nil {
		return nil, fmt.Errorf("invalid golden touch private key %s", encoding.GoldenTouchPrivKey)
	}

	return &AnchorTxConstructor{rpc, signer}, nil
}

// AssembleAnchorV3Tx assembles a signed TaikoAnchor.anchorV3 transaction.
func (c *AnchorTxConstructor) AssembleAnchorV3Tx(
	ctx context.Context,
	// Parameters of the TaikoAnchor.anchorV3 transaction.
	anchorBlockID *big.Int,
	anchorBlockHash common.Hash,
	parentGasUsed uint64,
	publicationId *big.Int,
	blockHeader types.Header,
	// Height of the L2 block which including the TaikoAnchor.anchorV3 transaction.
	l2Height *big.Int,
	baseFee *big.Int,
) (*types.Transaction, error) {
	opts, err := c.transactOpts(ctx, l2Height, baseFee)
	if err != nil {
		return nil, err
	}

	log.Info(
		"AnchorV3 arguments",
		"l2Height", l2Height,
		"anchorBlockId", anchorBlockID,
		"anchorBlockHash", anchorBlockHash,
		"parentGasUsed", parentGasUsed,
		"baseFee", utils.WeiToGWei(baseFee),
	)

	taikoBlockHeader := ConvertToITaikoAnchorBlockHeader(&blockHeader)
	blockHeaderHash := blockHeader.Hash()
	if blockHeaderHash != anchorBlockHash {
		return nil, fmt.Errorf("block header hash %s is not the same as the anchor block hash %s", blockHeaderHash, anchorBlockHash)
	}

	return c.rpc.MinimalRollupClients.TaikoAnchor.Anchor(opts, publicationId, anchorBlockID, anchorBlockHash, *taikoBlockHeader, uint32(parentGasUsed))
}

func ConvertToITaikoAnchorBlockHeader(h *types.Header) *minimal.ITaikoAnchorBlockHeader {
	header := &minimal.ITaikoAnchorBlockHeader{
		ParentHash:       h.ParentHash,
		OmnersHash:       h.UncleHash,
		Coinbase:         h.Coinbase,
		StateRoot:        h.Root,
		TransactionsRoot: h.TxHash,
		ReceiptsRoot:     h.ReceiptHash,
		LogsBloom:        h.Bloom[:],
		Difficulty:       h.Difficulty,
		Number:           h.Number,
		GasLimit:         h.GasLimit,
		GasUsed:          h.GasUsed,
		Timestamp:        h.Time,
		ExtraData:        h.Extra,
		MixedHash:        h.MixDigest,
		Nonce:            h.Nonce.Uint64(),
		BaseFeePerGas:    h.BaseFee,
	}

	if h.BaseFee != nil {
		header.BaseFeePerGas = *&h.BaseFee
	} else {
		header.BaseFeePerGas = big.NewInt(0)
	}

	// These fields are sometimes nil
	if h.WithdrawalsHash != nil {
		header.WithdrawalsRoot = *h.WithdrawalsHash
	} else {
		header.WithdrawalsRoot = common.Hash{}
	}

	if h.BlobGasUsed != nil {
		header.BlobGasUsed = *h.BlobGasUsed
	} else {
		header.BlobGasUsed = 0
	}

	if h.ExcessBlobGas != nil {
		header.ExcessBlobGas = *h.ExcessBlobGas
	} else {
		header.ExcessBlobGas = 0
	}

	if h.ParentBeaconRoot != nil {
		header.ParentBeaconBlockRoot = *h.ParentBeaconRoot
	} else {
		header.ParentBeaconBlockRoot = common.Hash{}
	}

	if h.RequestsHash != nil {
		header.RequestsHash = *h.RequestsHash
	} else {
		header.RequestsHash = common.Hash{}
	}

	return header
}

// transactOpts is a utility method to create some transact options of the anchor transaction in given L2 block with
// golden touch account's private key.
func (c *AnchorTxConstructor) transactOpts(
	ctx context.Context,
	l2Height *big.Int,
	baseFee *big.Int,
) (*bind.TransactOpts, error) {
	var (
		signer       = types.LatestSignerForChainID(c.rpc.L2.ChainID)
		parentHeight = new(big.Int).Sub(l2Height, common.Big1)
	)

	// Get the nonce of golden touch account at the specified parentHeight.
	nonce, err := c.rpc.L2AccountNonce(ctx, consensus.GoldenTouchAccount, parentHeight)
	if err != nil {
		return nil, err
	}

	log.Info(
		"Golden touch account nonce",
		"address", consensus.GoldenTouchAccount,
		"nonce", nonce,
		"parent", parentHeight,
	)

	gasLimit := consensus.AnchorGasLimit
	// if l2Height.Uint64() >= c.rpc.PacayaClients.ForkHeights.Pacaya {
	// 	gasLimit = consensus.AnchorV3GasLimit
	// }

	return &bind.TransactOpts{
		From: consensus.GoldenTouchAccount,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != consensus.GoldenTouchAccount {
				return nil, bind.ErrNotAuthorized
			}
			signature, err := c.signTxPayload(signer.Hash(tx).Bytes())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
		Nonce:     new(big.Int).SetUint64(nonce),
		Context:   ctx,
		GasFeeCap: baseFee,
		GasTipCap: common.Big0,
		GasLimit:  gasLimit,
		NoSend:    true,
	}, nil
}

// signTxPayload calculates an ECDSA signature for an anchor transaction.
func (c *AnchorTxConstructor) signTxPayload(hash []byte) ([]byte, error) {
	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}

	// Try k = 1.
	sig, ok := c.signer.SignWithK(new(secp256k1.ModNScalar).SetInt(1))(hash)
	if !ok {
		// Try k = 2.
		sig, ok = c.signer.SignWithK(new(secp256k1.ModNScalar).SetInt(2))(hash)
		if !ok {
			log.Crit("Failed to sign TaikoAnchor.anchorV3 transaction using K = 1 and K = 2")
		}
	}

	return sig[:], nil
}
