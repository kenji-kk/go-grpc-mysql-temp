// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	GetAllCountry(ctx context.Context, in *AllCountryRequest, opts ...grpc.CallOption) (*AllCountryResponse, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) GetAllCountry(ctx context.Context, in *AllCountryRequest, opts ...grpc.CallOption) (*AllCountryResponse, error) {
	out := new(AllCountryResponse)
	err := c.cc.Invoke(ctx, "/api.AppService/GetAllCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations should embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	GetAllCountry(context.Context, *AllCountryRequest) (*AllCountryResponse, error)
}

// UnimplementedAppServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) GetAllCountry(context.Context, *AllCountryRequest) (*AllCountryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCountry not implemented")
}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_GetAllCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetAllCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AppService/GetAllCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetAllCountry(ctx, req.(*AllCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCountry",
			Handler:    _AppService_GetAllCountry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/pb/app.proto",
}
