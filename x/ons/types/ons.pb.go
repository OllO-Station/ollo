// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/ons/v1/ons.proto

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

type NameStatus int32

const (
	NameStatusUnspecified     NameStatus = 0
	NameStatusClaimedInactive NameStatus = 1
	NameStatusClaimedActive   NameStatus = 2
	NameStatusUnclaimed       NameStatus = 3
	NameStatusListed          NameStatus = 4
	NameStatusDisabled        NameStatus = 5
)

var NameStatus_name = map[int32]string{
	0: "NAME_STATUS_UNSPECIFIED",
	1: "NAME_STATUS_CLAIMED_INACTIVE",
	2: "NAME_STATUS_CLAIMED_ACTIVE",
	3: "NAME_STATUS_UNCLAIMED",
	4: "NAME_STATUS_LISTED",
	5: "NAME_STATUS_DISABLED",
}

var NameStatus_value = map[string]int32{
	"NAME_STATUS_UNSPECIFIED":      0,
	"NAME_STATUS_CLAIMED_INACTIVE": 1,
	"NAME_STATUS_CLAIMED_ACTIVE":   2,
	"NAME_STATUS_UNCLAIMED":        3,
	"NAME_STATUS_LISTED":           4,
	"NAME_STATUS_DISABLED":         5,
}

func (NameStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4f5dab6f73e07b86, []int{0}
}

type Name struct {
	Name     string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	Value    string     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" yaml:"value"`
	Metadata string     `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty" yaml:"metadata"`
	Threads  []*Name    `protobuf:"bytes,4,rep,name=threads,proto3" json:"threads,omitempty" yaml:"threads"`
	Owner    string     `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	Status   NameStatus `protobuf:"varint,6,opt,name=status,proto3,enum=ollo.ons.v1.NameStatus" json:"status,omitempty" yaml:"status"`
	Expiry   time.Time  `protobuf:"bytes,7,opt,name=expiry,proto3,stdtime" json:"expiry" yaml:"expiry"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f5dab6f73e07b86, []int{0}
}
func (m *Name) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Name.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(m, src)
}
func (m *Name) XXX_Size() int {
	return m.Size()
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Name) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Name) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Name) GetThreads() []*Name {
	if m != nil {
		return m.Threads
	}
	return nil
}

func (m *Name) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Name) GetStatus() NameStatus {
	if m != nil {
		return m.Status
	}
	return NameStatusUnspecified
}

