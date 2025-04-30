package eventiterator

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	minimalBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
	chainIterator "github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/chain_iterator"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/pkg/rpc"
)

// EndPublishedEventIterFunc ends the current iteration.
type EndPublishedEventIterFunc func()

// OnPublishedEvent represents the callback function which will be called when a PublicationFeed.Published event is
// iterated.
type OnPublishedEvent func(
	context.Context,
	*minimalBindings.IPublicationFeedPublished,
	EndPublishedEventIterFunc,
) error

// PublishedIterator iterates the emitted PublicationFeed.Published events in the chain,
// with the awareness of reorganization.
type PublishedIterator struct {
	ctx                context.Context
	publicationFeed    *minimalBindings.IPublicationFeed
	blockBatchIterator *chainIterator.BlockBatchIterator
	isEnd              bool
}

// PublishedIteratorConfig represents the configs of a Published event iterator.
type PublishedIteratorConfig struct {
	Client                *rpc.EthClient
	PublicationFeed       *minimalBindings.IPublicationFeed
	MaxBlocksReadPerEpoch *uint64
	StartHeight           *big.Int
	EndHeight             *big.Int
	OnPublishedEvent      OnPublishedEvent
	BlockConfirmations    *uint64
}

// NewPublishedIterator creates a new instance of Published event iterator.
func NewPublishedIterator(ctx context.Context, cfg *PublishedIteratorConfig) (*PublishedIterator, error) {
	if cfg.OnPublishedEvent == nil {
		return nil, errors.New("invalid callback")
	}

	iterator := &PublishedIterator{ctx: ctx, publicationFeed: cfg.PublicationFeed}

	// Initialize the inner block iterator.
	blockIterator, err := chainIterator.NewBlockBatchIterator(ctx, &chainIterator.BlockBatchIteratorConfig{
		Client:                cfg.Client,
		MaxBlocksReadPerEpoch: cfg.MaxBlocksReadPerEpoch,
		StartHeight:           cfg.StartHeight,
		EndHeight:             cfg.EndHeight,
		BlockConfirmations:    cfg.BlockConfirmations,
		OnBlocks: assemblePublishedIteratorCallback(
			cfg.Client,
			cfg.PublicationFeed,
			cfg.OnPublishedEvent,
			iterator,
		),
	})
	if err != nil {
		return nil, err
	}

	iterator.blockBatchIterator = blockIterator

	return iterator, nil
}

// Iter iterates the given chain between the given start and end heights,
// will call the callback when a Published event is iterated.
func (i *PublishedIterator) Iter() error {
	return i.blockBatchIterator.Iter()
}

// end ends the current iteration.
func (i *PublishedIterator) end() {
	i.isEnd = true
}

// assemblePublishedIteratorCallback assembles the callback which will be used
// by a event iterator's inner block iterator.
func assemblePublishedIteratorCallback(
	client *rpc.EthClient,
	publicationFeed *minimalBindings.IPublicationFeed,
	callback OnPublishedEvent,
	eventIter *PublishedIterator,
) chainIterator.OnBlocksFunc {
	return func(
		ctx context.Context,
		start, end *types.Header,
		updateCurrentFunc chainIterator.UpdateCurrentFunc,
		endFunc chainIterator.EndIterFunc,
	) error {
		var (
			endHeight  = end.Number.Uint64()
			lastPubID  uint64
		)

		// Iterate the Published events.
		iter, err := publicationFeed.FilterPublished(
			&bind.FilterOpts{Start: start.Number.Uint64(), End: &endHeight, Context: ctx},
			nil, // pubHash indexed param (we're not filtering by it)
		)
		if err != nil {
			return err
		}
		defer iter.Close()

		for iter.Next() {
			event := iter.Event
			log.Debug("Processing Published event", "pubId", event.Header.Id, "l1BlockHeight", event.Raw.BlockNumber)

			if lastPubID != 0 && event.Header.Id.Uint64() != lastPubID+1 {
				log.Warn(
					"Published event is not continuous, rescan the L1 chain",
					"fromL1Block", start.Number,
					"toL1Block", endHeight,
					"lastScannedPubID", lastPubID,
					"currentScannedPubID", event.Header.Id.Uint64(),
				)
				return fmt.Errorf(
					"Published event is not continuous, lastScannedPubID: %d, currentScannedPubID: %d",
					lastPubID, event.Header.Id.Uint64(),
				)
			}

			//TODO(gustavo): We might need to convert the event to a metadata.TaikoProposalMetaData
			// just like in the batch proposed iterator
			if err := callback(ctx, event, eventIter.end); err != nil {
				log.Warn("Error while processing Published events, keep retrying", "error", err)
				return err
			}

			if eventIter.isEnd {
				log.Debug("PublishedIterator is ended", "start", start.Number, "end", endHeight)
				endFunc()
				return nil
			}

			current, err := client.HeaderByHash(ctx, event.Raw.BlockHash)
			if err != nil {
				return err
			}

			log.Debug("Updating current block cursor for processing Published events", "block", current.Number)

			lastPubID = event.Header.Id.Uint64()

			updateCurrentFunc(current)
		}

		// Check if there is any error during the iteration.
		if iter.Error() != nil {
			return iter.Error()
		}

		return nil
	}
} 