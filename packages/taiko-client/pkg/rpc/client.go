package rpc

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	minimalBindings "github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
)

const (
	defaultTimeout = 1 * time.Minute
)

// PacayaClients contains all smart contract clients for Pacaya fork.
// type PacayaClients struct {
// 	TaikoInbox           *pacayaBindings.TaikoInboxClient
// 	TaikoWrapper         *pacayaBindings.TaikoWrapperClient
// 	ForcedInclusionStore *pacayaBindings.ForcedInclusionStore
// 	TaikoAnchor          *pacayaBindings.TaikoAnchorClient
// 	TaikoToken           *pacayaBindings.TaikoToken
// 	ProverSet            *pacayaBindings.ProverSet
// 	ForkRouter           *pacayaBindings.ForkRouter
// 	ComposeVerifier      *pacayaBindings.ComposeVerifier
// 	PreconfWhitelist     *pacayaBindings.PreconfWhitelist
// ForkHeights          *pacayaBindings.ITaikoInboxForkHeights
// }

// MinimalRollupClients contains all smart contract clients for minimal-rollup interfaces.
type MinimalRollupClients struct {
	NewTaikoInbox *minimalBindings.IInbox
	TaikoAnchor   *minimalBindings.ITaikoAnchor
	// ProverManager   *minimalBindings.IProverManager
	ProposerFees          *minimalBindings.IProposerFees
	Verifier              *minimalBindings.IVerifier
	Lookahead             *minimalBindings.ILookahead
	Inbox                 *minimalBindings.IInbox
	CheckpointTracker     *minimalBindings.ICheckpointTracker
	DelayedInclusionStore *minimalBindings.IDelayedInclusionStore
}

// Client contains all L1/L2 RPC clients that a driver needs.
type Client struct {
	// Geth ethclient clients
	L1           *EthClient
	L2           *EthClient
	L2CheckPoint *EthClient
	// Geth Engine API clients
	L2Engine *EngineClient
	// Beacon clients
	L1Beacon *BeaconClient
	// Protocol contracts clients
	MinimalRollupClients *MinimalRollupClients
}

// ClientConfig contains all configs which will be used to initializing an
// RPC client. If not providing L2EngineEndpoint or JwtSecret, then the L2Engine client
// won't be initialized.
type ClientConfig struct {
	L1Endpoint                  string
	L2Endpoint                  string
	L1BeaconEndpoint            string
	L2CheckPoint                string
	NewTaikoInboxAddress        common.Address
	TaikoWrapperAddress         common.Address
	TaikoAnchorAddress          common.Address
	TaikoTokenAddress           common.Address
	ForcedInclusionStoreAddress common.Address
	PreconfWhitelistAddress     common.Address
	ProverSetAddress            common.Address
	// Minimal Rollup Addresses
	ProverManagerAddress common.Address
	ProposerFeesAddress  common.Address
	VerifierAddress      common.Address
	LookaheadAddress     common.Address
	InboxAddress         common.Address
	L2EngineEndpoint     string
	JwtSecret            string
	Timeout              time.Duration
}

