// Code generated by protoc-gen-go. DO NOT EDIT.
// source: microservices-example-go/details/internal/proto-files/service/details-service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	domain "microservices-example-go/details/internal/grpc/domain"
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

type GetDetailsResponse struct {
	Details              *domain.Details `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetDetailsResponse) Reset()         { *m = GetDetailsResponse{} }
func (m *GetDetailsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDetailsResponse) ProtoMessage()    {}
func (*GetDetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_beb6d9aa221348ac, []int{0}
}

func (m *GetDetailsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDetailsResponse.Unmarshal(m, b)
}
func (m *GetDetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDetailsResponse.Marshal(b, m, deterministic)
}
func (m *GetDetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDetailsResponse.Merge(m, src)
}
func (m *GetDetailsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDetailsResponse.Size(m)
}
func (m *GetDetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDetailsResponse proto.InternalMessageInfo

func (m *GetDetailsResponse) GetDetails() *domain.Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDetailsResponse)(nil), "service.GetDetailsResponse")
}

func init() {
	proto.RegisterFile("microservices-example-go/details/internal/proto-files/service/details-service.proto", fileDescriptor_beb6d9aa221348ac)
}

var fileDescriptor_beb6d9aa221348ac = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0xce, 0xcd, 0x4c, 0x2e,
	0xca, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4d, 0xad, 0x48, 0xcc, 0x2d, 0xc8,
	0x49, 0xd5, 0x4d, 0xcf, 0xd7, 0x4f, 0x49, 0x2d, 0x49, 0xcc, 0xcc, 0x29, 0xd6, 0xcf, 0xcc, 0x2b,
	0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4d, 0xcb, 0xcc, 0x49,
	0x2d, 0xd6, 0x87, 0x6a, 0x80, 0x29, 0xd2, 0x85, 0xf2, 0xf5, 0xc0, 0x6a, 0x84, 0xd8, 0xa1, 0x5c,
	0x29, 0x2f, 0xf2, 0x4c, 0x4f, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x83, 0xa9, 0x81, 0x18, 0xaa, 0x64,
	0xcf, 0x25, 0xe4, 0x9e, 0x5a, 0xe2, 0x02, 0x11, 0x0b, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e,
	0x15, 0xd2, 0xe4, 0x62, 0x87, 0x2a, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd7, 0x83,
	0xe8, 0xd6, 0x83, 0x0a, 0x07, 0xc1, 0xe4, 0x8d, 0x5c, 0xb9, 0xf8, 0xa0, 0xba, 0x83, 0x21, 0x0e,
	0x12, 0x32, 0xe6, 0x62, 0x4e, 0x4f, 0x2d, 0x11, 0x42, 0xd7, 0x22, 0x25, 0xad, 0x07, 0xf3, 0x0f,
	0xa6, 0x8d, 0x4e, 0x16, 0x51, 0x66, 0xc4, 0xfb, 0x2a, 0xbd, 0xa8, 0x20, 0x19, 0x16, 0x58, 0x49,
	0x6c, 0x60, 0x8f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x23, 0x21, 0x7f, 0xe2, 0x74, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DetailsServiceClient is the client API for DetailsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DetailsServiceClient interface {
	Get(ctx context.Context, in *domain.Details, opts ...grpc.CallOption) (*GetDetailsResponse, error)
}

type detailsServiceClient struct {
	cc *grpc.ClientConn
}

func NewDetailsServiceClient(cc *grpc.ClientConn) DetailsServiceClient {
	return &detailsServiceClient{cc}
}

func (c *detailsServiceClient) Get(ctx context.Context, in *domain.Details, opts ...grpc.CallOption) (*GetDetailsResponse, error) {
	out := new(GetDetailsResponse)
	err := c.cc.Invoke(ctx, "/service.DetailsService/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailsServiceServer is the server API for DetailsService service.
type DetailsServiceServer interface {
	Get(context.Context, *domain.Details) (*GetDetailsResponse, error)
}

// UnimplementedDetailsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDetailsServiceServer struct {
}

func (*UnimplementedDetailsServiceServer) Get(ctx context.Context, req *domain.Details) (*GetDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterDetailsServiceServer(s *grpc.Server, srv DetailsServiceServer) {
	s.RegisterService(&_DetailsService_serviceDesc, srv)
}

func _DetailsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.Details)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.DetailsService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailsServiceServer).Get(ctx, req.(*domain.Details))
	}
	return interceptor(ctx, in, info, handler)
}

var _DetailsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.DetailsService",
	HandlerType: (*DetailsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _DetailsService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservices-example-go/details/internal/proto-files/service/details-service.proto",
}