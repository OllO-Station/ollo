package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyContractAddress                     = []byte("contract")
	_                  paramtypes.ParamSet = &Params{}
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})

}

func NewParams(contractAddress sdk.AccAddress) (Params, error) {
	return Params{
		ContractAddress: contractAddress,
	}, nil
}

func (p Params) Validate() error {
	if err := validateContractAddress(p.ContractAddress); err != nil {
		return err
	}
	return nil
}

func DefaultParams() Params {
	return Params{
		ContractAddress: sdk.AccAddress{},
	}
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyContractAddress, &p.ContractAddress, validateContractAddress),
	}
}

func validateContractAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v == "" {
		return nil
	}
	bech32, err := sdk.AccAddressFromBech32(v)
	if err != nil {
		return err
	}

	err = sdk.VerifyAddressFormat(bech32)
	if err != nil {
		return err
	}

	return nil
}
