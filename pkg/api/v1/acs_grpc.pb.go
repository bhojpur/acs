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

// AcsServiceClient is the client API for AcsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AcsServiceClient interface {
	// StartLocalAccount starts a Account on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the acs/config.yaml
	//   3. all bytes constituting the Account YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalAccount(ctx context.Context, opts ...grpc.CallOption) (AcsService_StartLocalAccountClient, error)
	// StartFromPreviousAccount starts a new Account based on a previous one.
	// If the previous Account does not have the can-replay condition set this call will result in an error.
	StartFromPreviousAccount(ctx context.Context, in *StartFromPreviousAccountRequest, opts ...grpc.CallOption) (*StartAccountResponse, error)
	// StartAccountRequest starts a new Account based on its specification.
	StartAccount(ctx context.Context, in *StartAccountRequest, opts ...grpc.CallOption) (*StartAccountResponse, error)
	// Searches for Account(s) known to this instance
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	// Subscribe listens to new Account(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (AcsService_SubscribeClient, error)
	// GetAccount retrieves details of a single Account
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	// Listen listens to Account updates and log output of a running Account
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (AcsService_ListenClient, error)
	// StopAccount stops a currently running Account
	StopAccount(ctx context.Context, in *StopAccountRequest, opts ...grpc.CallOption) (*StopAccountResponse, error)
}

type acsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAcsServiceClient(cc grpc.ClientConnInterface) AcsServiceClient {
	return &acsServiceClient{cc}
}

func (c *acsServiceClient) StartLocalAccount(ctx context.Context, opts ...grpc.CallOption) (AcsService_StartLocalAccountClient, error) {
	stream, err := c.cc.NewStream(ctx, &AcsService_ServiceDesc.Streams[0], "/v1.AcsService/StartLocalAccount", opts...)
	if err != nil {
		return nil, err
	}
	x := &acsServiceStartLocalAccountClient{stream}
	return x, nil
}

type AcsService_StartLocalAccountClient interface {
	Send(*StartLocalAccountRequest) error
	CloseAndRecv() (*StartAccountResponse, error)
	grpc.ClientStream
}

type acsServiceStartLocalAccountClient struct {
	grpc.ClientStream
}

