// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ollo/marketplace/v1/auction.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type AuctionStatus int32

const (
	AUCTION_STATUS_UNSPECIFIED AuctionStatus = 0
	AUCTION_STATUS_INACTIVE    AuctionStatus = 1
	AUCTION_STATUS_ACTIVE      AuctionStatus = 2
)

var AuctionStatus_name = map[int32]string{
	0: "AUCTION_STATUS_UNSPECIFIED",
	1: "AUCTION_STATUS_INACTIVE",
	2: "AUCTION_STATUS_ACTIVE",
}

var AuctionStatus_value = map[string]int32{
	"AUCTION_STATUS_UNSPECIFIED": 0,
	"AUCTION_STATUS_INACTIVE":    1,
	"AUCTION_STATUS_ACTIVE":      2,
}

func (x AuctionStatus) String() string {
	return proto.EnumName(AuctionStatus_name, int32(x))
}

func (AuctionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_118374a82457c143, []int{0}
}

type AuctionListing struct {
	Id                  uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NftId               string                                 `protobuf:"bytes,2,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty" yaml:"nft_id"`
	DenomId             string                                 `protobuf:"bytes,3,opt,name=denom_id,json=denomId,proto3" json:"denom_id,omitempty" yaml:"denom_id"`
	StartPrice          types.Coin                             `protobuf:"bytes,4,opt,name=start_price,json=startPrice,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"start_price" yaml:"start_price"`
	StartTime           time.Time                              `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	EndTime             *time.Time                             `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time,omitempty" yaml:"end_time"`
	Owner               string                                 `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	IncrementPercentage github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=increment_percentage,json=incrementPercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"increment_percentage" yaml:"increment_percentage"`
	WhitelistAccounts   []string                               `protobuf:"bytes,9,rep,name=whitelist_accounts,json=whitelistAccounts,proto3" json:"whitelist_accounts,omitempty" yaml:"whitelist_accounts"`
	SplitShares         []WeightedAddress                      `protobuf:"bytes,10,rep,name=split_shares,json=splitShares,proto3" json:"split_shares" yaml:"split_shares"`
}

func (m *AuctionListing) Reset()         { *m = AuctionListing{} }
func (m *AuctionListing) String() string { return proto.CompactTextString(m) }
func (*AuctionListing) ProtoMessage()    {}
func (*AuctionListing) Descriptor() ([]byte, []int) {
	return fileDescriptor_118374a82457c143, []int{0}
}
func (m *AuctionListing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionListing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuctionListing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuctionListing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionListing.Merge(m, src)
}
func (m *AuctionListing) XXX_Size() int {
	return m.Size()
}
func (m *AuctionListing) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionListing.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionListing proto.InternalMessageInfo

type Bid struct {
	AuctionId uint64     `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty" yaml:"auction_id"`
	Bidder    string     `protobuf:"bytes,2,opt,name=bidder,proto3" json:"bidder,omitempty"`
	Amount    types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	Time      time.Time  `protobuf:"bytes,4,opt,name=time,proto3,stdtime" json:"time"`
}

