package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInsufficientDepositAmount = sdkerrors.New(ModuleName, 2, "insufficient deposit amount")
	ErrPairAlreadyExists         = sdkerrors.New(ModuleName, 3, "pair already exists")
	ErrPoolAlreadyExists         = sdkerrors.New(ModuleName, 4, "pool already exists")
	ErrWrongPoolCoinDenom        = sdkerrors.New(ModuleName, 5, "wrong pool coin denom")
	ErrInvalidCoinDenom          = sdkerrors.New(ModuleName, 6, "invalid coin denom")
	ErrNoLastPrice               = sdkerrors.New(ModuleName, 8, "cannot make a market order to a pair with no last price")
	ErrInsufficientOfferCoin     = sdkerrors.New(ModuleName, 9, "insufficient offer coin")
	ErrPriceOutOfRange           = sdkerrors.New(ModuleName, 10, "price out of range limit")
	ErrTooLongOrderLifespan      = sdkerrors.New(ModuleName, 11, "order lifespan is too long")
	ErrInactivePool              = sdkerrors.New(ModuleName, 12, "inactive pool")
	ErrWrongPair                 = sdkerrors.New(ModuleName, 13, "wrong denom pair")
	ErrSameBatch                 = sdkerrors.New(ModuleName, 14, "cannot cancel an order within the same batch")
	ErrAlreadyCanceled           = sdkerrors.New(ModuleName, 15, "the order is already canceled")
	ErrDuplicatePairId           = sdkerrors.New(ModuleName, 16, "duplicate pair id presents in the pair id list")
	ErrTooSmallOrder             = sdkerrors.New(ModuleName, 17, "too small order")
	ErrTooLargePool              = sdkerrors.New(ModuleName, 18, "too large pool")
	ErrTooManyPools              = sdkerrors.New(ModuleName, 19, "too many pools in the pair")
	ErrPriceNotOnTicks           = sdkerrors.New(ModuleName, 20, "price is not on ticks")
)
