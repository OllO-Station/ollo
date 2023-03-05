package types

const (
	//
	ModuleName = "epoch"
	//
	QuerierRoute = ModuleName
	//
	StoreKey = ModuleName
	//
	RouterKey = ModuleName
)

var (
	//
	EpochPrefix = []byte{0x01}
)

func GetEpochPrefix(pre string) []byte {
	return []byte(pre)
}
