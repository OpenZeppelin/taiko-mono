package rpc

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"

	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
	pacayaBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/pacaya"
)

// TODO: ADD ALL SUSBSCRIPTIONS FOR TAIKO INBOX

// SubscribeEvent creates a event subscription, will retry if the established subscription failed.
func SubscribeEvent(
	eventName string,
	handler func(ctx context.Context) (event.Subscription, error),
) event.Subscription {
	return event.ResubscribeErr(
		backoff.DefaultMaxInterval,
		func(ctx context.Context, err error) (event.Subscription, error) {
			if err != nil {
				log.Warn("Failed to subscribe protocol event, try resubscribing", "event", eventName, "error", err)
			}

			return handler(ctx)
		},
	)
}

func SubscribePublishedAlethia(
	taikoInbox *minimal.IInboxFilterer,
	ch chan *minimal.IInboxPublished,
	pubHash [][32]byte,
) event.Subscription {
	return SubscribeEvent("Published", func(ctx context.Context) (event.Subscription, error) {
		sub, err := taikoInbox.WatchPublished(nil, ch, pubHash)
		if err != nil {
			log.Error("Create TaikoInbox.Published subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// SubscribeBatchesVerifiedPacaya subscribes the Pacaya protocol's BatchesVerified events.
func SubscribeCheckpointVerifiedPacaya(
	taikoInbox *pacayaBindings.TaikoInboxClient,
	ch chan *pacayaBindings.TaikoInboxClientBatchesVerified,
) event.Subscription {
	return SubscribeEvent("BatchesVerified", func(ctx context.Context) (event.Subscription, error) {
		sub, err := taikoInbox.WatchBatchesVerified(nil, ch)
		if err != nil {
			log.Error("Create TaikoInbox.BatchesVerified subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

//
// // SubscribeBatchProposedPacaya subscribes the Pacaya protocol's BatchProposed events.
// func SubscribeBatchProposedPacaya(
// 	taikoInbox *pacayaBindings.TaikoInboxClient,
// 	ch chan *pacayaBindings.TaikoInboxClientBatchProposed,
// ) event.Subscription {
// 	return SubscribeEvent("BatchProposed", func(ctx context.Context) (event.Subscription, error) {
// 		sub, err := taikoInbox.WatchBatchProposed(nil, ch)
// 		if err != nil {
// 			log.Error("Create TaikoInbox.BatchProposed subscription error", "error", err)
// 			return nil, err
// 		}
//
// 		defer sub.Unsubscribe()
//
// 		return waitSubErr(ctx, sub)
// 	})
// }

// SubscribeCheckpointUpdatedAlethia subscribes the Alethia protocol's CheckpointUpdated events.
func SubscribeCheckpointUpdatedAlethia(
	taikoCheckpoint *minimal.ICheckpointTracker,
	ch chan *minimal.ICheckpointTrackerCheckpointUpdated,
) event.Subscription {
	return SubscribeEvent("CheckpointUpdated", func(ctx context.Context) (event.Subscription, error) {
		sub, err := taikoCheckpoint.WatchCheckpointUpdated(nil, ch)
		if err != nil {
			log.Error("Create CheckpointTracker.CheckpointUpdated subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// SubscribeChainHead subscribes the new chain heads.
func SubscribeChainHead(
	client *EthClient,
	ch chan *types.Header,
) event.Subscription {
	return SubscribeEvent("ChainHead", func(ctx context.Context) (event.Subscription, error) {
		sub, err := client.SubscribeNewHead(ctx, ch)
		if err != nil {
			log.Error("Create chain head subscription error", "error", err)
			return nil, err
		}

		defer sub.Unsubscribe()

		return waitSubErr(ctx, sub)
	})
}

// waitSubErr keeps waiting until the given subscription failed.
func waitSubErr(ctx context.Context, sub event.Subscription) (event.Subscription, error) {
	select {
	case err := <-sub.Err():
		return sub, err
	case <-ctx.Done():
		return sub, ctx.Err()
	}
}
