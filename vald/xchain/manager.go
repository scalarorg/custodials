package xchain

import (
	"context"
	"fmt"

	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/scalar-core/sdk-utils/broadcast"
	"github.com/scalarorg/scalar-core/utils/log"
	xcommon "github.com/scalarorg/scalar-core/vald/xchain/common"

	sdkClient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/chains/types"
	vote "github.com/scalarorg/scalar-core/x/vote/exported"
)

// Manager manages all communication with Ethereum
type Manager struct {
	rpcs                      map[chain.ChainInfoBytes]xcommon.Client
	broadcaster               broadcast.Broadcaster
	validator                 sdk.ValAddress
	proxy                     sdk.AccAddress
	latestFinalizedBlockCache xcommon.LatestFinalizedBlockCache
}

// NewManager returns a new Manager instance
func NewManager(
	clientCtx sdkClient.Context,
	rpcs map[chain.ChainInfoBytes]xcommon.Client,
	broadcaster broadcast.Broadcaster,
	valAddr sdk.ValAddress,
) *Manager {
	return &Manager{
		rpcs:        rpcs,
		broadcaster: broadcaster,
		validator:   valAddr,
		proxy:       clientCtx.FromAddress,
	}
}

func (mgr Manager) ProcessSourceTxsConfirmation(event *types.EventConfirmSourceTxsStarted) error {
	if !mgr.isParticipantOf(event.Participants) {
		pollIDs := slices.Map(event.PollMappings, func(m types.PollMapping) vote.PollID { return m.PollID })
		mgr.logger("poll_ids", pollIDs).Debug("ignoring staking txs confirmation poll: not a participant")
		return nil
	}

	mgr.logger("event", event).Debug("processing staking txs confirmation poll")

	chainInfoBytes := chain.ChainInfoBytes{}

	err := chainInfoBytes.FromString(event.Chain.String())
	if err != nil {
		return err
	}

	client, ok := mgr.rpcs[chainInfoBytes]
	if !ok {
		return fmt.Errorf("rpc client not found for chain %s", event.Chain.String())
	}

	votes, err := client.ProcessSourceTxsConfirmation(event, mgr.proxy)
	if err != nil {
		return err
	}

	_, err = mgr.broadcaster.Broadcast(context.TODO(), votes...)

	return err
}

// isParticipantOf checks if the validator is in the poll participants list
func (mgr Manager) isParticipantOf(participants []sdk.ValAddress) bool {
	return slices.Any(participants, func(v sdk.ValAddress) bool { return v.Equals(mgr.validator) })
}

func (mgr Manager) logger(keyvals ...any) log.Logger {
	keyvals = append([]any{"listener", "btc"}, keyvals...)
	return log.WithKeyVals(keyvals...)
}
