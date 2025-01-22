package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/scalarorg/scalar-core/app/params"
	"github.com/scalarorg/scalar-core/testutils/fake"
	"github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/utils"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsKeepr "github.com/scalarorg/scalar-core/x/chains/keeper"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/x/nexus/types"
	testutils "github.com/scalarorg/scalar-core/x/nexus/types/testutils"
	scalarnet "github.com/scalarorg/scalar-core/x/scalarnet/exported"
	ScalarnetKeeper "github.com/scalarorg/scalar-core/x/scalarnet/keeper"
	scalarnetTypes "github.com/scalarorg/scalar-core/x/scalarnet/types"
	"github.com/scalarorg/scalar-core/x/scalarnet/types/mock"
)

func setup() (sdk.Context, Keeper) {
	ctx := sdk.NewContext(fake.NewMultiStore(), tmproto.Header{}, false, log.TestingLogger())
	encodingConfig := params.MakeEncodingConfig()
	types.RegisterLegacyAminoCodec(encodingConfig.Amino)
	types.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	subspace := paramstypes.NewSubspace(encodingConfig.Codec, encodingConfig.Amino, sdk.NewKVStoreKey("paramsKey"), sdk.NewKVStoreKey("tparamsKey"), "nexus")

	keeper := NewKeeper(
		encodingConfig.Codec,
		sdk.NewKVStoreKey(types.StoreKey),
		subspace,
	)

	scalarnetK := &mock.BaseKeeperMock{
		GetCosmosChainByNameFunc: func(ctx sdk.Context, chain exported.ChainName) (scalarnetTypes.CosmosChain, bool) {
			return scalarnetTypes.CosmosChain{Name: scalarnet.Scalarnet.Name, AddrPrefix: ""}, true
		},
	}

	addressValidators := types.NewAddressValidators()
	addressValidators.AddAddressValidator(chainsTypes.ModuleName, chainsKeepr.NewAddressValidator()).
		AddAddressValidator(scalarnetTypes.ModuleName, ScalarnetKeeper.NewAddressValidator(scalarnetK))

	addressValidators.Seal()
	keeper.SetAddressValidators(addressValidators)

	return ctx, keeper
}

func getRandomScalarnetAddress() exported.CrossChainAddress {
	sdk.GetConfig().SetBech32PrefixForAccount("", "")
	return exported.CrossChainAddress{
		Chain:   scalarnet.Scalarnet,
		Address: rand.AccAddr().String(),
	}
}

func getRandomEthereumAddress() exported.CrossChainAddress {
	return exported.CrossChainAddress{
		Chain:   chains.Ethereum,
		Address: common.BytesToAddress(rand.Bytes(common.AddressLength)).Hex(),
	}
}

func getRandomMessage(id string) exported.GeneralMessage {

	return exported.GeneralMessage{
		ID:          id,
		Sender:      getRandomScalarnetAddress(),
		Recipient:   getRandomEthereumAddress(),
		Status:      exported.Processing,
		PayloadHash: crypto.Keccak256Hash(rand.Bytes(int(rand.I64Between(1, 100)))).Bytes(),
		SourceTxID:  rand.Bytes(32),
		Asset:       nil,
	}

}

func assertChainStatesEqual(t *testing.T, expected, actual *types.GenesisState) {
	assert.Equal(t, expected.Params, actual.Params)
	assert.Equal(t, expected.Nonce, actual.Nonce)
	assert.ElementsMatch(t, expected.Chains, actual.Chains)
	assert.ElementsMatch(t, expected.ChainStates, actual.ChainStates)
	assert.ElementsMatch(t, expected.LinkedAddresses, actual.LinkedAddresses)
	assert.ElementsMatch(t, expected.Transfers, actual.Transfers)
	assert.Equal(t, expected.Fee, actual.Fee)
	assert.ElementsMatch(t, expected.FeeInfos, actual.FeeInfos)
	assert.ElementsMatch(t, expected.RateLimits, actual.RateLimits)
	assert.ElementsMatch(t, expected.Messages, actual.Messages)
	assert.Equal(t, expected.MessageNonce, actual.MessageNonce)
	// TODO: Track this with some random transfers
	// assert.ElementsMatch(t, expected.TransferEpochs, actual.TransferEpochs)
}

