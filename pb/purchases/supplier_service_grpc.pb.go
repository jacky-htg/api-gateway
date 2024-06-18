// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: purchases/supplier_service.proto

package purchases

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SupplierService_SupplierCreate_FullMethodName = "/purchases.SupplierService/SupplierCreate"
	SupplierService_SupplierUpdate_FullMethodName = "/purchases.SupplierService/SupplierUpdate"
	SupplierService_SupplierView_FullMethodName   = "/purchases.SupplierService/SupplierView"
	SupplierService_SupplierList_FullMethodName   = "/purchases.SupplierService/SupplierList"
)

// SupplierServiceClient is the client API for SupplierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupplierServiceClient interface {
	SupplierCreate(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*Supplier, error)
	SupplierUpdate(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*Supplier, error)
	SupplierView(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Supplier, error)
	SupplierList(ctx context.Context, in *ListSupplierRequest, opts ...grpc.CallOption) (SupplierService_SupplierListClient, error)
}

type supplierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupplierServiceClient(cc grpc.ClientConnInterface) SupplierServiceClient {
	return &supplierServiceClient{cc}
}

func (c *supplierServiceClient) SupplierCreate(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, SupplierService_SupplierCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) SupplierUpdate(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, SupplierService_SupplierUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) SupplierView(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, SupplierService_SupplierView_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) SupplierList(ctx context.Context, in *ListSupplierRequest, opts ...grpc.CallOption) (SupplierService_SupplierListClient, error) {
	stream, err := c.cc.NewStream(ctx, &SupplierService_ServiceDesc.Streams[0], SupplierService_SupplierList_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &supplierServiceSupplierListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SupplierService_SupplierListClient interface {
	Recv() (*ListSupplierResponse, error)
	grpc.ClientStream
}

type supplierServiceSupplierListClient struct {
	grpc.ClientStream
}

func (x *supplierServiceSupplierListClient) Recv() (*ListSupplierResponse, error) {
	m := new(ListSupplierResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SupplierServiceServer is the server API for SupplierService service.
// All implementations must embed UnimplementedSupplierServiceServer
// for forward compatibility
type SupplierServiceServer interface {
	SupplierCreate(context.Context, *Supplier) (*Supplier, error)
	SupplierUpdate(context.Context, *Supplier) (*Supplier, error)
	SupplierView(context.Context, *Id) (*Supplier, error)
	SupplierList(*ListSupplierRequest, SupplierService_SupplierListServer) error
	mustEmbedUnimplementedSupplierServiceServer()
}

// UnimplementedSupplierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupplierServiceServer struct {
}

func (UnimplementedSupplierServiceServer) SupplierCreate(context.Context, *Supplier) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupplierCreate not implemented")
}
func (UnimplementedSupplierServiceServer) SupplierUpdate(context.Context, *Supplier) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupplierUpdate not implemented")
}
func (UnimplementedSupplierServiceServer) SupplierView(context.Context, *Id) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupplierView not implemented")
}
func (UnimplementedSupplierServiceServer) SupplierList(*ListSupplierRequest, SupplierService_SupplierListServer) error {
	return status.Errorf(codes.Unimplemented, "method SupplierList not implemented")
}
func (UnimplementedSupplierServiceServer) mustEmbedUnimplementedSupplierServiceServer() {}

// UnsafeSupplierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupplierServiceServer will
// result in compilation errors.
type UnsafeSupplierServiceServer interface {
	mustEmbedUnimplementedSupplierServiceServer()
}

func RegisterSupplierServiceServer(s grpc.ServiceRegistrar, srv SupplierServiceServer) {
	s.RegisterService(&SupplierService_ServiceDesc, srv)
}

func _SupplierService_SupplierCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supplier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).SupplierCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SupplierService_SupplierCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).SupplierCreate(ctx, req.(*Supplier))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_SupplierUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supplier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).SupplierUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SupplierService_SupplierUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).SupplierUpdate(ctx, req.(*Supplier))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_SupplierView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).SupplierView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SupplierService_SupplierView_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).SupplierView(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_SupplierList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListSupplierRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SupplierServiceServer).SupplierList(m, &supplierServiceSupplierListServer{stream})
}

type SupplierService_SupplierListServer interface {
	Send(*ListSupplierResponse) error
	grpc.ServerStream
}

type supplierServiceSupplierListServer struct {
	grpc.ServerStream
}

func (x *supplierServiceSupplierListServer) Send(m *ListSupplierResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SupplierService_ServiceDesc is the grpc.ServiceDesc for SupplierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupplierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "purchases.SupplierService",
	HandlerType: (*SupplierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SupplierCreate",
			Handler:    _SupplierService_SupplierCreate_Handler,
		},
		{
			MethodName: "SupplierUpdate",
			Handler:    _SupplierService_SupplierUpdate_Handler,
		},
		{
			MethodName: "SupplierView",
			Handler:    _SupplierService_SupplierView_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SupplierList",
			Handler:       _SupplierService_SupplierList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "purchases/supplier_service.proto",
}
