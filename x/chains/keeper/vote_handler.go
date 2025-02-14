package keeper

import (
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/funcs"

	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
	vote "github.com/scalarorg/scalar-core/x/vote/exported"
)

var _ vote.VoteHandler = &voteHandler{}

type voteHandler struct {
	cdc      codec.Codec
	keeper   types.BaseKeeper
	nexus    types.Nexus
	rewarder types.Rewarder
}

// NewVoteHandler returns the handler for processing vote delivered by the vote module
func NewVoteHandler(cdc codec.Codec, keeper types.BaseKeeper, nexus types.Nexus, rewarder types.Rewarder) vote.VoteHandler {
	return voteHandler{
		cdc:      cdc,
		keeper:   keeper,
		nexus:    nexus,
		rewarder: rewarder,
	}
}

func (v voteHandler) HandleFailedPoll(ctx sdk.Context, poll vote.Poll) error {
	md := mustGetMetadata(poll)
	events.Emit(ctx, &types.PollFailed{
		TxID:   md.TxID,
		Chain:  md.Chain,
		PollID: poll.GetID(),
	})

	return nil
}

func (v voteHandler) IsFalsyResult(result codec.ProtoMarshaler) bool {
	return len(result.(*types.VoteEvents).Events) == 0
}

func (v voteHandler) HandleExpiredPoll(ctx sdk.Context, poll vote.Poll) error {
	rewardPoolName, ok := poll.GetRewardPoolName()
	if !ok {
		return fmt.Errorf("reward pool not set for poll %s", poll.GetID().String())
	}

	md := mustGetMetadata(poll)
	rewardPool := v.rewarder.GetPool(ctx, rewardPoolName)
	chain, ok := v.nexus.GetChain(ctx, md.Chain)
	if !ok {
		return fmt.Errorf("%s is not a registered chain", md.Chain)
	}
	// Penalize voters who failed to vote
	for _, voter := range poll.GetVoters() {
		hasVoted := poll.HasVoted(voter)
		if maintainerState, ok := v.nexus.GetChainMaintainerState(ctx, chain, voter); ok {
			maintainerState.MarkMissingVote(!hasVoted)
			funcs.MustNoErr(v.nexus.SetChainMaintainerState(ctx, maintainerState))

			msg := fmt.Sprintf("marked voter %s behaviour", voter.String())
			clog.Red("HandleExpiredPoll", msg)
			v.keeper.Logger(ctx).Debug(msg,
				"voter", voter.String(),
				"missing_vote", !hasVoted,
				"poll", poll.GetID().String(),
			)
		}

		if !hasVoted {
			rewardPool.ClearRewards(voter)
			msg := fmt.Sprintf("penalized voter %s due to timeout", voter.String())
			clog.Red("HandleExpiredPoll", msg)
			v.keeper.Logger(ctx).Debug(msg,
				"voter", voter.String(),
				"poll", poll.GetID().String())
		}
	}

	events.Emit(ctx, &types.PollExpired{
		TxID:   md.TxID,
		Chain:  md.Chain,
		PollID: poll.GetID(),
	})

	return nil
}

