package types

import "time"

const (
	DefaultId uint64 = 1
)

func NewGenesisState(epochs []Epoch) *GenesisState {
	return &GenesisState{
		Epochs: epochs,
	}
}

func NewGenesisEpoch(id string, dur time.Duration) Epoch {
	return Epoch{
		Id:                      id,
		CurrentEpochStart:       time.Time{},
		Start:                   time.Time{},
		EpochStarted:            false,
		Duration:                dur,
		CurrentEpochNumber:      0,
		CurrentEpochStartHeight: 0,
	}
}

func DefaultGenesis() *GenesisState {
	epochs := []Epoch{
		NewGenesisEpoch("hour", time.Hour*1),
		NewGenesisEpoch("day", time.Hour*24),
		NewGenesisEpoch("week", time.Hour*24*7),
		NewGenesisEpoch("month", time.Hour*24*7*30),
	}
	return NewGenesisState(epochs)
}