func (m *Name) GetExpiry() time.Time {
	if m != nil {
		return m.Expiry
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("ollo.ons.v1.NameStatus", NameStatus_name, NameStatus_value)
	proto.RegisterType((*Name)(nil), "ollo.ons.v1.Name")
}

func init() { proto.RegisterFile("ollo/ons/v1/ons.proto", fileDescriptor_4f5dab6f73e07b86) }

var fileDescriptor_4f5dab6f73e07b86 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x14, 0x85, 0xed, 0x26, 0x4d, 0xfb, 0x4f, 0x7e, 0x5a, 0x77, 0xda, 0x12, 0x77, 0x00, 0xdb, 0x32,
	0x12, 0x8a, 0x10, 0xd8, 0x34, 0x48, 0x2c, 0x60, 0x51, 0xc5, 0x89, 0x91, 0x2c, 0xa5, 0x11, 0x72,
	0x52, 0x16, 0x6c, 0xaa, 0x49, 0x32, 0x4d, 0x2d, 0xd9, 0x99, 0x28, 0x9e, 0xa4, 0xed, 0x8e, 0x25,
	0xca, 0xaa, 0x2f, 0x90, 0x15, 0x2c, 0x78, 0x94, 0x2e, 0xbb, 0x64, 0x15, 0xaa, 0xf6, 0x0d, 0xf2,
	0x04, 0xc8, 0x1e, 0x07, 0x1b, 0xca, 0xca, 0xbe, 0xf7, 0x9c, 0xef, 0xce, 0xd1, 0x1d, 0x1b, 0xec,
	0x52, 0xdf, 0xa7, 0x26, 0x1d, 0x84, 0xe6, 0x64, 0x3f, 0x7a, 0x18, 0xc3, 0x11, 0x65, 0x14, 0x16,
	0xa3, 0xb6, 0x11, 0xd5, 0x93, 0x7d, 0xb4, 0xd3, 0xa7, 0x7d, 0x1a, 0xf7, 0xcd, 0xe8, 0x8d, 0x5b,
	0x90, 0x9c, 0x25, 0x87, 0x78, 0x84, 0x83, 0x04, 0x46, 0xa5, 0xac, 0x72, 0x76, 0x4a, 0xbd, 0xa5,
	0xa0, 0xf6, 0x29, 0xed, 0xfb, 0xc4, 0x8c, 0xab, 0xce, 0xf8, 0xc4, 0x64, 0x5e, 0x40, 0x42, 0x86,
	0x83, 0x21, 0x37, 0xe8, 0x9f, 0x73, 0x20, 0xdf, 0xc4, 0x01, 0x81, 0x4f, 0x41, 0x7e, 0x80, 0x03,
	0x22, 0x8b, 0x9a, 0x58, 0xfe, 0xcf, 0xda, 0x5c, 0xcc, 0xd5, 0xe2, 0x05, 0x0e, 0xfc, 0xb7, 0x7a,
	0xd4, 0xd5, 0xdd, 0x58, 0x84, 0xcf, 0xc0, 0xea, 0x04, 0xfb, 0x63, 0x22, 0xaf, 0xc4, 0x2e, 0x69,
	0x31, 0x57, 0xff, 0xe7, 0xae, 0xb8, 0xad, 0xbb, 0x5c, 0x86, 0x26, 0x58, 0x0f, 0x08, 0xc3, 0x3d,
	0xcc, 0xb0, 0x9c, 0x8b, 0xad, 0xdb, 0x8b, 0xb9, 0xba, 0xc9, 0xad, 0x4b, 0x45, 0x77, 0x7f, 0x9b,
	0xe0, 0x01, 0x58, 0x63, 0xa7, 0x23, 0x82, 0x7b, 0xa1, 0x9c, 0xd7, 0x72, 0xe5, 0x62, 0x65, 0xcb,
	0xc8, 0xec, 0xc3, 0x88, 0x12, 0x5a, 0x70, 0x31, 0x57, 0x37, 0xf8, 0x88, 0xc4, 0xab, 0xbb, 0x4b,
	0x2a, 0x4a, 0x46, 0xcf, 0x06, 0x64, 0x24, 0xaf, 0xfe, 0x9d, 0x2c, 0x6e, 0xeb, 0x2e, 0x97, 0xa1,
	0x05, 0x0a, 0x21, 0xc3, 0x6c, 0x1c, 0xca, 0x05, 0x4d, 0x2c, 0x6f, 0x54, 0x4a, 0xf7, 0xce, 0x69,
	0xc5, 0xb2, 0xb5, 0xb5, 0x98, 0xab, 0x0f, 0xf8, 0x04, 0x0e, 0xe8, 0x6e, 0x42, 0xc2, 0x43, 0x50,
	0x20, 0xe7, 0x43, 0x6f, 0x74, 0x21, 0xaf, 0x69, 0x62, 0xb9, 0x58, 0x41, 0x06, 0xdf, 0xb2, 0xb1,
	0xdc, 0xb2, 0xd1, 0x5e, 0x6e, 0xd9, 0xda, 0xbb, 0x9a, 0xab, 0x42, 0x3a, 0x8a, 0x73, 0xfa, 0xe5,
	0x4f, 0x55, 0x74, 0x93, 0x21, 0xcf, 0x6f, 0x56, 0x00, 0x48, 0x0f, 0x86, 0x6f, 0x40, 0xa9, 0x59,
	0x3d, 0xb4, 0x8f, 0x5b, 0xed, 0x6a, 0xfb, 0xa8, 0x75, 0x7c, 0xd4, 0x6c, 0x7d, 0xb0, 0x6b, 0xce,
	0x7b, 0xc7, 0xae, 0x4b, 0x02, 0xda, 0x9b, 0xce, 0xb4, 0xdd, 0xd4, 0x7c, 0x34, 0x08, 0x87, 0xa4,
	0xeb, 0x9d, 0x78, 0xa4, 0x07, 0x0f, 0xc0, 0xe3, 0x2c, 0x57, 0x6b, 0x54, 0x9d, 0x43, 0xbb, 0x7e,
	0xec, 0x34, 0xab, 0xb5, 0xb6, 0xf3, 0xd1, 0x96, 0x44, 0xf4, 0x64, 0x3a, 0xd3, 0xf6, 0x52, 0xb8,
	0xe6, 0x63, 0x2f, 0x20, 0x3d, 0x67, 0x80, 0xbb, 0xcc, 0x9b, 0x10, 0xf8, 0x0e, 0xa0, 0x7f, 0x0d,
	0x48, 0xf0, 0x15, 0xf4, 0x68, 0x3a, 0xd3, 0x4a, 0xf7, 0xf0, 0x2a, 0x87, 0x2b, 0x60, 0xf7, 0xcf,
	0xd4, 0x09, 0x2e, 0xe5, 0x50, 0x69, 0x3a, 0xd3, 0xb6, 0xb3, 0x99, 0xbb, 0x9c, 0x84, 0x2f, 0x00,
	0xcc, 0x32, 0x0d, 0xa7, 0xd5, 0xb6, 0xeb, 0x52, 0x1e, 0xed, 0x4c, 0x67, 0x9a, 0x94, 0x02, 0x0d,
	0x2f, 0x64, 0xa4, 0x07, 0x5f, 0x81, 0x9d, 0xac, 0xbb, 0xee, 0xb4, 0xaa, 0x56, 0xc3, 0xae, 0x4b,
	0xab, 0xe8, 0xe1, 0x74, 0xa6, 0xc1, 0xd4, 0x5f, 0xf7, 0x42, 0xdc, 0xf1, 0x49, 0x0f, 0xad, 0x7f,
	0xf9, 0xaa, 0x08, 0xdf, 0xbf, 0x29, 0x82, 0x65, 0x5d, 0xdd, 0x2a, 0xe2, 0xf5, 0xad, 0x22, 0xde,
	0xdc, 0x2a, 0xe2, 0xe5, 0x9d, 0x22, 0x5c, 0xdf, 0x29, 0xc2, 0x8f, 0x3b, 0x45, 0xf8, 0x54, 0xee,
	0x7b, 0xec, 0x74, 0xdc, 0x31, 0xba, 0x34, 0x30, 0xa3, 0x2f, 0xe1, 0x65, 0x74, 0xc7, 0x1e, 0x1d,
	0xc4, 0x85, 0x79, 0x1e, 0xff, 0x53, 0xec, 0x62, 0x48, 0xc2, 0x4e, 0x21, 0xbe, 0xdd, 0xd7, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x09, 0xbe, 0x91, 0xc0, 0x03, 0x00, 0x00,
}