func (v voteHandler) HandleCompletedPoll(ctx sdk.Context, poll vote.Poll) error {

	clog.Red("HandleCompletedPoll", "poll", poll.GetID().String())

	voteEvents := poll.GetResult().(*types.VoteEvents)

	chain, ok := v.nexus.GetChain(ctx, voteEvents.Chain)
	if !ok {
		return fmt.Errorf("%s is not a registered chain", voteEvents.Chain)
	}

	rewardPoolName, ok := poll.GetRewardPoolName()
	if !ok {
		return fmt.Errorf("reward pool not set for poll %s", poll.GetID().String())
	}

	rewardPool := v.rewarder.GetPool(ctx, rewardPoolName)

	for _, voter := range poll.GetVoters() {
		maintainerState, ok := v.nexus.GetChainMaintainerState(ctx, chain, voter)
		if !ok {
			continue // voter is no longer a chain maintainer, so recording the state is irrelevant
		}

		hasVoted := poll.HasVoted(voter)
		hasVotedIncorrectly := hasVoted && !poll.HasVotedCorrectly(voter)

		maintainerState.MarkMissingVote(!hasVoted)
		maintainerState.MarkIncorrectVote(hasVotedIncorrectly)
		funcs.MustNoErr(v.nexus.SetChainMaintainerState(ctx, maintainerState))

		msg := fmt.Sprintf("marked voter %s behaviour", voter.String())
		clog.Red("HandleCompletedPoll", msg)
		v.keeper.Logger(ctx).Debug(msg,
			"voter", voter.String(),
			"missing_vote", !hasVoted,
			"incorrect_vote", hasVotedIncorrectly,
			"poll", poll.GetID().String(),
		)

		switch {
		case hasVotedIncorrectly, !hasVoted:
			rewardPool.ClearRewards(voter)
			msg := fmt.Sprintf("penalized voter %s due to incorrect vote or missing vote", voter.String())
			clog.Red("HandleCompletedPoll", msg)
			v.keeper.Logger(ctx).Debug(msg,
				"voter", voter.String(),
				"poll", poll.GetID().String())
		default:
			if err := rewardPool.ReleaseRewards(voter); err != nil {
				return err
			}
			msg := fmt.Sprintf("released rewards for voter %s", voter.String())
			clog.Red("HandleCompletedPoll", msg)
			v.keeper.Logger(ctx).Debug(msg,
				"voter", voter.String(),
				"poll", poll.GetID().String())
		}
	}

	md := mustGetMetadata(poll)
	if v.IsFalsyResult(voteEvents) {
		events.Emit(ctx, &types.NoEventsConfirmed{
			TxID:   md.TxID,
			Chain:  md.Chain,
			PollID: poll.GetID(),
		})
	}

	event := &types.PollCompleted{
		TxID:   md.TxID,
		Chain:  md.Chain,
		PollID: poll.GetID(),
	}

	clog.Red("Poll Completed Event", event)

	events.Emit(ctx, event)

	return nil
}

func (v voteHandler) HandleResult(ctx sdk.Context, result codec.ProtoMarshaler) error {
	voteEvents := result.(*types.VoteEvents)

	if v.IsFalsyResult(result) {
		return nil
	}

	chain, ok := v.nexus.GetChain(ctx, voteEvents.Chain)
	if !ok {
		return fmt.Errorf("%s is not a registered chain", voteEvents.Chain)
	}

	ck, err := v.keeper.ForChain(ctx, chain.Name)
	if err != nil {
		return fmt.Errorf("%s is not an evm chain", voteEvents.Chain)
	}

	for _, event := range voteEvents.Events {
		if err := v.handleEvent(ctx, ck, event, chain); err != nil {
			return err
		}
	}

	return nil
}

func (v voteHandler) handleEvent(ctx sdk.Context, ck types.ChainKeeper, event types.Event, chain nexus.Chain) error {
	if err := ck.SetConfirmedEvent(ctx, event); err != nil {
		return err
	}

	// if not match the event type, the event also set confirmed in the abci of this module, and also enqueued to the nexus message
	// but it was not routed to the destination chain in the nexus keeper, so it will be ignored in this module

	eventType := event.GetEvent()
	switch eventType.(type) {
	case *types.Event_SourceTxConfirmationEvent:
		if err := v.handleCrossChainEvent(ctx, ck, event); err != nil {
			return err
		}
	default:
		funcs.MustNoErr(ck.EnqueueConfirmedEvent(ctx, event.GetID()))
	}

	ck.Logger(ctx).Info(fmt.Sprintf("confirmed %s event %s in transaction %s", chain.Name, event.GetID(), event.TxID.Hex()))

	return nil
}

