package types

import (
	fmt "fmt"

	"github.com/scalarorg/bitcoin-vault/go-utils/types"
	utils "github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/log"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	KeyChainName           = []byte("chainName")
	KeyConfirmationHeight  = []byte("confirmationHeight")
	KeyNetworkKind         = []byte("networkKind")
	KeyRevoteLockingPeriod = []byte("revoteLockingPeriod")
	KeyChainID             = []byte("chainId")
	KeyVotingThreshold     = []byte("votingThreshold")
	KeyToken               = []byte("token")
	KeyBurnable            = []byte("burnable")
	KeyMinVoterCount       = []byte("minVoterCount")
	KeyCommandsGasLimit    = []byte("commandsGasLimit")
	KeyVotingGracePeriod   = []byte("votingGracePeriod")
	KeyEndBlockerLimit     = []byte("endBlockerLimit")
	KeyTransferLimit       = []byte("transferLimit")
	KeyMetadata            = []byte("metadata")
)

func KeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns the module's parameter set initialized with default values
func DefaultChainParams(chainId sdk.Int, chain nexus.ChainName, networkKind types.NetworkKind, metadata map[string]string) Params {
	bzToken, err := utils.HexDecode(Token)
	if err != nil {
		panic(err)
	}

	bzBurnable, err := utils.HexDecode(Burnable)
	if err != nil {
		panic(err)
	}
	return Params{
		ChainID:             chainId,
		Chain:               chain,
		ConfirmationHeight:  2,
		NetworkKind:         networkKind,
		TokenCode:           bzToken,
		Burnable:            bzBurnable,
		RevoteLockingPeriod: 50,
		VotingThreshold:     utils.Threshold{Numerator: 51, Denominator: 100},
		MinVoterCount:       1,
		CommandsGasLimit:    5000000,
		VotingGracePeriod:   50,
		EndBlockerLimit:     50,
		TransferLimit:       1000,
		Metadata:            metadata,
	}
}

func (Params) Validate() error {
	// TODO: Implement validation
	log.Debug("Not implemented params validation")
	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(KeyChainName, &m.Chain, validateChainName),
		params.NewParamSetPair(KeyConfirmationHeight, &m.ConfirmationHeight, validateConfirmationHeight),
		params.NewParamSetPair(KeyNetworkKind, &m.NetworkKind, validateNetworkKind),
		params.NewParamSetPair(KeyToken, &m.TokenCode, validateBytes),
		params.NewParamSetPair(KeyBurnable, &m.Burnable, validateBurnable),
		params.NewParamSetPair(KeyRevoteLockingPeriod, &m.RevoteLockingPeriod, validateRevoteLockingPeriod),
		params.NewParamSetPair(KeyChainID, &m.ChainID, validateChainId),
		params.NewParamSetPair(KeyVotingThreshold, &m.VotingThreshold, validateVotingThreshold),
		params.NewParamSetPair(KeyMinVoterCount, &m.MinVoterCount, validateMinVoterCount),
		params.NewParamSetPair(KeyCommandsGasLimit, &m.CommandsGasLimit, validateCommandsGasLimit),
		params.NewParamSetPair(KeyVotingGracePeriod, &m.VotingGracePeriod, validateVotingGracePeriod),
		params.NewParamSetPair(KeyEndBlockerLimit, &m.EndBlockerLimit, validateEndBlockerLimit),
		params.NewParamSetPair(KeyTransferLimit, &m.TransferLimit, validateTransferLimit),
		params.NewParamSetPair(KeyMetadata, &m.Metadata, validateMetadata),
	}
}

func validateChainName(i interface{}) error {
	chainName, ok := i.(nexus.ChainName)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return chainName.Validate()
}

func validateConfirmationHeight(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateNetworkKind(i interface{}) error {
	networkKind, ok := i.(types.NetworkKind)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if !networkKind.Valid() {
		return fmt.Errorf("invalid network kind: %v", networkKind)
	}
	return nil
}

func validateBytes(bytes interface{}) error {
	b, ok := bytes.([]byte)
	if !ok {
		return fmt.Errorf("invalid parameter type for byte slice: %T", bytes)
	}

	if len(b) == 0 {
		// Todo: Enable check burnable code it not empty
		// return fmt.Errorf("byte slice cannot be empty")
	}

	return nil
}

func validateRevoteLockingPeriod(i interface{}) error {
	period, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if period < 1 {
		return fmt.Errorf("revote locking period must be greater than 0")
	}
	return nil
}

func validateChainId(i interface{}) error {
	chainId, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if !chainId.IsPositive() {
		return fmt.Errorf("chain id must be positive")
	}
	return nil
}

func validateVotingThreshold(i interface{}) error {
	threshold, ok := i.(utils.Threshold)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return threshold.Validate()
}

func validateMinVoterCount(i interface{}) error {
	count, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if count < 1 {
		return fmt.Errorf("min voter count must be greater than 0")
	}
	return nil
}

func validateCommandsGasLimit(commandsGasLimit interface{}) error {
	val, ok := commandsGasLimit.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type for commands gas limit: %T", commandsGasLimit)
	}

	if val <= 0 {
		return fmt.Errorf("commands gas limit must be >0")
	}

	return nil
}

func validateVotingGracePeriod(i interface{}) error {
	period, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if period < 1 {
		return fmt.Errorf("voting grace period must be greater than 0")
	}
	return nil
}

func validateBurnable(i interface{}) error {
	if err := validateBytes(i); err != nil {
		return err
	}
	// Todo: Turnon validateBurnerCode
	// if err := validateBurnerCode(i.([]byte)); err != nil {
	// 	return err
	// }

	return nil
}

func validateEndBlockerLimit(i interface{}) error {
	limit, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if limit < 1 {
		return fmt.Errorf("end blocker limit must be greater than 0")
	}
	return nil
}

func validateTransferLimit(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateMetadata(i interface{}) error {
	_, ok := i.(map[string]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
