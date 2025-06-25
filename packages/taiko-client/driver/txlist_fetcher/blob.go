package txlistfetcher

import (
	"context"
	"crypto/sha256"
	"math/big"

	// "math/big"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/metadata"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
)

// BlobFetcher is responsible for fetching the txList blob from the L1 block sidecar.
type BlobFetcher struct {
	cli        *rpc.Client
	dataSource *rpc.BlobDataSource
}

// NewBlobTxListFetcher creates a new BlobFetcher instance based on the given rpc client.
func NewBlobTxListFetcher(cli *rpc.Client, ds *rpc.BlobDataSource) *BlobFetcher {
	return &BlobFetcher{cli, ds}
}

func convertToHashes(data [][32]uint8) []common.Hash {
	hashes := make([]common.Hash, len(data))
	for i, item := range data {
		hashes[i] = item
	}
	return hashes
}

// FetchPacaya implements the TxListFetcher interface.
func (d *BlobFetcher) FetchPacaya(
	ctx context.Context,
	meta metadata.TaikoPublicationData,
) ([]byte, error) {
	if len(meta.Attributes().BlobRef.Blobhashes) == 0 {
		return nil, pkg.ErrBlobUnused
	}

	var blockNum uint64
	blockNum = uint64(meta.Attributes().BlobRef.BlockNumber.Uint64())

	// Fetch the L1 block header with the given blob.
	l1Header, err := d.cli.L1.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNum))
	if err != nil {
		return nil, err
	}

	var b []byte
	a := convertToHashes(meta.Attributes().BlobRef.Blobhashes)
	log.Info("a", "a", a)
	// Fetch the L1 block sidecars.
	sidecars, err := d.dataSource.GetBlobs(
		ctx,
		l1Header.Time,
		convertToHashes(meta.Attributes().BlobRef.Blobhashes),
	)
	if err != nil {
		return nil, err
	}

	log.Info(
		"Fetch sidecars",
		"blockNumber", blockNum,
		"sidecars", len(sidecars),
	)

	// NOTE: Ignoring this this for now as its failing
	// This just checks that the blob hash matches the kzg commitment
	// TODO: Fix comparing blob hash with kzg commitment

	for _, blobHash := range meta.Attributes().BlobRef.Blobhashes {
		// Compare the blob hash with the sidecar's kzg commitment.
		for j, sidecar := range sidecars {
			log.Debug(
				"Block sidecar",
				"index", j,
				"KzgCommitment", sidecar.KzgCommitment,
				"blobHash", blobHash,
			)
			commitment := kzg4844.Commitment(common.FromHex(sidecar.KzgCommitment))
			if kzg4844.CalcBlobHashV1(sha256.New(), &commitment) == blobHash {
				blob := eth.Blob(common.FromHex(sidecar.Blob))
				bytes, err := blob.ToData()
				if err != nil {
					return nil, err
				}

				b = append(b, bytes...)
			}
		}
	}
	if len(b) == 0 {
		return nil, pkg.ErrSidecarNotFound
	}

	return sliceTxList(meta.Header().Id, b, 0, uint32(len(meta.Attributes().BlobRef.Blobhashes)))
}
