package vald

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"golang.org/x/sync/errgroup"

	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/scalar-core/app"
	"github.com/scalarorg/scalar-core/cmd/scalard/cmd/utils"
	"github.com/scalarorg/scalar-core/sdk-utils/broadcast"
	tmEvents "github.com/scalarorg/scalar-core/tm-events/events"
	"github.com/scalarorg/scalar-core/tm-events/pubsub"
	"github.com/scalarorg/scalar-core/tm-events/tendermint"
	scalarUtils "github.com/scalarorg/scalar-core/utils"
	errors2 "github.com/scalarorg/scalar-core/utils/errors"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/jobs"
	"github.com/scalarorg/scalar-core/utils/log"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/vald/config"
	"github.com/scalarorg/scalar-core/vald/evm"
	evmRPC "github.com/scalarorg/scalar-core/vald/evm/rpc"
	"github.com/scalarorg/scalar-core/vald/multisig"
	grpc "github.com/scalarorg/scalar-core/vald/tofnd_grpc"
	"github.com/scalarorg/scalar-core/vald/tss"
	"github.com/scalarorg/scalar-core/vald/xchain"
	"github.com/scalarorg/scalar-core/vald/xchain/btc"
	btcTypes "github.com/scalarorg/scalar-core/x/btc/types"
	evmTypes "github.com/scalarorg/scalar-core/x/evm/types"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	scalarnet "github.com/scalarorg/scalar-core/x/scalarnet/exported"
	"github.com/scalarorg/scalar-core/x/tss/tofnd"
	tssTypes "github.com/scalarorg/scalar-core/x/tss/types"
)

// RW grants -rw------- file permissions
const RW = 0600

// RWX grants -rwx------ file permissions
const RWX = 0700

var once sync.Once
var cleanupCommands []func()

// GetValdCommand returns the command to start vald
func GetValdCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "vald-start",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			if !cmd.Flags().Changed(flags.FlagFrom) {
				if err := cmd.Flags().Set(flags.FlagFrom, serverCtx.Viper.GetString("broadcast.broadcaster-account")); err != nil {
					return err
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			logger := serverCtx.Logger.With("module", "vald")
			log.Setup(logger)
			v := serverCtx.Viper

			if err := v.BindPFlag("tss.tofnd-host", cmd.PersistentFlags().Lookup("tofnd-host")); err != nil {
				return err
			}
			if err := v.BindPFlag("tss.tofnd-port", cmd.PersistentFlags().Lookup("tofnd-port")); err != nil {
				return err
			}
			if err := v.BindPFlag("tss.tofnd-dial-timeout", cmd.PersistentFlags().Lookup("tofnd-dial-timeout")); err != nil {
				return err
			}

			cliCtx, err := sdkClient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// dynamically adjust gas limit by simulating the tx first
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithSimulateAndExecute(true)

			return runVald(cliCtx, txf, v)
		},
	}
	setPersistentFlags(cmd)
	flags.AddTxFlagsToCmd(cmd)
	values := map[string]string{
		flags.FlagGasAdjustment: "4",
	}
	utils.OverwriteFlagDefaults(cmd, values, true)

	// Only set default, not actual value, so it can be overwritten by env variable
	utils.OverwriteFlagDefaults(cmd, map[string]string{
		flags.FlagBroadcastMode:  flags.BroadcastSync,
		flags.FlagChainID:        app.Name,
		flags.FlagGasPrices:      "0.007" + scalarnet.BaseAsset,
		flags.FlagKeyringBackend: "file",
	}, false)

	return cmd
}

