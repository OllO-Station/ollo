package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBandPricesData = "band_prices_data"

var (
	_ sdk.Msg = &MsgBandPricesData{}

	// BandPricesResultStoreKeyPrefix is a prefix for storing result
	BandPricesResultStoreKeyPrefix = "band_prices_result"

	// LastBandPricesIDKey is the key for the last request id
	LastBandPricesIDKey = "band_prices_last_id"

	// BandPricesClientIDKey is query request identifier
	BandPricesClientIDKey = "band_prices_id"
)

// NewMsgBandPricesData creates a new BandPrices message
func NewMsgBandPricesData(
	creator string,
	oracleScriptID OracleScriptID,
	sourceChannel string,
	calldata *BandPricesCallData,
	askCount uint64,
	minCount uint64,
	feeLimit sdk.Coins,
	prepareGas uint64,
	executeGas uint64,
) *MsgBandPricesData {
	return &MsgBandPricesData{
		ClientID:       				BandPricesClientIDKey,
		Creator:	creator,
		OracleScriptID: 				uint64(oracleScriptID),
		SourceChannel:  				sourceChannel,
		Calldata:       				calldata,
		AskCount:       				askCount,
		MinCount:       				minCount,
		FeeLimit:       				feeLimit,
		PrepareGas:     				prepareGas,
		ExecuteGas:     				executeGas,
	}
}

// Route returns the message route
func (m *MsgBandPricesData) Route() string {
	return RouterKey
}

// Type returns the message type
func (m *MsgBandPricesData) Type() string {
	return TypeMsgBandPricesData
}

// GetSigners returns the message signers
func (m *MsgBandPricesData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes returns the signed bytes from the message
func (m *MsgBandPricesData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic check the basic message validation
func (m *MsgBandPricesData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if m.SourceChannel == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid source channel")
	}
	return nil
}

// BandPricesResultStoreKey is a function to generate key for each result in store
func BandPricesResultStoreKey(requestID OracleRequestID) []byte {
	return append(KeyPrefix(BandPricesResultStoreKeyPrefix), int64ToBytes(int64(requestID))...)
}
