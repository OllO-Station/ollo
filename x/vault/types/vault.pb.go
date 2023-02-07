// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/vault/v1/vault.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Vault struct {
	Id                    uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner                 string                                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	AmountIn              github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=amount_in,json=amountIn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_in" yaml:"amount_in"`
	AmountOut             github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=amount_out,json=amountOut,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_out" yaml:"amount_out"`
	CreatedAt             time.Time                              `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at" yaml:"created_at"`
	InterestAccumulated   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=interest_accumulated,json=interestAccumulated,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"interest_accumulated" yaml:"interest_accumulated"`
	ClosingFeeAccumulated github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=closing_fee_accumulated,json=closingFeeAccumulated,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"closing_fee_accumulated" yaml:"interest_accumulated"`
	BlockHeight           int64                                  `protobuf:"varint,10,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty" yaml:"block_height"`
	BlockTime             time.Time                              `protobuf:"bytes,11,opt,name=block_time,json=blockTime,proto3,stdtime" json:"block_time" yaml:"block_time"`
}

func (m *Vault) Reset()         { *m = Vault{} }
func (m *Vault) String() string { return proto.CompactTextString(m) }
func (*Vault) ProtoMessage()    {}
func (*Vault) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd351dc1da3b282, []int{0}
}
func (m *Vault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vault.Merge(m, src)
}
func (m *Vault) XXX_Size() int {
	return m.Size()
}
func (m *Vault) XXX_DiscardUnknown() {
	xxx_messageInfo_Vault.DiscardUnknown(m)
}

var xxx_messageInfo_Vault proto.InternalMessageInfo

type EmissionsVault struct {
	Id        uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	AmountIn  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount_in,json=amountIn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_in" yaml:"amount_in"`
	AmountOut github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount_out,json=amountOut,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_out" yaml:"amount_out"`
	CreatedAt time.Time                              `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at" yaml:"created_at"`
}

func (m *EmissionsVault) Reset()         { *m = EmissionsVault{} }
func (m *EmissionsVault) String() string { return proto.CompactTextString(m) }
func (*EmissionsVault) ProtoMessage()    {}
func (*EmissionsVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd351dc1da3b282, []int{1}
}
func (m *EmissionsVault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmissionsVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmissionsVault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmissionsVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmissionsVault.Merge(m, src)
}
func (m *EmissionsVault) XXX_Size() int {
	return m.Size()
}
func (m *EmissionsVault) XXX_DiscardUnknown() {
	xxx_messageInfo_EmissionsVault.DiscardUnknown(m)
}

var xxx_messageInfo_EmissionsVault proto.InternalMessageInfo

type EmissionsVaultRewards struct {
	User        string                                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" yaml:"user"`
	BlockHeight uint64                                 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty" yaml:"block_height"`
	Amount      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount" yaml:"amount"`
}

func (m *EmissionsVaultRewards) Reset()         { *m = EmissionsVaultRewards{} }
func (m *EmissionsVaultRewards) String() string { return proto.CompactTextString(m) }
func (*EmissionsVaultRewards) ProtoMessage()    {}
func (*EmissionsVaultRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd351dc1da3b282, []int{2}
}
func (m *EmissionsVaultRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmissionsVaultRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmissionsVaultRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmissionsVaultRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmissionsVaultRewards.Merge(m, src)
}
func (m *EmissionsVaultRewards) XXX_Size() int {
	return m.Size()
}
func (m *EmissionsVaultRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_EmissionsVaultRewards.DiscardUnknown(m)
}

var xxx_messageInfo_EmissionsVaultRewards proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Vault)(nil), "ollo.vault.v1.Vault")
	proto.RegisterType((*EmissionsVault)(nil), "ollo.vault.v1.EmissionsVault")
	proto.RegisterType((*EmissionsVaultRewards)(nil), "ollo.vault.v1.EmissionsVaultRewards")
}

func init() { proto.RegisterFile("ollo/vault/v1/vault.proto", fileDescriptor_2cd351dc1da3b282) }