func (m *Name) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Name) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Name) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiry, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintOns(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	if m.Status != 0 {
		i = encodeVarintOns(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintOns(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Threads) > 0 {
		for iNdEx := len(m.Threads) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Threads[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOns(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintOns(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintOns(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintOns(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOns(dAtA []byte, offset int, v uint64) int {
	offset -= sovOns(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Name) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOns(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovOns(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovOns(uint64(l))
	}
	if len(m.Threads) > 0 {
		for _, e := range m.Threads {
			l = e.Size()
			n += 1 + l + sovOns(uint64(l))
		}
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovOns(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovOns(uint64(m.Status))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry)
	n += 1 + l + sovOns(uint64(l))
	return n
}

func sovOns(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOns(x uint64) (n int) {
	return sovOns(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Name) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOns
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
			return fmt.Errorf("proto: Name: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Name: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOns
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
				return ErrInvalidLengthOns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOns
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
				return ErrInvalidLengthOns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOns
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
				return ErrInvalidLengthOns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threads", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOns
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
				return ErrInvalidLengthOns
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Threads = append(m.Threads, &Name{})
			if err := m.Threads[len(m.Threads)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOns
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
				return ErrInvalidLengthOns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= NameStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOns
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
				return ErrInvalidLengthOns
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOns
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
func skipOns(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOns
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
					return 0, ErrIntOverflowOns
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
					return 0, ErrIntOverflowOns
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
				return 0, ErrInvalidLengthOns
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOns
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOns
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOns        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOns          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOns = fmt.Errorf("proto: unexpected end of group")
)
