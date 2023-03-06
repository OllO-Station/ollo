package types

import (
	"bytes"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	ethaccounts "github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

type (
	EthAccount struct {
		*authtypes.BaseAccount `       protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
		CodeHash               string `protobuf:"bytes,2,opt,name=code_hash,json=codeHash,proto3"                             json:"code_hash,omitempty"    yaml:"code_hash"`
	}

	EthAccountI interface {
		authtypes.AccountI
		EthAddress() ethcommon.Address
		GetCodeHash() ethcommon.Hash
		SetCodeHash(codeHash ethcommon.Hash) error
		Type() int8
	}

	HdPathIterator func() ethaccounts.DerivationPath
)

const (
	//
	AccountTypeEOA = int8(iota + 1)
	//
	AccountTypeContract
)

var (
	//
	emptyCodeHash = ethcrypto.Keccak256(nil)
	//
	Bip44CoinType uint32 = 60

	Bip44HdPath = ethaccounts.DefaultBaseDerivationPath.String()
)

var (
	//
	_ authtypes.AccountI = (*EthAccount)(nil)
	//
	_ EthAccountI = (*EthAccount)(nil)
	//
	_ authtypes.GenesisAccount = (*EthAccount)(nil)
	//
	_ codectypes.UnpackInterfacesMessage = (*EthAccount)(nil)
)

func ProtoAccount() authtypes.AccountI {
	return &EthAccount{
		CodeHash:    ethcommon.BytesToHash(emptyCodeHash).String(),
		BaseAccount: &authtypes.BaseAccount{},
	}
}

func (a EthAccount) GetBaseAccount() *authtypes.BaseAccount {
	return a.BaseAccount
}

func (a EthAccount) EthAddress() ethcommon.Address {
	return ethcommon.BytesToAddress(a.GetAddress().Bytes())
}

func (a EthAccount) GetCodeHash() ethcommon.Hash {
	return ethcommon.HexToHash(a.CodeHash)
}

func (a *EthAccount) SetCodeHash(codeHash ethcommon.Hash) error {
	a.CodeHash = codeHash.Hex()
	return nil
}

func (a EthAccount) Type() int8 {
	if bytes.Equal(emptyCodeHash, ethcommon.HexToHash(a.CodeHash).Bytes()) {
		return AccountTypeEOA
	}
	return AccountTypeContract
}

func NewHdPathIterator(basePath string, ledgerIter bool) (HdPathIterator, error) {
	hdPath, e := ethaccounts.ParseDerivationPath(basePath)
	if e != nil {
		return nil, e
	}
	if ledgerIter {
		return ethaccounts.LedgerLiveIterator(hdPath), nil
	}
	return ethaccounts.DefaultIterator(hdPath), nil
}