var fileDescriptor_2cd351dc1da3b282 = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xb1, 0x6e, 0xd4, 0x40,
	0x10, 0xf5, 0x3a, 0xc9, 0x11, 0xef, 0x25, 0x11, 0x38, 0x89, 0xe2, 0x04, 0xc5, 0x3e, 0x6d, 0x01,
	0xd7, 0xc4, 0x56, 0x80, 0x2a, 0x48, 0xa0, 0x18, 0x81, 0x48, 0x81, 0x90, 0x0c, 0x02, 0x44, 0x63,
	0xf9, 0xec, 0x8d, 0x63, 0xc5, 0xf6, 0x46, 0xde, 0x75, 0x42, 0x3a, 0x1a, 0xfa, 0x74, 0xfc, 0x02,
	0x1f, 0xc0, 0x47, 0x5c, 0x85, 0x22, 0x2a, 0x44, 0x61, 0xe0, 0xae, 0xa2, 0xbd, 0x2f, 0x40, 0xde,
	0xdd, 0xe3, 0xee, 0x04, 0x48, 0x9c, 0x44, 0xa8, 0xec, 0x79, 0xb3, 0xf3, 0x66, 0xe6, 0xcd, 0xec,
	0xc2, 0x75, 0x92, 0xa6, 0xc4, 0x39, 0x0e, 0xca, 0x94, 0x39, 0xc7, 0xdb, 0xe2, 0xc7, 0x3e, 0x2a,
	0x08, 0x23, 0xfa, 0x62, 0xed, 0xb2, 0x05, 0x72, 0xbc, 0xbd, 0xb1, 0x12, 0x93, 0x98, 0x70, 0x8f,
	0x53, 0xff, 0x89, 0x43, 0x1b, 0x56, 0x4c, 0x48, 0x9c, 0x62, 0x87, 0x5b, 0x9d, 0x72, 0xdf, 0x61,
	0x49, 0x86, 0x29, 0x0b, 0xb2, 0x23, 0x79, 0x60, 0x3d, 0x24, 0x34, 0x23, 0xd4, 0x17, 0x91, 0xc2,
	0x10, 0x2e, 0xf4, 0xb6, 0x01, 0xe7, 0x9e, 0xd5, 0xf4, 0xfa, 0x12, 0x54, 0x93, 0xc8, 0x00, 0x2d,
	0xd0, 0x9e, 0xf5, 0xd4, 0x24, 0xd2, 0xef, 0xc0, 0x39, 0x72, 0x92, 0xe3, 0xc2, 0x50, 0x5b, 0xa0,
	0xad, 0xb9, 0xed, 0x41, 0x65, 0x2d, 0x9c, 0x06, 0x59, 0xba, 0x83, 0x38, 0x8c, 0x3e, 0xbe, 0xdf,
	0x5a, 0x91, 0x54, 0xbb, 0x51, 0x54, 0x60, 0x4a, 0x9f, 0xb0, 0x22, 0xc9, 0x63, 0x4f, 0x84, 0xe9,
	0x3e, 0xd4, 0x82, 0x8c, 0x94, 0x39, 0xf3, 0x93, 0xdc, 0x98, 0xe3, 0x1c, 0x6e, 0xb7, 0xb2, 0x94,
	0xcf, 0x95, 0x75, 0x2d, 0x4e, 0xd8, 0x41, 0xd9, 0xb1, 0x43, 0x92, 0xc9, 0x6a, 0xe4, 0x67, 0x8b,
	0x46, 0x87, 0x0e, 0x3b, 0x3d, 0xc2, 0xd4, 0xde, 0xcb, 0xd9, 0xa0, 0xb2, 0x2e, 0x8b, 0x8c, 0x3f,
	0x89, 0x90, 0x37, 0x2f, 0xfe, 0xf7, 0x72, 0xbd, 0x03, 0xa1, 0xc4, 0x49, 0xc9, 0x8c, 0x06, 0xcf,
	0x70, 0x6f, 0xea, 0x0c, 0x57, 0x26, 0x32, 0x90, 0x92, 0x21, 0x4f, 0xd6, 0xfd, 0xb8, 0x64, 0xfa,
	0x0b, 0x08, 0xc3, 0x02, 0x07, 0x0c, 0x47, 0x7e, 0xc0, 0x8c, 0x4b, 0x2d, 0xd0, 0x6e, 0xde, 0xd8,
	0xb0, 0x85, 0xde, 0xf6, 0x50, 0x6f, 0xfb, 0xe9, 0x50, 0x6f, 0x77, 0xb3, 0xce, 0x3f, 0x62, 0x1d,
	0xc5, 0xa2, 0xb3, 0x2f, 0x16, 0xf0, 0x34, 0x09, 0xec, 0x32, 0xfd, 0x35, 0x80, 0x2b, 0x49, 0xce,
	0x70, 0x81, 0x29, 0xf3, 0x83, 0x30, 0x2c, 0xb3, 0x32, 0xad, 0x5d, 0xc6, 0x3c, 0x6f, 0xe4, 0xd1,
	0xd4, 0x8d, 0x5c, 0x15, 0x29, 0x7f, 0xc7, 0x89, 0xbc, 0xe5, 0x21, 0xbc, 0x3b, 0x42, 0xf5, 0x37,
	0x00, 0xae, 0x85, 0x29, 0xa1, 0x49, 0x1e, 0xfb, 0xfb, 0x18, 0x4f, 0x54, 0xa1, 0x5d, 0x44, 0x15,
	0xab, 0x32, 0xdb, 0x03, 0x8c, 0xc7, 0xeb, 0xd8, 0x81, 0x0b, 0x9d, 0x94, 0x84, 0x87, 0xfe, 0x01,
	0x4e, 0xe2, 0x03, 0x66, 0xc0, 0x16, 0x68, 0xcf, 0xb8, 0x6b, 0x83, 0xca, 0x5a, 0x16, 0x6c, 0xe3,
	0x5e, 0xe4, 0x35, 0xb9, 0xf9, 0x90, 0x5b, 0xf5, 0x80, 0x84, 0xb7, 0xde, 0x79, 0xa3, 0x39, 0xed,
	0x80, 0x46, 0xb1, 0x72, 0x40, 0x1c, 0xa8, 0x8f, 0xa3, 0x0f, 0x2a, 0x5c, 0xba, 0x9f, 0x25, 0x94,
	0x26, 0x24, 0xa7, 0xe2, 0x8a, 0x6c, 0x8e, 0xae, 0x88, 0xbb, 0x38, 0xa8, 0x2c, 0x4d, 0x36, 0x1b,
	0x21, 0x7e, 0x63, 0x26, 0x36, 0x5e, 0xbd, 0xf0, 0x8d, 0x9f, 0xf9, 0x0f, 0x1b, 0x3f, 0xfb, 0xef,
	0x36, 0x1e, 0x7d, 0x07, 0x70, 0x75, 0x52, 0x50, 0x0f, 0x9f, 0x04, 0x45, 0x44, 0xf5, 0xdb, 0x70,
	0xb6, 0xa4, 0xb8, 0xe0, 0xca, 0x6a, 0xee, 0xf5, 0x41, 0x65, 0x35, 0x05, 0x5b, 0x8d, 0xfe, 0xf9,
	0xa1, 0xe1, 0x41, 0xbf, 0x6c, 0x8f, 0xca, 0xc7, 0xf3, 0x77, 0xdb, 0xf3, 0x1c, 0x36, 0x44, 0xe7,
	0x52, 0xcc, 0xbb, 0x53, 0x8b, 0xb9, 0x38, 0x2e, 0x26, 0xf2, 0x24, 0x9d, 0x7b, 0xab, 0xfb, 0xcd,
	0x54, 0xde, 0xf5, 0x4c, 0xa5, 0xdb, 0x33, 0xc1, 0x79, 0xcf, 0x04, 0x5f, 0x7b, 0x26, 0x38, 0xeb,
	0x9b, 0xca, 0x79, 0xdf, 0x54, 0x3e, 0xf5, 0x4d, 0xe5, 0xa5, 0xce, 0x1f, 0xfd, 0x57, 0xf2, 0xd9,
	0xe7, 0x94, 0x9d, 0x06, 0xd7, 0xf7, 0xe6, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x90, 0x9a,
	0xae, 0x11, 0x06, 0x00, 0x00,
}

