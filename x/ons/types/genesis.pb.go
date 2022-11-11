// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/ons/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the ons module's genesis state.
type GenesisState struct {
	Params Params    `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	State  NameState `protobuf:"bytes,2,opt,name=state,proto3" json:"state"`
	PortId string    `protobuf:"bytes,3,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f995f9cebb1a26, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetState() NameState {
	if m != nil {
		return m.State
	}
	return NameState{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

type NameState struct {
	Names             []Name             `protobuf:"bytes,1,rep,name=names,proto3" json:"names"`
	Threads           []Thread           `protobuf:"bytes,2,rep,name=threads,proto3" json:"threads"`
	ThreadTags        []ThreadTag        `protobuf:"bytes,3,rep,name=thread_tags,json=threadTags,proto3" json:"thread_tags"`
	BuyOffers         []BuyNameOffer     `protobuf:"bytes,4,rep,name=buy_offers,json=buyOffers,proto3" json:"buy_offers"`
	SellOffers        []SellNameOffer    `protobuf:"bytes,5,rep,name=sell_offers,json=sellOffers,proto3" json:"sell_offers"`
	ActiveLoans       []LoanName         `protobuf:"bytes,6,rep,name=active_loans,json=activeLoans,proto3" json:"active_loans"`
	NameTags          []NameTag          `protobuf:"bytes,7,rep,name=name_tags,json=nameTags,proto3" json:"name_tags"`
	ThreadMessage     []ThreadMessage    `protobuf:"bytes,8,rep,name=thread_message,json=threadMessage,proto3" json:"thread_message"`
	ThreadMessageTags []ThreadMessageTag `protobuf:"bytes,9,rep,name=thread_message_tags,json=threadMessageTags,proto3" json:"thread_message_tags"`
	ActionTag         []ActionTag        `protobuf:"bytes,10,rep,name=action_tag,json=actionTag,proto3" json:"action_tag"`
}

func (m *NameState) Reset()         { *m = NameState{} }
func (m *NameState) String() string { return proto.CompactTextString(m) }
func (*NameState) ProtoMessage()    {}
func (*NameState) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f995f9cebb1a26, []int{1}
}
func (m *NameState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NameState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NameState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NameState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameState.Merge(m, src)
}
func (m *NameState) XXX_Size() int {
	return m.Size()
}
func (m *NameState) XXX_DiscardUnknown() {
	xxx_messageInfo_NameState.DiscardUnknown(m)
}

var xxx_messageInfo_NameState proto.InternalMessageInfo

func (m *NameState) GetNames() []Name {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *NameState) GetThreads() []Thread {
	if m != nil {
		return m.Threads
	}
	return nil
}

func (m *NameState) GetThreadTags() []ThreadTag {
	if m != nil {
		return m.ThreadTags
	}
	return nil
}

func (m *NameState) GetBuyOffers() []BuyNameOffer {
	if m != nil {
		return m.BuyOffers
	}
	return nil
}

func (m *NameState) GetSellOffers() []SellNameOffer {
	if m != nil {
		return m.SellOffers
	}
	return nil
}

func (m *NameState) GetActiveLoans() []LoanName {
	if m != nil {
		return m.ActiveLoans
	}
	return nil
}

func (m *NameState) GetNameTags() []NameTag {
	if m != nil {
		return m.NameTags
	}
	return nil
}

func (m *NameState) GetThreadMessage() []ThreadMessage {
	if m != nil {
		return m.ThreadMessage
	}
	return nil
}

func (m *NameState) GetThreadMessageTags() []ThreadMessageTag {
	if m != nil {
		return m.ThreadMessageTags
	}
	return nil
}

func (m *NameState) GetActionTag() []ActionTag {
	if m != nil {
		return m.ActionTag
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ollo.ons.v1beta1.GenesisState")
	proto.RegisterType((*NameState)(nil), "ollo.ons.v1beta1.NameState")
}

func init() { proto.RegisterFile("ollo/ons/v1beta1/genesis.proto", fileDescriptor_38f995f9cebb1a26) }

var fileDescriptor_38f995f9cebb1a26 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xed, 0xa6, 0xf9, 0xe3, 0x71, 0x7f, 0x3f, 0xc1, 0x82, 0x60, 0x49, 0x84, 0x53, 0xe5,
	0x14, 0x71, 0xb0, 0xd5, 0x20, 0x01, 0x42, 0x1c, 0x20, 0x95, 0x40, 0x48, 0xe5, 0x8f, 0xda, 0x1e,
	0x10, 0x97, 0x68, 0x43, 0xb6, 0x26, 0xc2, 0xf6, 0x46, 0xde, 0x4d, 0x45, 0xde, 0x80, 0x23, 0x8f,
	0x50, 0xf1, 0x12, 0xbc, 0x42, 0x8f, 0x3d, 0x72, 0x42, 0x28, 0xb9, 0xf0, 0x18, 0x68, 0x67, 0x37,
	0x4d, 0x89, 0x63, 0x71, 0xdb, 0x9d, 0xf9, 0x7e, 0x3f, 0x3b, 0x33, 0xeb, 0x35, 0x04, 0x22, 0x49,
	0x44, 0x24, 0x32, 0x19, 0x9d, 0xee, 0x0d, 0xb9, 0x62, 0x7b, 0x51, 0xcc, 0x33, 0x2e, 0xc7, 0x32,
	0x9c, 0xe4, 0x42, 0x09, 0x72, 0x4d, 0xe7, 0x43, 0x91, 0xc9, 0xd0, 0xe6, 0x9b, 0x37, 0x63, 0x11,
	0x0b, 0x4c, 0x46, 0x7a, 0x65, 0x74, 0xcd, 0xbb, 0x05, 0xce, 0x84, 0xe5, 0x2c, 0xb5, 0x98, 0x66,
	0xab, 0x90, 0xce, 0x58, 0xca, 0x4b, 0xbd, 0x29, 0xcb, 0x3f, 0x71, 0x65, 0xd3, 0xc5, 0x12, 0xd9,
	0x07, 0x35, 0xd6, 0x25, 0x95, 0xd9, 0xd5, 0xc7, 0x9c, 0xb3, 0x91, 0x49, 0x77, 0xbe, 0xb9, 0xb0,
	0xf3, 0xc2, 0xf4, 0x74, 0xa4, 0x98, 0xe2, 0xe4, 0x01, 0xd4, 0x4c, 0x6d, 0xd4, 0xdd, 0x75, 0xbb,
	0x7e, 0x8f, 0x86, 0xeb, 0x3d, 0x86, 0x6f, 0x31, 0xdf, 0xdf, 0x3e, 0xff, 0xd9, 0x76, 0x0e, 0xad,
	0x9a, 0x3c, 0x84, 0xaa, 0xd4, 0x00, 0xba, 0x85, 0xb6, 0x56, 0xd1, 0xf6, 0x9a, 0xa5, 0x1c, 0xcf,
	0xb0, 0x4e, 0xa3, 0x27, 0xb7, 0xa1, 0x3e, 0x11, 0xb9, 0x1a, 0x8c, 0x47, 0xb4, 0xb2, 0xeb, 0x76,
	0xbd, 0xc3, 0x9a, 0xde, 0xbe, 0x1c, 0x3d, 0x6e, 0x7c, 0x39, 0x6b, 0xbb, 0xbf, 0xcf, 0xda, 0x4e,
	0xe7, 0x7b, 0x15, 0xbc, 0x4b, 0x37, 0xe9, 0x41, 0x55, 0x8f, 0x47, 0x17, 0x58, 0xe9, 0xfa, 0xbd,
	0x5b, 0x9b, 0x4f, 0x5a, 0x1e, 0x82, 0x52, 0xf2, 0x08, 0xea, 0xa6, 0x6d, 0x49, 0xb7, 0xd0, 0xb5,
	0xa1, 0xad, 0x63, 0x14, 0x58, 0xdf, 0x52, 0x4e, 0xfa, 0xe0, 0x9b, 0xe5, 0x40, 0xb1, 0x58, 0xd2,
	0x0a, 0xba, 0x5b, 0x65, 0xee, 0x63, 0x16, 0x5b, 0x00, 0xa8, 0x65, 0x40, 0x92, 0x7d, 0x80, 0xe1,
	0x74, 0x36, 0x10, 0x27, 0x27, 0x3c, 0x97, 0x74, 0x1b, 0x11, 0x41, 0x11, 0xd1, 0x9f, 0xce, 0x74,
	0xe5, 0x6f, 0xb4, 0xcc, 0x52, 0xbc, 0xe1, 0x74, 0x86, 0x7b, 0x49, 0x9e, 0x83, 0x2f, 0x79, 0x92,
	0x2c, 0x29, 0x55, 0xa4, 0xb4, 0x8b, 0x94, 0x23, 0x9e, 0x24, 0xeb, 0x18, 0xd0, 0x4e, 0xcb, 0xd9,
	0x87, 0x1d, 0xfd, 0x85, 0x9c, 0xf2, 0x41, 0x22, 0x58, 0x26, 0x69, 0x0d, 0x41, 0xcd, 0x22, 0xe8,
	0x40, 0xb0, 0xec, 0xca, 0x24, 0x7d, 0xe3, 0xd2, 0x51, 0x49, 0x9e, 0x80, 0xa7, 0x07, 0x6b, 0x66,
	0x52, 0x47, 0xc2, 0x9d, 0xcd, 0xf7, 0xb0, 0x9a, 0x48, 0x23, 0x33, 0x5b, 0x49, 0x0e, 0xe0, 0x7f,
	0x3b, 0xd3, 0x94, 0x4b, 0xc9, 0x62, 0x4e, 0x1b, 0x65, 0xdd, 0x98, 0xb1, 0xbe, 0x32, 0x32, 0x0b,
	0xfa, 0x4f, 0x5d, 0x0d, 0x92, 0x77, 0x70, 0xe3, 0x6f, 0x9a, 0xa9, 0xca, 0x43, 0x64, 0xe7, 0x1f,
	0xc8, 0x55, 0x79, 0xd7, 0xd5, 0x5a, 0x5c, 0x92, 0xa7, 0x00, 0xe6, 0x31, 0x69, 0x22, 0x85, 0xb2,
	0xab, 0x7f, 0x86, 0x9a, 0x15, 0xc9, 0x63, 0x97, 0x81, 0x7b, 0xe7, 0xf3, 0xc0, 0xbd, 0x98, 0x07,
	0xee, 0xaf, 0x79, 0xe0, 0x7e, 0x5d, 0x04, 0xce, 0xc5, 0x22, 0x70, 0x7e, 0x2c, 0x02, 0xe7, 0x3d,
	0xfe, 0x3a, 0xa2, 0xcf, 0xf8, 0x32, 0xd5, 0x6c, 0xc2, 0xe5, 0xb0, 0x86, 0x2f, 0xf2, 0xfe, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0xf0, 0x87, 0x65, 0x75, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.State.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *NameState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NameState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NameState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionTag) > 0 {
		for iNdEx := len(m.ActionTag) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActionTag[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.ThreadMessageTags) > 0 {
		for iNdEx := len(m.ThreadMessageTags) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ThreadMessageTags[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.ThreadMessage) > 0 {
		for iNdEx := len(m.ThreadMessage) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ThreadMessage[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.NameTags) > 0 {
		for iNdEx := len(m.NameTags) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NameTags[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ActiveLoans) > 0 {
		for iNdEx := len(m.ActiveLoans) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActiveLoans[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.SellOffers) > 0 {
		for iNdEx := len(m.SellOffers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SellOffers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BuyOffers) > 0 {
		for iNdEx := len(m.BuyOffers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BuyOffers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ThreadTags) > 0 {
		for iNdEx := len(m.ThreadTags) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ThreadTags[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Threads) > 0 {
		for iNdEx := len(m.Threads) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Threads[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Names) > 0 {
		for iNdEx := len(m.Names) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Names[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.State.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *NameState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Names) > 0 {
		for _, e := range m.Names {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Threads) > 0 {
		for _, e := range m.Threads {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ThreadTags) > 0 {
		for _, e := range m.ThreadTags {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BuyOffers) > 0 {
		for _, e := range m.BuyOffers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SellOffers) > 0 {
		for _, e := range m.SellOffers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ActiveLoans) > 0 {
		for _, e := range m.ActiveLoans {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NameTags) > 0 {
		for _, e := range m.NameTags {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ThreadMessage) > 0 {
		for _, e := range m.ThreadMessage {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ThreadMessageTags) > 0 {
		for _, e := range m.ThreadMessageTags {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ActionTag) > 0 {
		for _, e := range m.ActionTag {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.State.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *NameState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: NameState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NameState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Names", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Names = append(m.Names, Name{})
			if err := m.Names[len(m.Names)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threads", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Threads = append(m.Threads, Thread{})
			if err := m.Threads[len(m.Threads)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThreadTags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ThreadTags = append(m.ThreadTags, ThreadTag{})
			if err := m.ThreadTags[len(m.ThreadTags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyOffers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuyOffers = append(m.BuyOffers, BuyNameOffer{})
			if err := m.BuyOffers[len(m.BuyOffers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellOffers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellOffers = append(m.SellOffers, SellNameOffer{})
			if err := m.SellOffers[len(m.SellOffers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveLoans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActiveLoans = append(m.ActiveLoans, LoanName{})
			if err := m.ActiveLoans[len(m.ActiveLoans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NameTags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NameTags = append(m.NameTags, NameTag{})
			if err := m.NameTags[len(m.NameTags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThreadMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ThreadMessage = append(m.ThreadMessage, ThreadMessage{})
			if err := m.ThreadMessage[len(m.ThreadMessage)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThreadMessageTags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ThreadMessageTags = append(m.ThreadMessageTags, ThreadMessageTag{})
			if err := m.ThreadMessageTags[len(m.ThreadMessageTags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionTag", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActionTag = append(m.ActionTag, ActionTag{})
			if err := m.ActionTag[len(m.ActionTag)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
