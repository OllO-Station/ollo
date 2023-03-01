package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewPostHandler() sdk.AnteHandler {
	return sdk.ChainAnteDecorators()
}
