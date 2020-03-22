// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeting.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Greeter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeter) Reset()         { *m = Greeter{} }
func (m *Greeter) String() string { return proto.CompactTextString(m) }
func (*Greeter) ProtoMessage()    {}
func (*Greeter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acac03ccd168a87, []int{0}
}

func (m *Greeter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeter.Unmarshal(m, b)
}
func (m *Greeter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeter.Marshal(b, m, deterministic)
}
func (m *Greeter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeter.Merge(m, src)
}
func (m *Greeter) XXX_Size() int {
	return xxx_messageInfo_Greeter.Size(m)
}
func (m *Greeter) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeter.DiscardUnknown(m)
}

var xxx_messageInfo_Greeter proto.InternalMessageInfo

func (m *Greeter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetingResponse struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingResponse) Reset()         { *m = GreetingResponse{} }
func (m *GreetingResponse) String() string { return proto.CompactTextString(m) }
func (*GreetingResponse) ProtoMessage()    {}
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acac03ccd168a87, []int{1}
}

func (m *GreetingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingResponse.Unmarshal(m, b)
}
func (m *GreetingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingResponse.Marshal(b, m, deterministic)
}
func (m *GreetingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingResponse.Merge(m, src)
}
func (m *GreetingResponse) XXX_Size() int {
	return xxx_messageInfo_GreetingResponse.Size(m)
}
func (m *GreetingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingResponse proto.InternalMessageInfo

func (m *GreetingResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func (m *GreetingResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeter)(nil), "protocol.Greeter")
	proto.RegisterType((*GreetingResponse)(nil), "protocol.GreetingResponse")
}

func init() {
	proto.RegisterFile("greeting.proto", fileDescriptor_6acac03ccd168a87)
}

var fileDescriptor_6acac03ccd168a87 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2f, 0x4a, 0x4d,
	0x2d, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9,
	0x39, 0x4a, 0xb2, 0x5c, 0xec, 0xee, 0x20, 0xb9, 0xd4, 0x22, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4,
	0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x89, 0x4b, 0xc0, 0x1d,
	0xaa, 0x35, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8a, 0x8b, 0x03, 0x66, 0x1c,
	0x54, 0x2d, 0x9c, 0x0f, 0x37, 0x83, 0x09, 0x61, 0x86, 0x91, 0x0b, 0x17, 0x07, 0xcc, 0x0c, 0x21,
	0x0b, 0x2e, 0x56, 0x8f, 0xd4, 0x9c, 0x9c, 0x7c, 0x21, 0x41, 0x3d, 0x98, 0x13, 0xf4, 0xa0, 0xf6,
	0x4b, 0x49, 0xa1, 0x09, 0x21, 0xd9, 0xa9, 0xc4, 0x90, 0xc4, 0x06, 0x96, 0x34, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xe9, 0x24, 0x9f, 0xa8, 0xcb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetingClient is the client API for Greeting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetingClient interface {
	Hello(ctx context.Context, in *Greeter, opts ...grpc.CallOption) (*GreetingResponse, error)
}

type greetingClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingClient(cc grpc.ClientConnInterface) GreetingClient {
	return &greetingClient{cc}
}

func (c *greetingClient) Hello(ctx context.Context, in *Greeter, opts ...grpc.CallOption) (*GreetingResponse, error) {
	out := new(GreetingResponse)
	err := c.cc.Invoke(ctx, "/protocol.Greeting/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServer is the server API for Greeting service.
type GreetingServer interface {
	Hello(context.Context, *Greeter) (*GreetingResponse, error)
}

// UnimplementedGreetingServer can be embedded to have forward compatible implementations.
type UnimplementedGreetingServer struct {
}

func (*UnimplementedGreetingServer) Hello(ctx context.Context, req *Greeter) (*GreetingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}

func RegisterGreetingServer(s *grpc.Server, srv GreetingServer) {
	s.RegisterService(&_Greeting_serviceDesc, srv)
}

func _Greeting_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Greeter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Greeting/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServer).Hello(ctx, req.(*Greeter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Greeting",
	HandlerType: (*GreetingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Greeting_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeting.proto",
}
