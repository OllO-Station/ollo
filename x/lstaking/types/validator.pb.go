// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/lstaking/v1/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ValidatorStatus defines the status of a validator.
type ValidatorStatus int32

const (
	// VALIDATOR_STATUS_UNSPECIFIED defines an invalid validator status.
	ValidatorStatusUnspecified ValidatorStatus = 0
	// VALIDATOR_STATUS_ACTIVE defines an active validator.
	ValidatorStatusActive ValidatorStatus = 1
	// VALIDATOR_STATUS_INACTIVE defines an inactive validator.
	ValidatorStatusInactive ValidatorStatus = 2
)

var ValidatorStatus_name = map[int32]string{
	0: "VALIDATOR_STATUS_UNSPECIFIED",
	1: "VALIDATOR_STATUS_ACTIVE",
	2: "VALIDATOR_STATUS_INACTIVE",
}

var ValidatorStatus_value = map[string]int32{
	"VALIDATOR_STATUS_UNSPECIFIED": 0,
	"VALIDATOR_STATUS_ACTIVE":      1,
	"VALIDATOR_STATUS_INACTIVE":    2,
}

func (x ValidatorStatus) String() string {
	return proto.EnumName(ValidatorStatus_name, int32(x))
}

func (ValidatorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_befdb61b51919114, []int{0}
}

// LiquidValidator defines a liquid validator
type LiquidValidator struct {
	// operator_address is the address of the validator.
	OperatorAddress github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"operator_address,omitempty" yaml:"operator_address"`
}

func (m *LiquidValidator) Reset()         { *m = LiquidValidator{} }
func (m *LiquidValidator) String() string { return proto.CompactTextString(m) }
func (*LiquidValidator) ProtoMessage()    {}
func (*LiquidValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_befdb61b51919114, []int{0}
}
func (m *LiquidValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidValidator.Merge(m, src)
}
func (m *LiquidValidator) XXX_Size() int {
	return m.Size()
}
func (m *LiquidValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidValidator.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidValidator proto.InternalMessageInfo

// WhitelistValidator defines a whitelisted validator
type WhitelistedValidator struct {
	// operator_address is the address of the validator.
	OperatorAddress github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"operator_address,omitempty" yaml:"operator_address"`
	// target weight for liquid staking
	TargetWeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=target_weight,json=targetWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"target_weight" yaml:"target_weight"`
}

func (m *WhitelistedValidator) Reset()         { *m = WhitelistedValidator{} }
func (m *WhitelistedValidator) String() string { return proto.CompactTextString(m) }
func (*WhitelistedValidator) ProtoMessage()    {}
func (*WhitelistedValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_befdb61b51919114, []int{1}
}
func (m *WhitelistedValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhitelistedValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhitelistedValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhitelistedValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistedValidator.Merge(m, src)
}
func (m *WhitelistedValidator) XXX_Size() int {
	return m.Size()
}
func (m *WhitelistedValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistedValidator.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistedValidator proto.InternalMessageInfo

// liquid validator with added state info
type LiquidValidatorState struct {
	// operator_address is the address of the validator.
	OperatorAddress github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"operator_address,omitempty" yaml:"operator_address"`
	// target weight for liquid staking
	Weight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"weight" yaml:"weight"`
	// status of the validator
	Status ValidatorStatus `protobuf:"varint,3,opt,name=status,proto3,enum=ollo.lstaking.v1.ValidatorStatus" json:"status,omitempty" yaml:"status"`
	// delegation shares of the validator
	DelegationShares github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=delegation_shares,json=delegationShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"delegation_shares" yaml:"delegation_shares"`
	// liquid tokens defines the worth of delegation shares
	LiquidTokens github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=liquid_tokens,json=liquidTokens,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liquid_tokens" yaml:"liquid_tokens"`
}

