// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/crypto/crypto.proto

package crypto

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
	CryptoService_GetTopCoins_FullMethodName = "/crypto.CryptoService/GetTopCoins"
)

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoServiceClient interface {
	GetTopCoins(ctx context.Context, in *TopCoinsRequest, opts ...grpc.CallOption) (*TopCoinsResponse, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) GetTopCoins(ctx context.Context, in *TopCoinsRequest, opts ...grpc.CallOption) (*TopCoinsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TopCoinsResponse)
	err := c.cc.Invoke(ctx, CryptoService_GetTopCoins_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
// All implementations should embed UnimplementedCryptoServiceServer
// for forward compatibility.
type CryptoServiceServer interface {
	GetTopCoins(context.Context, *TopCoinsRequest) (*TopCoinsResponse, error)
}

// UnimplementedCryptoServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCryptoServiceServer struct{}

func (UnimplementedCryptoServiceServer) GetTopCoins(context.Context, *TopCoinsRequest) (*TopCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopCoins not implemented")
}
func (UnimplementedCryptoServiceServer) testEmbeddedByValue() {}

// UnsafeCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoServiceServer will
// result in compilation errors.
type UnsafeCryptoServiceServer interface {
	mustEmbedUnimplementedCryptoServiceServer()
}

func RegisterCryptoServiceServer(s grpc.ServiceRegistrar, srv CryptoServiceServer) {
	// If the following call pancis, it indicates UnimplementedCryptoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CryptoService_ServiceDesc, srv)
}

func _CryptoService_GetTopCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetTopCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CryptoService_GetTopCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetTopCoins(ctx, req.(*TopCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptoService_ServiceDesc is the grpc.ServiceDesc for CryptoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crypto.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopCoins",
			Handler:    _CryptoService_GetTopCoins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/crypto/crypto.proto",
}