func runVald(cliCtx sdkClient.Context, txf tx.Factory, viper *viper.Viper) error {
	// in case of panic we still want to try and cleanup resources,
	// but we have to make sure it's not called more than once if the program is stopped by an interrupt signal
	defer once.Do(cleanUp)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Infof("captured signal \"%s\"", sig)
		once.Do(cleanUp)
	}()

	valdConf := config.DefaultValdConfig()
	viper.RegisterAlias("broadcast.max_timeout", "rpc.timeout_broadcast_tx_commit")
	if err := viper.Unmarshal(&valdConf, config.AddDecodeHooks); err != nil {
		panic(err)
	}

	valAddr, err := sdk.ValAddressFromBech32(viper.GetString("validator-addr"))
	if err != nil {
		return sdkerrors.Wrap(err, "invalid validator operator address")
	}

	valdHome := filepath.Join(cliCtx.HomeDir, "vald")
	if _, err := os.Stat(valdHome); os.IsNotExist(err) {
		log.Infof("folder %s does not exist, creating...", valdHome)
		err := os.Mkdir(valdHome, RWX)
		if err != nil {
			return err
		}
	}

	fPath := filepath.Join(valdHome, "state.json")
	stateSource := NewRWFile(fPath)

	log.Info("start listening to events")
	listen(cliCtx, txf, valdConf, valAddr, stateSource)
	log.Info("shutting down")
	return nil
}

func cleanUp() {
	for _, cmd := range cleanupCommands {
		cmd()
	}
}

func setPersistentFlags(cmd *cobra.Command) {
	defaultConf := tssTypes.DefaultConfig()
	cmd.PersistentFlags().String("tofnd-host", defaultConf.Host, "host name for tss daemon")
	cmd.PersistentFlags().String("tofnd-port", defaultConf.Port, "port for tss daemon")
	cmd.PersistentFlags().String("tofnd-dial-timeout", defaultConf.DialTimeout.String(), "dialup timeout to the tss daemon")
	cmd.PersistentFlags().String("validator-addr", "", "the address of the validator operator, i.e scalarvaloper1..")
	cmd.PersistentFlags().String(flags.FlagChainID, app.Name, "The network chain ID")
}

