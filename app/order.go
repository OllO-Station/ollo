package app

import (
	abci "github.com/tendermint/tendermint/abci/types"
	// solomachine "github.com/cosmos/ibc-go/v6/modules/light-clients/06-solomachine"
	// minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	// // tmint "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint"
	//
	// ibcmock "github.com/cosmos/ibc-go/v6/testing/mock"
	//
	// // "github.com/cosmos/ibc-go/v6/modules/core/04-channel"
	//
	// authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	// vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	// "github.com/cosmos/cosmos-sdk/x/authz"
	// banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	// capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	// distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	// evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	// "github.com/cosmos/cosmos-sdk/x/feegrant"
	// genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	// govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	// "github.com/cosmos/cosmos-sdk/x/group"
	//
	// // "github.com/cosmos/cosmos-sdk/x/mint"
	// // mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	// // minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	// "github.com/cosmos/cosmos-sdk/x/nft"
	//
	// // nftmodule "github.com/cosmos/cosmos-sdk/x/nft/client/cli"
	//
	// paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	// slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	// stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	// upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	// icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
	// ibcfeetypes "github.com/cosmos/ibc-go/v6/modules/apps/29-fee/types"
	// ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	// ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	//
	// // v6 "github.com/cosmos/ibc-go/v6/testing/simapp/upgrades/v6"
	//
	// claimmoduletypes "github.com/ollo-station/ollo/x/claim/types"
	// liquiditymoduletypes "github.com/ollo-station/ollo/x/liquidity/types"
	// loanmoduletypes "github.com/ollo-station/ollo/x/loan/types"
	// marketmoduletypes "github.com/ollo-station/ollo/x/market/types"
	// onsmoduletypes "github.com/ollo-station/ollo/x/ons/types"
	// reservemoduletypes "github.com/ollo-station/ollo/x/reserve/types"
	// mintmoduletypes "github.com/ollo-station/ollo/x/mint/types"
)

func (app *App) DeliverTx(req abci.RequestDeliverTx) (res abci.ResponseDeliverTx) {
	defer func() {
		if res.IsErr() {
			// app.tpsCounter.incrementFailed()
		} else {
			// app.tpsCounter.incrementSuccess()
		}
	}()
	return app.BaseApp.DeliverTx(req)
}
