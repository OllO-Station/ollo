package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		Airdrops:     []Airdrop{},
		ClaimRecords: []ClaimRecord{},
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	for _, a := range p.Airdrops {
		if err := a.Validate(); err != nil {
			return err
		}
	}

	for _, r := range p.ClaimRecords {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (a Airdrop) Validate() error {
	if _, err := sdk.AccAddressFromBech32(a.SourceAddress); err != nil {
		return err
	}

	if a.StartTime.After(a.EndTime) {
		return errors.New("end time must be greater than start time")
	}

	for _, c := range a.Conditions {
		switch c {
		case ConditionTypeDeposit, ConditionTypeSwap,
			ConditionTypeStake, ConditionTypeVote:
		default:
			return fmt.Errorf("unknown condition type %T", c)
		}
	}
	return nil
}

// Validate validates claim record object.
func (r ClaimRecord) Validate() error {
	if _, err := sdk.AccAddressFromBech32(r.Recipient); err != nil {
		return err
	}

	if !r.InitialClaimableCoins.IsAllPositive() {
		return fmt.Errorf("initial claimable amount must be positive: %s", r.InitialClaimableCoins.String())
	}

	if err := r.ClaimableCoins.Validate(); err != nil {
		return fmt.Errorf("invalid claimable coins: %w", err)
	}
	return nil
}
