// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/claim/v1/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
type Params struct {
	DecayInformation DecayInformation `protobuf:"bytes,1,opt,name=decayInformation,proto3" json:"decayInformation"`
	AirdropStart     time.Time        `protobuf:"bytes,2,opt,name=airdropStart,proto3,stdtime" json:"airdropStart"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff6b687bb67a6686, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDecayInformation() DecayInformation {
	if m != nil {
		return m.DecayInformation
	}
	return DecayInformation{}
}

func (m *Params) GetAirdropStart() time.Time {
	if m != nil {
		return m.AirdropStart
	}
	return time.Time{}
}

// DecayInformation defines the information about decay for the airdrop
// when claimable airdrop amount starts to decrease and when it ends
type DecayInformation struct {
	Enabled    bool      `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DecayStart time.Time `protobuf:"bytes,2,opt,name=decayStart,proto3,stdtime" json:"decayStart"`
	DecayEnd   time.Time `protobuf:"bytes,3,opt,name=decayEnd,proto3,stdtime" json:"decayEnd"`
}

func (m *DecayInformation) Reset()         { *m = DecayInformation{} }
func (m *DecayInformation) String() string { return proto.CompactTextString(m) }
func (*DecayInformation) ProtoMessage()    {}
func (*DecayInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff6b687bb67a6686, []int{1}
}
func (m *DecayInformation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DecayInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DecayInformation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DecayInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecayInformation.Merge(m, src)
}
func (m *DecayInformation) XXX_Size() int {
	return m.Size()
}
func (m *DecayInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_DecayInformation.DiscardUnknown(m)
}

var xxx_messageInfo_DecayInformation proto.InternalMessageInfo

func (m *DecayInformation) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *DecayInformation) GetDecayStart() time.Time {
	if m != nil {
		return m.DecayStart
	}
	return time.Time{}
}

func (m *DecayInformation) GetDecayEnd() time.Time {
	if m != nil {
		return m.DecayEnd
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Params)(nil), "ollo.claim.v1.Params")
	proto.RegisterType((*DecayInformation)(nil), "ollo.claim.v1.DecayInformation")
}

func init() { proto.RegisterFile("ollo/claim/v1/params.proto", fileDescriptor_ff6b687bb67a6686) }

var fileDescriptor_ff6b687bb67a6686 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xcf, 0xc9, 0xc9,
	0xd7, 0x4f, 0xce, 0x49, 0xcc, 0xcc, 0xd5, 0x2f, 0x33, 0xd4, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x05, 0xc9, 0xe9, 0x81, 0xe5, 0xf4, 0xca, 0x0c,
	0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x32, 0xfa, 0x20, 0x16, 0x44, 0x91, 0x94, 0x7c, 0x7a,
	0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x5f, 0x92, 0x99, 0x9b, 0x5a,
	0x5c, 0x92, 0x98, 0x5b, 0x00, 0x51, 0xa0, 0xb4, 0x9a, 0x91, 0x8b, 0x2d, 0x00, 0x6c, 0xac, 0x50,
	0x20, 0x97, 0x40, 0x4a, 0x6a, 0x72, 0x62, 0xa5, 0x67, 0x5e, 0x5a, 0x7e, 0x51, 0x6e, 0x62, 0x49,
	0x66, 0x7e, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xbc, 0x1e, 0x8a, 0x5d, 0x7a, 0x2e,
	0x68, 0xca, 0x9c, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0xc2, 0xd0, 0x2e, 0xe4, 0xc1, 0xc5, 0x93,
	0x98, 0x59, 0x94, 0x52, 0x94, 0x5f, 0x10, 0x5c, 0x92, 0x58, 0x54, 0x22, 0xc1, 0x04, 0x36, 0x4e,
	0x4a, 0x0f, 0xe2, 0x2a, 0x3d, 0x98, 0xab, 0xf4, 0x42, 0x60, 0xae, 0x72, 0xe2, 0x00, 0x99, 0x34,
	0xe1, 0xbe, 0x3c, 0x63, 0x10, 0x8a, 0x4e, 0x2b, 0x96, 0x19, 0x0b, 0xe4, 0x19, 0x94, 0xb6, 0x30,
	0x72, 0x09, 0xa0, 0x5b, 0x2e, 0x24, 0xc1, 0xc5, 0x9e, 0x9a, 0x97, 0x98, 0x94, 0x93, 0x9a, 0x02,
	0x76, 0x2e, 0x47, 0x10, 0x8c, 0x2b, 0xe4, 0xc2, 0xc5, 0x05, 0x76, 0x12, 0xe9, 0x96, 0x23, 0xe9,
	0x13, 0x72, 0xe0, 0xe2, 0x00, 0xf3, 0x5c, 0xf3, 0x52, 0x24, 0x98, 0x49, 0x30, 0x03, 0xae, 0xcb,
	0x49, 0xe7, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0,
	0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x84, 0xc0, 0x11, 0x5c,
	0x01, 0x8d, 0xe2, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xa9, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd4, 0xb8, 0x76, 0x3f, 0xfd, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStart, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStart):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.DecayInformation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DecayInformation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DecayInformation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DecayInformation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.DecayEnd, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.DecayEnd):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.DecayStart, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.DecayStart):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintParams(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x12
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DecayInformation.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStart)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *DecayInformation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.DecayStart)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.DecayEnd)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DecayInformation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DecayInformation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStart, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *DecayInformation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: DecayInformation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DecayInformation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DecayStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.DecayStart, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DecayEnd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.DecayEnd, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
