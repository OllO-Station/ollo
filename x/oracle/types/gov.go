package types

import (
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalPrices = "FetchPrices"
)

func init() {
	govtypes.RegisterProposalType(ProposalPrices)
	// govtypes.RegisterProposalTypeCodec(&PricesProposal{}, "comdex/FetchPricesProposal")
}

var _ govtypes.Content = &PricesProposal{}

func NewPricesProposal(title, description string, fetchPrice MsgPricesData) govtypes.Content {
	return &PricesProposal{
		Title:       title,
		Description: description,
		FetchPrice:  fetchPrice,
	}
}

func (p *PricesProposal) ProposalRoute() string { return RouterKey }

func (p *PricesProposal) ProposalType() string { return ProposalPrices }

func (p *PricesProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	// if p.FetchPrice.TwaBatchSize == 0 {
		// return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid batch size")
	// }
	if err != nil {
		return err
	}
	return nil
}
