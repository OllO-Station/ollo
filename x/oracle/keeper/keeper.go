package keeper

import (
	"fmt"

	"ollo/x/oracle/types"

	"github.com/cosmos/cosmos-sdk/codec"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"github.com/ignite/cli/ignite/pkg/cosmosibckeeper"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		*cosmosibckeeper.Keeper
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper cosmosibckeeper.ChannelKeeper,
	portKeeper cosmosibckeeper.PortKeeper,
	scopedKeeper cosmosibckeeper.ScopedKeeper,
	accountKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		Keeper: cosmosibckeeper.NewKeeper(
			types.PortKey,
			storeKey,
			channelKeeper,
			portKeeper,
			scopedKeeper,
		),
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper, bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// IsBound checks if the market module is already bound to the desired port.
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.ScopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the ort Keeper's function in
// order to expose it to module's InitGenesis function.
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	capPort := k.PortKeeper.BindPort(ctx, portID)
	return k.ScopedKeeper.ClaimCapability(ctx, capPort, host.PortPath(portID))
}

// GetPort returns the portID for the market module. Used in ExportGenesis.
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the market module. Used in InitGenesis.
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function.
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.ScopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the transfer module that can claim a capability that IBC module.
// passes to it.
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.ScopedKeeper.ClaimCapability(ctx, cap, name)
}
// )

// func (k Keeper) SetFetchPriceResult(ctx sdk.Context, requestID types.OracleRequestID, result types.FetchPriceResult) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Set(types.FetchPriceResultStoreKey(requestID), k.cdc.MustMarshal(&result))
// }

// // GetFetchPriceResult returns the FetchPrice by requestId.
// func (k Keeper) GetFetchPriceResult(ctx sdk.Context, id types.OracleRequestID) (types.FetchPriceResult, error) {
// 	bz := ctx.KVStore(k.storeKey).Get(types.FetchPriceResultStoreKey(id))
// 	if bz == nil {
// 		return types.FetchPriceResult{}, sdkerrors.Wrapf(types.ErrSample,
// 			"GetResult: Result for request ID %d is not available.", id,
// 		)
// 	}
// 	var result types.FetchPriceResult
// 	k.cdc.MustUnmarshal(bz, &result)
// 	return result, nil
// }

// // GetLastFetchPriceID return the id from the last FetchPrice request.
// func (k Keeper) GetLastFetchPriceID(ctx sdk.Context) int64 {
// 	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.LastFetchPriceIDKey))
// 	intV := gogotypes.Int64Value{}
// 	k.cdc.MustUnmarshalLengthPrefixed(bz, &intV)
// 	return intV.GetValue()
// }

// // SetLastFetchPriceID saves the id from the last FetchPrice request.
// func (k Keeper) SetLastFetchPriceID(ctx sdk.Context, id types.OracleRequestID) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Set(types.KeyPrefix(types.LastFetchPriceIDKey),
// 		k.cdc.MustMarshalLengthPrefixed(&gogotypes.Int64Value{Value: int64(id)}))
// }

// func (k Keeper) FetchPrice(ctx sdk.Context, msg types.MsgFetchPriceData) (*types.MsgFetchPriceDataResponse, error) {
// 	sourcePort := types.PortID
// 	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, msg.SourceChannel)
// 	if !found {
// 		return nil, nil
// 	}
// 	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
// 	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

// 	// get the next sequence
// 	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, msg.SourceChannel)
// 	if !found {
// 		return nil, nil
// 	}

// 	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, msg.SourceChannel))
// 	if !ok {
// 		return nil, nil
// 	}

// 	var symbol []string
// 	assets := k.GetAssets(ctx)
// 	for _, asset := range assets {
// 		if asset.IsOraclePriceRequired {
// 			symbol = append(symbol, asset.Name)
// 		}
// 	}

// 	encodedCallData := obi.MustEncode(types.FetchPriceCallData{Symbols: symbol, Multiplier: 1000000})

// 	packetData := packet.NewOracleRequestPacketData(
// 		msg.ClientID,
// 		msg.OracleScriptID,
// 		encodedCallData,
// 		msg.AskCount,
// 		msg.MinCount,
// 		msg.FeeLimit,
// 		msg.PrepareGas,
// 		msg.ExecuteGas,
// 	)
// 	err := k.channelKeeper.SendPacket(ctx, channelCap, channeltypes.NewPacket(
// 		packetData.GetBytes(),
// 		sequence,
// 		sourcePort,
// 		msg.SourceChannel,
// 		destinationPort,
// 		destinationChannel,
// 		clienttypes.NewHeight(0, 0),
// 		uint64(ctx.BlockTime().UnixNano()+int64(10*time.Minute)), // Arbitrary timestamp timeout for now
// 	))
// 	if err != nil {
// 		return nil, nil
// 	}

