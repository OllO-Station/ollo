package types

import "fmt"

func ValidateEpochId(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid type for epoch id")
	}
	if err := ValidateEpochIdString(v); err != nil {
		return err
	}
	return nil
}
func ValidateEpochIdString(s string) error {
	if s == "" {
		return fmt.Errorf("epoch id cannot be empty")
	}
	return nil
}
func (gs GenesisState) Validate() error {
	epochIds := map[string]bool{}
	for _, epoch := range gs.Epochs {
		if err := ValidateEpochId(epoch.Id); err != nil {
			return err
		}
		if epochIds[epoch.Id] {
			return fmt.Errorf("duplicate epoch id: %s", epoch.Id)
		}
		epochIds[epoch.Id] = true
	}
	return nil
}

func (e Epoch) Validate() error {
	if e.Id == "" {
		return fmt.Errorf("epoch id cannot be empty")
	}
	// if e.Start.IsZero() {
	// 	return fmt.Errorf("epoch start cannot be zero")
	// }
	if e.Duration == 0 {
		return fmt.Errorf("epoch end cannot be zero")
	}
	if e.CurrentEpochNumber < 0 {
		return fmt.Errorf("epoch number cannot be negative")
	}
	if e.CurrentEpochStartHeight < 0 {
		return fmt.Errorf("epoch start height cannot be negative")
	}
	return nil
}