func (x *acsServiceStartLocalAccountClient) Send(m *StartLocalAccountRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *acsServiceStartLocalAccountClient) CloseAndRecv() (*StartAccountResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartAccountResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *acsServiceClient) StartFromPreviousAccount(ctx context.Context, in *StartFromPreviousAccountRequest, opts ...grpc.CallOption) (*StartAccountResponse, error) {
	out := new(StartAccountResponse)
	err := c.cc.Invoke(ctx, "/v1.AcsService/StartFromPreviousAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acsServiceClient) StartAccount(ctx context.Context, in *StartAccountRequest, opts ...grpc.CallOption) (*StartAccountResponse, error) {
	out := new(StartAccountResponse)
	err := c.cc.Invoke(ctx, "/v1.AcsService/StartAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acsServiceClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, "/v1.AcsService/ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acsServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (AcsService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &AcsService_ServiceDesc.Streams[1], "/v1.AcsService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &acsServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AcsService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type acsServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *acsServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *acsServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/v1.AcsService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *acsServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (AcsService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &AcsService_ServiceDesc.Streams[2], "/v1.AcsService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &acsServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AcsService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type acsServiceListenClient struct {
	grpc.ClientStream
}

func (x *acsServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *acsServiceClient) StopAccount(ctx context.Context, in *StopAccountRequest, opts ...grpc.CallOption) (*StopAccountResponse, error) {
	out := new(StopAccountResponse)
	err := c.cc.Invoke(ctx, "/v1.AcsService/StopAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AcsServiceServer is the server API for AcsService service.
// All implementations must embed UnimplementedAcsServiceServer
// for forward compatibility
type AcsServiceServer interface {
	// StartLocalAccount starts a Account on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the acs/config.yaml
	//   3. all bytes constituting the Account YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalAccount(AcsService_StartLocalAccountServer) error
	// StartFromPreviousAccount starts a new Account based on a previous one.
	// If the previous Account does not have the can-replay condition set this call will result in an error.
	StartFromPreviousAccount(context.Context, *StartFromPreviousAccountRequest) (*StartAccountResponse, error)
	// StartAccountRequest starts a new Account based on its specification.
	StartAccount(context.Context, *StartAccountRequest) (*StartAccountResponse, error)
	// Searches for Account(s) known to this instance
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	// Subscribe listens to new Account(s) updates
	Subscribe(*SubscribeRequest, AcsService_SubscribeServer) error
	// GetAccount retrieves details of a single Account
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	// Listen listens to Account updates and log output of a running Account
	Listen(*ListenRequest, AcsService_ListenServer) error
	// StopAccount stops a currently running Account
	StopAccount(context.Context, *StopAccountRequest) (*StopAccountResponse, error)
	mustEmbedUnimplementedAcsServiceServer()
}

// UnimplementedAcsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAcsServiceServer struct {
}

func (UnimplementedAcsServiceServer) StartLocalAccount(AcsService_StartLocalAccountServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalAccount not implemented")
}
func (UnimplementedAcsServiceServer) StartFromPreviousAccount(context.Context, *StartFromPreviousAccountRequest) (*StartAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousAccount not implemented")
}
func (UnimplementedAcsServiceServer) StartAccount(context.Context, *StartAccountRequest) (*StartAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartAccount not implemented")
}
func (UnimplementedAcsServiceServer) ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedAcsServiceServer) Subscribe(*SubscribeRequest, AcsService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedAcsServiceServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAcsServiceServer) Listen(*ListenRequest, AcsService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedAcsServiceServer) StopAccount(context.Context, *StopAccountRequest) (*StopAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopAccount not implemented")
}
func (UnimplementedAcsServiceServer) mustEmbedUnimplementedAcsServiceServer() {}

// UnsafeAcsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AcsServiceServer will
// result in compilation errors.
type UnsafeAcsServiceServer interface {
	mustEmbedUnimplementedAcsServiceServer()
}

func RegisterAcsServiceServer(s grpc.ServiceRegistrar, srv AcsServiceServer) {
	s.RegisterService(&AcsService_ServiceDesc, srv)
}

func _AcsService_StartLocalAccount_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AcsServiceServer).StartLocalAccount(&acsServiceStartLocalAccountServer{stream})
}

type AcsService_StartLocalAccountServer interface {
	SendAndClose(*StartAccountResponse) error
	Recv() (*StartLocalAccountRequest, error)
	grpc.ServerStream
}

type acsServiceStartLocalAccountServer struct {
	grpc.ServerStream
}

func (x *acsServiceStartLocalAccountServer) SendAndClose(m *StartAccountResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *acsServiceStartLocalAccountServer) Recv() (*StartLocalAccountRequest, error) {
	m := new(StartLocalAccountRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AcsService_StartFromPreviousAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcsServiceServer).StartFromPreviousAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AcsService/StartFromPreviousAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcsServiceServer).StartFromPreviousAccount(ctx, req.(*StartFromPreviousAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcsService_StartAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcsServiceServer).StartAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AcsService/StartAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcsServiceServer).StartAccount(ctx, req.(*StartAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcsService_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcsServiceServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AcsService/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcsServiceServer).ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcsService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AcsServiceServer).Subscribe(m, &acsServiceSubscribeServer{stream})
}

type AcsService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type acsServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *acsServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AcsService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcsServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AcsService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcsServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AcsService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AcsServiceServer).Listen(m, &acsServiceListenServer{stream})
}

type AcsService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type acsServiceListenServer struct {
	grpc.ServerStream
}

func (x *acsServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AcsService_StopAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AcsServiceServer).StopAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AcsService/StopAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AcsServiceServer).StopAccount(ctx, req.(*StopAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AcsService_ServiceDesc is the grpc.ServiceDesc for AcsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AcsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AcsService",
	HandlerType: (*AcsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousAccount",
			Handler:    _AcsService_StartFromPreviousAccount_Handler,
		},
		{
			MethodName: "StartAccount",
			Handler:    _AcsService_StartAccount_Handler,
		},
		{
			MethodName: "ListAccounts",
			Handler:    _AcsService_ListAccounts_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AcsService_GetAccount_Handler,
		},
		{
			MethodName: "StopAccount",
			Handler:    _AcsService_StopAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalAccount",
			Handler:       _AcsService_StartLocalAccount_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _AcsService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _AcsService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "acs.proto",
}