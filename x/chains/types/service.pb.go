// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/chains/v1beta1/service.proto

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

func init() {
	proto.RegisterFile("scalar/chains/v1beta1/service.proto", fileDescriptor_bd3ac3d6d884da59)
}
func init() {
	golang_proto.RegisterFile("scalar/chains/v1beta1/service.proto", fileDescriptor_bd3ac3d6d884da59)
}

var fileDescriptor_bd3ac3d6d884da59 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0x33, 0x31,
	0x18, 0xc7, 0x9b, 0x77, 0x78, 0x87, 0xe3, 0x85, 0x57, 0x0e, 0x5d, 0x8a, 0x04, 0x5a, 0xc1, 0x41,
	0x6c, 0x42, 0xab, 0x38, 0x88, 0x53, 0x3b, 0x77, 0xd0, 0x3a, 0xb9, 0x94, 0x5c, 0x1a, 0xaf, 0x07,
	0xbd, 0x3c, 0xd7, 0x24, 0x57, 0xae, 0x94, 0x2e, 0x7e, 0x02, 0xc1, 0x4f, 0xe0, 0x17, 0x10, 0xbf,
	0x81, 0x8e, 0x8e, 0x05, 0x17, 0x47, 0xe9, 0xf9, 0x41, 0xa4, 0x97, 0x74, 0xb0, 0x9c, 0xd0, 0xed,
	0x81, 0xff, 0xef, 0x9f, 0xe7, 0x97, 0xc4, 0x3b, 0xd0, 0x9c, 0x8d, 0x98, 0xa2, 0x7c, 0xc8, 0x22,
	0xa9, 0xe9, 0xa4, 0x19, 0x08, 0xc3, 0x9a, 0x54, 0x0b, 0x35, 0x89, 0xb8, 0x20, 0x89, 0x02, 0x03,
	0xfe, 0x9e, 0x85, 0x88, 0x85, 0x88, 0x83, 0xaa, 0xbb, 0x21, 0x84, 0x50, 0x10, 0x74, 0x35, 0x59,
	0xb8, 0xba, 0x1f, 0x02, 0x84, 0x23, 0x41, 0x59, 0x12, 0x51, 0x26, 0x25, 0x18, 0x66, 0x22, 0x90,
	0xda, 0xa5, 0xb8, 0x7c, 0x9f, 0xc9, 0x5c, 0x5e, 0x2b, 0xcf, 0xc7, 0xa9, 0x50, 0x53, 0x8b, 0xb4,
	0x9e, 0x91, 0xe7, 0x75, 0x75, 0xd8, 0xb3, 0x8a, 0xfe, 0x23, 0xf2, 0x76, 0x3a, 0x20, 0x6f, 0x23,
	0x15, 0xf7, 0x20, 0x55, 0x5c, 0x5c, 0x67, 0xda, 0x27, 0xa4, 0x54, 0x99, 0x6c, 0x82, 0x57, 0x62,
	0x9c, 0x0a, 0x6d, 0xaa, 0x74, 0x6b, 0x5e, 0x27, 0x20, 0xb5, 0xa8, 0x1f, 0xdf, 0xbd, 0x7f, 0x3d,
	0xfc, 0x39, 0xac, 0xd7, 0xe8, 0x4f, 0x61, 0x6e, 0x0b, 0x7d, 0x5d, 0x34, 0xfa, 0x26, 0xd3, 0xe7,
	0xe8, 0xa8, 0xf5, 0x82, 0xbc, 0x7f, 0x97, 0xab, 0x2b, 0xac, 0xa5, 0x9f, 0x90, 0xf7, 0xbf, 0xcd,
	0x0c, 0x1f, 0x8a, 0x41, 0x07, 0xe2, 0x98, 0xc9, 0x81, 0xf6, 0x1b, 0xbf, 0x38, 0x6c, 0x70, 0x6b,
	0x65, 0xb2, 0x2d, 0xee, 0x8c, 0x2f, 0x0a, 0xe3, 0x33, 0xff, 0x94, 0x96, 0x3f, 0x71, 0x60, 0x7b,
	0x7d, 0xee, 0x8a, 0x74, 0x56, 0x00, 0x73, 0x3a, 0x8b, 0x06, 0xf3, 0x76, 0xf7, 0x6d, 0x89, 0xd1,
	0x62, 0x89, 0xd1, 0xe7, 0x12, 0xa3, 0xfb, 0x1c, 0x57, 0x5e, 0x73, 0x8c, 0x16, 0x39, 0xae, 0x7c,
	0xe4, 0xb8, 0x72, 0x43, 0xc3, 0xc8, 0x0c, 0xd3, 0x80, 0x70, 0x88, 0xdd, 0xe9, 0xa0, 0x42, 0x37,
	0x35, 0x38, 0x28, 0x41, 0xb3, 0xf5, 0x3a, 0x33, 0x4d, 0x84, 0x0e, 0xfe, 0x16, 0x5f, 0x79, 0xf2,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x6c, 0x54, 0x01, 0x7f, 0x02, 0x00, 0x00,
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
	ConfirmSourceTxs(ctx context.Context, in *ConfirmSourceTxsRequest, opts ...grpc.CallOption) (*ConfirmSourceTxsResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) ConfirmSourceTxs(ctx context.Context, in *ConfirmSourceTxsRequest, opts ...grpc.CallOption) (*ConfirmSourceTxsResponse, error) {
	out := new(ConfirmSourceTxsResponse)
	err := c.cc.Invoke(ctx, "/scalar.chains.v1beta1.MsgService/ConfirmSourceTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	ConfirmSourceTxs(context.Context, *ConfirmSourceTxsRequest) (*ConfirmSourceTxsResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) ConfirmSourceTxs(ctx context.Context, req *ConfirmSourceTxsRequest) (*ConfirmSourceTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmSourceTxs not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_ConfirmSourceTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmSourceTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConfirmSourceTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.chains.v1beta1.MsgService/ConfirmSourceTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConfirmSourceTxs(ctx, req.(*ConfirmSourceTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.chains.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfirmSourceTxs",
			Handler:    _MsgService_ConfirmSourceTxs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/chains/v1beta1/service.proto",
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
	err := c.cc.Invoke(ctx, "/scalar.chains.v1beta1.QueryService/BatchedCommands", in, out, opts...)
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
		FullMethod: "/scalar.chains.v1beta1.QueryService/BatchedCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).BatchedCommands(ctx, req.(*BatchedCommandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.chains.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchedCommands",
			Handler:    _QueryService_BatchedCommands_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/chains/v1beta1/service.proto",
}