func (m *Vault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.BlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.BlockTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintVault(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x5a
	if m.BlockHeight != 0 {
		i = encodeVarintVault(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.ClosingFeeAccumulated.Size()
		i -= size
		if _, err := m.ClosingFeeAccumulated.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.InterestAccumulated.Size()
		i -= size
		if _, err := m.InterestAccumulated.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintVault(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x3a
	{
		size := m.AmountOut.Size()
		i -= size
		if _, err := m.AmountOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.AmountIn.Size()
		i -= size
		if _, err := m.AmountIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintVault(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintVault(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EmissionsVault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmissionsVault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmissionsVault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintVault(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	{
		size := m.AmountOut.Size()
		i -= size
		if _, err := m.AmountOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.AmountIn.Size()
		i -= size
		if _, err := m.AmountIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintVault(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EmissionsVaultRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmissionsVaultRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmissionsVaultRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVault(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.BlockHeight != 0 {
		i = encodeVarintVault(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintVault(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVault(dAtA []byte, offset int, v uint64) int {
	offset -= sovVault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Vault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovVault(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovVault(uint64(l))
	}
	l = m.AmountIn.Size()
	n += 1 + l + sovVault(uint64(l))
	l = m.AmountOut.Size()
	n += 1 + l + sovVault(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovVault(uint64(l))
	l = m.InterestAccumulated.Size()
	n += 1 + l + sovVault(uint64(l))
	l = m.ClosingFeeAccumulated.Size()
	n += 1 + l + sovVault(uint64(l))
	if m.BlockHeight != 0 {
		n += 1 + sovVault(uint64(m.BlockHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.BlockTime)
	n += 1 + l + sovVault(uint64(l))
	return n
}

func (m *EmissionsVault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovVault(uint64(m.Id))
	}
	l = m.AmountIn.Size()
	n += 1 + l + sovVault(uint64(l))
	l = m.AmountOut.Size()
	n += 1 + l + sovVault(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovVault(uint64(l))
	return n
}

func (m *EmissionsVaultRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovVault(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovVault(uint64(m.BlockHeight))
	}
	l = m.Amount.Size()
	n += 1 + l + sovVault(uint64(l))
	return n
}

func sovVault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVault(x uint64) (n int) {
	return sovVault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Vault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestAccumulated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InterestAccumulated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosingFeeAccumulated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClosingFeeAccumulated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.BlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EmissionsVault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmissionsVault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmissionsVault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EmissionsVaultRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmissionsVaultRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmissionsVaultRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVault
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVault
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVault
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthVault
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVault
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVault
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVault        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVault          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVault = fmt.Errorf("proto: unexpected end of group")
)
