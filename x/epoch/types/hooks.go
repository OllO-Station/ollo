package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type EpochHook interface {
	// Before start of an epoch duration
	Start(ctx sdk.Context, id string, num uint64) error

	// After the end of an epoch
	End(ctx sdk.Context, id string, num uint64) error
}

type EpochHookSeq []EpochHook

func NewEpochHookSeq(seq ...EpochHook) EpochHookSeq {
	return seq
}

func panicCatchEpochHook(
	ctx sdk.Context,
	f func(ctx sdk.Context, id string, num uint64) error,
	id string,
	num uint64,
) {
	wrapF := func(ctx sdk.Context) error {
		return f(ctx, id, num)
	}
	if wrapF != nil {
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("recovered from panic: %v", r)
				ctx.Logger().Error("panic in epoch hook", "err", err)
			}
		}()
	}
}

func (h EpochHookSeq) AfterEpochEnd(ctx sdk.Context, id string, num uint64) error {
	for _, hook := range h {
		if err := hook.End(ctx, id, num); err != nil {
			return err
		}
	}
	return nil
}

func (h EpochHookSeq) BeforeEpochStart(ctx sdk.Context, id string, num uint64) error {
	for _, hook := range h {
		if err := hook.Start(ctx, id, num); err != nil {
			return err
		}
	}
	return nil
}