func (m *LiquidValidatorState) Reset()         { *m = LiquidValidatorState{} }
func (m *LiquidValidatorState) String() string { return proto.CompactTextString(m) }
func (*LiquidValidatorState) ProtoMessage()    {}
func (*LiquidValidatorState) Descriptor() ([]byte, []int) {
	return fileDescriptor_befdb61b51919114, []int{2}
}
func (m *LiquidValidatorState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidValidatorState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidValidatorState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidValidatorState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidValidatorState.Merge(m, src)
}
func (m *LiquidValidatorState) XXX_Size() int {
	return m.Size()
}
func (m *LiquidValidatorState) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidValidatorState.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidValidatorState proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ollo.lstaking.v1.ValidatorStatus", ValidatorStatus_name, ValidatorStatus_value)
	proto.RegisterType((*LiquidValidator)(nil), "ollo.lstaking.v1.LiquidValidator")
	proto.RegisterType((*WhitelistedValidator)(nil), "ollo.lstaking.v1.WhitelistedValidator")
	proto.RegisterType((*LiquidValidatorState)(nil), "ollo.lstaking.v1.LiquidValidatorState")
}

func init() { proto.RegisterFile("ollo/lstaking/v1/validator.proto", fileDescriptor_befdb61b51919114) }

var fileDescriptor_befdb61b51919114 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4f, 0x6b, 0xd4, 0x4e,
	0x18, 0x4e, 0xfa, 0xeb, 0xaf, 0xe8, 0xd0, 0xda, 0x6d, 0x58, 0x69, 0x1a, 0x6b, 0x12, 0x23, 0x48,
	0x11, 0x9a, 0x50, 0x05, 0x0f, 0xbd, 0x68, 0xfa, 0x4f, 0x22, 0x4b, 0x95, 0xec, 0x76, 0x0b, 0x5e,
	0x42, 0xba, 0x19, 0xb3, 0xc3, 0xa6, 0x99, 0x35, 0x33, 0xbb, 0xb5, 0x5f, 0x40, 0xca, 0x9e, 0x3c,
	0x78, 0x5d, 0x28, 0x08, 0x7e, 0x96, 0x5e, 0x84, 0x1e, 0xc5, 0x43, 0x90, 0xdd, 0x8b, 0xe7, 0x3d,
	0x7a, 0x92, 0x64, 0xb2, 0x6c, 0x37, 0xb9, 0x58, 0x3c, 0xe8, 0x29, 0xf3, 0xce, 0xfb, 0x3e, 0xcf,
	0xbc, 0xf3, 0xbc, 0x4f, 0x06, 0xa8, 0x38, 0x08, 0xb0, 0x11, 0x10, 0xea, 0xb6, 0x50, 0xe8, 0x1b,
	0xdd, 0x0d, 0xa3, 0xeb, 0x06, 0xc8, 0x73, 0x29, 0x8e, 0xf4, 0x76, 0x84, 0x29, 0x16, 0x4a, 0x49,
	0x85, 0x3e, 0xae, 0xd0, 0xbb, 0x1b, 0x52, 0xd9, 0xc7, 0x3e, 0x4e, 0x93, 0x46, 0xb2, 0x62, 0x75,
	0xd2, 0x4a, 0x03, 0x93, 0x63, 0x4c, 0x1c, 0x96, 0x60, 0x41, 0x96, 0xba, 0x5b, 0x3c, 0x04, 0xd3,
	0x84, 0x8c, 0xa5, 0x57, 0x0b, 0x69, 0x42, 0x5d, 0x0a, 0xb3, 0xec, 0xfd, 0x42, 0x36, 0x82, 0x1e,
	0x0c, 0xa0, 0xef, 0x52, 0x84, 0x43, 0x56, 0xa4, 0x7d, 0xe4, 0xc1, 0x62, 0x05, 0xbd, 0xed, 0x20,
	0xaf, 0x3e, 0x6e, 0x5f, 0x38, 0x01, 0x25, 0xdc, 0x86, 0x51, 0xb2, 0x76, 0x5c, 0xcf, 0x8b, 0x20,
	0x21, 0x22, 0xaf, 0xf2, 0x6b, 0x37, 0xb7, 0x2a, 0xa3, 0x58, 0x59, 0x3e, 0x75, 0x8f, 0x83, 0x4d,
	0x2d, 0x5f, 0xa1, 0xfd, 0x8c, 0x95, 0x75, 0x1f, 0xd1, 0x66, 0xe7, 0x48, 0x6f, 0xe0, 0xe3, 0xec,
	0x1e, 0xd9, 0x67, 0x9d, 0x78, 0x2d, 0x83, 0x9e, 0xb6, 0x21, 0xd1, 0xeb, 0x6e, 0x60, 0x32, 0x84,
	0xbd, 0x38, 0xe6, 0xc8, 0x36, 0x36, 0x6f, 0x9c, 0x9d, 0x2b, 0xdc, 0x8f, 0x73, 0x85, 0xd7, 0xde,
	0xcf, 0x80, 0xf2, 0x61, 0x13, 0x51, 0x18, 0x20, 0x42, 0xe1, 0x3f, 0xd0, 0x9b, 0xd0, 0x02, 0x0b,
	0xd4, 0x8d, 0x7c, 0x48, 0x9d, 0x13, 0x88, 0xfc, 0x26, 0x15, 0x67, 0xd2, 0x53, 0xf7, 0x2e, 0x62,
	0x85, 0xfb, 0x16, 0x2b, 0x0f, 0x7e, 0x83, 0xde, 0x0a, 0xe9, 0x28, 0x56, 0xca, 0xac, 0xc7, 0x29,
	0x32, 0xcd, 0x9e, 0x67, 0xf1, 0x61, 0x1a, 0x5e, 0x11, 0xe2, 0xf3, 0x2c, 0x28, 0xe7, 0xe6, 0x53,
	0x4d, 0x66, 0xfc, 0xf7, 0x84, 0x38, 0x04, 0x73, 0x53, 0x0a, 0x3c, 0xbd, 0xb6, 0x02, 0x0b, 0xac,
	0xb9, 0xf1, 0xd5, 0x33, 0x3a, 0xa1, 0x02, 0xe6, 0x12, 0xfb, 0x76, 0x88, 0xf8, 0x9f, 0xca, 0xaf,
	0xdd, 0x7a, 0x74, 0x4f, 0xcf, 0xff, 0x40, 0xfa, 0x94, 0x06, 0x1d, 0xb2, 0xb5, 0x34, 0x61, 0x63,
	0x50, 0xcd, 0xce, 0x38, 0x84, 0x13, 0xb0, 0x34, 0x31, 0xbb, 0x43, 0x9a, 0x6e, 0x04, 0x89, 0x38,
	0x9b, 0x76, 0xfc, 0xe2, 0x1a, 0x1d, 0xef, 0xc0, 0xc6, 0x28, 0x56, 0x44, 0x76, 0x46, 0x81, 0x50,
	0xb3, 0x4b, 0x93, 0xbd, 0x6a, 0xba, 0x95, 0x18, 0x25, 0x48, 0x07, 0xe6, 0x50, 0xdc, 0x82, 0x21,
	0x11, 0xff, 0xff, 0x33, 0xa3, 0x4c, 0x91, 0x69, 0xf6, 0x3c, 0x8b, 0x6b, 0x69, 0x38, 0x31, 0xca,
	0xc3, 0x2f, 0x3c, 0x58, 0xcc, 0xc9, 0x23, 0x3c, 0x03, 0xab, 0x75, 0xb3, 0x62, 0xed, 0x98, 0xb5,
	0x97, 0xb6, 0x53, 0xad, 0x99, 0xb5, 0x83, 0xaa, 0x73, 0xb0, 0x5f, 0x7d, 0xb5, 0xbb, 0x6d, 0xed,
	0x59, 0xbb, 0x3b, 0x25, 0x4e, 0x92, 0x7b, 0x7d, 0x55, 0xca, 0xc1, 0x0e, 0x42, 0xd2, 0x86, 0x0d,
	0xf4, 0x06, 0x41, 0x4f, 0x78, 0x02, 0x96, 0x0b, 0x0c, 0xe6, 0x76, 0xcd, 0xaa, 0xef, 0x96, 0x78,
	0x69, 0xa5, 0xd7, 0x57, 0x6f, 0xe7, 0xc0, 0x66, 0x83, 0xa2, 0x2e, 0x14, 0x36, 0xc1, 0x4a, 0x01,
	0x67, 0xed, 0x67, 0xc8, 0x19, 0xe9, 0x4e, 0xaf, 0xaf, 0x2e, 0xe7, 0x90, 0x56, 0xe8, 0xa6, 0x58,
	0x69, 0xf6, 0xec, 0x93, 0xcc, 0x6d, 0x3d, 0xbf, 0x18, 0xc8, 0xfc, 0xe5, 0x40, 0xe6, 0xbf, 0x0f,
	0x64, 0xfe, 0xc3, 0x50, 0xe6, 0x2e, 0x87, 0x32, 0xf7, 0x75, 0x28, 0x73, 0xaf, 0xaf, 0x1a, 0x38,
	0x71, 0xc8, 0x7a, 0x32, 0x71, 0x84, 0xc3, 0x34, 0x30, 0xde, 0x4d, 0x5e, 0xbc, 0x54, 0xcc, 0xa3,
	0xb9, 0xf4, 0xa1, 0x7b, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xed, 0x8b, 0x51, 0x75, 0xb1, 0x05,
	0x00, 0x00,
}