// 	return &types.MsgFetchPriceDataResponse{}, nil
// }

// func (k Keeper) SetFetchPriceMsg(ctx sdk.Context, msg types.MsgFetchPriceData) {
// 	var (
// 		store = ctx.KVStore(k.storeKey)
// 		key   = types.MsgDataKey
// 		v     = types.NewMsgFetchPriceData(
// 			types.ModuleName,
// 			types.OracleScriptID(msg.OracleScriptID),
// 			msg.SourceChannel,
// 			nil,
// 			msg.AskCount,
// 			msg.MinCount,
// 			msg.FeeLimit,
// 			msg.PrepareGas,
// 			msg.ExecuteGas,
// 		)
// 		value = k.cdc.MustMarshal(v)
// 	)

// 	store.Set(key, value)
// }

// func (k Keeper) GetFetchPriceMsg(ctx sdk.Context) types.MsgFetchPriceData {
// 	var (
// 		store = ctx.KVStore(k.storeKey)
// 		key   = types.MsgDataKey
// 		value = store.Get(key)
// 	)
// 	var msg types.MsgFetchPriceData
// 	if value != nil {
// 		k.cdc.MustUnmarshal(value, &msg)
// 	}

// 	return msg
// }

// func (k Keeper) GetLastBlockHeight(ctx sdk.Context) int64 {
// 	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.LastBlockHeightKey))
// 	intV := gogotypes.Int64Value{}
// 	k.cdc.MustUnmarshalLengthPrefixed(bz, &intV)
// 	return intV.GetValue()
// }

// func (k Keeper) SetLastBlockHeight(ctx sdk.Context, height int64) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Set(types.KeyPrefix(types.LastBlockHeightKey),
// 		k.cdc.MustMarshalLengthPrefixed(&gogotypes.Int64Value{Value: height}))
// }

// func (k Keeper) AddFetchPriceRecords(ctx sdk.Context, price types.MsgFetchPriceData) error {
// 	k.SetFetchPriceMsg(ctx, price)
// 	k.SetLastBlockHeight(ctx, ctx.BlockHeight())
// 	k.SetCheckFlag(ctx, false)
// 	return nil
// }

// func (k Keeper) OraclePriceValidationByRequestID(ctx sdk.Context, req int64) bool {
// 	currentReqID := k.GetLastFetchPriceID(ctx)

// 	return currentReqID != req
// }

// func (k Keeper) SetOracleValidationResult(ctx sdk.Context, res bool) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Set(types.KeyPrefix(types.OracleValidationResultKey),
// 		k.cdc.MustMarshalLengthPrefixed(&gogotypes.BoolValue{Value: res}))
// }

// func (k Keeper) GetOracleValidationResult(ctx sdk.Context) bool {
// 	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.OracleValidationResultKey))
// 	boolV := gogotypes.BoolValue{}
// 	k.cdc.MustUnmarshalLengthPrefixed(bz, &boolV)
// 	return boolV.GetValue()
// }

// func (k Keeper) SetTempFetchPriceID(ctx sdk.Context, id int64) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Set(types.KeyPrefix(types.TempFetchPriceIDKey),
// 		k.cdc.MustMarshalLengthPrefixed(&gogotypes.Int64Value{Value: id}))
// }

// func (k Keeper) GetTempFetchPriceID(ctx sdk.Context) int64 {
// 	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.TempFetchPriceIDKey))
// 	intV := gogotypes.Int64Value{}
// 	k.cdc.MustUnmarshalLengthPrefixed(bz, &intV)
// 	return intV.GetValue()
// }

// func (k Keeper) SetCheckFlag(ctx sdk.Context, flag bool) {
// 	var (
// 		store = ctx.KVStore(k.storeKey)
// 		key   = types.CheckFlagKey
// 		value = k.cdc.MustMarshal(
// 			&protobuftypes.BoolValue{
// 				Value: flag,
// 			},
// 		)
// 	)

// 	store.Set(key, value)
// }

// func (k Keeper) GetCheckFlag(ctx sdk.Context) bool {
// 	var (
// 		store = ctx.KVStore(k.storeKey)
// 		key   = types.CheckFlagKey
// 		value = store.Get(key)
// 	)

// 	var id protobuftypes.BoolValue
// 	k.cdc.MustUnmarshal(value, &id)

// 	return id.GetValue()
// }
