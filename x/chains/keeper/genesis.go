package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/funcs"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"

	"github.com/scalarorg/scalar-core/x/chains/types"
)

// InitGenesis initializes the state from a genesis file
func (k BaseKeeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	// TODO: add btc genesis
	for _, chain := range state.Chains {
		clog.Red("InitGenesis", "chain", chain.Params.Chain)
		funcs.MustNoErr(k.CreateChain(ctx, chain.Params))
		ck := funcs.Must(k.ForChain(ctx, chain.Params.Chain)).(chainKeeper)

		if err := ck.validateCommandQueueState(chain.CommandQueue, commandQueueName); err != nil {
			panic(err)
		}

		ck.getCommandQueue(ctx).ImportState(chain.CommandQueue)

		for _, sourceTx := range chain.ConfirmedSourceTxs {
			ck.SetSourceTx(ctx, sourceTx, types.SourceTxStatus_Confirmed)
		}

		var latestBatch types.CommandBatchMetadata
		for _, batch := range chain.CommandBatches {
			ck.setCommandBatchMetadata(ctx, batch)
			latestBatch = batch
		}

		if latestBatch.Status != types.BatchNonExistent {
			ck.setLatestBatchMetadata(ctx, latestBatch)
			ck.SetLatestSignedCommandBatchID(ctx, latestBatch.ID)
		}

		if chain.Gateway != nil {
			ck.setGateway(ctx, *chain.Gateway)
		}

		for _, token := range chain.Tokens {
			ck.setTokenMetadata(ctx, token)
		}

		for _, event := range chain.Events {
			ck.setEvent(ctx, event)
		}

		if err := ck.validateConfirmedEventQueueState(chain.ConfirmedEventQueue, confirmedEventQueueName); err != nil {
			panic(err)
		}
		ck.GetConfirmedEventQueue(ctx).(utils.GeneralKVQueue).ImportState(chain.ConfirmedEventQueue)
	}
}

// ExportGenesis generates a genesis file from the state
func (k BaseKeeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	return *types.NewGenesisState(k.getChains(ctx))
}

func (k BaseKeeper) getChains(ctx sdk.Context) []types.GenesisState_Chain {
	iter := k.getBaseStore(ctx).Iterator(utils.KeyFromStr(subspacePrefix))
	defer utils.CloseLogError(iter, k.Logger(ctx))

	var chains []types.GenesisState_Chain
	for ; iter.Valid(); iter.Next() {
		ck := funcs.Must(k.ForChain(ctx, nexus.ChainName(iter.Value()))).(chainKeeper)

		chain := types.GenesisState_Chain{
			Params:              ck.GetParams(ctx),
			CommandQueue:        ck.getCommandQueue(ctx).ExportState(),
			ConfirmedSourceTxs:  ck.getConfirmedSourceTxs(ctx),
			CommandBatches:      ck.getCommandBatchesMetadata(ctx),
			Events:              ck.getEvents(ctx),
			ConfirmedEventQueue: ck.GetConfirmedEventQueue(ctx).(utils.GeneralKVQueue).ExportState(),
		}
		chains = append(chains, chain)
	}

	return chains
}