func TestExportGenesisInitGenesis(t *testing.T) {
	ctx, keeper := setup()

	keeper.InitGenesis(ctx, types.DefaultGenesisState())

	expected := types.DefaultGenesisState()

	if err := keeper.RegisterFee(ctx, scalarnet.Scalarnet, testutils.RandFee(scalarnet.Scalarnet.Name, scalarnet.NativeAsset)); err != nil {
		panic(err)
	}

	if err := keeper.RegisterAsset(ctx, chains.Ethereum, exported.NewAsset(scalarnet.NativeAsset, false), utils.MaxUint, time.Hour); err != nil {
		panic(err)
	}
	if err := keeper.RegisterFee(ctx, chains.Ethereum, testutils.RandFee(chains.Ethereum.Name, scalarnet.NativeAsset)); err != nil {
		panic(err)
	}

	rateLimit := testutils.RandRateLimit(scalarnet.Scalarnet.Name, scalarnet.NativeAsset)
	funcs.MustNoErr(keeper.SetRateLimit(ctx, rateLimit.Chain, rateLimit.Limit, rateLimit.Window))
	expected.RateLimits = keeper.getRateLimits(ctx)

	for _, chain := range expected.Chains {
		keeper.ActivateChain(ctx, chain)
	}

	linkedAddressesCount := rand.I64Between(100, 200)
	expectedLinkedAddresses := make([]types.LinkedAddresses, linkedAddressesCount)
	for i := 0; i < int(linkedAddressesCount); i++ {
		depositAddress := getRandomScalarnetAddress()
		recipientAddress := getRandomEthereumAddress()

		if err := keeper.LinkAddresses(ctx, depositAddress, recipientAddress); err != nil {
			panic(err)
		}
		expectedLinkedAddresses[i] = types.NewLinkedAddresses(depositAddress, recipientAddress)
	}
	expected.LinkedAddresses = expectedLinkedAddresses

	expected.Nonce = uint64(linkedAddressesCount)
	for i, linkedAddress := range expectedLinkedAddresses {
		depositAddress := linkedAddress.DepositAddress
		recipientAddress := linkedAddress.RecipientAddress
		_, minFee, maxFee, err := keeper.getCrossChainFees(ctx, depositAddress.Chain, recipientAddress.Chain, scalarnet.NativeAsset)
		assert.Nil(t, err)

		asset := sdk.NewCoin(scalarnet.NativeAsset, testutils.RandInt(minFee.Int64()/2, maxFee.Int64()*2))
		fees, err := keeper.ComputeTransferFee(ctx, depositAddress.Chain, recipientAddress.Chain, asset)
		assert.Nil(t, err)

		_, err = keeper.EnqueueForTransfer(
			ctx,
			depositAddress,
			asset,
		)
		if err != nil {
			panic(err)
		}

		if asset.Amount.LTE(fees.Amount) {
			expectedTransfer := exported.NewCrossChainTransfer(uint64(i), recipientAddress, asset, exported.InsufficientAmount)
			expected.Transfers = append(expected.Transfers, expectedTransfer)
			continue
		}

		expectedTransfer := exported.NewPendingCrossChainTransfer(uint64(i), recipientAddress, asset.Sub(fees))
		if rand.Bools(0.5).Next() {
			keeper.ArchivePendingTransfer(ctx, expectedTransfer)
			expectedTransfer.State = exported.Archived
		}

		expected.Transfers = append(expected.Transfers, expectedTransfer)

		expected.Fee.Coins = expected.Fee.Coins.Add(fees)
	}

	expected.ChainStates = []types.ChainState{
		{
			Chain:     scalarnet.Scalarnet,
			Assets:    []exported.Asset{exported.NewAsset(scalarnet.NativeAsset, true)},
			Activated: true,
		},
		{
			Chain:     chains.Ethereum,
			Assets:    []exported.Asset{exported.NewAsset(scalarnet.NativeAsset, false)},
			Activated: true,
		},
	}

	for _, chainState := range expected.ChainStates {
		for _, asset := range chainState.Assets {
			feeInfo, found := keeper.getFeeInfo(ctx, chainState.Chain, asset.Denom)
			if found {
				expected.FeeInfos = append(expected.FeeInfos, feeInfo)
			}
		}
	}

	messageCount := rand.I64Between(100, 256)
	for i := 0; i < int(messageCount); i++ {
		id, _, _ := keeper.GenerateMessageID(ctx)
		msg := getRandomMessage(id)
		expected.Messages = append(expected.Messages, msg)
		funcs.MustNoErr(keeper.setMessage(ctx, msg))
	}
	expected.MessageNonce = uint64(messageCount)

	actual := keeper.ExportGenesis(ctx)

	assert.NoError(t, actual.Validate())
	assertChainStatesEqual(t, expected, actual)

	ctx, keeper = setup()
	keeper.InitGenesis(ctx, expected)
	actual = keeper.ExportGenesis(ctx)

	assert.NoError(t, actual.Validate())
	assertChainStatesEqual(t, expected, actual)
}
