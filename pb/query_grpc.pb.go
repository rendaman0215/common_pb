// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: proto/query.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CategoryQueryClient is the client API for CategoryQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryQueryClient interface {
	// すべての商品カテゴリを取得する
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoriesResult, error)
	// 指定されたIDの商品カテゴリを取得する
	ById(ctx context.Context, in *CategoryParam, opts ...grpc.CallOption) (*CategoryResult, error)
}

type categoryQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryQueryClient(cc grpc.ClientConnInterface) CategoryQueryClient {
	return &categoryQueryClient{cc}
}

func (c *categoryQueryClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoriesResult, error) {
	out := new(CategoriesResult)
	err := c.cc.Invoke(ctx, "/proto.CategoryQuery/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryQueryClient) ById(ctx context.Context, in *CategoryParam, opts ...grpc.CallOption) (*CategoryResult, error) {
	out := new(CategoryResult)
	err := c.cc.Invoke(ctx, "/proto.CategoryQuery/ById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryQueryServer is the server API for CategoryQuery service.
// All implementations must embed UnimplementedCategoryQueryServer
// for forward compatibility
type CategoryQueryServer interface {
	// すべての商品カテゴリを取得する
	List(context.Context, *emptypb.Empty) (*CategoriesResult, error)
	// 指定されたIDの商品カテゴリを取得する
	ById(context.Context, *CategoryParam) (*CategoryResult, error)
	mustEmbedUnimplementedCategoryQueryServer()
}

// UnimplementedCategoryQueryServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryQueryServer struct {
}

func (UnimplementedCategoryQueryServer) List(context.Context, *emptypb.Empty) (*CategoriesResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCategoryQueryServer) ById(context.Context, *CategoryParam) (*CategoryResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ById not implemented")
}
func (UnimplementedCategoryQueryServer) mustEmbedUnimplementedCategoryQueryServer() {}

// UnsafeCategoryQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryQueryServer will
// result in compilation errors.
type UnsafeCategoryQueryServer interface {
	mustEmbedUnimplementedCategoryQueryServer()
}

func RegisterCategoryQueryServer(s grpc.ServiceRegistrar, srv CategoryQueryServer) {
	s.RegisterService(&CategoryQuery_ServiceDesc, srv)
}

func _CategoryQuery_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryQueryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CategoryQuery/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryQueryServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryQuery_ById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryQueryServer).ById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CategoryQuery/ById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryQueryServer).ById(ctx, req.(*CategoryParam))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryQuery_ServiceDesc is the grpc.ServiceDesc for CategoryQuery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryQuery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CategoryQuery",
	HandlerType: (*CategoryQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CategoryQuery_List_Handler,
		},
		{
			MethodName: "ById",
			Handler:    _CategoryQuery_ById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/query.proto",
}

// ProductQueryClient is the client API for ProductQuery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductQueryClient interface {
	// すべての商品を取得する
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ProductsResult, error)
	// 指定されたIDの商品を取得する
	ById(ctx context.Context, in *ProductParam, opts ...grpc.CallOption) (*ProductResult, error)
	// 指定されたキーワードを含む商品を取得する
	ByKeyword(ctx context.Context, in *ProductParam, opts ...grpc.CallOption) (*ProductsResult, error)
}

type productQueryClient struct {
	cc grpc.ClientConnInterface
}

func NewProductQueryClient(cc grpc.ClientConnInterface) ProductQueryClient {
	return &productQueryClient{cc}
}

func (c *productQueryClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ProductsResult, error) {
	out := new(ProductsResult)
	err := c.cc.Invoke(ctx, "/proto.ProductQuery/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productQueryClient) ById(ctx context.Context, in *ProductParam, opts ...grpc.CallOption) (*ProductResult, error) {
	out := new(ProductResult)
	err := c.cc.Invoke(ctx, "/proto.ProductQuery/ById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productQueryClient) ByKeyword(ctx context.Context, in *ProductParam, opts ...grpc.CallOption) (*ProductsResult, error) {
	out := new(ProductsResult)
	err := c.cc.Invoke(ctx, "/proto.ProductQuery/ByKeyword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductQueryServer is the server API for ProductQuery service.
// All implementations must embed UnimplementedProductQueryServer
// for forward compatibility
type ProductQueryServer interface {
	// すべての商品を取得する
	List(context.Context, *emptypb.Empty) (*ProductsResult, error)
	// 指定されたIDの商品を取得する
	ById(context.Context, *ProductParam) (*ProductResult, error)
	// 指定されたキーワードを含む商品を取得する
	ByKeyword(context.Context, *ProductParam) (*ProductsResult, error)
	mustEmbedUnimplementedProductQueryServer()
}

// UnimplementedProductQueryServer must be embedded to have forward compatible implementations.
type UnimplementedProductQueryServer struct {
}

func (UnimplementedProductQueryServer) List(context.Context, *emptypb.Empty) (*ProductsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProductQueryServer) ById(context.Context, *ProductParam) (*ProductResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ById not implemented")
}
func (UnimplementedProductQueryServer) ByKeyword(context.Context, *ProductParam) (*ProductsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByKeyword not implemented")
}
func (UnimplementedProductQueryServer) mustEmbedUnimplementedProductQueryServer() {}

// UnsafeProductQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductQueryServer will
// result in compilation errors.
type UnsafeProductQueryServer interface {
	mustEmbedUnimplementedProductQueryServer()
}

func RegisterProductQueryServer(s grpc.ServiceRegistrar, srv ProductQueryServer) {
	s.RegisterService(&ProductQuery_ServiceDesc, srv)
}

func _ProductQuery_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductQueryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductQuery/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductQueryServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductQuery_ById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductQueryServer).ById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductQuery/ById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductQueryServer).ById(ctx, req.(*ProductParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductQuery_ByKeyword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductQueryServer).ByKeyword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductQuery/ByKeyword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductQueryServer).ByKeyword(ctx, req.(*ProductParam))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductQuery_ServiceDesc is the grpc.ServiceDesc for ProductQuery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductQuery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductQuery",
	HandlerType: (*ProductQueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ProductQuery_List_Handler,
		},
		{
			MethodName: "ById",
			Handler:    _ProductQuery_ById_Handler,
		},
		{
			MethodName: "ByKeyword",
			Handler:    _ProductQuery_ByKeyword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/query.proto",
}
