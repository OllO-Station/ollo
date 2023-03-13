package exported

import (
	"time"
)

type EpochI interface {
	GetId() uint64
	GetNum() uint64
	GetStartTime() time.Time
	GetEndTime() time.Time
}
