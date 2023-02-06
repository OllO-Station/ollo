package types

import (
	"encoding/binary"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetGoalIDBytes returns the byte representation of the ID
func GetGoalIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// Validate checks the mission is valid
func (m Goal) Validate() error {
	if m.Weight.LT(sdk.ZeroDec()) || m.Weight.GT(sdk.OneDec()) {
		return errors.New("mission weight must be in range [0:1]")
	}

	return nil
}
