package types

import (
	"fmt"
	"time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	_ paramtypes.ParamSet = (*Params)(nil)

	KeyDecayInformation = []byte("DecayInformation")
	KeyAirdropStart     = []byte("AirdropStart")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(di DecayInformation, airdropStart time.Time) Params {
	return Params{
		DecayInformation: di,
		AirdropStart:     airdropStart,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		NewDisabledDecay(),
		time.Unix(0, 0),
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDecayInformation, &p.DecayInformation, validateDecayInformation),
		paramtypes.NewParamSetPair(KeyAirdropStart, &p.AirdropStart, validateAirdropStart),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateDecayInformation(p.DecayInformation); err != nil {
		return err
	}
	return validateAirdropStart(p.AirdropStart)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateDecayInformation validates the DecayInformation param
func validateDecayInformation(v interface{}) error {
	decayInfo, ok := v.(DecayInformation)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	return decayInfo.Validate()
}

func validateAirdropStart(i interface{}) error {
	if _, ok := i.(time.Time); !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
