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

// AcsUIClient is the client API for AcsUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AcsUIClient interface {
	// ListAccountSpecs returns a list of Account(s) that can be started through the UI.
	ListAccountSpecs(ctx context.Context, in *ListAccountSpecsRequest, opts ...grpc.CallOption) (AcsUI_ListAccountSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type acsUIClient struct {
	cc grpc.ClientConnInterface
}

func NewAcsUIClient(cc grpc.ClientConnInterface) AcsUIClient {
	return &acsUIClient{cc}
}

func (c *acsUIClient) ListAccountSpecs(ctx context.Context, in *ListAccountSpecsRequest, opts ...grpc.CallOption) (AcsUI_ListAccountSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AcsUI_ServiceDesc.Streams[0], "/v1.AcsUI/ListAccountSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &acsUIListAccountSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AcsUI_ListAccountSpecsClient interface {
	Recv() (*ListAccountSpecsResponse, error)
	grpc.ClientStream
}

type acsUIListAccountSpecsClient struct {
	grpc.ClientStream
}

func (x *acsUIListAccountSpecsClient) Recv() (*ListAccountSpecsResponse, error) {
	m := new(ListAccountSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *acsUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.AcsUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AcsUIServer is the server API for AcsUI service.
// All implementations must embed UnimplementedAcsUIServer
// for forward compatibility
type AcsUIServer interface {
	// ListAccountSpecs returns a list of Account(s) that can be started through the UI.
	ListAccountSpecs(*ListAccountSpecsRequest, AcsUI_ListAccountSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedAcsUIServer()
}

// UnimplementedAcsUIServer must be embedded to have forward compatible implementations.
type UnimplementedAcsUIServer struct {
}

func (UnimplementedAcsUIServer) ListAccountSpecs(*ListAccountSpecsRequest, AcsUI_ListAccountSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAccountSpecs not implemented")
}
func (UnimplementedAcsUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedAcsUIServer) mustEmbedUnimplementedAcsUIServer() {}

// UnsafeAcsUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AcsUIServer will
// result in compilation errors.
type UnsafeAcsUIServer interface {
	mustEmbedUnimplementedAcsUIServer()
}

func RegisterAcsUIServer(s grpc.ServiceRegistrar, srv AcsUIServer) {
	s.RegisterService(&AcsUI_ServiceDesc, srv)
}

func _AcsUI_ListAccountSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAccountSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AcsUIServer).ListAccountSpecs(m, &acsUIListAccountSpecsServer{stream})
}

type AcsUI_ListAccountSpecsServer interface {
	Send(*ListAccountSpecsResponse) error
	grpc.ServerStream
}

type acsUIListAccountSpecsServer struct {
	grpc.ServerStream
}

func (x *acsUIListAccountSpecsServer) Send(m *ListAccountSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AcsUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcsUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AcsUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcsUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AcsUI_ServiceDesc is the grpc.ServiceDesc for AcsUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AcsUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AcsUI",
	HandlerType: (*AcsUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _AcsUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAccountSpecs",
			Handler:       _AcsUI_ListAccountSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "acs-ui.proto",
}