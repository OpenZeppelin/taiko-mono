package txlistdecompressor

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/utils"
)

// TxListDecompressor is responsible for validating and decompressing
// the transactions list in a TaikoInbox.proposeBatch transaction.
type TxListDecompressor struct {
	blockMaxGasLimit  uint64
	maxBytesPerTxList uint64
	chainID           *big.Int
}

// NewTxListDecompressor creates a new TxListDecompressor instance based on giving configurations.
func NewTxListDecompressor(
	blockMaxGasLimit uint64,
	maxBytesPerTxList uint64,
	chainID *big.Int,
) *TxListDecompressor {
	return &TxListDecompressor{
		blockMaxGasLimit:  blockMaxGasLimit,
		maxBytesPerTxList: maxBytesPerTxList,
		chainID:           chainID,
	}
}

type DecompressionResult struct {
	Transactions types.Transactions
	Field1       uint16
	Field2       uint16
}

// New method that returns both transactions and metadata
func (v *TxListDecompressor) TryDecompressWithMetadata(
	chainID *big.Int,
	txListBytes []byte,
	blobUsed bool,
) DecompressionResult {
	return v.tryDecompressWithMetadata(txListBytes)
}

// TryDecompress validates and decompresses whether the transactions list in the TaikoInbox.proposeBatch transaction's
// input data is valid, the rules are:
// - If the transaction list is empty, it's valid.
// - If the transaction list is not empty:
//  1. If the transaction list is using calldata, the compressed bytes of the transaction list must be
//     less than or equal to maxBytesPerTxList.
//  2. The transaction list bytes must be able to be RLP decoded into a list of transactions.
func (v *TxListDecompressor) TryDecompress(
	chainID *big.Int,
	txListBytes []byte,
	blobUsed bool,
) types.Transactions {
	return v.tryDecompress(txListBytes)
}

// tryDecompress is the inner implementation of TryDecompress.
func (v *TxListDecompressor) tryDecompress(
	txListBytes []byte,
) types.Transactions {
	// If the transaction list is empty, it's valid.
	if len(txListBytes) == 0 {
		return types.Transactions{}
	}

	// NOTE: Calldata is never used

	// // If calldata is used, the compressed bytes of the transaction list must be
	// // less than or equal to maxBytesPerTxList.
	// if !blobUsed && (len(txListBytes) > int(v.maxBytesPerTxList)) {
	// 	log.Info(
	// 		"Compressed transactions list binary too large",
	// 		"length", len(txListBytes),
	// 	)
	// 	return types.Transactions{}
	// }

	var (
		txs types.Transactions
		err error
	)

	// Decompress the transaction list bytes.
	if txListBytes, err = utils.DecompressPacaya(txListBytes); err != nil {
		log.Info("Failed to decompress tx list bytes", "error", err)
		return types.Transactions{}
	}

	// Try to RLP decode the transaction list bytes.
	if err = rlp.DecodeBytes(txListBytes, &txs); err != nil {
		log.Info("Failed to decode transactions list bytes", "error", err)
		return types.Transactions{}
	}

	return txs
}

// tryDecompressWithMetadata is the inner implementation that extracts metadata
func (v *TxListDecompressor) tryDecompressWithMetadata(
	txListBytes []byte,
) DecompressionResult {
	const metadataSize = 4

	if len(txListBytes) == 0 {
		return DecompressionResult{
			Transactions: types.Transactions{},
			Field1:       0,
			Field2:       0,
		}
	}

	if len(txListBytes) < metadataSize {
		log.Info("Transaction list bytes too small for metadata", "length", len(txListBytes))
		return DecompressionResult{
			Transactions: types.Transactions{},
			Field1:       0,
			Field2:       0,
		}
	}

	Field1 := binary.BigEndian.Uint16(txListBytes[0:2])
	field2 := binary.BigEndian.Uint16(txListBytes[2:4])

	log.Debug("Extracted blob metadata",
		"Field1", Field1,
		"ield2", field2,
	)

	actualTxData := txListBytes[metadataSize:]

	var (
		txs types.Transactions
		err error
	)

	if actualTxData, err = utils.DecompressPacaya(actualTxData); err != nil {
		log.Info("Failed to decompress tx list bytes", "error", err)
		return DecompressionResult{
			Transactions: types.Transactions{},
			Field1:       Field1,
			Field2:       field2,
		}
	}

	if err = rlp.DecodeBytes(actualTxData, &txs); err != nil {
		log.Info("Failed to decode transactions list bytes", "error", err)
		return DecompressionResult{
			Transactions: types.Transactions{},
			Field1:       Field1,
			Field2:       field2,
		}
	}

	return DecompressionResult{
		Transactions: txs,
		Field1:       Field1,
		Field2:       field2,
	}
}
