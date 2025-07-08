package builder

import (
	"context"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ProposeBlocksTransactionBuilder is an interface for building a TaikoInbox.proposeBatch
// transaction.
type ProposeBlocksTransactionBuilder interface {
	BuildPacaya(
		ctx context.Context,
		txBatch []types.Transactions,
		anchorBlockId uint64,
		// forcedInclusion *pacayaBindings.IForcedInclusionStoreForcedInclusion,
		// minTxsPerForcedInclusion *big.Int,
		parentMetahash common.Hash,
	) (*txmgr.TxCandidate, error)
}
