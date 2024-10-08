// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: pokemon/calc/v1/calc.proto

package calcv1

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
	DamageCalcService_DamageCalc_FullMethodName = "/pokemon.calc.v1.DamageCalcService/DamageCalc"
)

// DamageCalcServiceClient is the client API for DamageCalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ダメージ計算サービス
type DamageCalcServiceClient interface {
	// ダメージ計算
	DamageCalc(ctx context.Context, in *DamageCalcRequest, opts ...grpc.CallOption) (*DamageCalcResponse, error)
}

type damageCalcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDamageCalcServiceClient(cc grpc.ClientConnInterface) DamageCalcServiceClient {
	return &damageCalcServiceClient{cc}
}

func (c *damageCalcServiceClient) DamageCalc(ctx context.Context, in *DamageCalcRequest, opts ...grpc.CallOption) (*DamageCalcResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DamageCalcResponse)
	err := c.cc.Invoke(ctx, DamageCalcService_DamageCalc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DamageCalcServiceServer is the server API for DamageCalcService service.
// All implementations must embed UnimplementedDamageCalcServiceServer
// for forward compatibility.
//
// ダメージ計算サービス
type DamageCalcServiceServer interface {
	// ダメージ計算
	DamageCalc(context.Context, *DamageCalcRequest) (*DamageCalcResponse, error)
	mustEmbedUnimplementedDamageCalcServiceServer()
}

// UnimplementedDamageCalcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDamageCalcServiceServer struct{}

func (UnimplementedDamageCalcServiceServer) DamageCalc(context.Context, *DamageCalcRequest) (*DamageCalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DamageCalc not implemented")
}
func (UnimplementedDamageCalcServiceServer) mustEmbedUnimplementedDamageCalcServiceServer() {}
func (UnimplementedDamageCalcServiceServer) testEmbeddedByValue()                           {}

// UnsafeDamageCalcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DamageCalcServiceServer will
// result in compilation errors.
type UnsafeDamageCalcServiceServer interface {
	mustEmbedUnimplementedDamageCalcServiceServer()
}

func RegisterDamageCalcServiceServer(s grpc.ServiceRegistrar, srv DamageCalcServiceServer) {
	// If the following call pancis, it indicates UnimplementedDamageCalcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DamageCalcService_ServiceDesc, srv)
}

func _DamageCalcService_DamageCalc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DamageCalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DamageCalcServiceServer).DamageCalc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DamageCalcService_DamageCalc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DamageCalcServiceServer).DamageCalc(ctx, req.(*DamageCalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DamageCalcService_ServiceDesc is the grpc.ServiceDesc for DamageCalcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DamageCalcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.calc.v1.DamageCalcService",
	HandlerType: (*DamageCalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DamageCalc",
			Handler:    _DamageCalcService_DamageCalc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokemon/calc/v1/calc.proto",
}
