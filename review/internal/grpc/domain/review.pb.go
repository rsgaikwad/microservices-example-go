// Code generated by protoc-gen-go. DO NOT EDIT.
// source: microservices-example-go/review/internal/proto-files/domain/review.proto

package domain

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Review struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Review               string   `protobuf:"bytes,2,opt,name=review,proto3" json:"review,omitempty"`
	Rating               string   `protobuf:"bytes,3,opt,name=rating,proto3" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Review) Reset()         { *m = Review{} }
func (m *Review) String() string { return proto.CompactTextString(m) }
func (*Review) ProtoMessage()    {}
func (*Review) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e8442d25abb14d, []int{0}
}

func (m *Review) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Review.Unmarshal(m, b)
}
func (m *Review) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Review.Marshal(b, m, deterministic)
}
func (m *Review) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Review.Merge(m, src)
}
func (m *Review) XXX_Size() int {
	return xxx_messageInfo_Review.Size(m)
}
func (m *Review) XXX_DiscardUnknown() {
	xxx_messageInfo_Review.DiscardUnknown(m)
}

var xxx_messageInfo_Review proto.InternalMessageInfo

func (m *Review) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Review) GetReview() string {
	if m != nil {
		return m.Review
	}
	return ""
}

func (m *Review) GetRating() string {
	if m != nil {
		return m.Rating
	}
	return ""
}

func init() {
	proto.RegisterType((*Review)(nil), "domain.review")
}

func init() {
	proto.RegisterFile("microservices-example-go/review/internal/proto-files/domain/review.proto", fileDescriptor_73e8442d25abb14d)
}

var fileDescriptor_73e8442d25abb14d = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0xc8, 0xcd, 0x4c, 0x2e,
	0xca, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4d, 0xad, 0x48, 0xcc, 0x2d, 0xc8,
	0x49, 0xd5, 0x4d, 0xcf, 0xd7, 0x2f, 0x4a, 0x2d, 0xcb, 0x4c, 0x2d, 0xd7, 0xcf, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4d, 0xcb, 0xcc, 0x49, 0x2d,
	0xd6, 0x4f, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x83, 0x2a, 0xd1, 0x03, 0xcb, 0x08, 0xb1, 0x41, 0x04,
	0x95, 0x3c, 0xb8, 0xd8, 0x20, 0xe2, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xcc, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x62, 0x30, 0x19, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x98, 0x3a, 0x90, 0x78, 0x62, 0x49, 0x66, 0x5e, 0xba, 0x04, 0x33, 0x54, 0x1c, 0xcc, 0x73,
	0x32, 0x8b, 0x32, 0x21, 0xda, 0x75, 0xe9, 0x45, 0x05, 0xc9, 0x50, 0x67, 0x25, 0xb1, 0x81, 0x1d,
	0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x24, 0x30, 0x24, 0xd8, 0xdc, 0x00, 0x00, 0x00,
}