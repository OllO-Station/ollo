package keeper
import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/ons/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetName(c sdk.Context, req *types.QueryGetNameRequest) (*types.QueryGetNameResponse, error) {
  ctx := sdk.UnwrapSDKContext(c)

  name := types.Name{Id: 0, NameType: types.NAME_TYPE_NAME, Value: "", OwnerAddr: "", Status: types.NAME_STATUS_ACTIVE, Name: req.Name,ColorHint: "",Prefix: &types.NameEnd{Letter: []*types.NameEndCharacterFormat{}},Suffix: &types.NameEnd{Letter: *&[]*types.NameEndCharacterFormat{}},Tag: []string{},BoughtAt: ctx.BlockTime().String()}
  return &types.QueryGetNameResponse{Name: name}, nil
}

// // Airdrop queries the specific airdrop.
// func (k Querier) Airdrop(c context.Context, req *types.QueryAirdropRequest) (*types.QueryAirdropResponse, error) {
// 	if req == nil {
// 		return nil, status.Error(codes.InvalidArgument, "empty request")
// 	}

// 	ctx := sdk.UnwrapSDKContext(c)

// 	airdrop, found := k.Keeper.GetAirdrop(ctx, req.AirdropId)
// 	if !found {
// 		return nil, status.Error(codes.NotFound, "airdrop not found")
// 	}

// 	return &types.QueryAirdropResponse{Airdrop: airdrop}, nil
// }

// // ClaimRecord queries the specific claim record.
// func (k Querier) ClaimRecord(c context.Context, req *types.QueryClaimRecordRequest) (*types.QueryClaimRecordResponse, error) {
// 	if req == nil {
// 		return nil, status.Error(codes.InvalidArgument, "invalid request")
// 	}

// 	ctx := sdk.UnwrapSDKContext(c)

// 	recipientAddr, err := sdk.AccAddressFromBech32(req.Recipient)
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, "invalid address")
// 	}

// 	record, found := k.GetClaimRecordByRecipient(ctx, req.AirdropId, recipientAddr)
// 	if !found {
// 		return nil, status.Error(codes.NotFound, "claim record not found")
// 	}

// 	return &types.QueryClaimRecordResponse{ClaimRecord: record}, nil
//   
//   return 
// }

// // SetParams set the params
// func (k Keeper) SetName(ctx sdk.Context, params types.Params) {
// 	k.paramstore.SetParamSet(ctx, &params)
// }
