package exported

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NFT non fungible token interface
type NFT interface {
	GetID() string
	GetName() string
	GetOwner() sdk.AccAddress
	GetURI() string
	GetURIHash() string
	GetData() string
	IsTransferable() bool
	GetTimeCreated() time.Time
	GetRoyaltyShare() sdk.Dec
}