func (this *LiquidValidator) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LiquidValidator)
	if !ok {
		that2, ok := that.(LiquidValidator)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.OperatorAddress != that1.OperatorAddress {
		return false
	}
	return true
}
func (this *WhitelistedValidator) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WhitelistedValidator)
	if !ok {
		that2, ok := that.(WhitelistedValidator)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.OperatorAddress != that1.OperatorAddress {
		return false
	}
	if !this.TargetWeight.Equal(that1.TargetWeight) {
		return false
	}
	return true
}
func (this *LiquidValidatorState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LiquidValidatorState)
	if !ok {
		that2, ok := that.(LiquidValidatorState)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.OperatorAddress != that1.OperatorAddress {
		return false
	}
	if !this.Weight.Equal(that1.Weight) {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !this.DelegationShares.Equal(that1.DelegationShares) {
		return false
	}
	if !this.LiquidTokens.Equal(that1.LiquidTokens) {
		return false
	}
	return true
}
func (m *LiquidValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WhitelistedValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhitelistedValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhitelistedValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TargetWeight.Size()
		i -= size
		if _, err := m.TargetWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LiquidValidatorState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidValidatorState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidValidatorState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LiquidTokens.Size()
		i -= size
		if _, err := m.LiquidTokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.DelegationShares.Size()
		i -= size
		if _, err := m.DelegationShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Status != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LiquidValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func (m *WhitelistedValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = m.TargetWeight.Size()
	n += 1 + l + sovValidator(uint64(l))
	return n
}

func (m *LiquidValidatorState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.Status != 0 {
		n += 1 + sovValidator(uint64(m.Status))
	}
	l = m.DelegationShares.Size()
	n += 1 + l + sovValidator(uint64(l))
	l = m.LiquidTokens.Size()
	n += 1 + l + sovValidator(uint64(l))
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LiquidValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: LiquidValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = github_com_cosmos_cosmos_sdk_types.ValAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *WhitelistedValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: WhitelistedValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhitelistedValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = github_com_cosmos_cosmos_sdk_types.ValAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TargetWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *LiquidValidatorState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: LiquidValidatorState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidValidatorState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = github_com_cosmos_cosmos_sdk_types.ValAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ValidatorStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegationShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
