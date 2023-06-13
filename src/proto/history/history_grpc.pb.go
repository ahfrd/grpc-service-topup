// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: src/proto/history/history.proto

package history

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

// HistoryServiceClient is the client API for HistoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoryServiceClient interface {
	InsertHistoryLog(ctx context.Context, in *InsertHistoryLogRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	UpdateStatusLog(ctx context.Context, in *UpdateStatusLogRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
}

type historyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryServiceClient(cc grpc.ClientConnInterface) HistoryServiceClient {
	return &historyServiceClient{cc}
}

func (c *historyServiceClient) InsertHistoryLog(ctx context.Context, in *InsertHistoryLogRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/history.HistoryService/InsertHistoryLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyServiceClient) UpdateStatusLog(ctx context.Context, in *UpdateStatusLogRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/history.HistoryService/UpdateStatusLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryServiceServer is the server API for HistoryService service.
// All implementations must embed UnimplementedHistoryServiceServer
// for forward compatibility
type HistoryServiceServer interface {
	InsertHistoryLog(context.Context, *InsertHistoryLogRequest) (*GeneralResponse, error)
	UpdateStatusLog(context.Context, *UpdateStatusLogRequest) (*GeneralResponse, error)
	mustEmbedUnimplementedHistoryServiceServer()
}

// UnimplementedHistoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHistoryServiceServer struct {
}

func (UnimplementedHistoryServiceServer) InsertHistoryLog(context.Context, *InsertHistoryLogRequest) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertHistoryLog not implemented")
}
func (UnimplementedHistoryServiceServer) UpdateStatusLog(context.Context, *UpdateStatusLogRequest) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatusLog not implemented")
}
func (UnimplementedHistoryServiceServer) mustEmbedUnimplementedHistoryServiceServer() {}

// UnsafeHistoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoryServiceServer will
// result in compilation errors.
type UnsafeHistoryServiceServer interface {
	mustEmbedUnimplementedHistoryServiceServer()
}

func RegisterHistoryServiceServer(s grpc.ServiceRegistrar, srv HistoryServiceServer) {
	s.RegisterService(&HistoryService_ServiceDesc, srv)
}

func _HistoryService_InsertHistoryLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertHistoryLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServiceServer).InsertHistoryLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/history.HistoryService/InsertHistoryLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServiceServer).InsertHistoryLog(ctx, req.(*InsertHistoryLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoryService_UpdateStatusLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServiceServer).UpdateStatusLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/history.HistoryService/UpdateStatusLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServiceServer).UpdateStatusLog(ctx, req.(*UpdateStatusLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HistoryService_ServiceDesc is the grpc.ServiceDesc for HistoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HistoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "history.HistoryService",
	HandlerType: (*HistoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertHistoryLog",
			Handler:    _HistoryService_InsertHistoryLog_Handler,
		},
		{
			MethodName: "UpdateStatusLog",
			Handler:    _HistoryService_UpdateStatusLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/history/history.proto",
}
