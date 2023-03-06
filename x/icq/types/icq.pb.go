// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/icq/v1/icq.proto

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

type InterchainQuery struct {
	Id           string                                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ConnectionId string                                 `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	ChainId      string                                 `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	QueryType    string                                 `protobuf:"bytes,4,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	Request      []byte                                 `protobuf:"bytes,5,opt,name=request,proto3" json:"request,omitempty"`
	Period       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=period,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"period,omitempty" yaml:"period"`
	LastHeight   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=last_height,json=lastHeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_height,omitempty" yaml:"last_height"`
	CallbackId   string                                 `protobuf:"bytes,8,opt,name=callback_id,json=callbackId,proto3" json:"callback_id,omitempty"`
	Ttl          uint64                                 `protobuf:"varint,9,opt,name=ttl,proto3" json:"ttl,omitempty"`
	LastEmission github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=last_emission,json=lastEmission,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_emission,omitempty" yaml:"last_emission"`
}

func (m *InterchainQuery) Reset()         { *m = InterchainQuery{} }
func (m *InterchainQuery) String() string { return proto.CompactTextString(m) }
func (*InterchainQuery) ProtoMessage()    {}
func (*InterchainQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef067951d76fd899, []int{0}
}
func (m *InterchainQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterchainQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterchainQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterchainQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterchainQuery.Merge(m, src)
}
func (m *InterchainQuery) XXX_Size() int {
	return m.Size()
}
func (m *InterchainQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_InterchainQuery.DiscardUnknown(m)
}

var xxx_messageInfo_InterchainQuery proto.InternalMessageInfo

func (m *InterchainQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InterchainQuery) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *InterchainQuery) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *InterchainQuery) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *InterchainQuery) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *InterchainQuery) GetCallbackId() string {
	if m != nil {
		return m.CallbackId
	}
	return ""
}

func (m *InterchainQuery) GetTtl() uint64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type DataPoint struct {
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RemoteHeight uint64 `protobuf:"varint,2,opt,name=remote_height,json=remoteHeight,proto3" json:"remote_height,omitempty" yaml:"remote_height"`
	LocalHeight  string `protobuf:"bytes,3,opt,name=local_height,json=localHeight,proto3" json:"local_height,omitempty" yaml:"local_height"`
	Value        []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty" yaml:"value"`
}

func (m *DataPoint) Reset()         { *m = DataPoint{} }
func (m *DataPoint) String() string { return proto.CompactTextString(m) }
func (*DataPoint) ProtoMessage()    {}
func (*DataPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef067951d76fd899, []int{1}
}
func (m *DataPoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataPoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPoint.Merge(m, src)
}
func (m *DataPoint) XXX_Size() int {
	return m.Size()
}
func (m *DataPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPoint.DiscardUnknown(m)
}

var xxx_messageInfo_DataPoint proto.InternalMessageInfo

func (m *DataPoint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataPoint) GetRemoteHeight() uint64 {
	if m != nil {
		return m.RemoteHeight
	}
	return 0
}

func (m *DataPoint) GetLocalHeight() string {
	if m != nil {
		return m.LocalHeight
	}
	return ""
}

func (m *DataPoint) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*InterchainQuery)(nil), "ollo.icq.v1.InterchainQuery")
	proto.RegisterType((*DataPoint)(nil), "ollo.icq.v1.DataPoint")
}

func init() { proto.RegisterFile("ollo/icq/v1/icq.proto", fileDescriptor_ef067951d76fd899) }