func listen(clientCtx sdkClient.Context, txf tx.Factory, scalarCfg config.ValdConfig, valAddr sdk.ValAddress, stateSource ReadWriter) {
	btcConfigs := scalarCfg.BTCConfig
	evmConfigs := scalarCfg.EVMConfig

	for _, btcConfig := range btcConfigs {
		err := btcConfig.ValidateBasic()
		if err != nil {
			panic(err)
		}
	}

	for _, evmConfig := range evmConfigs {
		err := evmConfig.ValidateBasic()
		if err != nil {
			panic(err)
		}
	}

	encCfg := app.MakeEncodingConfig()
	cdc := encCfg.Amino
	sender, err := clientCtx.Keyring.Key(clientCtx.From)
	if err != nil {
		panic(sdkerrors.Wrap(err, "failed to read broadcaster account info from keyring"))
	}
	clientCtx = clientCtx.
		WithFromAddress(sender.GetAddress()).
		WithFromName(sender.GetName())

	bc := createRefundableBroadcaster(txf, clientCtx, scalarCfg)

	robustClient := tendermint.NewRobustClient(func() (rpcclient.Client, error) {
		cl, err := sdkClient.NewClientFromNode(clientCtx.NodeURI)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create a new client")
		}

		err = cl.Start()
		if err != nil {
			return nil, errors.Wrap(err, "unable to start client")
		}
		return cl, nil
	})

	tssMgr := createTSSMgr(bc, clientCtx, scalarCfg, valAddr.String(), cdc)

	evmMgr := createEVMMgr(scalarCfg, clientCtx, bc, valAddr)

	xMgr := createXChainMgr(scalarCfg, clientCtx, bc, valAddr)

	multisigMgr := createMultisigMgr(bc, clientCtx, scalarCfg, valAddr)

	nodeHeight, err := waitUntilNetworkSync(scalarCfg, robustClient)
	if err != nil {
		panic(err)
	}

	stateStore := NewStateStore(stateSource)
	startBlock, err := getStartBlock(scalarCfg, stateStore, nodeHeight)
	if err != nil {
		panic(err)
	}

	eventBus := createEventBus(robustClient, startBlock, scalarCfg.EventNotificationsMaxRetries, scalarCfg.EventNotificationsBackOff)
	var blockHeight int64
	blockHeaderSub := eventBus.Subscribe(func(event tmEvents.ABCIEventWithHeight) bool {
		if event.Height != blockHeight {
			blockHeight = event.Height
			return true
		}
		return false
	})

	heartbeat := eventBus.Subscribe(func(e tmEvents.ABCIEventWithHeight) bool {
		event := tmEvents.Map(e)
		return event.Type == tssTypes.EventTypeHeartBeat &&
			event.Attributes[sdk.AttributeKeyModule] == tssTypes.ModuleName &&
			event.Attributes[sdk.AttributeKeyAction] == tssTypes.AttributeValueSend
	})

	evmNewChain := eventBus.Subscribe(tmEvents.Filter[*evmTypes.ChainAdded]())
	evmDepConf := eventBus.Subscribe(tmEvents.Filter[*evmTypes.ConfirmDepositStarted]())
	evmTokConf := eventBus.Subscribe(tmEvents.Filter[*evmTypes.ConfirmTokenStarted]())
	evmTraConf := eventBus.Subscribe(tmEvents.Filter[*evmTypes.ConfirmKeyTransferStarted]())
	evmGatewayTxConf := eventBus.Subscribe(tmEvents.Filter[*evmTypes.ConfirmGatewayTxStarted]())
	evmGatewayTxsConf := eventBus.Subscribe(tmEvents.Filter[*evmTypes.ConfirmGatewayTxsStarted]())

	multisigKeygen := eventBus.Subscribe(tmEvents.Filter[*multisigTypes.KeygenStarted]())
	multisigSigning := eventBus.Subscribe(tmEvents.Filter[*multisigTypes.SigningStarted]())

	// TODO: Version2: handle staking and unstaking events for multiple chains, currently it uses type of btc, we need to change it to more generic type
	v2StakingConf := eventBus.Subscribe(tmEvents.Filter[*btcTypes.EventConfirmStakingTxsStarted]())
	v2UnstakingConf := eventBus.Subscribe(tmEvents.Filter[*btcTypes.EventConfirmUnstakingTxsStarted]())

	eventCtx, cancelEventCtx := context.WithCancel(context.Background())
	eGroup, eventCtx := errgroup.WithContext(eventCtx)

	// stop the jobs if process gets interrupted/terminated
	cleanupCommands = append(cleanupCommands, func() {
		log.Info("stop listening for events...")
		cancelEventCtx()
		<-eventBus.Done()
		log.Info("event listener stopped")
		log.Info("stopping subscribers...")
		if err := eGroup.Wait(); err != nil {
			log.Error(err.Error())
		}
		log.Info("subscriptions stopped")
	})

	fetchEvents := func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return nil
		case err := <-eventBus.FetchEvents(ctx):
			cancelEventCtx()
			return err
		}
	}

	timer := time.AfterFunc(0, func() {})
	defer timer.Stop()
	blockTimeout, timeoutCancel := context.WithCancel(context.Background())
	processBlockHeader := func(event tmEvents.Event) error {
		timer.Stop()
		timer = time.AfterFunc(scalarCfg.NoNewBlockPanicTimeout, timeoutCancel)

		return stateStore.SetState(event.Height)
	}

	failOnTimeout := func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return nil
		case <-blockTimeout.Done():
			return errors.New("no new blocks received from the node")
		}
	}

	js := []jobs.Job{
		createJob(blockHeaderSub, processBlockHeader, cancelEventCtx),
		fetchEvents,
		failOnTimeout,
		createJob(heartbeat, tssMgr.ProcessHeartBeatEvent, cancelEventCtx),
		createJobTyped(evmNewChain, evmMgr.ProcessNewChain, cancelEventCtx),
		createJobTyped(evmDepConf, evmMgr.ProcessDepositConfirmation, cancelEventCtx),
		createJobTyped(evmTokConf, evmMgr.ProcessTokenConfirmation, cancelEventCtx),
		createJobTyped(evmTraConf, evmMgr.ProcessTransferKeyConfirmation, cancelEventCtx),
		createJobTyped(evmGatewayTxConf, evmMgr.ProcessGatewayTxConfirmation, cancelEventCtx),
		createJobTyped(evmGatewayTxsConf, evmMgr.ProcessGatewayTxsConfirmation, cancelEventCtx),
		createJobTyped(multisigKeygen, multisigMgr.ProcessKeygenStarted, cancelEventCtx),
		createJobTyped(multisigSigning, multisigMgr.ProcessSigningStarted, cancelEventCtx),

		createJobTyped(v2StakingConf, xMgr.ProcessStakingTxsConfirmation, cancelEventCtx),
		// createJobTyped(v2UnstakingConf, xMgr.ProcessUnstakingTxsConfirmation, cancelEventCtx),
	}

	_ = v2UnstakingConf

	slices.ForEach(js, func(job jobs.Job) {
		eGroup.Go(func() error { return job(eventCtx) })
	})

	if err := eGroup.Wait(); err != nil {
		log.Error(err.Error())
	}
}

