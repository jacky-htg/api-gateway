// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: users/feature_service.proto

package users

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
	FeatureService_List_FullMethodName = "/wiradata.users.FeatureService/List"
)

// FeatureServiceClient is the client API for FeatureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeatureServiceClient interface {
	List(ctx context.Context, in *MyEmpty, opts ...grpc.CallOption) (FeatureService_ListClient, error)
}

type featureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureServiceClient(cc grpc.ClientConnInterface) FeatureServiceClient {
	return &featureServiceClient{cc}
}

func (c *featureServiceClient) List(ctx context.Context, in *MyEmpty, opts ...grpc.CallOption) (FeatureService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &FeatureService_ServiceDesc.Streams[0], FeatureService_List_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &featureServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FeatureService_ListClient interface {
	Recv() (*ListFeatureResponse, error)
	grpc.ClientStream
}

type featureServiceListClient struct {
	grpc.ClientStream
}

func (x *featureServiceListClient) Recv() (*ListFeatureResponse, error) {
	m := new(ListFeatureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FeatureServiceServer is the server API for FeatureService service.
// All implementations must embed UnimplementedFeatureServiceServer
// for forward compatibility
type FeatureServiceServer interface {
	List(*MyEmpty, FeatureService_ListServer) error
	mustEmbedUnimplementedFeatureServiceServer()
}

// UnimplementedFeatureServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeatureServiceServer struct {
}

func (UnimplementedFeatureServiceServer) List(*MyEmpty, FeatureService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFeatureServiceServer) mustEmbedUnimplementedFeatureServiceServer() {}

// UnsafeFeatureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeatureServiceServer will
// result in compilation errors.
type UnsafeFeatureServiceServer interface {
	mustEmbedUnimplementedFeatureServiceServer()
}

func RegisterFeatureServiceServer(s grpc.ServiceRegistrar, srv FeatureServiceServer) {
	s.RegisterService(&FeatureService_ServiceDesc, srv)
}

func _FeatureService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MyEmpty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FeatureServiceServer).List(m, &featureServiceListServer{stream})
}

type FeatureService_ListServer interface {
	Send(*ListFeatureResponse) error
	grpc.ServerStream
}

type featureServiceListServer struct {
	grpc.ServerStream
}

func (x *featureServiceListServer) Send(m *ListFeatureResponse) error {
	return x.ServerStream.SendMsg(m)
}

// FeatureService_ServiceDesc is the grpc.ServiceDesc for FeatureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeatureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.users.FeatureService",
	HandlerType: (*FeatureServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _FeatureService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users/feature_service.proto",
}

const (
	PackageFeatureService_View_FullMethodName = "/wiradata.users.PackageFeatureService/View"
	PackageFeatureService_List_FullMethodName = "/wiradata.users.PackageFeatureService/List"
)

// PackageFeatureServiceClient is the client API for PackageFeatureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackageFeatureServiceClient interface {
	View(ctx context.Context, in *Id, opts ...grpc.CallOption) (*PackageOfFeature, error)
	List(ctx context.Context, in *MyEmpty, opts ...grpc.CallOption) (PackageFeatureService_ListClient, error)
}

type packageFeatureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageFeatureServiceClient(cc grpc.ClientConnInterface) PackageFeatureServiceClient {
	return &packageFeatureServiceClient{cc}
}

func (c *packageFeatureServiceClient) View(ctx context.Context, in *Id, opts ...grpc.CallOption) (*PackageOfFeature, error) {
	out := new(PackageOfFeature)
	err := c.cc.Invoke(ctx, PackageFeatureService_View_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageFeatureServiceClient) List(ctx context.Context, in *MyEmpty, opts ...grpc.CallOption) (PackageFeatureService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &PackageFeatureService_ServiceDesc.Streams[0], PackageFeatureService_List_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &packageFeatureServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PackageFeatureService_ListClient interface {
	Recv() (*ListPackageFeatureResponse, error)
	grpc.ClientStream
}

type packageFeatureServiceListClient struct {
	grpc.ClientStream
}

func (x *packageFeatureServiceListClient) Recv() (*ListPackageFeatureResponse, error) {
	m := new(ListPackageFeatureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PackageFeatureServiceServer is the server API for PackageFeatureService service.
// All implementations must embed UnimplementedPackageFeatureServiceServer
// for forward compatibility
type PackageFeatureServiceServer interface {
	View(context.Context, *Id) (*PackageOfFeature, error)
	List(*MyEmpty, PackageFeatureService_ListServer) error
	mustEmbedUnimplementedPackageFeatureServiceServer()
}

// UnimplementedPackageFeatureServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPackageFeatureServiceServer struct {
}

func (UnimplementedPackageFeatureServiceServer) View(context.Context, *Id) (*PackageOfFeature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method View not implemented")
}
func (UnimplementedPackageFeatureServiceServer) List(*MyEmpty, PackageFeatureService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPackageFeatureServiceServer) mustEmbedUnimplementedPackageFeatureServiceServer() {}

// UnsafePackageFeatureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackageFeatureServiceServer will
// result in compilation errors.
type UnsafePackageFeatureServiceServer interface {
	mustEmbedUnimplementedPackageFeatureServiceServer()
}

func RegisterPackageFeatureServiceServer(s grpc.ServiceRegistrar, srv PackageFeatureServiceServer) {
	s.RegisterService(&PackageFeatureService_ServiceDesc, srv)
}

func _PackageFeatureService_View_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageFeatureServiceServer).View(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PackageFeatureService_View_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageFeatureServiceServer).View(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageFeatureService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MyEmpty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PackageFeatureServiceServer).List(m, &packageFeatureServiceListServer{stream})
}

type PackageFeatureService_ListServer interface {
	Send(*ListPackageFeatureResponse) error
	grpc.ServerStream
}

type packageFeatureServiceListServer struct {
	grpc.ServerStream
}

func (x *packageFeatureServiceListServer) Send(m *ListPackageFeatureResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PackageFeatureService_ServiceDesc is the grpc.ServiceDesc for PackageFeatureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PackageFeatureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wiradata.users.PackageFeatureService",
	HandlerType: (*PackageFeatureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "View",
			Handler:    _PackageFeatureService_View_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _PackageFeatureService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users/feature_service.proto",
}
