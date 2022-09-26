package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	channelkeeper "github.com/cosmos/ibc-go/v3/modules/core/04-channel/keeper"
	// ibcante "github.com/cosmos/ibc-go/v3/modules/core/ante"
)

type HandlerOptions struct {
	IBCChannelKeeper channelkeeper.Keeper
	ante.HandlerOptions
}

func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, errors.Wrap(errors.ErrLogic, "account keeper req for ante handler")
	}
	if options.BankKeeper == nil {
		return nil, errors.Wrap(errors.ErrLogic, "bank keeper req for ante handler")

	}
	if options.SignModeHandler == nil {
		return nil, errors.Wrap(errors.ErrLogic, "signer req for ante handler")
	}
	if options.SigGasConsumer == nil {
		return nil, errors.Wrap(errors.ErrLogic, "gas signer req for ante handler")
	}
	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(),
		ante.NewRejectExtensionOptionsDecorator(),
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		// ibcante.NewAnteDecorator(options.IBCChannelKeeper.ChanCloseConfirm),
	}
	return sdk.ChainAnteDecorators(anteDecorators...), nil

}
