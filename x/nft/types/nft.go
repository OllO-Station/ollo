package types

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ollo-station/ollo/x/nft/exported"
)

var _ exported.NFT = BaseNFT{}

// NewBaseNFT creates a new NFT instance
func NewBaseNFT(id, name string, owner sdk.AccAddress, uri, uriHash, data string, created time.Time, royaltyShare sdk.Dec, transferable bool) BaseNFT {
	return BaseNFT{
		Id:           id,
		Name:         name,
		Owner:        owner.String(),
		URI:          uri,
		UriHash:      uriHash,
		Data:         data,
		CreatedAt:    created,
		RoyaltyShare: royaltyShare,
		Transferable: true,
	}
}

// GetID return the id of BaseNFT
func (bnft BaseNFT) GetID() string {
	return bnft.Id
}

// GetName return the name of BaseNFT
func (bnft BaseNFT) GetName() string {
	return bnft.Name
}

// GetCreatedAt return the created time of BaseNFT
func (bnft BaseNFT) GetTimeCreated() time.Time {
	return bnft.CreatedAt
}

// GetRoyaltyShare return the royalty share of BaseNFT
func (bnft BaseNFT) GetRoyaltyShare() sdk.Dec {
	return bnft.RoyaltyShare
}

// GetTransferable return the transferable of BaseNFT
func (bnft BaseNFT) IsTransferable() bool {
	return bnft.Transferable
}

// GetOwner return the owner of BaseNFT
func (bnft BaseNFT) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(bnft.Owner)
	return owner
}

// GetURI return the URI of BaseNFT
func (bnft BaseNFT) GetURI() string {
	return bnft.URI
}

// GetURIHash return the UriHash of BaseNFT
func (bnft BaseNFT) GetURIHash() string {
	return bnft.UriHash
}

// GetData return the Data of BaseNFT
func (bnft BaseNFT) GetData() string {
	return bnft.Data
}

// ----------------------------------------------------------------------------
// NFT

// NFTs define a list of NFT
type NFTs []exported.NFT

func UnmarshalNFTMetadata(cdc codec.Codec, bz []byte) (NFTMetadata, error) {
	var nftMetadata NFTMetadata
	if len(bz) == 0 {
		return nftMetadata, nil
	}

	if err := cdc.Unmarshal(bz, &nftMetadata); err != nil {
		return nftMetadata, err
	}
	return nftMetadata, nil
}
