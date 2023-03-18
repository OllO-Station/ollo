package types

import (
	fmt "fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeMarketMaker = "MarketMaker"
)

var _ gov.Content = &MarketMakerProposal{}

func init() {
	gov.RegisterProposalType(ProposalTypeMarketMaker)
}

func NewMarketMakerProposal(
	title string,
	description string,
	inclusions []MarketMakerHandle,
	exclusions []MarketMakerHandle,
	rejections []MarketMakerHandle,
	distributions []IncentiveDistribution,
) *MarketMakerProposal {
	return &MarketMakerProposal{
		Title:         title,
		Description:   description,
		Inclusions:    inclusions,
		Exclusions:    exclusions,
		Rejections:    rejections,
		Distributions: distributions,
	}
}

func (p *MarketMakerProposal) GetTitle() string {
	return p.Title
}

func (p *MarketMakerProposal) GetDescription() string {
	return p.Description
}

func (p *MarketMakerProposal) ProposalRoute() string {
	return RouterKey
}

func (p *MarketMakerProposal) ProposalType() string {
	return ProposalTypeMarketMaker
}

func (p *MarketMakerProposal) ValidateInclusions() error {
	return nil
}

func (p *MarketMakerProposal) ValidateBasic() error {
	if len(p.Inclusions) == 0 &&
		len(p.Exclusions) == 0 &&
		len(p.Rejections) == 0 &&
		len(p.Distributions) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "proposal must not be empty")
	}
	addrMap := make(map[MarketMakerHandle]struct{})
	for _, mm := range p.Inclusions {
		if _, ok := addrMap[mm]; ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "duplicate market maker inclusion")
		}
		addrMap[mm] = struct{}{}
		if err := mm.Validate(); err != nil {
			return err
		}
	}
	for _, mm := range p.Exclusions {
		if _, ok := addrMap[mm]; ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "market maker can't be duplicated")
		}
		addrMap[mm] = struct{}{}
		if err := mm.Validate(); err != nil {
			return err
		}
	}

	for _, mm := range p.Rejections {
		if _, ok := addrMap[mm]; ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "market maker can't be duplicated")
		}
		addrMap[mm] = struct{}{}
		if err := mm.Validate(); err != nil {
			return err
		}
	}

	for _, dp := range p.Distributions {
		if err := dp.Validate(); err != nil {
			return err
		}
	}
	return gov.ValidateAbstract(p)

}

func (p MarketMakerProposal) String() string {
	return fmt.Sprintf(`Market Maker Proposal:
  Title:         %s
  Description:   %s
  Inclusions:    %v
  Exclusions:    %v
  Rejections:    %v
  Distributions: %v
`, p.Title, p.Description, p.Inclusions, p.Exclusions, p.Rejections, p.Distributions)
}