// NewClient initializes all RPC clients used by Taiko client software.
func NewClient(ctx context.Context, cfg *ClientConfig) (*Client, error) {
	var (
		l1Client       *EthClient
		l2Client       *EthClient
		l1BeaconClient *BeaconClient
		l2CheckPoint   *EthClient
		err            error
	)

	// Keep retrying to connect to the RPC endpoints until success or context is cancelled.
	if err := backoff.Retry(func() error {
		ctxWithTimeout, cancel := CtxWithTimeoutOrDefault(ctx, defaultTimeout)
		defer cancel()

		if l1Client, err = NewEthClient(ctxWithTimeout, cfg.L1Endpoint, cfg.Timeout); err != nil {
			log.Error("Failed to connect to L1 endpoint, retrying", "endpoint", cfg.L1Endpoint, "err", err)
			return err
		}

		if l2Client, err = NewEthClient(ctxWithTimeout, cfg.L2Endpoint, cfg.Timeout); err != nil {
			log.Error("Failed to connect to L2 endpoint, retrying", "endpoint", cfg.L2Endpoint, "err", err)
			return err
		}

		// NOTE: when running tests, we do not have a L1 beacon endpoint.
		if cfg.L1BeaconEndpoint != "" && os.Getenv("RUN_TESTS") == "" {
			if l1BeaconClient, err = NewBeaconClient(cfg.L1BeaconEndpoint, defaultTimeout); err != nil {
				log.Error("Failed to connect to L1 beacon endpoint, retrying", "endpoint", cfg.L1BeaconEndpoint, "err", err)
				return err
			}
		}

		if cfg.L2CheckPoint != "" {
			l2CheckPoint, err = NewEthClient(ctxWithTimeout, cfg.L2CheckPoint, cfg.Timeout)
			if err != nil {
				log.Error("Failed to connect to L2 checkpoint endpoint, retrying", "endpoint", cfg.L2CheckPoint, "err", err)
				return err
			}
		}

		return nil
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx)); err != nil {
		return nil, err
	}

	// If not providing L2EngineEndpoint or JwtSecret, then the L2Engine client
	// won't be initialized.
	var l2AuthClient *EngineClient
	if len(cfg.L2EngineEndpoint) != 0 && len(cfg.JwtSecret) != 0 {
		l2AuthClient, err = NewJWTEngineClient(cfg.L2EngineEndpoint, cfg.JwtSecret)
		if err != nil {
			return nil, err
		}
	}

	c := &Client{
		L1:           l1Client,
		L1Beacon:     l1BeaconClient,
		L2:           l2Client,
		L2CheckPoint: l2CheckPoint,
		L2Engine:     l2AuthClient,
	}

	// Initialize minimal rollup clients
	if err := c.initMinimalRollupClients(cfg); err != nil {
		return nil, fmt.Errorf("failed to initialize Minimal Rollup clients: %w", err)
	}

	// ctxWithTimeout, cancel := CtxWithTimeoutOrDefault(ctx, defaultTimeout)
	// defer cancel()
	// Initialize the fork height numbers.
	// if err := c.initForkHeightConfigs(ctxWithTimeout); err != nil {
	// 	return nil, fmt.Errorf("failed to initialize fork height configs: %w", err)
	// }

	// Ensure that the genesis block hash of L1 and L2 match.
	//TODO(gustavo): Check if we should add this back
	// if err := c.ensureGenesisMatched(ctxWithTimeout, cfg.NewTaikoInboxAddress); err != nil {
	// 	return nil, fmt.Errorf("failed to ensure genesis block matched: %w", err)
	// }

	return c, nil
}

// initMinimalRollupClients initializes all minimal-rollup smart contract clients.
func (c *Client) initMinimalRollupClients(cfg *ClientConfig) error {
	minimalClients := &MinimalRollupClients{}
	var err error

	if minimalClients.NewTaikoInbox, err = minimalBindings.NewIInbox(cfg.NewTaikoInboxAddress, c.L1); err != nil {
		return fmt.Errorf("failed to initialize PublicationFeed client: %w", err)
	}

	// if cfg.ProverManagerAddress.Hex() != ZeroAddress.Hex() {
	// 	if minimalClients.ProverManager, err = minimalBindings.NewIProverManager(cfg.ProverManagerAddress, c.L1); err != nil {
	// 		return fmt.Errorf("failed to initialize ProverManager client: %w", err)
	// 	}
	// }

	if cfg.ProposerFeesAddress.Hex() != ZeroAddress.Hex() {
		if minimalClients.ProposerFees, err = minimalBindings.NewIProposerFees(cfg.ProposerFeesAddress, c.L1); err != nil {
			return fmt.Errorf("failed to initialize ProposerFees client: %w", err)
		}
	}

	if cfg.VerifierAddress.Hex() != ZeroAddress.Hex() {
		if minimalClients.Verifier, err = minimalBindings.NewIVerifier(cfg.VerifierAddress, c.L1); err != nil {
			return fmt.Errorf("failed to initialize Verifier client: %w", err)
		}
	}

	if cfg.LookaheadAddress.Hex() != ZeroAddress.Hex() {
		if minimalClients.Lookahead, err = minimalBindings.NewILookahead(cfg.LookaheadAddress, c.L1); err != nil {
			return fmt.Errorf("failed to initialize Lookahead client: %w", err)
		}
	}

	if cfg.InboxAddress.Hex() != ZeroAddress.Hex() {
		if minimalClients.Inbox, err = minimalBindings.NewIInbox(cfg.InboxAddress, c.L1); err != nil {
			return fmt.Errorf("failed to initialize Inbox client: %w", err)
		}
	}

	c.MinimalRollupClients = minimalClients
	return nil
}

// initForkHeightConfigs initializes the fork heights in protocol.
// func (c *Client) initForkHeightConfigs(ctx context.Context) error {
// 	protocolConfigs, err := c.PacayaClients.TaikoInbox.PacayaConfig(&bind.CallOpts{Context: ctx})
// 	if err != nil {
// 		return err
// 	}
//
// 	c.PacayaClients.ForkHeights = &pacayaBindings.ITaikoInboxForkHeights{
// 		Ontake: protocolConfigs.ForkHeights.Ontake,
// 		Pacaya: protocolConfigs.ForkHeights.Pacaya,
// 	}
//
// 	return nil
// }
