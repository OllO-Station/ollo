package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"ollo/x/liquidity/types"
)

// ExecuteRequests executes all orders, deposit requests and withdraw requests.
// ExecuteRequests also handles order expiration.
func (k Keeper) ExecuteRequests(ctx sdk.Context) {
	if err := k.IterateAllPairs(ctx, func(pair types.Pair) (stop bool, err error) {
		if err := k.ExecuteMatching(ctx, pair); err != nil {
			return false, err
		}
		return false, nil
	}); err != nil {
		panic(err)
	}
	if err := k.IterateAllOrders(ctx, func(order types.Order) (stop bool, err error) {
		if order.Status.CanBeExpired() && order.ExpiredAt(ctx.BlockTime()) {
			if err := k.FinishOrder(ctx, order, types.OrderStatusExpired); err != nil {
				return false, err
			}
		} else if types.IsTooSmallOrderAmount(order.OpenAmt, order.Price) {
			// TODO: should we introduce new order status for this type of expiration?
			if err := k.FinishOrder(ctx, order, types.OrderStatusExpired); err != nil {
				return false, err
			}
		}
		return false, nil
	}); err != nil {
		panic(err)
	}
	if err := k.IterateAllRequestDeposits(ctx, func(req types.RequestDeposit) (stop bool, err error) {
		if req.Status == types.RequestStatusPending {
			if err := k.ExecuteRequestDeposit(ctx, req); err != nil {
				return false, err
			}
		}
		return false, nil
	}); err != nil {
		panic(err)
	}
	if err := k.IterateAllRequestWithdraws(ctx, func(req types.RequestWithdraw) (stop bool, err error) {
		if req.Status == types.RequestStatusPending {
			if err := k.ExecuteRequestWithdraw(ctx, req); err != nil {
				return false, err
			}
		}
		return false, nil
	}); err != nil {
		panic(err)
	}
}

// DeleteOutdatedRequests deletes outdated(should be deleted) requests.
// Determining if a request should be deleted is based on its status.
func (k Keeper) DeleteOutdatedRequests(ctx sdk.Context) {
	_ = k.IterateAllRequestDeposits(ctx, func(req types.RequestDeposit) (stop bool, err error) {
		if req.Status.ShouldBeDeleted() {
			k.DeleteRequestDeposit(ctx, req)
		}
		return false, nil
	})
	_ = k.IterateAllRequestWithdraws(ctx, func(req types.RequestWithdraw) (stop bool, err error) {
		if req.Status.ShouldBeDeleted() {
			k.DeleteRequestWithdraw(ctx, req)
		}
		return false, nil
	})
	_ = k.IterateAllOrders(ctx, func(order types.Order) (stop bool, err error) {
		if order.Status.ShouldBeDeleted() {
			k.DeleteOrder(ctx, order)
		}
		return false, nil
	})
}