func createJob(sub <-chan tmEvents.ABCIEventWithHeight, processor func(event tmEvents.Event) error, cancel context.CancelFunc) jobs.Job {
	return func(ctx context.Context) error {
		processWithLog := func(e tmEvents.ABCIEventWithHeight) {
			err := processor(tmEvents.Map(e))
			if err != nil {
				ctx = log.AppendKeyVals(ctx, errors2.KeyVals(err)...)
				log.FromCtx(ctx).Error(err.Error())
			}
		}
		consume := tmEvents.Consume(sub, processWithLog)
		err := consume(ctx)
		if err != nil {
			cancel()
			return err
		}
		return nil
	}
}

func createJobTyped[T proto.Message](sub <-chan tmEvents.ABCIEventWithHeight, processor func(event T) error, cancel context.CancelFunc) jobs.Job {
	return func(ctx context.Context) error {
		processWithLog := func(e tmEvents.ABCIEventWithHeight) {
			event := funcs.Must(sdk.ParseTypedEvent(e.Event)).(T)
			err := processor(event)
			if err != nil {
				ctx = log.AppendKeyVals(ctx, errors2.KeyVals(err)...)
				log.FromCtx(ctx).Error(err.Error())
			}
		}

		consume := tmEvents.Consume(sub, processWithLog)
		err := consume(ctx)
		if err != nil {
			cancel()
			return err
		}

		return nil
	}
}

// Wait until the node has synced with the network and return the node height
func waitUntilNetworkSync(cfg config.ValdConfig, tmClient tmEvents.SyncInfoClient) (int64, error) {
	for {
		rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		syncInfo, err := tmClient.LatestSyncInfo(rpcCtx)
		cancel()
		if err != nil {
			return 0, err
		}

		// If the block height is older than the allowed time, then wait for the node to sync
		if syncInfo.LatestBlockTime.Add(cfg.MaxLatestBlockAge).After(time.Now()) {
			return syncInfo.LatestBlockHeight, nil
		}

		log.Infof("node height %d is old, waiting for a recent block", syncInfo.LatestBlockHeight)
		time.Sleep(cfg.MaxLatestBlockAge)
	}
}