func (m *Bid) Reset()         { *m = Bid{} }
func (m *Bid) String() string { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()    {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_118374a82457c143, []int{1}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return m.Size()
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("ollo.marketplace.v1.AuctionStatus", AuctionStatus_name, AuctionStatus_value)
	proto.RegisterType((*AuctionListing)(nil), "ollo.marketplace.v1.AuctionListing")
	proto.RegisterType((*Bid)(nil), "ollo.marketplace.v1.Bid")
}

func init() { proto.RegisterFile("ollo/marketplace/v1/auction.proto", fileDescriptor_118374a82457c143) }

var fileDescriptor_118374a82457c143 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x6f, 0xfb, 0x44,
	0x10, 0x8d, 0xf3, 0xaf, 0xc9, 0x86, 0x5f, 0x69, 0x37, 0x2d, 0x75, 0x53, 0x61, 0xa7, 0x16, 0x7f,
	0x22, 0x24, 0x6c, 0xb5, 0x20, 0x81, 0x7a, 0x8b, 0xd3, 0x20, 0x59, 0x82, 0x10, 0x39, 0x29, 0x20,
	0x2e, 0x96, 0xe3, 0xdd, 0x3a, 0xab, 0xc6, 0xde, 0xc8, 0xbb, 0x69, 0xe9, 0x0d, 0xbe, 0x41, 0x3f,
	0x02, 0x67, 0x3e, 0x49, 0x0f, 0x1c, 0x7a, 0x44, 0x1c, 0x52, 0xda, 0x5e, 0x38, 0xe7, 0xc6, 0x0d,
	0x79, 0xbd, 0x69, 0x43, 0xa9, 0x28, 0x27, 0xef, 0xcc, 0xbc, 0xf7, 0x3c, 0x7e, 0x3b, 0x63, 0xb0,
	0x4f, 0x27, 0x13, 0x6a, 0x45, 0x7e, 0x72, 0x86, 0xf9, 0x74, 0xe2, 0x07, 0xd8, 0x3a, 0x3f, 0xb0,
	0xfc, 0x59, 0xc0, 0x09, 0x8d, 0xcd, 0x69, 0x42, 0x39, 0x85, 0xf5, 0x14, 0x62, 0xae, 0x40, 0xcc,
	0xf3, 0x83, 0x86, 0x16, 0x50, 0x16, 0x51, 0x66, 0x8d, 0x7c, 0x96, 0x52, 0x46, 0x98, 0xfb, 0x07,
	0x56, 0x40, 0x89, 0x24, 0x35, 0xb6, 0x42, 0x1a, 0x52, 0x71, 0xb4, 0xd2, 0x93, 0xcc, 0xea, 0x21,
	0xa5, 0xe1, 0x04, 0x5b, 0x22, 0x1a, 0xcd, 0x4e, 0x2d, 0x4e, 0x22, 0xcc, 0xb8, 0x1f, 0x4d, 0x25,
	0xe0, 0xfd, 0x97, 0xda, 0x59, 0x7d, 0xb5, 0x80, 0x19, 0x7f, 0x95, 0xc0, 0x7a, 0x3b, 0x6b, 0xf2,
	0x4b, 0xc2, 0x38, 0x89, 0x43, 0xb8, 0x0e, 0xf2, 0x04, 0xa9, 0x4a, 0x53, 0x69, 0x15, 0xdd, 0x3c,
	0x41, 0xb0, 0x05, 0xca, 0xf1, 0x29, 0xf7, 0x08, 0x52, 0xf3, 0x4d, 0xa5, 0x55, 0xb5, 0x37, 0x17,
	0x73, 0xfd, 0xcd, 0xa5, 0x1f, 0x4d, 0x8e, 0x8c, 0x2c, 0x6f, 0xb8, 0xa5, 0xf8, 0x94, 0x3b, 0x08,
	0x9a, 0xa0, 0x82, 0x70, 0x4c, 0xa3, 0x14, 0x5b, 0x10, 0xd8, 0xfa, 0x62, 0xae, 0xbf, 0x9d, 0x61,
	0x97, 0x15, 0xc3, 0x5d, 0x13, 0x47, 0x07, 0xc1, 0x9f, 0x14, 0x50, 0x63, 0xdc, 0x4f, 0xb8, 0x37,
	0x4d, 0x48, 0x80, 0xd5, 0x62, 0x53, 0x69, 0xd5, 0x0e, 0x77, 0xcd, 0xcc, 0x11, 0x33, 0x75, 0xc4,
	0x94, 0x8e, 0x98, 0x1d, 0x4a, 0x62, 0xbb, 0x7b, 0x3d, 0xd7, 0x73, 0x8b, 0xb9, 0x0e, 0x33, 0xc9,
	0x15, 0xae, 0xf1, 0xcb, 0xad, 0xfe, 0x61, 0x48, 0xf8, 0x78, 0x36, 0x32, 0x03, 0x1a, 0x59, 0xd2,
	0xd4, 0xec, 0xf1, 0x31, 0x43, 0x67, 0x16, 0xbf, 0x9c, 0x62, 0x26, 0x64, 0x5c, 0x20, 0x88, 0xfd,
	0x94, 0x07, 0xbf, 0x03, 0x59, 0xe4, 0xa5, 0x06, 0xaa, 0x25, 0xd1, 0x41, 0xc3, 0xcc, 0xdc, 0x35,
	0x97, 0xee, 0x9a, 0xc3, 0xa5, 0xbb, 0xf6, 0xbb, 0xb2, 0x85, 0xcd, 0xd5, 0x16, 0x52, 0xae, 0x71,
	0x75, 0xab, 0x2b, 0x6e, 0x55, 0x24, 0x52, 0x38, 0xec, 0x81, 0x0a, 0x8e, 0x51, 0xa6, 0x5b, 0x7e,
	0x55, 0x77, 0xe7, 0xc9, 0xa9, 0x25, 0x2b, 0x53, 0x5c, 0xc3, 0x31, 0x12, 0x7a, 0x5b, 0xa0, 0x44,
	0x2f, 0x62, 0x9c, 0xa8, 0x6b, 0xa9, 0xb5, 0x6e, 0x16, 0xc0, 0x1f, 0x15, 0xb0, 0x45, 0xe2, 0x20,
	0xc1, 0x11, 0x8e, 0xb9, 0x37, 0xc5, 0x49, 0x80, 0x63, 0xee, 0x87, 0x58, 0xad, 0x88, 0x0b, 0xf8,
	0x2a, 0x6d, 0xf7, 0xf7, 0xb9, 0xfe, 0xc1, 0xff, 0xf0, 0xe6, 0x18, 0x07, 0x8b, 0xb9, 0xbe, 0x97,
	0x35, 0xf1, 0x92, 0xa6, 0xe1, 0xd6, 0x1f, 0xd3, 0xfd, 0xc7, 0x2c, 0xec, 0x03, 0x78, 0x31, 0x26,
	0x1c, 0x4f, 0x08, 0xe3, 0x9e, 0x1f, 0x04, 0x74, 0x16, 0x73, 0xa6, 0x56, 0x9b, 0x85, 0x56, 0xd5,
	0xde, 0x97, 0x76, 0xed, 0x66, 0xaa, 0xff, 0xc6, 0x19, 0xee, 0xe6, 0x63, 0xb2, 0x2d, 0x73, 0x10,
	0x81, 0xb7, 0xd8, 0x74, 0x42, 0xb8, 0xc7, 0xc6, 0x7e, 0x82, 0x99, 0x0a, 0x9a, 0x85, 0x56, 0xed,
	0xf0, 0x3d, 0xf3, 0x85, 0xfd, 0x31, 0xbf, 0xc5, 0x24, 0x1c, 0x73, 0x8c, 0xda, 0x08, 0x25, 0x98,
	0x31, 0x7b, 0x4f, 0xbe, 0xb1, 0x2e, 0x2f, 0x68, 0x45, 0xc7, 0x70, 0x6b, 0x22, 0x1c, 0x64, 0xd1,
	0xaf, 0x0a, 0x28, 0xd8, 0x04, 0xc1, 0x4f, 0x01, 0x90, 0x7b, 0xea, 0x2d, 0x07, 0xdf, 0xde, 0x7e,
	0xba, 0xe2, 0xa7, 0x9a, 0xe1, 0x56, 0x65, 0xe0, 0x20, 0xf8, 0x0e, 0x28, 0x8f, 0x08, 0x42, 0x38,
	0xc9, 0xd6, 0xc2, 0x95, 0x11, 0xfc, 0x0c, 0x94, 0xfd, 0x28, 0xfd, 0x0c, 0xb1, 0x02, 0xff, 0x39,
	0xce, 0xc5, 0xb4, 0x55, 0x57, 0xc2, 0xe1, 0xe7, 0xa0, 0x28, 0x66, 0xa5, 0xf8, 0xea, 0xac, 0x54,
	0x52, 0x9e, 0x18, 0x0e, 0xc1, 0x38, 0x2a, 0xfe, 0xf9, 0xb3, 0xae, 0x7c, 0x14, 0x82, 0x37, 0x72,
	0x93, 0x07, 0xdc, 0xe7, 0x33, 0x06, 0x35, 0xd0, 0x68, 0x9f, 0x74, 0x86, 0xce, 0xd7, 0x3d, 0x6f,
	0x30, 0x6c, 0x0f, 0x4f, 0x06, 0xde, 0x49, 0x6f, 0xd0, 0xef, 0x76, 0x9c, 0x2f, 0x9c, 0xee, 0xf1,
	0x46, 0x0e, 0xee, 0x81, 0x9d, 0x67, 0x75, 0xa7, 0xd7, 0xee, 0x0c, 0x9d, 0x6f, 0xba, 0x1b, 0x0a,
	0xdc, 0x05, 0xdb, 0xcf, 0x8a, 0xb2, 0x94, 0xb7, 0x8f, 0xae, 0xef, 0xb4, 0xdc, 0xcd, 0x9d, 0x96,
	0xbb, 0xbe, 0xd7, 0x94, 0x9b, 0x7b, 0x4d, 0xf9, 0xe3, 0x5e, 0x53, 0xae, 0x1e, 0xb4, 0xdc, 0xcd,
	0x83, 0x96, 0xfb, 0xed, 0x41, 0xcb, 0x7d, 0xaf, 0x8a, 0x9f, 0xcf, 0x0f, 0xff, 0xf8, 0xfd, 0x88,
	0x19, 0x1b, 0x95, 0xc5, 0xe7, 0x7c, 0xf2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x39, 0xff,
	0x63, 0x2e, 0x05, 0x00, 0x00,
}

