// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/lend/v1/proprosal.proto

package types

import (
	fmt "fmt"
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

type ProprosalSetLendingFee struct {
}

func (m *ProprosalSetLendingFee) Reset()         { *m = ProprosalSetLendingFee{} }
func (m *ProprosalSetLendingFee) String() string { return proto.CompactTextString(m) }
func (*ProprosalSetLendingFee) ProtoMessage()    {}
func (*ProprosalSetLendingFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fce6bd84832659f, []int{0}
}
func (m *ProprosalSetLendingFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProprosalSetLendingFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProprosalSetLendingFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProprosalSetLendingFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProprosalSetLendingFee.Merge(m, src)
}
func (m *ProprosalSetLendingFee) XXX_Size() int {
	return m.Size()
}
func (m *ProprosalSetLendingFee) XXX_DiscardUnknown() {
	xxx_messageInfo_ProprosalSetLendingFee.DiscardUnknown(m)
}

var xxx_messageInfo_ProprosalSetLendingFee proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProprosalSetLendingFee)(nil), "ollo.lend.v1.ProprosalSetLendingFee")
}

func init() { proto.RegisterFile("ollo/lend/v1/proprosal.proto", fileDescriptor_5fce6bd84832659f) }

var fileDescriptor_5fce6bd84832659f = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0xcf, 0xc9, 0xc9,
	0xd7, 0xcf, 0x49, 0xcd, 0x4b, 0xd1, 0x2f, 0x33, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0x28, 0xca, 0x2f,
	0x4e, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0xc9, 0xea, 0x81, 0x64, 0xf5,
	0xca, 0x0c, 0x95, 0x24, 0xb8, 0xc4, 0x02, 0x60, 0x0a, 0x82, 0x53, 0x4b, 0x7c, 0x52, 0xf3, 0x52,
	0x32, 0xf3, 0xd2, 0xdd, 0x52, 0x53, 0x9d, 0x9c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e,
	0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58,
	0x8e, 0x21, 0x4a, 0x33, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x64,
	0x98, 0x6e, 0x71, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x1e, 0x98, 0xa3, 0x5f, 0x01, 0xb1, 0xb9, 0xa4,
	0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x6c, 0xa7, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xda,
	0x1a, 0xab, 0x93, 0x00, 0x00, 0x00,
}

func (m *ProprosalSetLendingFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProprosalSetLendingFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProprosalSetLendingFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintProprosal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProprosal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProprosalSetLendingFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovProprosal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProprosal(x uint64) (n int) {
	return sovProprosal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProprosalSetLendingFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProprosal
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
			return fmt.Errorf("proto: ProprosalSetLendingFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProprosalSetLendingFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProprosal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProprosal
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
func skipProprosal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProprosal
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
					return 0, ErrIntOverflowProprosal
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
					return 0, ErrIntOverflowProprosal
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
				return 0, ErrInvalidLengthProprosal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProprosal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProprosal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProprosal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProprosal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProprosal = fmt.Errorf("proto: unexpected end of group")
)