func (v voteHandler) handleCrossChainEvent(ctx sdk.Context, ck types.ChainKeeper, event types.Event) error {
	// msg := mustToGeneralMessage(ctx, v.nexus, event)
	msg := mustToGeneralMessageWithPayload(ctx, v.nexus, event)

	clog.Bluef("msg: %+v", msg)
	if err := v.nexus.SetNewMessage(ctx, msg); err != nil {
		clog.Redf("[Chain] SetNewMessage error: %+v", err)
		return err
	}

	clog.Blue("handleCrossChainEvent, enqueueRouteMessage", "msg", msg.ID)
	clog.Blue("handleCrossChainEvent, setEventCompleted", "event", event.GetID())

	// this enqueues the message to be routed to the destination chain, it will be handled at abci of the destination module, for example evm module
	funcs.MustNoErr(v.nexus.EnqueueRouteMessage(ctx, msg.ID))
	funcs.MustNoErr(ck.SetEventCompleted(ctx, event.GetID()))

	return nil
}

// createGeneralMessageBase creates the common components of a general message
func createGeneralMessageBase(ctx sdk.Context, n types.Nexus, event types.Event) (
	string,
	nexus.CrossChainAddress,
	nexus.CrossChainAddress,
	*types.SourceTxConfirmationEvent,
	error) {

	id := string(event.GetID())

	confirmEvent, ok := event.GetEvent().(*types.Event_SourceTxConfirmationEvent)
	if !ok {
		return "", nexus.CrossChainAddress{}, nexus.CrossChainAddress{}, nil,
			fmt.Errorf("invalid event type: expected SourceTxConfirmationEvent")
	}
	confirmationEvent := confirmEvent.SourceTxConfirmationEvent

	sourceChain, ok := n.GetChain(ctx, event.Chain)
	if !ok {
		return "", nexus.CrossChainAddress{}, nexus.CrossChainAddress{}, nil,
			fmt.Errorf("source chain not found: %s", event.Chain)
	}
	sender := nexus.CrossChainAddress{Chain: sourceChain, Address: confirmationEvent.Sender}

	destinationChain, ok := n.GetChain(ctx, confirmationEvent.DestinationChain)
	if !ok {
		// Default to wasm router for unregistered chains
		destinationChain = nexus.Chain{
			Name:                  nexus.ChainName(confirmationEvent.DestinationChain.String()),
			SupportsForeignAssets: false,
			KeyType:               tss.None,
			Module:                wasm.ModuleName,
		}
	}
	recipient := nexus.CrossChainAddress{Chain: destinationChain, Address: confirmationEvent.DestinationContractAddress}

	return id, sender, recipient, confirmationEvent, nil
}

func mustToGeneralMessageWithPayload(ctx sdk.Context, n types.Nexus, event types.Event) nexus.GeneralMessage {
	id, sender, recipient, confirmationEvent, err := createGeneralMessageBase(ctx, n, event)
	if err != nil {
		panic(err)
	}

	return nexus.NewGeneralMessageWithPayload(
		id,
		sender,
		recipient,
		confirmationEvent.PayloadHash.Bytes(),
		event.TxID.Bytes(),
		event.Index,
		nil,
		confirmationEvent.Payload,
	)
}

func mustToGeneralMessage(ctx sdk.Context, n types.Nexus, event types.Event) nexus.GeneralMessage {
	id, sender, recipient, confirmationEvent, err := createGeneralMessageBase(ctx, n, event)
	if err != nil {
		panic(err)
	}

	return nexus.NewGeneralMessage(
		id,
		sender,
		recipient,
		confirmationEvent.PayloadHash.Bytes(),
		event.TxID.Bytes(),
		event.Index,
		nil,
	)
}

func mustGetMetadata(poll vote.Poll) types.PollMetadata {
	md := funcs.MustOk(poll.GetMetaData())
	metadata, ok := md.(*types.PollMetadata)
	if !ok {
		panic(fmt.Sprintf("poll metadata should be of type %T", &types.PollMetadata{}))
	}
	return *metadata
}
