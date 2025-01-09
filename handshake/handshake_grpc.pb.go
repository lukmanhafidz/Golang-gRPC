// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: main.proto

package handshake

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Handshake_Handshake_FullMethodName = "/golanggrpc.Handshake/Handshake"
)

// HandshakeClient is the client API for Handshake service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandshakeClient interface {
	Handshake(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*HandshakeReply, error)
}

type handshakeClient struct {
	cc grpc.ClientConnInterface
}

func NewHandshakeClient(cc grpc.ClientConnInterface) HandshakeClient {
	return &handshakeClient{cc}
}

func (c *handshakeClient) Handshake(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*HandshakeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HandshakeReply)
	err := c.cc.Invoke(ctx, Handshake_Handshake_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandshakeServer is the server API for Handshake service.
// All implementations must embed UnimplementedHandshakeServer
// for forward compatibility.
type HandshakeServer interface {
	Handshake(context.Context, *HandshakeRequest) (*HandshakeReply, error)
	mustEmbedUnimplementedHandshakeServer()
}

// UnimplementedHandshakeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHandshakeServer struct{}

func (UnimplementedHandshakeServer) Handshake(context.Context, *HandshakeRequest) (*HandshakeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}
func (UnimplementedHandshakeServer) mustEmbedUnimplementedHandshakeServer() {}
func (UnimplementedHandshakeServer) testEmbeddedByValue()                   {}

// UnsafeHandshakeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandshakeServer will
// result in compilation errors.
type UnsafeHandshakeServer interface {
	mustEmbedUnimplementedHandshakeServer()
}

func RegisterHandshakeServer(s grpc.ServiceRegistrar, srv HandshakeServer) {
	// If the following call pancis, it indicates UnimplementedHandshakeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Handshake_ServiceDesc, srv)
}

func _Handshake_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandshakeServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Handshake_Handshake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandshakeServer).Handshake(ctx, req.(*HandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Handshake_ServiceDesc is the grpc.ServiceDesc for Handshake service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Handshake_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "golanggrpc.Handshake",
	HandlerType: (*HandshakeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _Handshake_Handshake_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
