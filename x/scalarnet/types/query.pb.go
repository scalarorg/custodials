// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/scalarnet/v1beta1/query.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_scalarorg_scalar_core_x_nexus_exported "github.com/scalarorg/scalar-core/x/nexus/exported"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PendingIBCTransferCountRequest struct {
}

func (m *PendingIBCTransferCountRequest) Reset()         { *m = PendingIBCTransferCountRequest{} }
func (m *PendingIBCTransferCountRequest) String() string { return proto.CompactTextString(m) }
func (*PendingIBCTransferCountRequest) ProtoMessage()    {}
func (*PendingIBCTransferCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dfbde35d33aceb6, []int{0}
}
func (m *PendingIBCTransferCountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingIBCTransferCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingIBCTransferCountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingIBCTransferCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingIBCTransferCountRequest.Merge(m, src)
}
func (m *PendingIBCTransferCountRequest) XXX_Size() int {
	return m.Size()
}
func (m *PendingIBCTransferCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingIBCTransferCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PendingIBCTransferCountRequest proto.InternalMessageInfo

type PendingIBCTransferCountResponse struct {
	TransfersByChain map[string]uint32 `protobuf:"bytes,1,rep,name=transfers_by_chain,json=transfersByChain,proto3" json:"transfers_by_chain" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *PendingIBCTransferCountResponse) Reset()         { *m = PendingIBCTransferCountResponse{} }
func (m *PendingIBCTransferCountResponse) String() string { return proto.CompactTextString(m) }
func (*PendingIBCTransferCountResponse) ProtoMessage()    {}
func (*PendingIBCTransferCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dfbde35d33aceb6, []int{1}
}
func (m *PendingIBCTransferCountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingIBCTransferCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingIBCTransferCountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingIBCTransferCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingIBCTransferCountResponse.Merge(m, src)
}
func (m *PendingIBCTransferCountResponse) XXX_Size() int {
	return m.Size()
}
func (m *PendingIBCTransferCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingIBCTransferCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PendingIBCTransferCountResponse proto.InternalMessageInfo

// ParamsRequest represents a message that queries the params
type ParamsRequest struct {
}

func (m *ParamsRequest) Reset()         { *m = ParamsRequest{} }
func (m *ParamsRequest) String() string { return proto.CompactTextString(m) }
func (*ParamsRequest) ProtoMessage()    {}
func (*ParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dfbde35d33aceb6, []int{2}
}
func (m *ParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsRequest.Merge(m, src)
}
func (m *ParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsRequest proto.InternalMessageInfo

type ParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *ParamsResponse) Reset()         { *m = ParamsResponse{} }
func (m *ParamsResponse) String() string { return proto.CompactTextString(m) }
func (*ParamsResponse) ProtoMessage()    {}
func (*ParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dfbde35d33aceb6, []int{3}
}
func (m *ParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsResponse.Merge(m, src)
}
func (m *ParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsResponse proto.InternalMessageInfo

// IBCPathRequest represents a message that queries the IBC path registered for
// a given chain
type IBCPathRequest struct {
	Chain string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (m *IBCPathRequest) Reset()         { *m = IBCPathRequest{} }
func (m *IBCPathRequest) String() string { return proto.CompactTextString(m) }
func (*IBCPathRequest) ProtoMessage()    {}
func (*IBCPathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dfbde35d33aceb6, []int{4}
}
func (m *IBCPathRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCPathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCPathRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCPathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCPathRequest.Merge(m, src)
}
func (m *IBCPathRequest) XXX_Size() int {
	return m.Size()
}
func (m *IBCPathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCPathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IBCPathRequest proto.InternalMessageInfo

type IBCPathResponse struct {
	IBCPath string `protobuf:"bytes,1,opt,name=ibc_path,json=ibcPath,proto3" json:"ibc_path,omitempty"`
}

func (m *IBCPathResponse) Reset()         { *m = IBCPathResponse{} }
func (m *IBCPathResponse) String() string { return proto.CompactTextString(m) }
func (*IBCPathResponse) ProtoMessage()    {}
func (*IBCPathResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dfbde35d33aceb6, []int{5}
}
func (m *IBCPathResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCPathResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCPathResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCPathResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCPathResponse.Merge(m, src)
}
func (m *IBCPathResponse) XXX_Size() int {
	return m.Size()
}
func (m *IBCPathResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCPathResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IBCPathResponse proto.InternalMessageInfo

// ChainByIBCPathRequest represents a message that queries the chain that an IBC
// path is registered to
type ChainByIBCPathRequest struct {
	IbcPath string `protobuf:"bytes,1,opt,name=ibc_path,json=ibcPath,proto3" json:"ibc_path,omitempty"`
}

func (m *ChainByIBCPathRequest) Reset()         { *m = ChainByIBCPathRequest{} }
func (m *ChainByIBCPathRequest) String() string { return proto.CompactTextString(m) }
func (*ChainByIBCPathRequest) ProtoMessage()    {}
func (*ChainByIBCPathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dfbde35d33aceb6, []int{6}
}
func (m *ChainByIBCPathRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainByIBCPathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainByIBCPathRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainByIBCPathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainByIBCPathRequest.Merge(m, src)
}
func (m *ChainByIBCPathRequest) XXX_Size() int {
	return m.Size()
}
func (m *ChainByIBCPathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainByIBCPathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChainByIBCPathRequest proto.InternalMessageInfo

type ChainByIBCPathResponse struct {
	Chain github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
}

func (m *ChainByIBCPathResponse) Reset()         { *m = ChainByIBCPathResponse{} }
func (m *ChainByIBCPathResponse) String() string { return proto.CompactTextString(m) }
func (*ChainByIBCPathResponse) ProtoMessage()    {}
func (*ChainByIBCPathResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dfbde35d33aceb6, []int{7}
}
func (m *ChainByIBCPathResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainByIBCPathResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainByIBCPathResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainByIBCPathResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainByIBCPathResponse.Merge(m, src)
}
func (m *ChainByIBCPathResponse) XXX_Size() int {
	return m.Size()
}
func (m *ChainByIBCPathResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainByIBCPathResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChainByIBCPathResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PendingIBCTransferCountRequest)(nil), "scalar.scalarnet.v1beta1.PendingIBCTransferCountRequest")
	proto.RegisterType((*PendingIBCTransferCountResponse)(nil), "scalar.scalarnet.v1beta1.PendingIBCTransferCountResponse")
	proto.RegisterMapType((map[string]uint32)(nil), "scalar.scalarnet.v1beta1.PendingIBCTransferCountResponse.TransfersByChainEntry")
	proto.RegisterType((*ParamsRequest)(nil), "scalar.scalarnet.v1beta1.ParamsRequest")
	proto.RegisterType((*ParamsResponse)(nil), "scalar.scalarnet.v1beta1.ParamsResponse")
	proto.RegisterType((*IBCPathRequest)(nil), "scalar.scalarnet.v1beta1.IBCPathRequest")
	proto.RegisterType((*IBCPathResponse)(nil), "scalar.scalarnet.v1beta1.IBCPathResponse")
	proto.RegisterType((*ChainByIBCPathRequest)(nil), "scalar.scalarnet.v1beta1.ChainByIBCPathRequest")
	proto.RegisterType((*ChainByIBCPathResponse)(nil), "scalar.scalarnet.v1beta1.ChainByIBCPathResponse")
}

func init() {
	proto.RegisterFile("scalar/scalarnet/v1beta1/query.proto", fileDescriptor_6dfbde35d33aceb6)
}

var fileDescriptor_6dfbde35d33aceb6 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xbd, 0x2d, 0xb4, 0xb0, 0x51, 0x3f, 0x64, 0xb5, 0x28, 0xe4, 0xe0, 0x58, 0x16, 0x54,
	0x11, 0x12, 0xb6, 0x1a, 0x2e, 0x7c, 0x48, 0x20, 0xd9, 0xe2, 0xd0, 0x0b, 0x58, 0x11, 0x5c, 0xb8,
	0x44, 0x6b, 0x77, 0x70, 0x2c, 0x9a, 0x5d, 0x77, 0x77, 0x5d, 0xc5, 0x67, 0x5e, 0x80, 0xc7, 0xca,
	0xb1, 0x47, 0x4e, 0x11, 0x24, 0x57, 0x9e, 0x80, 0x13, 0xf2, 0xee, 0x3a, 0x15, 0x11, 0x06, 0x4e,
	0x9e, 0xdd, 0xfd, 0xcd, 0xff, 0x3f, 0x33, 0x1a, 0xe3, 0x07, 0x22, 0x25, 0x17, 0x84, 0x07, 0xfa,
	0x43, 0x41, 0x06, 0x57, 0xa7, 0x09, 0x48, 0x72, 0x1a, 0x5c, 0x96, 0xc0, 0x2b, 0xbf, 0xe0, 0x4c,
	0x32, 0xbb, 0xab, 0x9f, 0xfd, 0x35, 0xe5, 0x1b, 0xaa, 0x77, 0x94, 0xb1, 0x8c, 0x29, 0x28, 0xa8,
	0x23, 0xcd, 0xf7, 0xda, 0x55, 0x65, 0x55, 0x80, 0x30, 0xd4, 0xa3, 0x94, 0x89, 0x29, 0x13, 0x41,
	0x42, 0x04, 0x68, 0xbb, 0x35, 0x56, 0x90, 0x2c, 0xa7, 0x44, 0xe6, 0x8c, 0x1a, 0xf6, 0x61, 0xab,
	0x62, 0x41, 0x38, 0x99, 0x1a, 0x49, 0xcf, 0xc5, 0x4e, 0x0c, 0xf4, 0x3c, 0xa7, 0xd9, 0x59, 0x18,
	0xbd, 0xe3, 0x84, 0x8a, 0x8f, 0xc0, 0x23, 0x56, 0x52, 0x39, 0x82, 0xcb, 0x12, 0x84, 0xf4, 0x7e,
	0x20, 0xdc, 0x6f, 0x45, 0x44, 0xc1, 0xa8, 0x00, 0xfb, 0x33, 0xc2, 0xb6, 0x34, 0x2f, 0x62, 0x9c,
	0x54, 0xe3, 0x74, 0x42, 0x72, 0xda, 0x45, 0xee, 0xf6, 0xa0, 0x33, 0x7c, 0xeb, 0xb7, 0x0d, 0xc3,
	0xff, 0x87, 0xae, 0xdf, 0xdc, 0x8a, 0xb0, 0x8a, 0x6a, 0xc5, 0xd7, 0x54, 0xf2, 0x2a, 0xbc, 0x35,
	0x5f, 0xf4, 0xad, 0xd1, 0xa1, 0xdc, 0x78, 0xec, 0x45, 0xf8, 0xf8, 0x8f, 0x09, 0xf6, 0x21, 0xde,
	0xfe, 0x04, 0x55, 0x17, 0xb9, 0x68, 0x70, 0x77, 0x54, 0x87, 0xf6, 0x11, 0xbe, 0x7d, 0x45, 0x2e,
	0x4a, 0xe8, 0x6e, 0xb9, 0x68, 0xb0, 0x37, 0xd2, 0x87, 0xe7, 0x5b, 0x4f, 0x91, 0x77, 0x80, 0xf7,
	0x62, 0x35, 0xa0, 0xa6, 0xff, 0x18, 0xef, 0x37, 0x17, 0xa6, 0xdb, 0x97, 0x78, 0x47, 0xcf, 0x50,
	0x29, 0x76, 0x86, 0xee, 0x5f, 0x1a, 0x54, 0x9c, 0xa9, 0xd8, 0x64, 0x79, 0x27, 0x78, 0xff, 0x2c,
	0x8c, 0x62, 0x22, 0x27, 0xc6, 0xa3, 0x2e, 0xa7, 0x99, 0x58, 0x5d, 0xa2, 0x3e, 0x78, 0xcf, 0xf0,
	0xc1, 0x9a, 0x33, 0xd6, 0x27, 0xf8, 0x4e, 0x9e, 0xa4, 0xe3, 0x82, 0xc8, 0x89, 0x66, 0xc3, 0xce,
	0x72, 0xd1, 0xdf, 0x6d, 0xb0, 0xdd, 0x3c, 0x49, 0xeb, 0xc0, 0x1b, 0xe2, 0x63, 0xd5, 0x7f, 0x58,
	0x6d, 0x38, 0xdd, 0xdf, 0x14, 0xb8, 0xc9, 0x61, 0xf8, 0xde, 0x66, 0x8e, 0x71, 0x7d, 0xff, 0x5b,
	0x79, 0xe1, 0xab, 0x9f, 0x8b, 0xfe, 0x8b, 0x2c, 0x97, 0x93, 0x32, 0xf1, 0x53, 0x36, 0x35, 0x2b,
	0xc6, 0x78, 0x66, 0xa2, 0xc7, 0x29, 0xe3, 0x10, 0xcc, 0x02, 0x0a, 0xb3, 0x52, 0x04, 0x30, 0x2b,
	0x18, 0x97, 0x70, 0xee, 0x2b, 0xf1, 0x37, 0x64, 0x0a, 0xa6, 0xbf, 0x30, 0x9e, 0x7f, 0x77, 0xac,
	0xf9, 0xd2, 0x41, 0xd7, 0x4b, 0x07, 0x7d, 0x5b, 0x3a, 0xe8, 0xcb, 0xca, 0xb1, 0xae, 0x57, 0x8e,
	0xf5, 0x75, 0xe5, 0x58, 0x1f, 0x86, 0xff, 0xe1, 0x70, 0xb3, 0xdc, 0xea, 0x37, 0x49, 0x76, 0xd4,
	0x52, 0x3f, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x97, 0xe8, 0x03, 0xa5, 0x03, 0x00, 0x00,
}

func (m *PendingIBCTransferCountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingIBCTransferCountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingIBCTransferCountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PendingIBCTransferCountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingIBCTransferCountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingIBCTransferCountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TransfersByChain) > 0 {
		for k := range m.TransfersByChain {
			v := m.TransfersByChain[k]
			baseI := i
			i = encodeVarintQuery(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintQuery(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintQuery(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IBCPathRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCPathRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCPathRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IBCPathResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCPathResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCPathResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IBCPath) > 0 {
		i -= len(m.IBCPath)
		copy(dAtA[i:], m.IBCPath)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.IBCPath)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChainByIBCPathRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainByIBCPathRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainByIBCPathRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcPath) > 0 {
		i -= len(m.IbcPath)
		copy(dAtA[i:], m.IbcPath)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.IbcPath)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChainByIBCPathResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainByIBCPathResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainByIBCPathResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PendingIBCTransferCountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PendingIBCTransferCountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TransfersByChain) > 0 {
		for k, v := range m.TransfersByChain {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovQuery(uint64(len(k))) + 1 + sovQuery(uint64(v))
			n += mapEntrySize + 1 + sovQuery(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *IBCPathRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *IBCPathResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IBCPath)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ChainByIBCPathRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IbcPath)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ChainByIBCPathResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PendingIBCTransferCountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PendingIBCTransferCountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingIBCTransferCountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PendingIBCTransferCountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PendingIBCTransferCountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingIBCTransferCountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransfersByChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TransfersByChain == nil {
				m.TransfersByChain = make(map[string]uint32)
			}
			var mapkey string
			var mapvalue uint32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQuery
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthQuery
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQuery(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthQuery
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.TransfersByChain[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IBCPathRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IBCPathRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCPathRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IBCPathResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IBCPathResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCPathResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IBCPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IBCPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChainByIBCPathRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainByIBCPathRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainByIBCPathRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChainByIBCPathResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainByIBCPathResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainByIBCPathResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