var fileDescriptor_ef067951d76fd899 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xe3, 0x34, 0x4d, 0x9a, 0x17, 0x87, 0x56, 0x47, 0x0b, 0x6e, 0x25, 0xec, 0xc8, 0x95,
	0x50, 0x06, 0x92, 0x28, 0xea, 0x04, 0xa3, 0x05, 0x12, 0xde, 0xc0, 0xea, 0x80, 0x58, 0x22, 0xc7,
	0x3e, 0x25, 0xa7, 0xd8, 0xbe, 0xc4, 0xbe, 0x44, 0xf8, 0x0b, 0x30, 0x31, 0x30, 0xf0, 0x51, 0xf8,
	0x10, 0x1d, 0x2b, 0x26, 0xc4, 0x60, 0xa1, 0x64, 0xcb, 0xc0, 0xc0, 0x27, 0x40, 0x77, 0x67, 0x0b,
	0xbb, 0xb0, 0x74, 0xb2, 0xdf, 0xfb, 0xdd, 0xfd, 0xfd, 0xfe, 0xef, 0xde, 0x19, 0xce, 0x68, 0x10,
	0xd0, 0x11, 0xf1, 0x56, 0xa3, 0xcd, 0x98, 0x3f, 0x86, 0xcb, 0x98, 0x32, 0x8a, 0x3a, 0x3c, 0x3d,
	0xe4, 0xf1, 0x66, 0x7c, 0x71, 0x3a, 0xa3, 0x33, 0x2a, 0xf2, 0x23, 0xfe, 0x26, 0x97, 0x5c, 0x9c,
	0x7b, 0x34, 0x09, 0x69, 0x32, 0x91, 0x40, 0x06, 0x12, 0x99, 0xbf, 0x1a, 0x70, 0x6c, 0x47, 0x0c,
	0xc7, 0xde, 0xdc, 0x25, 0xd1, 0xdb, 0x35, 0x8e, 0x53, 0xf4, 0x00, 0xea, 0xc4, 0xd7, 0x94, 0x9e,
	0xd2, 0x6f, 0x3b, 0x75, 0xe2, 0xa3, 0x4b, 0xe8, 0x7a, 0x34, 0x8a, 0xb0, 0xc7, 0x08, 0x8d, 0x26,
	0xc4, 0xd7, 0xea, 0x02, 0xa9, 0x7f, 0x93, 0xb6, 0x8f, 0xce, 0xe1, 0x48, 0x48, 0x70, 0x7e, 0x20,
	0x78, 0x4b, 0xc4, 0xb6, 0x8f, 0x9e, 0x00, 0xac, 0xb8, 0xf0, 0x84, 0xa5, 0x4b, 0xac, 0x35, 0x04,
	0x6c, 0x8b, 0xcc, 0x75, 0xba, 0xc4, 0x48, 0x83, 0x56, 0x8c, 0x57, 0x6b, 0x9c, 0x30, 0xed, 0xb0,
	0xa7, 0xf4, 0x55, 0xa7, 0x08, 0x51, 0x0a, 0xcd, 0x25, 0x8e, 0x09, 0xf5, 0xb5, 0x26, 0xdf, 0x64,
	0xb9, 0x37, 0x99, 0x51, 0xfb, 0x91, 0x19, 0x4f, 0x67, 0x84, 0xcd, 0xd7, 0xd3, 0xa1, 0x47, 0xc3,
	0xdc, 0x4d, 0xfe, 0x18, 0x24, 0xfe, 0x62, 0xc4, 0xbf, 0x92, 0x0c, 0xed, 0x88, 0xed, 0x33, 0xe3,
	0x44, 0xee, 0x7f, 0x46, 0x43, 0xc2, 0x70, 0xb8, 0x64, 0xe9, 0xef, 0xcc, 0xe8, 0xa6, 0x6e, 0x18,
	0xbc, 0x30, 0x25, 0x31, 0xbf, 0x7d, 0x1d, 0x40, 0xde, 0x11, 0x3b, 0x62, 0x4e, 0xfe, 0x41, 0xf4,
	0x49, 0x81, 0x4e, 0xe0, 0x26, 0x6c, 0x32, 0xc7, 0x64, 0x36, 0x67, 0x5a, 0x4b, 0x14, 0xb0, 0xb8,
	0x77, 0x01, 0x67, 0x25, 0x91, 0x4a, 0x15, 0x48, 0x56, 0x51, 0xc2, 0x77, 0x4b, 0x01, 0xce, 0x5e,
	0x0b, 0x84, 0x0c, 0xe8, 0x78, 0x6e, 0x10, 0x4c, 0x5d, 0x6f, 0xc1, 0x1b, 0x7c, 0x24, 0x7a, 0x08,
	0x45, 0xca, 0xf6, 0xd1, 0x09, 0x1c, 0x30, 0x16, 0x68, 0xed, 0x9e, 0xd2, 0x6f, 0x38, 0xfc, 0x15,
	0x7d, 0x51, 0xa0, 0x2b, 0xd4, 0x71, 0x48, 0x92, 0x84, 0xd0, 0x48, 0x03, 0xe1, 0x81, 0xde, 0xdb,
	0xc3, 0xe3, 0x8a, 0x4c, 0xc5, 0xc5, 0x69, 0xc9, 0x45, 0xb1, 0xe0, 0xae, 0x0f, 0x95, 0xd3, 0x57,
	0x05, 0xfc, 0x58, 0x87, 0xf6, 0x4b, 0x97, 0xb9, 0x6f, 0x28, 0x89, 0xd8, 0x3f, 0xa3, 0xf6, 0x0e,
	0xba, 0x31, 0x0e, 0x29, 0xc3, 0x45, 0xdf, 0xf9, 0xa8, 0x35, 0xac, 0x2b, 0x5e, 0x45, 0x05, 0xfc,
	0xaf, 0x8a, 0xca, 0x02, 0xd3, 0x51, 0x65, 0x9c, 0x77, 0xf0, 0x1a, 0xd4, 0x80, 0x7a, 0x6e, 0x50,
	0x08, 0x8b, 0x19, 0xb5, 0xc6, 0xfb, 0xcc, 0x78, 0x54, 0xce, 0x57, 0x74, 0x1f, 0xe6, 0xee, 0x4a,
	0xdc, 0x74, 0x3a, 0x22, 0xcc, 0x55, 0x9f, 0xc3, 0xe1, 0xc6, 0x0d, 0xd6, 0x72, 0xaa, 0x55, 0xeb,
	0x72, 0x9f, 0x19, 0xc7, 0x22, 0x51, 0xd1, 0x51, 0xa5, 0x8e, 0x00, 0xa6, 0x23, 0x77, 0x58, 0xd6,
	0xcd, 0x56, 0x57, 0x6e, 0xb7, 0xba, 0xf2, 0x73, 0xab, 0x2b, 0x9f, 0x77, 0x7a, 0xed, 0x76, 0xa7,
	0xd7, 0xbe, 0xef, 0xf4, 0xda, 0xfb, 0x7e, 0xe9, 0x64, 0xf8, 0xe5, 0x1e, 0x24, 0xcc, 0xe5, 0xb7,
	0x4c, 0x04, 0xa3, 0x0f, 0xe2, 0x17, 0x20, 0xce, 0x67, 0xda, 0x14, 0x97, 0xf8, 0xea, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x69, 0x1e, 0x3b, 0x6f, 0x1b, 0x04, 0x00, 0x00,
}