// Return the block height to start listening to TM events from
func getStartBlock(cfg config.ValdConfig, stateStore StateStore, nodeHeight int64) (int64, error) {
	storedHeight, err := stateStore.GetState()
	if err != nil {
		log.Infof("failed to retrieve the stored block height, using the latest: %s", err.Error())
		storedHeight = 0
	} else {
		log.Infof("retrieved stored block height %d", storedHeight)
	}

	// stored height must not be larger than node height
	if storedHeight > nodeHeight {
		return 0, fmt.Errorf("stored block height %d is ahead of the node height %d", storedHeight, nodeHeight)
	}

	log.Infof("node is synced, node height: %d", nodeHeight)

	startBlock := storedHeight
	if startBlock != 0 {
		// The block at the stored height might have already been processed by vald, so skip it
		startBlock++
	}

	if startBlock != 0 && nodeHeight-startBlock > cfg.MaxBlocksBehindLatest {
		log.Infof("stored block height %d is too old, starting from the latest block", startBlock)
		startBlock = 0
	}

	return startBlock, nil
}

func createEventBus(client *tendermint.RobustClient, startBlock int64, retries int, backOff time.Duration) *tmEvents.Bus {
	notifier := tmEvents.NewBlockNotifier(client, tmEvents.Retries(retries), tmEvents.BackOff(backOff)).StartingAt(startBlock)
	return tmEvents.NewEventBus(tmEvents.NewBlockSource(client, notifier, tmEvents.Retries(retries), tmEvents.BackOff(backOff)), pubsub.NewBus[tmEvents.ABCIEventWithHeight]())
}

func createRefundableBroadcaster(txf tx.Factory, ctx sdkClient.Context, valdCfg config.ValdConfig) broadcast.Broadcaster {
	broadcaster := broadcast.WithStateManager(ctx, txf, broadcast.WithResponseTimeout(valdCfg.BroadcastConfig.MaxTimeout))
	broadcaster = broadcast.WithRetry(broadcaster, valdCfg.MaxRetries, valdCfg.MinSleepBeforeRetry)
	broadcaster = broadcast.Batched(broadcaster, valdCfg.BatchThreshold, valdCfg.BatchSizeLimit)
	broadcaster = broadcast.WithRefund(broadcaster)
	broadcaster = broadcast.SuppressExecutionErrs(broadcaster)

	return broadcaster
}

func createMultisigMgr(broadcaster broadcast.Broadcaster, cliCtx sdkClient.Context, valdCfg config.ValdConfig, valAddr sdk.ValAddress) *multisig.Mgr {
	conn, err := grpc.Connect(valdCfg.TssConfig.Host, valdCfg.TssConfig.Port, valdCfg.TssConfig.DialTimeout)
	if err != nil {
		panic(sdkerrors.Wrap(err, "failed to create multisig manager"))
	}
	log.Debug("successful connection to tofnd gRPC server")

	return multisig.NewMgr(tofnd.NewMultisigClient(conn), cliCtx, valAddr, broadcaster, timeout)
}

func createTSSMgr(broadcaster broadcast.Broadcaster, cliCtx sdkClient.Context, valdCfg config.ValdConfig, valAddr string, cdc *codec.LegacyAmino) *tss.Mgr {
	create := func() (*tss.Mgr, error) {
		conn, err := tss.Connect(valdCfg.TssConfig.Host, valdCfg.TssConfig.Port, valdCfg.TssConfig.DialTimeout)
		if err != nil {
			return nil, err
		}
		log.Debug("successful connection to tofnd gRPC server")

		// creates client to communicate with the external tofnd process service
		multiSigClient := tofnd.NewMultisigClient(conn)

		tssMgr := tss.NewMgr(multiSigClient, cliCtx, 2*time.Hour, valAddr, broadcaster, cdc)

		return tssMgr, nil
	}

	mgr, err := create()
	if err != nil {
		panic(sdkerrors.Wrap(err, "failed to create tss manager"))
	}

	return mgr
}

func createEVMClient(config *evmTypes.EVMConfig) (evmRPC.Client, error) {
	return evmRPC.NewClient(config.RPCAddr, config.FinalityOverride)
}

