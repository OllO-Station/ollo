package keeper

import (
	"time"

	"github.com/bandprotocol/bandchain-packet/obi"
	"github.com/bandprotocol/bandchain-packet/packet"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	protobuftypes "github.com/gogo/protobuf/types"

	"ollo/x/oracle/types"
)

func (k Keeper) SetPriceResult(ctx sdk.Context, requestID types.OracleRequestID, result types.PricesResult) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PricesResultStoreKey(requestID), k.cdc.MustMarshal(&result))
}

// GetPriceResult returns the FetchPrice by requestId.
func (k Keeper) GetPriceResult(ctx sdk.Context, id types.OracleRequestID) (types.PricesResult, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.PricesResultStoreKey(id))
	if bz == nil {
		return types.PricesResult{}, sdkerrors.Wrapf(types.ErrSample,
			"GetResult: Result for request ID %d is not available.", id,
		)
	}
	var result types.PricesResult
	k.cdc.MustUnmarshal(bz, &result)
	return result, nil
}

// GetLastPriceID return the id from the last FetchPrice request.
func (k Keeper) GetLastPriceID(ctx sdk.Context) int64 {
	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.LastPricesIDKey))
	intV := protobuftypes.Int64Value{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intV)
	return intV.GetValue()
}

// SetLastPriceID saves the id from the last FetchPrice request.
func (k Keeper) SetLastPriceID(ctx sdk.Context, id types.OracleRequestID) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.LastPricesIDKey),
		k.cdc.MustMarshalLengthPrefixed(&protobuftypes.Int64Value{Value: int64(id)}))
}

func (k Keeper) Price(ctx sdk.Context, msg types.MsgPricesData) (*types.PricesCallData, error) {
	sourcePort := types.PortID
	sourceChannelEnd, found := k.ChannelKeeper.GetChannel(ctx, sourcePort, msg.SourceChannel)
	if !found {
		return nil, nil
	}
	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.ChannelKeeper.GetNextSequenceSend(ctx, sourcePort, msg.SourceChannel)
	if !found {
		return nil, nil
	}

	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, msg.SourceChannel))
	if !ok {
		return nil, nil
	}

	var symbol []string
	// assets := k.AssetKeeper.GetAssets(ctx)
	// for _, asset := range assets {
	// 	if asset.IsOraclePriceRequired {
	// 		symbol = append(symbol, asset.Name)
	// 	}
	// }

	encodedCallData := obi.MustEncode(types.PricesCallData{Symbols: symbol, Multiplier: 1000000})

	packetData := packet.NewOracleRequestPacketData(
		msg.ClientID,
		msg.OracleScriptID,
		encodedCallData,
		msg.AskCount,
		msg.MinCount,
		msg.FeeLimit,
		msg.PrepareGas,
		msg.ExecuteGas,
	)
	err := k.ChannelKeeper.SendPacket(ctx, channelCap, channeltypes.NewPacket(
		packetData.GetBytes(),
		sequence,
		sourcePort,
		msg.SourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.NewHeight(0, 0),
		uint64(ctx.BlockTime().UnixNano()+int64(10*time.Minute)), // Arbitrary timestamp timeout for now
	))
	if err != nil {
		return nil, nil
	}

	return &types.PricesCallData{}, nil
}

func (k Keeper) SetPriceMsg(ctx sdk.Context, msg types.MsgPricesData) {
	var (
		store = ctx.KVStore(k.storeKey)
		key   = types.MsgDataKey
		v     = types.NewMsgPricesData(
			types.ModuleName,
			types.OracleScriptID(msg.OracleScriptID),
			msg.SourceChannel,
			nil,
			msg.AskCount,
			msg.MinCount,
			msg.FeeLimit,
			msg.PrepareGas,
			msg.ExecuteGas,
		)
		value = k.cdc.MustMarshal(v)
	)

	store.Set(key, value)
}

func (k Keeper) GetPriceMsg(ctx sdk.Context) types.MsgPricesData {
	var (
		store = ctx.KVStore(k.storeKey)
		key   = types.MsgDataKey
		value = store.Get(key)
	)
	var msg types.MsgPricesData
	if value != nil {
		k.cdc.MustUnmarshal(value, &msg)
	}

	return msg
}

func (k Keeper) GetLastBlockHeight(ctx sdk.Context) int64 {
	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.LastBlockHeightKey))
	intV := protobuftypes.Int64Value{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intV)
	return intV.GetValue()
}

func (k Keeper) SetLastBlockHeight(ctx sdk.Context, height int64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.LastBlockHeightKey),
		k.cdc.MustMarshalLengthPrefixed(&protobuftypes.Int64Value{Value: height}))
}

func (k Keeper) AddPriceRecords(ctx sdk.Context, price types.MsgPricesData) error {
	k.SetPriceMsg(ctx, price)
	k.SetLastBlockHeight(ctx, ctx.BlockHeight())
	k.SetCheckFlag(ctx, false)
	return nil
}

func (k Keeper) OraclePriceValidationByRequestID(ctx sdk.Context, req int64) bool {
	currentReqID := k.GetLastPriceID(ctx)

	return currentReqID != req
}

func (k Keeper) SetOracleValidationResult(ctx sdk.Context, res bool) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.OracleValidationResultKey),
		k.cdc.MustMarshalLengthPrefixed(&protobuftypes.BoolValue{Value: res}))
}

func (k Keeper) GetOracleValidationResult(ctx sdk.Context) bool {
	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.OracleValidationResultKey))
	boolV := protobuftypes.BoolValue{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &boolV)
	return boolV.GetValue()
}

func (k Keeper) SetTempPriceID(ctx sdk.Context, id int64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.TempPricesIDKey),
		k.cdc.MustMarshalLengthPrefixed(&protobuftypes.Int64Value{Value: id}))
}

func (k Keeper) GetTempPriceID(ctx sdk.Context) int64 {
	bz := ctx.KVStore(k.storeKey).Get(types.KeyPrefix(types.TempPricesIDKey))
	intV := protobuftypes.Int64Value{}
	k.cdc.MustUnmarshalLengthPrefixed(bz, &intV)
	return intV.GetValue()
}

func (k Keeper) SetCheckFlag(ctx sdk.Context, flag bool) {
	var (
		store = ctx.KVStore(k.storeKey)
		key   = types.PortKey
		value = k.cdc.MustMarshal(
			&protobuftypes.BoolValue{
				Value: flag,
			},
		)
	)

	store.Set(key, value)
}

func (k Keeper) GetCheckFlag(ctx sdk.Context) bool {
	var (
		store = ctx.KVStore(k.storeKey)
		key   = types.CheckFlagKey
		value = store.Get(key)
	)

	var id protobuftypes.BoolValue
	k.cdc.MustUnmarshal(value, &id)

	return id.GetValue()
}
func (k Keeper) HandleProposalPrice(ctx sdk.Context, p *types.PricesProposal) error {
	return k.AddPriceRecords(ctx, p.FetchPrice)
}
