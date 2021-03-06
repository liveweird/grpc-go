// Code generated by protoc-gen-go.
// source: helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloServerStreamClient, error)
	SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloClientStreamClient, error)
	SayHelloBiDirectionalStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloBiDirectionalStreamClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloServerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[0], c.cc, "/helloworld.Greeter/SayHelloServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayHelloServerStreamClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloServerStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[1], c.cc, "/helloworld.Greeter/SayHelloClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloClientStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloClientStreamClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloClientStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloClientStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloClientStreamClient) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloBiDirectionalStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloBiDirectionalStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Greeter_serviceDesc.Streams[2], c.cc, "/helloworld.Greeter/SayHelloBiDirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloBiDirectionalStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloBiDirectionalStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloBiDirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloBiDirectionalStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloBiDirectionalStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloServerStream(*HelloRequest, Greeter_SayHelloServerStreamServer) error
	SayHelloClientStream(Greeter_SayHelloClientStreamServer) error
	SayHelloBiDirectionalStream(Greeter_SayHelloBiDirectionalStreamServer) error
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHelloServerStream(m, &greeterSayHelloServerStreamServer{stream})
}

type Greeter_SayHelloServerStreamServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterSayHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloServerStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_SayHelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloClientStream(&greeterSayHelloClientStreamServer{stream})
}

type Greeter_SayHelloClientStreamServer interface {
	SendAndClose(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloClientStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloClientStreamServer) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloClientStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_SayHelloBiDirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloBiDirectionalStream(&greeterSayHelloBiDirectionalStreamServer{stream})
}

type Greeter_SayHelloBiDirectionalStreamServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloBiDirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloBiDirectionalStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloBiDirectionalStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloServerStream",
			Handler:       _Greeter_SayHelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloClientStream",
			Handler:       _Greeter_SayHelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloBiDirectionalStream",
			Handler:       _Greeter_SayHelloBiDirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x91, 0x3f, 0x4b, 0x04, 0x31,
	0x10, 0xc5, 0x8d, 0x88, 0xa7, 0x83, 0xa0, 0x0c, 0x22, 0x8b, 0xd7, 0xc8, 0x16, 0x72, 0xd5, 0xb2,
	0x68, 0x6f, 0xb1, 0x0a, 0x5a, 0x58, 0x1c, 0xb7, 0x88, 0x75, 0x3c, 0x87, 0x33, 0x30, 0xbb, 0x89,
	0x93, 0xf8, 0x67, 0xbf, 0x9b, 0x1f, 0x4e, 0x36, 0x18, 0x4c, 0x61, 0xb5, 0x76, 0xef, 0x25, 0xef,
	0xfd, 0x12, 0x78, 0x70, 0xf4, 0x42, 0xcc, 0xf6, 0xc3, 0x0a, 0x3f, 0x57, 0x4e, 0x6c, 0xb0, 0x08,
	0xbf, 0x27, 0x65, 0x09, 0x07, 0x77, 0xa3, 0x5b, 0xd1, 0xeb, 0x1b, 0xf9, 0x80, 0x08, 0x3b, 0xbd,
	0xee, 0xa8, 0x50, 0x67, 0x6a, 0xb1, 0xbf, 0x8a, 0xba, 0x3c, 0x07, 0xf8, 0xc9, 0x38, 0x1e, 0xb0,
	0x80, 0x59, 0x47, 0xde, 0xeb, 0x4d, 0x0a, 0x25, 0x7b, 0xf1, 0xb5, 0x0d, 0xb3, 0x5b, 0x21, 0x0a,
	0x24, 0x78, 0x05, 0x7b, 0xad, 0x1e, 0x62, 0x0d, 0x8b, 0x2a, 0xfb, 0x42, 0xfe, 0xda, 0xe9, 0xc9,
	0x1f, 0x37, 0x8e, 0x87, 0x72, 0x0b, 0xef, 0xe1, 0x38, 0xf5, 0x5b, 0x92, 0x77, 0x92, 0x36, 0x08,
	0xe9, 0x6e, 0x0a, 0xab, 0x56, 0x39, 0xed, 0x9a, 0x0d, 0xf5, 0x61, 0x3a, 0x6d, 0xa1, 0xf0, 0x01,
	0xe6, 0x89, 0xd6, 0x98, 0x1b, 0x23, 0xb4, 0x0e, 0xc6, 0xf6, 0x9a, 0xff, 0x03, 0xad, 0x55, 0x53,
	0xc3, 0xdc, 0xd8, 0x6a, 0x23, 0x6e, 0x5d, 0xd1, 0xa7, 0xee, 0x1c, 0x93, 0xcf, 0xf2, 0xcd, 0x61,
	0x2c, 0x3c, 0x8e, 0x7a, 0x39, 0xce, 0xb8, 0x54, 0x4f, 0xbb, 0x71, 0xcf, 0xcb, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x25, 0x79, 0x5a, 0x44, 0xe3, 0x01, 0x00, 0x00,
}