func createEVMMgr(valdCfg config.ValdConfig, cliCtx sdkClient.Context, b broadcast.Broadcaster, valAddr sdk.ValAddress) *evm.Mgr {
	rpcs := make(map[string]evmRPC.Client)

	chainConfigs := slices.Filter(valdCfg.EVMConfig, func(config evmTypes.EVMConfig) bool {
		return config.WithBridge
	})

	slices.ForEach(chainConfigs, func(config evmTypes.EVMConfig) {
		chainName := strings.ToLower(config.Name)
		if _, ok := rpcs[chainName]; ok {
			err := fmt.Errorf("duplicate bridge configuration found for EVM chain %s", config.Name)
			log.Error(err.Error())
			panic(err)
		}

		if config.L1ChainName != nil {
			log.Infof("`l1_chain_name` config is deprecated for EVM chain '%s'. Please remove it from your RPC config", config.Name)
		}

		client, err := createEVMClient(&config)
		if err != nil {
			err = sdkerrors.Wrap(err, fmt.Sprintf("failed to create an RPC connection for EVM chain %s. Verify your RPC config.", config.Name))
			log.Error(err.Error())
			panic(err)
		}

		log.WithKeyVals("chain", config.Name, "url", config.RPCAddr).
			Debugf("created JSON-RPC client of type %T", client)

		// clean up evmRPC connection on process shutdown
		cleanupCommands = append(cleanupCommands, client.Close)

		rpcs[chainName] = client
		log.Infof("successfully connected to EVM bridge for chain %s", chainName)
	})

	return evm.NewMgr(rpcs, b, valAddr, cliCtx.FromAddress, evm.NewLatestFinalizedBlockCache())
}

// TODO: Refactor the name later, this mgr handles multiple chain confirmations
func createXChainMgr(valdCfg config.ValdConfig, cliCtx sdkClient.Context, b broadcast.Broadcaster, valAddr sdk.ValAddress) *xchain.Manager {
	rpcs := make(map[chain.ChainInfoBytes]xchain.Client)

	btcConfigs := slices.Filter(valdCfg.BTCConfig, func(config btcTypes.BTCConfig) bool {
		return config.WithBridge
	})

	// TODO: Add more chains here

	evmConfigs := slices.Filter(valdCfg.EVMConfig, func(config evmTypes.EVMConfig) bool {
		return config.WithBridge
	})

	_ = evmConfigs

	slices.ForEach(btcConfigs, func(config btcTypes.BTCConfig) {
		chainInfoBytes, err := scalarUtils.ChainInfoBytesFromID(config.ID)
		if err != nil {
			err := fmt.Errorf("invalid chain ID %s", config.ID)
			log.Error(err.Error())
			panic(err)
		}

		if _, ok := rpcs[chainInfoBytes]; ok {
			err := fmt.Errorf("duplicate bridge configuration found for BTC chain %s", config.Name)
			log.Error(err.Error())
			panic(err)
		}

		client, err := btc.NewClient(&config)
		if err != nil {
			err = sdkerrors.Wrap(err, fmt.Sprintf("failed to create an RPC connection for BTC chain %s. Verify your RPC config.", config.Name))
			log.Error(err.Error())
			panic(err)
		}

		log.WithKeyVals("chain", config.Name, "url", fmt.Sprintf("%s:%d", config.RpcHost, config.RpcPort)).
			Debugf("created JSON-RPC client of type %T", client)

		rpcs[chainInfoBytes] = client
	})
	return xchain.NewManager(cliCtx, rpcs, b, valAddr)
}

// RWFile implements the ReadWriter interface for an underlying file
type RWFile struct {
	path string
}

// NewRWFile returns a new RWFile instance for the given file path
func NewRWFile(path string) RWFile {
	return RWFile{path: path}
}

// ReadAll returns the full content of the file
func (f RWFile) ReadAll() ([]byte, error) { return os.ReadFile(f.path) }

// WriteAll writes the given bytes to a file. Creates a new file if it does not exist, overwrites the previous content otherwise.
func (f RWFile) WriteAll(bz []byte) error { return os.WriteFile(f.path, bz, RW) }
