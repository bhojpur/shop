// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ShopUIClient is the client API for ShopUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopUIClient interface {
	// ListBranchSpecs returns a list of Branch that can be started through the UI.
	ListBranchSpecs(ctx context.Context, in *ListBranchSpecsRequest, opts ...grpc.CallOption) (ShopUI_ListBranchSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type shopUIClient struct {
	cc grpc.ClientConnInterface
}

func NewShopUIClient(cc grpc.ClientConnInterface) ShopUIClient {
	return &shopUIClient{cc}
}

func (c *shopUIClient) ListBranchSpecs(ctx context.Context, in *ListBranchSpecsRequest, opts ...grpc.CallOption) (ShopUI_ListBranchSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ShopUI_ServiceDesc.Streams[0], "/v1.ShopUI/ListBranchSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &shopUIListBranchSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShopUI_ListBranchSpecsClient interface {
	Recv() (*ListBranchSpecsResponse, error)
	grpc.ClientStream
}

type shopUIListBranchSpecsClient struct {
	grpc.ClientStream
}

func (x *shopUIListBranchSpecsClient) Recv() (*ListBranchSpecsResponse, error) {
	m := new(ListBranchSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shopUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.ShopUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopUIServer is the server API for ShopUI service.
// All implementations must embed UnimplementedShopUIServer
// for forward compatibility
type ShopUIServer interface {
	// ListBranchSpecs returns a list of Branch that can be started through the UI.
	ListBranchSpecs(*ListBranchSpecsRequest, ShopUI_ListBranchSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedShopUIServer()
}

// UnimplementedShopUIServer must be embedded to have forward compatible implementations.
type UnimplementedShopUIServer struct {
}

func (UnimplementedShopUIServer) ListBranchSpecs(*ListBranchSpecsRequest, ShopUI_ListBranchSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBranchSpecs not implemented")
}
func (UnimplementedShopUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedShopUIServer) mustEmbedUnimplementedShopUIServer() {}

// UnsafeShopUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopUIServer will
// result in compilation errors.
type UnsafeShopUIServer interface {
	mustEmbedUnimplementedShopUIServer()
}

func RegisterShopUIServer(s grpc.ServiceRegistrar, srv ShopUIServer) {
	s.RegisterService(&ShopUI_ServiceDesc, srv)
}

func _ShopUI_ListBranchSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBranchSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShopUIServer).ListBranchSpecs(m, &shopUIListBranchSpecsServer{stream})
}

type ShopUI_ListBranchSpecsServer interface {
	Send(*ListBranchSpecsResponse) error
	grpc.ServerStream
}

type shopUIListBranchSpecsServer struct {
	grpc.ServerStream
}

func (x *shopUIListBranchSpecsServer) Send(m *ListBranchSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ShopUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ShopUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopUI_ServiceDesc is the grpc.ServiceDesc for ShopUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ShopUI",
	HandlerType: (*ShopUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _ShopUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBranchSpecs",
			Handler:       _ShopUI_ListBranchSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "shop-ui.proto",
}
