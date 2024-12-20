// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: products.proto

package proto

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
	Products_GetProducts_FullMethodName      = "/products.Products/GetProducts"
	Products_GetOneProduct_FullMethodName    = "/products.Products/GetOneProduct"
	Products_PostProduct_FullMethodName      = "/products.Products/PostProduct"
	Products_DeleteProduct_FullMethodName    = "/products.Products/DeleteProduct"
	Products_GetSavedProducts_FullMethodName = "/products.Products/GetSavedProducts"
	Products_GetUserProducts_FullMethodName  = "/products.Products/GetUserProducts"
)

// ProductsClient is the client API for Products service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsClient interface {
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	GetOneProduct(ctx context.Context, in *GetOneProductRequest, opts ...grpc.CallOption) (*GetOneProductResponse, error)
	PostProduct(ctx context.Context, in *PostProductRequest, opts ...grpc.CallOption) (*PostProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	GetSavedProducts(ctx context.Context, in *GetSavedProductsRequest, opts ...grpc.CallOption) (*GetSavedProductResponse, error)
	GetUserProducts(ctx context.Context, in *GetUserProductsRequest, opts ...grpc.CallOption) (*GetUserProductsResponse, error)
}

type productsClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsClient(cc grpc.ClientConnInterface) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, Products_GetProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetOneProduct(ctx context.Context, in *GetOneProductRequest, opts ...grpc.CallOption) (*GetOneProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOneProductResponse)
	err := c.cc.Invoke(ctx, Products_GetOneProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) PostProduct(ctx context.Context, in *PostProductRequest, opts ...grpc.CallOption) (*PostProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostProductResponse)
	err := c.cc.Invoke(ctx, Products_PostProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, Products_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetSavedProducts(ctx context.Context, in *GetSavedProductsRequest, opts ...grpc.CallOption) (*GetSavedProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSavedProductResponse)
	err := c.cc.Invoke(ctx, Products_GetSavedProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsClient) GetUserProducts(ctx context.Context, in *GetUserProductsRequest, opts ...grpc.CallOption) (*GetUserProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserProductsResponse)
	err := c.cc.Invoke(ctx, Products_GetUserProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServer is the server API for Products service.
// All implementations must embed UnimplementedProductsServer
// for forward compatibility.
type ProductsServer interface {
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
	GetOneProduct(context.Context, *GetOneProductRequest) (*GetOneProductResponse, error)
	PostProduct(context.Context, *PostProductRequest) (*PostProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	GetSavedProducts(context.Context, *GetSavedProductsRequest) (*GetSavedProductResponse, error)
	GetUserProducts(context.Context, *GetUserProductsRequest) (*GetUserProductsResponse, error)
	mustEmbedUnimplementedProductsServer()
}

// UnimplementedProductsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductsServer struct{}

func (UnimplementedProductsServer) GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProducts not implemented")
}
func (UnimplementedProductsServer) GetOneProduct(context.Context, *GetOneProductRequest) (*GetOneProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneProduct not implemented")
}
func (UnimplementedProductsServer) PostProduct(context.Context, *PostProductRequest) (*PostProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostProduct not implemented")
}
func (UnimplementedProductsServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductsServer) GetSavedProducts(context.Context, *GetSavedProductsRequest) (*GetSavedProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSavedProducts not implemented")
}
func (UnimplementedProductsServer) GetUserProducts(context.Context, *GetUserProductsRequest) (*GetUserProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProducts not implemented")
}
func (UnimplementedProductsServer) mustEmbedUnimplementedProductsServer() {}
func (UnimplementedProductsServer) testEmbeddedByValue()                  {}

// UnsafeProductsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsServer will
// result in compilation errors.
type UnsafeProductsServer interface {
	mustEmbedUnimplementedProductsServer()
}

func RegisterProductsServer(s grpc.ServiceRegistrar, srv ProductsServer) {
	// If the following call pancis, it indicates UnimplementedProductsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Products_ServiceDesc, srv)
}

func _Products_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Products_GetProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetOneProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetOneProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Products_GetOneProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetOneProduct(ctx, req.(*GetOneProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_PostProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).PostProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Products_PostProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).PostProduct(ctx, req.(*PostProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Products_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetSavedProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSavedProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetSavedProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Products_GetSavedProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetSavedProducts(ctx, req.(*GetSavedProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Products_GetUserProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).GetUserProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Products_GetUserProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).GetUserProducts(ctx, req.(*GetUserProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Products_ServiceDesc is the grpc.ServiceDesc for Products service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Products_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "products.Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProducts",
			Handler:    _Products_GetProducts_Handler,
		},
		{
			MethodName: "GetOneProduct",
			Handler:    _Products_GetOneProduct_Handler,
		},
		{
			MethodName: "PostProduct",
			Handler:    _Products_PostProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Products_DeleteProduct_Handler,
		},
		{
			MethodName: "GetSavedProducts",
			Handler:    _Products_GetSavedProducts_Handler,
		},
		{
			MethodName: "GetUserProducts",
			Handler:    _Products_GetUserProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "products.proto",
}
