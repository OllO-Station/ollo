// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/loan/v1/borrow.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
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

type Borrow struct {
	Borrower  string                                  `protobuf:"bytes,1,opt,name=borrower,proto3" json:"borrower,omitempty" yaml:"lender"`
	CreatedAt time.Time                               `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	AmountOut github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=amount_out,json=amountOut,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount_out" yaml:"amount_in"`
	asset_id  uint64                                  `protobuf:"varint,4,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
}

func (m *Borrow) Reset()         { *m = Borrow{} }
func (m *Borrow) String() string { return proto.CompactTextString(m) }
func (*Borrow) ProtoMessage()    {}
func (*Borrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_c523195fe0ca014c, []int{0}
}
func (m *Borrow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Borrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Borrow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Borrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Borrow.Merge(m, src)
}
func (m *Borrow) XXX_Size() int {
	return m.Size()
}
func (m *Borrow) XXX_DiscardUnknown() {
	xxx_messageInfo_Borrow.DiscardUnknown(m)
}

var xxx_messageInfo_Borrow proto.InternalMessageInfo

func (m *Borrow) GetBorrower() string {
	if m != nil {
		return m.Borrower
	}
	return ""
}

func (m *Borrow) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Borrow) GetAmountOut() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.AmountOut
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *Borrow) Getasset_id() uint64 {
	if m != nil {
		return m.asset_id
	}
	return 0
}

func init() {
	proto.RegisterType((*Borrow)(nil), "ollo.loan.v1.Borrow")
}

func init() { proto.RegisterFile("ollo/loan/v1/borrow.proto", fileDescriptor_c523195fe0ca014c) }

var fileDescriptor_c523195fe0ca014c = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x52, 0xb1, 0x8e, 0xd4, 0x30,
	0x10, 0x8d, 0x8f, 0xd3, 0xb1, 0x6b, 0x40, 0x40, 0x74, 0x45, 0x76, 0x8b, 0x78, 0x95, 0x86, 0x45,
	0xe8, 0x6c, 0x05, 0xba, 0xeb, 0x2e, 0x2b, 0x0a, 0x2a, 0xa4, 0x40, 0x45, 0x13, 0x39, 0x89, 0x09,
	0x11, 0x49, 0x66, 0x65, 0x3b, 0x0b, 0x57, 0xf0, 0x0f, 0xf7, 0x31, 0x7c, 0xc4, 0x75, 0x9c, 0xa8,
	0xa8, 0x02, 0xca, 0xfe, 0xc1, 0x96, 0x54, 0xa7, 0xd8, 0xbe, 0xab, 0xec, 0x99, 0x37, 0xef, 0xcd,
	0xbc, 0xd1, 0xe0, 0x05, 0x34, 0x0d, 0xb0, 0x06, 0x78, 0xc7, 0x76, 0x31, 0xcb, 0x41, 0x4a, 0xf8,
	0x46, 0xb7, 0x12, 0x34, 0xf8, 0x8f, 0x27, 0x88, 0x4e, 0x10, 0xdd, 0xc5, 0xcb, 0xd3, 0x0a, 0x2a,
	0x30, 0x00, 0x9b, 0x7e, 0xb6, 0x66, 0x49, 0x2a, 0x80, 0xaa, 0x11, 0xcc, 0x44, 0x79, 0xff, 0x99,
	0xe9, 0xba, 0x15, 0x4a, 0xf3, 0x76, 0xeb, 0x0a, 0x16, 0x05, 0xa8, 0x16, 0x54, 0x66, 0x99, 0x36,
	0x70, 0x50, 0x68, 0x23, 0x96, 0x73, 0x25, 0xd8, 0x2e, 0xce, 0x85, 0xe6, 0x31, 0x2b, 0xa0, 0xee,
	0x2c, 0x1e, 0xfd, 0x3a, 0xc2, 0x27, 0x89, 0x19, 0xc8, 0x7f, 0x8b, 0x67, 0x76, 0x34, 0x21, 0x03,
	0xb4, 0x42, 0xeb, 0x79, 0xf2, 0xf2, 0x30, 0x90, 0x27, 0x97, 0xbc, 0x6d, 0xce, 0xa3, 0x46, 0x74,
	0xa5, 0x90, 0xd1, 0xef, 0x9f, 0x67, 0xa7, 0x4e, 0xff, 0xa2, 0x2c, 0xa5, 0x50, 0xea, 0x83, 0x96,
	0x75, 0x57, 0xa5, 0xf7, 0x54, 0x7f, 0x83, 0x71, 0x21, 0x05, 0xd7, 0xa2, 0xcc, 0xb8, 0x0e, 0x8e,
	0x56, 0x68, 0xfd, 0xe8, 0xf5, 0x92, 0x5a, 0x0b, 0xf4, 0xce, 0x02, 0xfd, 0x78, 0x67, 0x21, 0x99,
	0x5d, 0x0f, 0xc4, 0xbb, 0xfa, 0x4b, 0x50, 0x3a, 0x77, 0xbc, 0x0b, 0xed, 0xff, 0xc0, 0x98, 0xb7,
	0xd0, 0x77, 0x3a, 0x83, 0x5e, 0x07, 0x0f, 0x8c, 0xc8, 0x82, 0xba, 0xce, 0x93, 0x17, 0xea, 0xbc,
	0xd0, 0x0d, 0xd4, 0x5d, 0xb2, 0x99, 0x34, 0x0e, 0x03, 0x79, 0x66, 0x87, 0x75, 0xd4, 0xba, 0x8b,
	0xfe, 0x0f, 0xe4, 0x45, 0x55, 0xeb, 0x2f, 0x7d, 0x4e, 0x0b, 0x68, 0xdd, 0x6a, 0xdc, 0x73, 0xa6,
	0xca, 0xaf, 0x4c, 0x5f, 0x6e, 0x85, 0x32, 0x22, 0xe9, 0xdc, 0xd2, 0xde, 0xf7, 0xda, 0x3f, 0xc7,
	0x33, 0xae, 0x94, 0xd0, 0x59, 0x5d, 0x06, 0xc7, 0x2b, 0xb4, 0x3e, 0x4e, 0xc8, 0x38, 0x90, 0xfb,
	0xdc, 0x61, 0x20, 0x4f, 0x5d, 0x27, 0x97, 0x89, 0xd2, 0x87, 0xe6, 0xfb, 0xae, 0x4c, 0x5e, 0x5d,
	0x8f, 0x21, 0xba, 0x19, 0x43, 0xf4, 0x6f, 0x0c, 0xd1, 0xd5, 0x3e, 0xf4, 0x6e, 0xf6, 0xa1, 0xf7,
	0x67, 0x1f, 0x7a, 0x9f, 0x9e, 0x9b, 0x33, 0xf8, 0x6e, 0x0f, 0xc1, 0x34, 0xce, 0x4f, 0xcc, 0x42,
	0xde, 0xdc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x0e, 0xf9, 0x52, 0x22, 0x02, 0x00, 0x00,
}

func (m *Borrow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Borrow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Borrow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.asset_id != 0 {
		i = encodeVarintBorrow(dAtA, i, uint64(m.asset_id))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.AmountOut.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBorrow(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintBorrow(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.Borrower) > 0 {
		i -= len(m.Borrower)
		copy(dAtA[i:], m.Borrower)
		i = encodeVarintBorrow(dAtA, i, uint64(len(m.Borrower)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBorrow(dAtA []byte, offset int, v uint64) int {
	offset -= sovBorrow(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Borrow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Borrower)
	if l > 0 {
		n += 1 + l + sovBorrow(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovBorrow(uint64(l))
	l = m.AmountOut.Size()
	n += 1 + l + sovBorrow(uint64(l))
	if m.asset_id != 0 {
		n += 1 + sovBorrow(uint64(m.asset_id))
	}
	return n
}

func sovBorrow(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBorrow(x uint64) (n int) {
	return sovBorrow(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Borrow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBorrow
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
			return fmt.Errorf("proto: Borrow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Borrow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Borrower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrow
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
				return ErrInvalidLengthBorrow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBorrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Borrower = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrow
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
				return ErrInvalidLengthBorrow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBorrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrow
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
				return ErrInvalidLengthBorrow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBorrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field asset_id", wireType)
			}
			m.asset_id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBorrow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.asset_id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBorrow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBorrow
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
func skipBorrow(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBorrow
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
					return 0, ErrIntOverflowBorrow
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
					return 0, ErrIntOverflowBorrow
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
				return 0, ErrInvalidLengthBorrow
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBorrow
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBorrow
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBorrow        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBorrow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBorrow = fmt.Errorf("proto: unexpected end of group")
)
