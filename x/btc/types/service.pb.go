// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/btc/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("scalar/btc/v1beta1/service.proto", fileDescriptor_6827246c7cbdb413) }
func init() {
	golang_proto.RegisterFile("scalar/btc/v1beta1/service.proto", fileDescriptor_6827246c7cbdb413)
}

var fileDescriptor_6827246c7cbdb413 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4b, 0x33, 0x31,
	0x18, 0xc7, 0x9b, 0x77, 0x78, 0x87, 0xe3, 0x85, 0x17, 0x0f, 0xa7, 0x2a, 0x51, 0x8a, 0x53, 0x6d,
	0x2f, 0xb4, 0x9d, 0x74, 0x6c, 0x07, 0x41, 0x70, 0xf0, 0xc7, 0xe4, 0x52, 0x72, 0x69, 0x4c, 0x0f,
	0x7a, 0x79, 0xae, 0x49, 0x5a, 0xaf, 0x94, 0x2e, 0xfe, 0x05, 0x82, 0xab, 0xab, 0x93, 0xbb, 0xb3,
	0xa3, 0x63, 0xc1, 0xc5, 0x51, 0x7a, 0xfe, 0x21, 0xd2, 0x4b, 0x0a, 0x62, 0x6f, 0xe8, 0xf6, 0xc0,
	0xf7, 0xf3, 0xe4, 0xfb, 0x21, 0x8f, 0xb7, 0xaf, 0x19, 0x1d, 0x50, 0x45, 0x42, 0xc3, 0xc8, 0xb8,
	0x11, 0x72, 0x43, 0x1b, 0x44, 0x73, 0x35, 0x8e, 0x18, 0x0f, 0x12, 0x05, 0x06, 0x7c, 0xdf, 0x12,
	0x41, 0x68, 0x58, 0xe0, 0x88, 0xf2, 0xb6, 0x00, 0x01, 0x79, 0x4c, 0x96, 0x93, 0x25, 0xcb, 0xbb,
	0x02, 0x40, 0x0c, 0x38, 0xa1, 0x49, 0x44, 0xa8, 0x94, 0x60, 0xa8, 0x89, 0x40, 0x6a, 0x97, 0xee,
	0x14, 0x34, 0x99, 0xd4, 0x85, 0xb8, 0x20, 0x1c, 0x8e, 0xb8, 0x9a, 0xd8, 0xbc, 0xf9, 0x8c, 0x3c,
	0xef, 0x4c, 0x8b, 0x4b, 0x6b, 0xe6, 0x3f, 0x22, 0x6f, 0xab, 0x03, 0xf2, 0x26, 0x52, 0xf1, 0x09,
	0x35, 0xfc, 0x96, 0x4e, 0xae, 0x52, 0xed, 0xd7, 0x82, 0x75, 0xd5, 0x60, 0x0d, 0xbb, 0xe0, 0xc3,
	0x11, 0xd7, 0xa6, 0x5c, 0xdf, 0x90, 0xd6, 0x09, 0x48, 0xcd, 0x2b, 0xd5, 0xbb, 0xf7, 0xaf, 0x87,
	0x3f, 0x07, 0x95, 0x3d, 0xf2, 0x43, 0x95, 0x59, 0xbc, 0x2b, 0x2c, 0xdf, 0x35, 0xa9, 0x3e, 0x46,
	0xd5, 0xe6, 0x0b, 0xf2, 0xfe, 0x9d, 0x2f, 0xed, 0x57, 0xbe, 0x4f, 0xc8, 0xfb, 0xdf, 0xa6, 0x86,
	0xf5, 0x79, 0xaf, 0x03, 0x71, 0x4c, 0x65, 0x4f, 0xfb, 0xd5, 0xa2, 0xfe, 0x5f, 0xd0, 0xca, 0xf5,
	0x70, 0x23, 0xd6, 0x99, 0x1e, 0xe5, 0xa6, 0x2d, 0xbf, 0x41, 0x0a, 0x3e, 0x35, 0xb4, 0x4b, 0x5d,
	0xe6, 0xb6, 0xc8, 0x94, 0xf5, 0x69, 0x24, 0x67, 0x64, 0x1a, 0xf5, 0x66, 0xed, 0xd3, 0xb7, 0x05,
	0x46, 0xf3, 0x05, 0x46, 0x9f, 0x0b, 0x8c, 0xee, 0x33, 0x5c, 0x7a, 0xcd, 0x30, 0x9a, 0x67, 0xb8,
	0xf4, 0x91, 0xe1, 0xd2, 0x75, 0x4d, 0x44, 0xa6, 0x3f, 0x0a, 0x03, 0x06, 0xb1, 0x7b, 0x1a, 0x94,
	0x70, 0x53, 0x9d, 0x81, 0xe2, 0x24, 0xcd, 0xbb, 0xcc, 0x24, 0xe1, 0x3a, 0xfc, 0x9b, 0x5f, 0xae,
	0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x59, 0xaa, 0x2f, 0x7e, 0x62, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	ConfirmGatewayTxs(ctx context.Context, in *ConfirmGatewayTxsRequest, opts ...grpc.CallOption) (*ConfirmGatewayTxsResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) ConfirmGatewayTxs(ctx context.Context, in *ConfirmGatewayTxsRequest, opts ...grpc.CallOption) (*ConfirmGatewayTxsResponse, error) {
	out := new(ConfirmGatewayTxsResponse)
	err := c.cc.Invoke(ctx, "/scalar.btc.v1beta1.MsgService/ConfirmGatewayTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	ConfirmGatewayTxs(context.Context, *ConfirmGatewayTxsRequest) (*ConfirmGatewayTxsResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) ConfirmGatewayTxs(ctx context.Context, req *ConfirmGatewayTxsRequest) (*ConfirmGatewayTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmGatewayTxs not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_ConfirmGatewayTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmGatewayTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConfirmGatewayTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.btc.v1beta1.MsgService/ConfirmGatewayTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConfirmGatewayTxs(ctx, req.(*ConfirmGatewayTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.btc.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfirmGatewayTxs",
			Handler:    _MsgService_ConfirmGatewayTxs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/btc/v1beta1/service.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// BatchedCommands queries the batched commands for a specified chain and
	// BatchedCommandsID if no BatchedCommandsID is specified, then it returns the
	// latest batched commands
	BatchedCommands(ctx context.Context, in *BatchedCommandsRequest, opts ...grpc.CallOption) (*BatchedCommandsResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) BatchedCommands(ctx context.Context, in *BatchedCommandsRequest, opts ...grpc.CallOption) (*BatchedCommandsResponse, error) {
	out := new(BatchedCommandsResponse)
	err := c.cc.Invoke(ctx, "/scalar.btc.v1beta1.QueryService/BatchedCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// BatchedCommands queries the batched commands for a specified chain and
	// BatchedCommandsID if no BatchedCommandsID is specified, then it returns the
	// latest batched commands
	BatchedCommands(context.Context, *BatchedCommandsRequest) (*BatchedCommandsResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) BatchedCommands(ctx context.Context, req *BatchedCommandsRequest) (*BatchedCommandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchedCommands not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_BatchedCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchedCommandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).BatchedCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.btc.v1beta1.QueryService/BatchedCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).BatchedCommands(ctx, req.(*BatchedCommandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.btc.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchedCommands",
			Handler:    _QueryService_BatchedCommands_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/btc/v1beta1/service.proto",
}
