// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package corev1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() { proto.RegisterFile("core.proto", fileDescriptor_f7e43720d1edc0fe) }

var fileDescriptor_f7e43720d1edc0fe = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x2f, 0x4a,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x02, 0x19, 0xa9, 0x39, 0x39,
	0xf9, 0xe5, 0xf9, 0x45, 0x39, 0x29, 0x10, 0x09, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54,
	0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62,
	0x88, 0xac, 0x51, 0x17, 0x13, 0x17, 0x97, 0x73, 0x51, 0x6a, 0x4a, 0x6a, 0x5e, 0x49, 0x66, 0x62,
	0x8e, 0x50, 0x10, 0x17, 0x47, 0x70, 0x62, 0xa5, 0x07, 0xc8, 0x0c, 0x21, 0x61, 0x88, 0x12, 0x3d,
	0x30, 0x2f, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x4a, 0x10, 0x55, 0xb0, 0x20, 0xa7, 0x52,
	0x49, 0xae, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x12, 0x4a, 0xc2, 0x60, 0x2b, 0xca, 0x0c, 0xf5, 0x93,
	0xe1, 0xe6, 0x19, 0x5a, 0x31, 0x6a, 0x09, 0x05, 0x73, 0x71, 0xc2, 0xcc, 0x34, 0xa2, 0x86, 0xa1,
	0x46, 0x68, 0x86, 0x1a, 0x53, 0xc3, 0x50, 0x63, 0x2b, 0x46, 0x2d, 0xa3, 0x48, 0x2e, 0x76, 0xdf,
	0xc4, 0xe4, 0x8c, 0xcc, 0xbc, 0x54, 0x21, 0x3f, 0x32, 0x02, 0x42, 0x0a, 0x6c, 0xbc, 0x88, 0x12,
	0x3f, 0xcc, 0xf8, 0x5c, 0x88, 0x61, 0x56, 0x8c, 0x5a, 0x4e, 0xbc, 0x51, 0xdc, 0x7a, 0xfa, 0x05,
	0x49, 0xfa, 0xa0, 0x18, 0x2b, 0x33, 0x4c, 0x62, 0x03, 0x6b, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x23, 0xef, 0xe0, 0x9f, 0xc2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CredentialClient is the client API for Credential service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CredentialClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHello2(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHello3(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type credentialClient struct {
	cc *grpc.ClientConn
}

func NewCredentialClient(cc *grpc.ClientConn) CredentialClient {
	return &credentialClient{cc}
}

func (c *credentialClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/proto.Credential/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialClient) SayHello2(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/proto.Credential/SayHello2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialClient) SayHello3(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/proto.Credential/SayHello3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialServer is the server API for Credential service.
type CredentialServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHello2(context.Context, *HelloRequest) (*HelloReply, error)
	SayHello3(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedCredentialServer can be embedded to have forward compatible implementations.
type UnimplementedCredentialServer struct {
}

func (*UnimplementedCredentialServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedCredentialServer) SayHello2(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello2 not implemented")
}
func (*UnimplementedCredentialServer) SayHello3(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello3 not implemented")
}

func RegisterCredentialServer(s *grpc.Server, srv CredentialServer) {
	s.RegisterService(&_Credential_serviceDesc, srv)
}

func _Credential_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Credential/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credential_SayHello2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServer).SayHello2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Credential/SayHello2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServer).SayHello2(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Credential_SayHello3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServer).SayHello3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Credential/SayHello3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServer).SayHello3(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Credential_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Credential",
	HandlerType: (*CredentialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Credential_SayHello_Handler,
		},
		{
			MethodName: "SayHello2",
			Handler:    _Credential_SayHello2_Handler,
		},
		{
			MethodName: "SayHello3",
			Handler:    _Credential_SayHello3_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}

// MachineClient is the client API for Machine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MachineClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type machineClient struct {
	cc *grpc.ClientConn
}

func NewMachineClient(cc *grpc.ClientConn) MachineClient {
	return &machineClient{cc}
}

func (c *machineClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/proto.Machine/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineServer is the server API for Machine service.
type MachineServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedMachineServer can be embedded to have forward compatible implementations.
type UnimplementedMachineServer struct {
}

func (*UnimplementedMachineServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterMachineServer(s *grpc.Server, srv MachineServer) {
	s.RegisterService(&_Machine_serviceDesc, srv)
}

func _Machine_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Machine/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Machine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Machine",
	HandlerType: (*MachineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Machine_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}
