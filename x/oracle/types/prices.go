package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPricesData = "prices_data"

var (
	_ sdk.Msg = &MsgPricesData{}

	// PricesResultStoreKeyPrefix is a prefix for storing result
	PricesResultStoreKeyPrefix = "prices_result"

	// LastPricesIDKey is the key for the last request id
	LastPricesIDKey = "prices_last_id"

	// PricesClientIDKey is query request identifier
	PricesClientIDKey = "prices_id"
)

// NewMsgPricesData creates a new Prices message
func NewMsgPricesData(
	creator string,
	oracleScriptID OracleScriptID,
	sourceChannel string,
	calldata *PricesCallData,
	askCount uint64,
	minCount uint64,
	feeLimit sdk.Coins,
	prepareGas uint64,
	executeGas uint64,
) *MsgPricesData {
	return &MsgPricesData{
		ClientID:       PricesClientIDKey,
		Creator:        creator,
		OracleScriptID: uint64(oracleScriptID),
		SourceChannel:  sourceChannel,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		FeeLimit:       feeLimit,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
	}
}

// Route returns the message route
func (m *MsgPricesData) Route() string {
	return RouterKey
}

// Type returns the message type
func (m *MsgPricesData) Type() string {
	return TypeMsgPricesData
}

// GetSigners returns the message signers
func (m *MsgPricesData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes returns the signed bytes from the message
func (m *MsgPricesData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic check the basic message validation
func (m *MsgPricesData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if m.SourceChannel == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid source channel")
	}
	return nil
}

// PricesResultStoreKey is a function to generate key for each result in store
func PricesResultStoreKey(requestID OracleRequestID) []byte {
	return append(KeyPrefix(PricesResultStoreKeyPrefix), int64ToBytes(int64(requestID))...)
}
