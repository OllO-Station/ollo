// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/prices/v1/prices.proto

package types

import (
	fmt "fmt"
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

// DataProvider is the type defined for feed data provider
type DataProvider struct {
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	PubKey  []byte                                        `protobuf:"bytes,2,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
}

func (m *DataProvider) Reset()         { *m = DataProvider{} }
func (m *DataProvider) String() string { return proto.CompactTextString(m) }
func (*DataProvider) ProtoMessage()    {}
func (*DataProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_407eea22785d1762, []int{0}
}
func (m *DataProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataProvider.Merge(m, src)
}
func (m *DataProvider) XXX_Size() int {
	return m.Size()
}
func (m *DataProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_DataProvider.DiscardUnknown(m)
}

var xxx_messageInfo_DataProvider proto.InternalMessageInfo

func (m *DataProvider) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *DataProvider) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type MsgModuleOwner struct {
	// address defines the address of the module owner
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty" yaml:"address"`
	// pubKey defined the public key of the module owner
	PubKey []byte `protobuf:"bytes,2,opt,name=pubKey,proto3" json:"pubKey,omitempty" yaml:"pub_key"`
	// the module owner who assigned this new module owner
	AssignerAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=assignerAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"assignerAddress,omitempty"`
}

func (m *MsgModuleOwner) Reset()         { *m = MsgModuleOwner{} }
func (m *MsgModuleOwner) String() string { return proto.CompactTextString(m) }
func (*MsgModuleOwner) ProtoMessage()    {}
func (*MsgModuleOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_407eea22785d1762, []int{1}
}
func (m *MsgModuleOwner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgModuleOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgModuleOwner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgModuleOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgModuleOwner.Merge(m, src)
}
func (m *MsgModuleOwner) XXX_Size() int {
	return m.Size()
}
func (m *MsgModuleOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgModuleOwner.DiscardUnknown(m)
}

var xxx_messageInfo_MsgModuleOwner proto.InternalMessageInfo

func (m *MsgModuleOwner) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *MsgModuleOwner) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *MsgModuleOwner) GetAssignerAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.AssignerAddress
	}
	return nil
}

// this will be the implementation used later will use pseudo OCR ABI encoded data instead
// because the structure of how the OCR will be generalized is still unknown
// OCRAbiEncoded implments the OCR data that is ABCI encoded. The use and form will conform to the
// Chainlink protocol specification.
type OCRAbiEncoded struct {
	// Context should be a 32-byte array struct.
	Context []byte `protobuf:"bytes,1,opt,name=Context,proto3" json:"Context,omitempty"`
	// Oracles should be a 32-byte record of all participating oracles. Assuming this is data provider address?
	Oracles []byte `protobuf:"bytes,2,opt,name=Oracles,proto3" json:"Oracles,omitempty"`
	// Observations should be an array on int192 containing the providers' independent observations.
	Observations []*Observation `protobuf:"bytes,3,rep,name=Observations,proto3" json:"Observations,omitempty"`
}

func (m *OCRAbiEncoded) Reset()         { *m = OCRAbiEncoded{} }
func (m *OCRAbiEncoded) String() string { return proto.CompactTextString(m) }
func (*OCRAbiEncoded) ProtoMessage()    {}
func (*OCRAbiEncoded) Descriptor() ([]byte, []int) {
	return fileDescriptor_407eea22785d1762, []int{2}
}
func (m *OCRAbiEncoded) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OCRAbiEncoded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OCRAbiEncoded.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OCRAbiEncoded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OCRAbiEncoded.Merge(m, src)
}
func (m *OCRAbiEncoded) XXX_Size() int {
	return m.Size()
}
func (m *OCRAbiEncoded) XXX_DiscardUnknown() {
	xxx_messageInfo_OCRAbiEncoded.DiscardUnknown(m)
}

var xxx_messageInfo_OCRAbiEncoded proto.InternalMessageInfo

func (m *OCRAbiEncoded) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *OCRAbiEncoded) GetOracles() []byte {
	if m != nil {
		return m.Oracles
	}
	return nil
}

func (m *OCRAbiEncoded) GetObservations() []*Observation {
	if m != nil {
		return m.Observations
	}
	return nil
}

type Observation struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Observation) Reset()         { *m = Observation{} }
func (m *Observation) String() string { return proto.CompactTextString(m) }
func (*Observation) ProtoMessage()    {}
func (*Observation) Descriptor() ([]byte, []int) {
	return fileDescriptor_407eea22785d1762, []int{3}
}
func (m *Observation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Observation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Observation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Observation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Observation.Merge(m, src)
}
func (m *Observation) XXX_Size() int {
	return m.Size()
}
func (m *Observation) XXX_DiscardUnknown() {
	xxx_messageInfo_Observation.DiscardUnknown(m)
}

var xxx_messageInfo_Observation proto.InternalMessageInfo

func (m *Observation) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type FeedRewardSchema struct {
	// amount is the base value that rewarded to each valid data provider before designated strategy applied
	// amount is not allowed to be zero
	Amount uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// reward strategy name, must be a registered strategy
	// this is allowed to be empty, in which case every data provider will be rewarded the same amount token
	Strategy string `protobuf:"bytes,2,opt,name=strategy,proto3" json:"strategy,omitempty"`
}

func (m *FeedRewardSchema) Reset()         { *m = FeedRewardSchema{} }
func (m *FeedRewardSchema) String() string { return proto.CompactTextString(m) }
func (*FeedRewardSchema) ProtoMessage()    {}
func (*FeedRewardSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_407eea22785d1762, []int{4}
}
func (m *FeedRewardSchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeedRewardSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeedRewardSchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeedRewardSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedRewardSchema.Merge(m, src)
}
func (m *FeedRewardSchema) XXX_Size() int {
	return m.Size()
}
func (m *FeedRewardSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedRewardSchema.DiscardUnknown(m)
}

var xxx_messageInfo_FeedRewardSchema proto.InternalMessageInfo

func (m *FeedRewardSchema) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *FeedRewardSchema) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func init() {
	proto.RegisterType((*DataProvider)(nil), "ollo.prices.v1.DataProvider")
	proto.RegisterType((*MsgModuleOwner)(nil), "ollo.prices.v1.MsgModuleOwner")
	proto.RegisterType((*OCRAbiEncoded)(nil), "ollo.prices.v1.OCRAbiEncoded")
	proto.RegisterType((*Observation)(nil), "ollo.prices.v1.Observation")
	proto.RegisterType((*FeedRewardSchema)(nil), "ollo.prices.v1.FeedRewardSchema")
}

func init() { proto.RegisterFile("ollo/prices/v1/prices.proto", fileDescriptor_407eea22785d1762) }

var fileDescriptor_407eea22785d1762 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xae, 0xd9, 0xb4, 0x81, 0x57, 0x0a, 0x32, 0x08, 0x55, 0x9b, 0x94, 0x82, 0x4f, 0x13, 0x52,
	0x13, 0x15, 0xc4, 0x85, 0x0b, 0x6a, 0x0b, 0xbb, 0x4c, 0x53, 0x91, 0xb9, 0x21, 0xa1, 0xc9, 0xb1,
	0x9f, 0xb2, 0x68, 0x49, 0x1c, 0xd9, 0x4e, 0xb6, 0xdc, 0xf9, 0x01, 0xfc, 0xac, 0x1d, 0x77, 0xe4,
	0x54, 0xa1, 0xf6, 0x1f, 0x70, 0x44, 0x1c, 0x50, 0x12, 0x07, 0xd6, 0xdd, 0xd0, 0x4e, 0x7e, 0x9f,
	0xbf, 0xf7, 0x3e, 0x7f, 0x7e, 0xef, 0xe1, 0x03, 0x95, 0x24, 0x2a, 0xc8, 0x75, 0x2c, 0xc0, 0x04,
	0xe5, 0xc4, 0x45, 0x7e, 0xae, 0x95, 0x55, 0x64, 0x50, 0x93, 0xbe, 0xbb, 0x2a, 0x27, 0xfb, 0x4f,
	0x23, 0x15, 0xa9, 0x86, 0x0a, 0xea, 0xa8, 0xcd, 0xa2, 0x06, 0xf7, 0xdf, 0x73, 0xcb, 0x3f, 0x6a,
	0x55, 0xc6, 0x12, 0x34, 0x39, 0xc6, 0xbb, 0x5c, 0x4a, 0x0d, 0xc6, 0x0c, 0xd1, 0x73, 0x74, 0xd8,
	0x9f, 0x4d, 0x7e, 0x2d, 0x47, 0xe3, 0x28, 0xb6, 0x67, 0x45, 0xe8, 0x0b, 0x95, 0x06, 0x42, 0x99,
	0x54, 0x19, 0x77, 0x8c, 0x8d, 0x3c, 0x0f, 0x6c, 0x95, 0x83, 0xf1, 0xa7, 0x42, 0x4c, 0xdb, 0x42,
	0xd6, 0x29, 0x90, 0x67, 0x78, 0x27, 0x2f, 0xc2, 0x63, 0xa8, 0x86, 0xf7, 0x6a, 0x2d, 0xe6, 0x10,
	0xfd, 0x8d, 0xf0, 0xe0, 0xc4, 0x44, 0x27, 0x4a, 0x16, 0x09, 0x2c, 0x2e, 0x32, 0xd0, 0xe4, 0xcb,
	0xed, 0x77, 0xe7, 0x3f, 0x97, 0xa3, 0x41, 0xc5, 0xd3, 0xe4, 0x2d, 0x75, 0x04, 0xbd, 0x83, 0x93,
	0x97, 0x9b, 0x4e, 0x66, 0xe4, 0x9f, 0x7a, 0x5e, 0x84, 0xa7, 0xe7, 0x50, 0xd1, 0xce, 0x1d, 0x39,
	0xc5, 0x8f, 0xb8, 0x31, 0x71, 0x94, 0x81, 0x76, 0x3a, 0xc3, 0xad, 0xa6, 0xe8, 0xcd, 0xd5, 0x72,
	0x84, 0xfe, 0xdf, 0xc4, 0x6d, 0x35, 0xfa, 0x15, 0xe1, 0x87, 0x8b, 0x39, 0x9b, 0x86, 0xf1, 0x87,
	0x4c, 0x28, 0x09, 0x92, 0x0c, 0xf1, 0xee, 0x5c, 0x65, 0x16, 0x2e, 0x6d, 0xfb, 0x7b, 0xd6, 0xc1,
	0x9a, 0x59, 0x68, 0x2e, 0x12, 0x30, 0xae, 0x87, 0x1d, 0x24, 0xef, 0x70, 0x7f, 0x11, 0x1a, 0xd0,
	0x25, 0xb7, 0xb1, 0xca, 0x6a, 0x8f, 0x5b, 0x87, 0x7b, 0xaf, 0x0e, 0xfc, 0xcd, 0xb1, 0xfb, 0x37,
	0x72, 0xd8, 0x46, 0x01, 0x7d, 0x81, 0xf7, 0x6e, 0x60, 0x42, 0xf0, 0xb6, 0xe4, 0x96, 0x3b, 0x03,
	0x4d, 0x4c, 0x8f, 0xf0, 0xe3, 0x23, 0x00, 0xc9, 0xe0, 0x82, 0x6b, 0xf9, 0x49, 0x9c, 0x41, 0xca,
	0xeb, 0xa1, 0xf2, 0x54, 0x15, 0x59, 0x6b, 0x75, 0x9b, 0x39, 0x44, 0xf6, 0xf1, 0x7d, 0x63, 0x35,
	0xb7, 0x10, 0xb5, 0x4d, 0x7e, 0xc0, 0xfe, 0xe2, 0xd9, 0xf8, 0x6a, 0xe5, 0xa1, 0xeb, 0x95, 0x87,
	0x7e, 0xac, 0x3c, 0xf4, 0x6d, 0xed, 0xf5, 0xae, 0xd7, 0x5e, 0xef, 0xfb, 0xda, 0xeb, 0x7d, 0x7e,
	0xd2, 0xac, 0xf0, 0x65, 0xb7, 0xc4, 0x4d, 0xf3, 0xc2, 0x9d, 0x66, 0x37, 0x5f, 0xff, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0x14, 0x30, 0x29, 0xe0, 0x02, 0x00, 0x00,
}

func (m *DataProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintPrices(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPrices(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgModuleOwner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgModuleOwner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgModuleOwner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssignerAddress) > 0 {
		i -= len(m.AssignerAddress)
		copy(dAtA[i:], m.AssignerAddress)
		i = encodeVarintPrices(dAtA, i, uint64(len(m.AssignerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintPrices(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPrices(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OCRAbiEncoded) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OCRAbiEncoded) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OCRAbiEncoded) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Observations) > 0 {
		for iNdEx := len(m.Observations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Observations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPrices(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Oracles) > 0 {
		i -= len(m.Oracles)
		copy(dAtA[i:], m.Oracles)
		i = encodeVarintPrices(dAtA, i, uint64(len(m.Oracles)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Context) > 0 {
		i -= len(m.Context)
		copy(dAtA[i:], m.Context)
		i = encodeVarintPrices(dAtA, i, uint64(len(m.Context)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Observation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Observation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Observation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPrices(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FeedRewardSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeedRewardSchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeedRewardSchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Strategy) > 0 {
		i -= len(m.Strategy)
		copy(dAtA[i:], m.Strategy)
		i = encodeVarintPrices(dAtA, i, uint64(len(m.Strategy)))
		i--
		dAtA[i] = 0x12
	}
	if m.Amount != 0 {
		i = encodeVarintPrices(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrices(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrices(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPrices(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovPrices(uint64(l))
	}
	return n
}

func (m *MsgModuleOwner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPrices(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovPrices(uint64(l))
	}
	l = len(m.AssignerAddress)
	if l > 0 {
		n += 1 + l + sovPrices(uint64(l))
	}
	return n
}

func (m *OCRAbiEncoded) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Context)
	if l > 0 {
		n += 1 + l + sovPrices(uint64(l))
	}
	l = len(m.Oracles)
	if l > 0 {
		n += 1 + l + sovPrices(uint64(l))
	}
	if len(m.Observations) > 0 {
		for _, e := range m.Observations {
			l = e.Size()
			n += 1 + l + sovPrices(uint64(l))
		}
	}
	return n
}

func (m *Observation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPrices(uint64(l))
	}
	return n
}

func (m *FeedRewardSchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Amount != 0 {
		n += 1 + sovPrices(uint64(m.Amount))
	}
	l = len(m.Strategy)
	if l > 0 {
		n += 1 + l + sovPrices(uint64(l))
	}
	return n
}

func sovPrices(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrices(x uint64) (n int) {
	return sovPrices(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataProvider) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrices
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
			return fmt.Errorf("proto: DataProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrices
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
func (m *MsgModuleOwner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrices
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
			return fmt.Errorf("proto: MsgModuleOwner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgModuleOwner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssignerAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssignerAddress = append(m.AssignerAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.AssignerAddress == nil {
				m.AssignerAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrices
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
func (m *OCRAbiEncoded) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrices
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
			return fmt.Errorf("proto: OCRAbiEncoded: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OCRAbiEncoded: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = append(m.Context[:0], dAtA[iNdEx:postIndex]...)
			if m.Context == nil {
				m.Context = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracles", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Oracles = append(m.Oracles[:0], dAtA[iNdEx:postIndex]...)
			if m.Oracles == nil {
				m.Oracles = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Observations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
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
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Observations = append(m.Observations, &Observation{})
			if err := m.Observations[len(m.Observations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrices
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
func (m *Observation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrices
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
			return fmt.Errorf("proto: Observation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Observation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrices
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
func (m *FeedRewardSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrices
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
			return fmt.Errorf("proto: FeedRewardSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeedRewardSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Strategy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrices
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
				return ErrInvalidLengthPrices
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrices
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Strategy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrices
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
func skipPrices(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrices
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
					return 0, ErrIntOverflowPrices
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
					return 0, ErrIntOverflowPrices
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
				return 0, ErrInvalidLengthPrices
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrices
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrices
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrices        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrices          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrices = fmt.Errorf("proto: unexpected end of group")
)
