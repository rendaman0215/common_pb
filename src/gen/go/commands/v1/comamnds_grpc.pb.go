// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: commands/v1/comamnds.proto

package commandsv1

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
	CategoryCommandService_Create_FullMethodName = "/commands.v1.CategoryCommandService/Create"
	CategoryCommandService_Update_FullMethodName = "/commands.v1.CategoryCommandService/Update"
	CategoryCommandService_Delete_FullMethodName = "/commands.v1.CategoryCommandService/Delete"
)

// CategoryCommandServiceClient is the client API for CategoryCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryCommandServiceClient interface {
	Create(ctx context.Context, in *CategoryCommandServiceCreateRequest, opts ...grpc.CallOption) (*CategoryCommandServiceCreateResponse, error)
	Update(ctx context.Context, in *CategoryCommandServiceUpdateRequest, opts ...grpc.CallOption) (*CategoryCommandServiceUpdateResponse, error)
	Delete(ctx context.Context, in *CategoryCommandServiceDeleteRequest, opts ...grpc.CallOption) (*CategoryCommandServiceDeleteResponse, error)
}

type categoryCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryCommandServiceClient(cc grpc.ClientConnInterface) CategoryCommandServiceClient {
	return &categoryCommandServiceClient{cc}
}

func (c *categoryCommandServiceClient) Create(ctx context.Context, in *CategoryCommandServiceCreateRequest, opts ...grpc.CallOption) (*CategoryCommandServiceCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryCommandServiceCreateResponse)
	err := c.cc.Invoke(ctx, CategoryCommandService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryCommandServiceClient) Update(ctx context.Context, in *CategoryCommandServiceUpdateRequest, opts ...grpc.CallOption) (*CategoryCommandServiceUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryCommandServiceUpdateResponse)
	err := c.cc.Invoke(ctx, CategoryCommandService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryCommandServiceClient) Delete(ctx context.Context, in *CategoryCommandServiceDeleteRequest, opts ...grpc.CallOption) (*CategoryCommandServiceDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryCommandServiceDeleteResponse)
	err := c.cc.Invoke(ctx, CategoryCommandService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryCommandServiceServer is the server API for CategoryCommandService service.
// All implementations must embed UnimplementedCategoryCommandServiceServer
// for forward compatibility.
type CategoryCommandServiceServer interface {
	Create(context.Context, *CategoryCommandServiceCreateRequest) (*CategoryCommandServiceCreateResponse, error)
	Update(context.Context, *CategoryCommandServiceUpdateRequest) (*CategoryCommandServiceUpdateResponse, error)
	Delete(context.Context, *CategoryCommandServiceDeleteRequest) (*CategoryCommandServiceDeleteResponse, error)
	mustEmbedUnimplementedCategoryCommandServiceServer()
}

// UnimplementedCategoryCommandServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCategoryCommandServiceServer struct{}

func (UnimplementedCategoryCommandServiceServer) Create(context.Context, *CategoryCommandServiceCreateRequest) (*CategoryCommandServiceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCategoryCommandServiceServer) Update(context.Context, *CategoryCommandServiceUpdateRequest) (*CategoryCommandServiceUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCategoryCommandServiceServer) Delete(context.Context, *CategoryCommandServiceDeleteRequest) (*CategoryCommandServiceDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCategoryCommandServiceServer) mustEmbedUnimplementedCategoryCommandServiceServer() {
}
func (UnimplementedCategoryCommandServiceServer) testEmbeddedByValue() {}

// UnsafeCategoryCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryCommandServiceServer will
// result in compilation errors.
type UnsafeCategoryCommandServiceServer interface {
	mustEmbedUnimplementedCategoryCommandServiceServer()
}

func RegisterCategoryCommandServiceServer(s grpc.ServiceRegistrar, srv CategoryCommandServiceServer) {
	// If the following call pancis, it indicates UnimplementedCategoryCommandServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CategoryCommandService_ServiceDesc, srv)
}

func _CategoryCommandService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryCommandServiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryCommandServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryCommandService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryCommandServiceServer).Create(ctx, req.(*CategoryCommandServiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryCommandService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryCommandServiceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryCommandServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryCommandService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryCommandServiceServer).Update(ctx, req.(*CategoryCommandServiceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryCommandService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryCommandServiceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryCommandServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryCommandService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryCommandServiceServer).Delete(ctx, req.(*CategoryCommandServiceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryCommandService_ServiceDesc is the grpc.ServiceDesc for CategoryCommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryCommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commands.v1.CategoryCommandService",
	HandlerType: (*CategoryCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CategoryCommandService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CategoryCommandService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CategoryCommandService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commands/v1/comamnds.proto",
}

const (
	ProductCommandService_Create_FullMethodName = "/commands.v1.ProductCommandService/Create"
	ProductCommandService_Update_FullMethodName = "/commands.v1.ProductCommandService/Update"
	ProductCommandService_Delete_FullMethodName = "/commands.v1.ProductCommandService/Delete"
)

// ProductCommandServiceClient is the client API for ProductCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCommandServiceClient interface {
	Create(ctx context.Context, in *ProductCommandServiceCreateRequest, opts ...grpc.CallOption) (*ProductCommandServiceCreateResponse, error)
	Update(ctx context.Context, in *ProductCommandServiceUpdateRequest, opts ...grpc.CallOption) (*ProductCommandServiceUpdateResponse, error)
	Delete(ctx context.Context, in *ProductCommandServiceDeleteRequest, opts ...grpc.CallOption) (*ProductCommandServiceDeleteResponse, error)
}

type productCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCommandServiceClient(cc grpc.ClientConnInterface) ProductCommandServiceClient {
	return &productCommandServiceClient{cc}
}

func (c *productCommandServiceClient) Create(ctx context.Context, in *ProductCommandServiceCreateRequest, opts ...grpc.CallOption) (*ProductCommandServiceCreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductCommandServiceCreateResponse)
	err := c.cc.Invoke(ctx, ProductCommandService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCommandServiceClient) Update(ctx context.Context, in *ProductCommandServiceUpdateRequest, opts ...grpc.CallOption) (*ProductCommandServiceUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductCommandServiceUpdateResponse)
	err := c.cc.Invoke(ctx, ProductCommandService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCommandServiceClient) Delete(ctx context.Context, in *ProductCommandServiceDeleteRequest, opts ...grpc.CallOption) (*ProductCommandServiceDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductCommandServiceDeleteResponse)
	err := c.cc.Invoke(ctx, ProductCommandService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCommandServiceServer is the server API for ProductCommandService service.
// All implementations must embed UnimplementedProductCommandServiceServer
// for forward compatibility.
type ProductCommandServiceServer interface {
	Create(context.Context, *ProductCommandServiceCreateRequest) (*ProductCommandServiceCreateResponse, error)
	Update(context.Context, *ProductCommandServiceUpdateRequest) (*ProductCommandServiceUpdateResponse, error)
	Delete(context.Context, *ProductCommandServiceDeleteRequest) (*ProductCommandServiceDeleteResponse, error)
	mustEmbedUnimplementedProductCommandServiceServer()
}

// UnimplementedProductCommandServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductCommandServiceServer struct{}

func (UnimplementedProductCommandServiceServer) Create(context.Context, *ProductCommandServiceCreateRequest) (*ProductCommandServiceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProductCommandServiceServer) Update(context.Context, *ProductCommandServiceUpdateRequest) (*ProductCommandServiceUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProductCommandServiceServer) Delete(context.Context, *ProductCommandServiceDeleteRequest) (*ProductCommandServiceDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProductCommandServiceServer) mustEmbedUnimplementedProductCommandServiceServer() {}
func (UnimplementedProductCommandServiceServer) testEmbeddedByValue()                               {}

// UnsafeProductCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCommandServiceServer will
// result in compilation errors.
type UnsafeProductCommandServiceServer interface {
	mustEmbedUnimplementedProductCommandServiceServer()
}

func RegisterProductCommandServiceServer(s grpc.ServiceRegistrar, srv ProductCommandServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductCommandServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductCommandService_ServiceDesc, srv)
}

func _ProductCommandService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCommandServiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCommandServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCommandService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCommandServiceServer).Create(ctx, req.(*ProductCommandServiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCommandService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCommandServiceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCommandServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCommandService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCommandServiceServer).Update(ctx, req.(*ProductCommandServiceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCommandService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCommandServiceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCommandServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCommandService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCommandServiceServer).Delete(ctx, req.(*ProductCommandServiceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCommandService_ServiceDesc is the grpc.ServiceDesc for ProductCommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commands.v1.ProductCommandService",
	HandlerType: (*ProductCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProductCommandService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductCommandService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProductCommandService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commands/v1/comamnds.proto",
}