func (m *InterchainQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterchainQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterchainQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LastEmission.Size()
		i -= size
		if _, err := m.LastEmission.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIcq(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.Ttl != 0 {
		i = encodeVarintIcq(dAtA, i, uint64(m.Ttl))
		i--
		dAtA[i] = 0x48
	}
	if len(m.CallbackId) > 0 {
		i -= len(m.CallbackId)
		copy(dAtA[i:], m.CallbackId)
		i = encodeVarintIcq(dAtA, i, uint64(len(m.CallbackId)))
		i--
		dAtA[i] = 0x42
	}
	{
		size := m.LastHeight.Size()
		i -= size
		if _, err := m.LastHeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIcq(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.Period.Size()
		i -= size
		if _, err := m.Period.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIcq(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Request) > 0 {
		i -= len(m.Request)
		copy(dAtA[i:], m.Request)
		i = encodeVarintIcq(dAtA, i, uint64(len(m.Request)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintIcq(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintIcq(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintIcq(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintIcq(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataPoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataPoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataPoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintIcq(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LocalHeight) > 0 {
		i -= len(m.LocalHeight)
		copy(dAtA[i:], m.LocalHeight)
		i = encodeVarintIcq(dAtA, i, uint64(len(m.LocalHeight)))
		i--
		dAtA[i] = 0x1a
	}
	if m.RemoteHeight != 0 {
		i = encodeVarintIcq(dAtA, i, uint64(m.RemoteHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintIcq(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIcq(dAtA []byte, offset int, v uint64) int {
	offset -= sovIcq(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InterchainQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovIcq(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovIcq(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovIcq(uint64(l))
	}
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovIcq(uint64(l))
	}
	l = len(m.Request)
	if l > 0 {
		n += 1 + l + sovIcq(uint64(l))
	}
	l = m.Period.Size()
	n += 1 + l + sovIcq(uint64(l))
	l = m.LastHeight.Size()
	n += 1 + l + sovIcq(uint64(l))
	l = len(m.CallbackId)
	if l > 0 {
		n += 1 + l + sovIcq(uint64(l))
	}
	if m.Ttl != 0 {
		n += 1 + sovIcq(uint64(m.Ttl))
	}
	l = m.LastEmission.Size()
	n += 1 + l + sovIcq(uint64(l))
	return n
}

func (m *DataPoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovIcq(uint64(l))
	}
	if m.RemoteHeight != 0 {
		n += 1 + sovIcq(uint64(m.RemoteHeight))
	}
	l = len(m.LocalHeight)
	if l > 0 {
		n += 1 + l + sovIcq(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovIcq(uint64(l))
	}
	return n
}

func sovIcq(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIcq(x uint64) (n int) {
	return sovIcq(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InterchainQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcq
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
			return fmt.Errorf("proto: InterchainQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterchainQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request[:0], dAtA[iNdEx:postIndex]...)
			if m.Request == nil {
				m.Request = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Period.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ttl |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEmission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastEmission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIcq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcq
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
func (m *DataPoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcq
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
			return fmt.Errorf("proto: DataPoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataPoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteHeight", wireType)
			}
			m.RemoteHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemoteHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocalHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIcq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcq
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
func skipIcq(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIcq
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
					return 0, ErrIntOverflowIcq
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
					return 0, ErrIntOverflowIcq
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
				return 0, ErrInvalidLengthIcq
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIcq
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIcq
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIcq        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIcq          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIcq = fmt.Errorf("proto: unexpected end of group")
)