func (this *Bid) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Bid)
	if !ok {
		that2, ok := that.(Bid)
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
	if this.AuctionId != that1.AuctionId {
		return false
	}
	if this.Bidder != that1.Bidder {
		return false
	}
	if !this.Amount.Equal(&that1.Amount) {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	return true
}
func (m *AuctionListing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionListing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionListing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SplitShares) > 0 {
		for iNdEx := len(m.SplitShares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SplitShares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuction(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.WhitelistAccounts) > 0 {
		for iNdEx := len(m.WhitelistAccounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WhitelistAccounts[iNdEx])
			copy(dAtA[i:], m.WhitelistAccounts[iNdEx])
			i = encodeVarintAuction(dAtA, i, uint64(len(m.WhitelistAccounts[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	{
		size := m.IncrementPercentage.Size()
		i -= size
		if _, err := m.IncrementPercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x3a
	}
	if m.EndTime != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndTime):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintAuction(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x32
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintAuction(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.StartPrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.DenomId) > 0 {
		i -= len(m.DenomId)
		copy(dAtA[i:], m.DenomId)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.DenomId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NftId) > 0 {
		i -= len(m.NftId)
		copy(dAtA[i:], m.NftId)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.NftId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Bid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintAuction(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x12
	}
	if m.AuctionId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuctionListing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAuction(uint64(m.Id))
	}
	l = len(m.NftId)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = len(m.DenomId)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = m.StartPrice.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovAuction(uint64(l))
	if m.EndTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndTime)
		n += 1 + l + sovAuction(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = m.IncrementPercentage.Size()
	n += 1 + l + sovAuction(uint64(l))
	if len(m.WhitelistAccounts) > 0 {
		for _, s := range m.WhitelistAccounts {
			l = len(s)
			n += 1 + l + sovAuction(uint64(l))
		}
	}
	if len(m.SplitShares) > 0 {
		for _, e := range m.SplitShares {
			l = e.Size()
			n += 1 + l + sovAuction(uint64(l))
		}
	}
	return n
}

func (m *Bid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovAuction(uint64(m.AuctionId))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovAuction(uint64(l))
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuctionListing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: AuctionListing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionListing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StartPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndTime == nil {
				m.EndTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncrementPercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IncrementPercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistAccounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistAccounts = append(m.WhitelistAccounts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SplitShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SplitShares = append(m.SplitShares, WeightedAddress{})
			if err := m.SplitShares[len(m.SplitShares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *Bid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Bid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)