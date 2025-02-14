package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/spf13/cobra"

	nexusTypes "github.com/scalarorg/scalar-core/x/nexus/types"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"

	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

// func AddGenesisBTCChainCmd(defaultNodeHome string) *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "add-genesis-btc-chain [name]",
// 		Short: "Adds an BTC chain in genesis.json",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx := client.GetClientContextFromCmd(cmd)
// 			cdc := clientCtx.Codec

// 			serverCtx := server.GetServerContextFromCmd(cmd)
// 			config := serverCtx.Config

// 			config.SetRoot(clientCtx.HomeDir)

// 			name := nexusExported.ChainName(args[0])
// 			err := name.Validate()
// 			if err != nil {
// 				return err
// 			}

// 			genFile := config.GenesisFile()
// 			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
// 			if err != nil {
// 				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
// 			}

// 			chain := nexusExported.Chain{
// 				Name:                  name,
// 				SupportsForeignAssets: true,
// 				KeyType:               tss.Multisig,
// 				Module:                btctypes.ModuleName,
// 			}

// 			if err := chain.Validate(); err != nil {
// 				return err
// 			}

// 			genesisState := nexusTypes.GetGenesisStateFromAppState(cdc, appState)
// 			genesisState.Chains = append(genesisState.Chains, chain)

// 			genesisStateBz, err := cdc.MarshalJSON(&genesisState)
// 			if err != nil {
// 				return fmt.Errorf("failed to marshal nexus genesis state: %w", err)
// 			}

// 			appState[nexusTypes.ModuleName] = genesisStateBz
// 			appStateJSON, err := json.Marshal(appState)
// 			if err != nil {
// 				return fmt.Errorf("failed to marshal application genesis state: %w", err)
// 			}

// 			genDoc.AppState = appStateJSON
// 			return genutil.ExportGenesisFile(genDoc, genFile)
// 		}}

// 	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "node's home directory")
// 	return cmd
// }

// AddGenesisEVMChainCmd returns set-genesis-chain cobra Command.
func AddGenesisEVMChainCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-evm-chain [name]",
		Short: "Adds an EVM chain in genesis.json",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			name := nexus.ChainName(args[0])
			err := name.Validate()
			if err != nil {
				return err
			}

			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			chain := nexus.Chain{
				Name:                  name,
				SupportsForeignAssets: true,
				KeyType:               tss.Multisig,
				Module:                chainsTypes.ModuleName,
			}

			if err := chain.Validate(); err != nil {
				return err
			}

			genesisState := nexusTypes.GetGenesisStateFromAppState(cdc, appState)
			genesisState.Chains = append(genesisState.Chains, chain)

			genesisStateBz, err := cdc.MarshalJSON(&genesisState)
			if err != nil {
				return fmt.Errorf("failed to marshal nexus genesis state: %w", err)
			}

			appState[nexusTypes.ModuleName] = genesisStateBz
			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		}}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "node's home directory")
	return cmd
}
