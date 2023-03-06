// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/liquidity/v1/genesis.proto

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

// GenesisState defines the liquidity module's genesis state.
type GenesisState struct {
	Params                       Params              `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	LastPairId                   uint64              `protobuf:"varint,2,opt,name=last_pair_id,json=lastPairId,proto3" json:"last_pair_id,omitempty"`
	LastPoolId                   uint64              `protobuf:"varint,3,opt,name=last_pool_id,json=lastPoolId,proto3" json:"last_pool_id,omitempty"`
	Pairs                        []Pair              `protobuf:"bytes,4,rep,name=pairs,proto3" json:"pairs"`
	Pools                        []Pool              `protobuf:"bytes,5,rep,name=pools,proto3" json:"pools"`
	DepositRequests              []DepositRequest    `protobuf:"bytes,6,rep,name=deposit_requests,json=depositRequests,proto3" json:"deposit_requests"`
	WithdrawRequests             []WithdrawRequest   `protobuf:"bytes,7,rep,name=withdraw_requests,json=withdrawRequests,proto3" json:"withdraw_requests"`
	Orders                       []Order             `protobuf:"bytes,8,rep,name=orders,proto3" json:"orders"`
	NumMarketMakingOrdersRecords []NumMMOrdersRecord `protobuf:"bytes,9,rep,name=num_market_making_orders_records,json=numMarketMakingOrdersRecords,proto3" json:"num_market_making_orders_records"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebf7278aa0db3fd2, []int{0}
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

// NumMMOrdersRecord holds information about how many MM orders an orderer
// ordered per pair.
type NumMMOrdersRecord struct {
	Orderer               string `protobuf:"bytes,1,opt,name=orderer,proto3" json:"orderer,omitempty"`
	PairId                uint64 `protobuf:"varint,2,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty"`
	NumMarketMakingOrders uint32 `protobuf:"varint,3,opt,name=num_market_making_orders,json=numMarketMakingOrders,proto3" json:"num_market_making_orders,omitempty"`
}

func (m *NumMMOrdersRecord) Reset()         { *m = NumMMOrdersRecord{} }
func (m *NumMMOrdersRecord) String() string { return proto.CompactTextString(m) }
func (*NumMMOrdersRecord) ProtoMessage()    {}
func (*NumMMOrdersRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebf7278aa0db3fd2, []int{1}
}
func (m *NumMMOrdersRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NumMMOrdersRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NumMMOrdersRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NumMMOrdersRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumMMOrdersRecord.Merge(m, src)
}
func (m *NumMMOrdersRecord) XXX_Size() int {
	return m.Size()
}
func (m *NumMMOrdersRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_NumMMOrdersRecord.DiscardUnknown(m)
}

var xxx_messageInfo_NumMMOrdersRecord proto.InternalMessageInfo

func (m *NumMMOrdersRecord) GetOrderer() string {
	if m != nil {
		return m.Orderer
	}
	return ""
}

func (m *NumMMOrdersRecord) GetPairId() uint64 {
	if m != nil {
		return m.PairId
	}
	return 0
}

func (m *NumMMOrdersRecord) GetNumMarketMakingOrders() uint32 {
	if m != nil {
		return m.NumMarketMakingOrders
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ollo.liquidity.v1.GenesisState")
	proto.RegisterType((*NumMMOrdersRecord)(nil), "ollo.liquidity.v1.NumMMOrdersRecord")
}

func init() { proto.RegisterFile("ollo/liquidity/v1/genesis.proto", fileDescriptor_ebf7278aa0db3fd2) }

var fileDescriptor_ebf7278aa0db3fd2 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x41, 0x4f, 0xd4, 0x40,
	0x14, 0xc7, 0x5b, 0x59, 0x76, 0x61, 0xc0, 0xc8, 0x4e, 0x34, 0x8c, 0x1b, 0xec, 0x96, 0x8d, 0x87,
	0xbd, 0xd8, 0x06, 0x48, 0x24, 0xf1, 0x48, 0x4c, 0x94, 0xc3, 0x2a, 0xa9, 0x31, 0x26, 0x5e, 0x9a,
	0x42, 0x27, 0x65, 0x42, 0xdb, 0x57, 0x66, 0xa6, 0xac, 0x9c, 0xbc, 0x7a, 0xf4, 0x23, 0xf0, 0x71,
	0x38, 0x72, 0x32, 0x9e, 0x8c, 0xd9, 0xbd, 0xf8, 0x31, 0xcc, 0xcc, 0x14, 0x76, 0x37, 0x2d, 0xdc,
	0xda, 0x79, 0xbf, 0xff, 0xef, 0xcd, 0x4b, 0x5f, 0x51, 0x1f, 0xd2, 0x14, 0xfc, 0x94, 0x9d, 0x97,
	0x2c, 0x66, 0xf2, 0xd2, 0xbf, 0xd8, 0xf1, 0x13, 0x9a, 0x53, 0xc1, 0x84, 0x57, 0x70, 0x90, 0x80,
	0xbb, 0x0a, 0xf0, 0xee, 0x00, 0xef, 0x62, 0xa7, 0xf7, 0x34, 0x81, 0x04, 0x74, 0xd5, 0x57, 0x4f,
	0x06, 0xec, 0x6d, 0xd5, 0x4d, 0x45, 0xc4, 0xf8, 0x03, 0x55, 0x80, 0xb4, 0xaa, 0x6e, 0xd7, 0xab,
	0xb3, 0x8e, 0x06, 0x71, 0x9a, 0xf4, 0x3c, 0xca, 0xaa, 0x7b, 0xf6, 0x5e, 0xd4, 0xeb, 0xc0, 0x63,
	0x5a, 0xf5, 0x1f, 0xfc, 0x6a, 0xa1, 0xf5, 0x77, 0x66, 0xb0, 0x4f, 0x32, 0x92, 0x14, 0xef, 0xa3,
	0xb6, 0xc9, 0x13, 0xdb, 0xb5, 0x87, 0x6b, 0xbb, 0xcf, 0xbd, 0xda, 0xa0, 0xde, 0x91, 0x06, 0x0e,
	0x5a, 0xd7, 0x7f, 0xfa, 0x56, 0x50, 0xe1, 0xd8, 0x45, 0xeb, 0x69, 0x24, 0x64, 0xa8, 0x86, 0x0b,
	0x59, 0x4c, 0x1e, 0xb9, 0xf6, 0xb0, 0x15, 0x20, 0x75, 0x76, 0x14, 0x31, 0x7e, 0x18, 0xcf, 0x08,
	0x80, 0x54, 0x11, 0x4b, 0x73, 0x04, 0x40, 0x7a, 0x18, 0xe3, 0x3d, 0xb4, 0xac, 0xe2, 0x82, 0xb4,
	0xdc, 0xa5, 0xe1, 0xda, 0xee, 0x66, 0x63, 0x6f, 0xc6, 0xab, 0xce, 0x86, 0xd5, 0x21, 0x80, 0x54,
	0x90, 0xe5, 0xfb, 0x43, 0x00, 0xe9, 0x5d, 0x48, 0xb1, 0x38, 0x40, 0x1b, 0x31, 0x2d, 0x40, 0x30,
	0x19, 0x72, 0x7a, 0x5e, 0x52, 0x21, 0x05, 0x69, 0xeb, 0xfc, 0x76, 0x43, 0xfe, 0xad, 0x41, 0x03,
	0x43, 0x56, 0xa6, 0x27, 0xf1, 0xc2, 0xa9, 0xc0, 0x9f, 0x51, 0x77, 0xcc, 0xe4, 0x69, 0xcc, 0xa3,
	0xf1, 0x4c, 0xda, 0xd1, 0xd2, 0x41, 0x83, 0xf4, 0x4b, 0xc5, 0x2e, 0x5a, 0x37, 0xc6, 0x8b, 0xc7,
	0x02, 0xbf, 0x46, 0x6d, 0xfd, 0xc5, 0x04, 0x59, 0xd1, 0x2e, 0xd2, 0xe0, 0xfa, 0xa8, 0x80, 0xdb,
	0x0f, 0x62, 0x68, 0xcc, 0x91, 0x9b, 0x97, 0x59, 0x98, 0x45, 0xfc, 0x8c, 0xca, 0x30, 0x8b, 0xce,
	0x58, 0x9e, 0x84, 0xa6, 0x16, 0x72, 0x7a, 0x02, 0x3c, 0x16, 0x64, 0x55, 0x1b, 0x5f, 0x36, 0x18,
	0x3f, 0x94, 0xd9, 0x68, 0xa4, 0xb5, 0x22, 0xd0, 0x70, 0x65, 0xdf, 0xca, 0xcb, 0x6c, 0xa4, 0x95,
	0x23, 0x6d, 0x9c, 0x47, 0xc4, 0x9b, 0x95, 0x1f, 0x57, 0x7d, 0xeb, 0xdf, 0x55, 0xdf, 0x1a, 0x7c,
	0x47, 0xdd, 0x9a, 0x02, 0x13, 0xd4, 0xd1, 0x17, 0xa0, 0x5c, 0x6f, 0xd7, 0x6a, 0x70, 0xfb, 0x8a,
	0x37, 0x51, 0x67, 0x71, 0x71, 0xda, 0x85, 0x59, 0x9a, 0x7d, 0x44, 0xee, 0x9b, 0x42, 0x2f, 0xd0,
	0xe3, 0xe0, 0x59, 0xe3, 0x8d, 0x0e, 0xde, 0x5f, 0x4f, 0x1c, 0xfb, 0x66, 0xe2, 0xd8, 0x7f, 0x27,
	0x8e, 0xfd, 0x73, 0xea, 0x58, 0x37, 0x53, 0xc7, 0xfa, 0x3d, 0x75, 0xac, 0xaf, 0x5e, 0xc2, 0xe4,
	0x69, 0x79, 0xec, 0x9d, 0x40, 0xe6, 0xab, 0xc1, 0x5f, 0x09, 0x19, 0x49, 0x06, 0xb9, 0x7e, 0xf1,
	0xbf, 0xcd, 0xfd, 0x2c, 0xf2, 0xb2, 0xa0, 0xe2, 0xb8, 0xad, 0x7f, 0x95, 0xbd, 0xff, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0xe6, 0x3a, 0x28, 0x14, 0x04, 0x00, 0x00,
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
	if len(m.NumMarketMakingOrdersRecords) > 0 {
		for iNdEx := len(m.NumMarketMakingOrdersRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NumMarketMakingOrdersRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Orders) > 0 {
		for iNdEx := len(m.Orders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Orders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.WithdrawRequests) > 0 {
		for iNdEx := len(m.WithdrawRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DepositRequests) > 0 {
		for iNdEx := len(m.DepositRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DepositRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Pairs) > 0 {
		for iNdEx := len(m.Pairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.LastPoolId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastPoolId))
		i--
		dAtA[i] = 0x18
	}
	if m.LastPairId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastPairId))
		i--
		dAtA[i] = 0x10
	}
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

func (m *NumMMOrdersRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NumMMOrdersRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NumMMOrdersRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumMarketMakingOrders != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NumMarketMakingOrders))
		i--
		dAtA[i] = 0x18
	}
	if m.PairId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PairId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Orderer) > 0 {
		i -= len(m.Orderer)
		copy(dAtA[i:], m.Orderer)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Orderer)))
		i--
		dAtA[i] = 0xa
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
	if m.LastPairId != 0 {
		n += 1 + sovGenesis(uint64(m.LastPairId))
	}
	if m.LastPoolId != 0 {
		n += 1 + sovGenesis(uint64(m.LastPoolId))
	}
	if len(m.Pairs) > 0 {
		for _, e := range m.Pairs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DepositRequests) > 0 {
		for _, e := range m.DepositRequests {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.WithdrawRequests) > 0 {
		for _, e := range m.WithdrawRequests {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Orders) > 0 {
		for _, e := range m.Orders {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NumMarketMakingOrdersRecords) > 0 {
		for _, e := range m.NumMarketMakingOrdersRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *NumMMOrdersRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Orderer)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.PairId != 0 {
		n += 1 + sovGenesis(uint64(m.PairId))
	}
	if m.NumMarketMakingOrders != 0 {
		n += 1 + sovGenesis(uint64(m.NumMarketMakingOrders))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastPairId", wireType)
			}
			m.LastPairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastPairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastPoolId", wireType)
			}
			m.LastPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pairs", wireType)
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
			m.Pairs = append(m.Pairs, Pair{})
			if err := m.Pairs[len(m.Pairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
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
			m.Pools = append(m.Pools, Pool{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositRequests", wireType)
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
			m.DepositRequests = append(m.DepositRequests, DepositRequest{})
			if err := m.DepositRequests[len(m.DepositRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawRequests", wireType)
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
			m.WithdrawRequests = append(m.WithdrawRequests, WithdrawRequest{})
			if err := m.WithdrawRequests[len(m.WithdrawRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orders", wireType)
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
			m.Orders = append(m.Orders, Order{})
			if err := m.Orders[len(m.Orders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumMarketMakingOrdersRecords", wireType)
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
			m.NumMarketMakingOrdersRecords = append(m.NumMarketMakingOrdersRecords, NumMMOrdersRecord{})
			if err := m.NumMarketMakingOrdersRecords[len(m.NumMarketMakingOrdersRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *NumMMOrdersRecord) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NumMMOrdersRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NumMMOrdersRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orderer", wireType)
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
			m.Orderer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			m.PairId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumMarketMakingOrders", wireType)
			}
			m.NumMarketMakingOrders = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumMarketMakingOrders |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
