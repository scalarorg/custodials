// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/protocol/v1beta1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/scalarorg/scalar-core/x/chains/btc/types"
	types1 "github.com/scalarorg/scalar-core/x/chains/types"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	github_com_scalarorg_scalar_core_x_nexus_exported "github.com/scalarorg/scalar-core/x/nexus/exported"
	exported "github.com/scalarorg/scalar-core/x/protocol/exported"
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

type Status int32

const (
	Unspecified Status = 0
	Activated   Status = 1
	Deactivated Status = 2
)

var Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ACTIVATED",
	2: "STATUS_DEACTIVATED",
}

var Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_ACTIVATED":   1,
	"STATUS_DEACTIVATED": 2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d53a37c7b7ae195, []int{0}
}

type ProtocolAttribute struct {
	Model exported.LiquidityModel `protobuf:"varint,1,opt,name=model,proto3,enum=scalar.protocol.exported.v1beta1.LiquidityModel" json:"model,omitempty"`
}

func (m *ProtocolAttribute) Reset()         { *m = ProtocolAttribute{} }
func (m *ProtocolAttribute) String() string { return proto.CompactTextString(m) }
func (*ProtocolAttribute) ProtoMessage()    {}
func (*ProtocolAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d53a37c7b7ae195, []int{0}
}
func (m *ProtocolAttribute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolAttribute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolAttribute.Merge(m, src)
}
func (m *ProtocolAttribute) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolAttribute proto.InternalMessageInfo

func (m *ProtocolAttribute) GetModel() exported.LiquidityModel {
	if m != nil {
		return m.Model
	}
	return exported.Pooling
}

// DestinationChain represents a blockchain where tokens can be sent
type SupportedChain struct {
	Chain   github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	Name    string                                                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string                                                      `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *SupportedChain) Reset()         { *m = SupportedChain{} }
func (m *SupportedChain) String() string { return proto.CompactTextString(m) }
func (*SupportedChain) ProtoMessage()    {}
func (*SupportedChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d53a37c7b7ae195, []int{1}
}
func (m *SupportedChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupportedChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupportedChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupportedChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedChain.Merge(m, src)
}
func (m *SupportedChain) XXX_Size() int {
	return m.Size()
}
func (m *SupportedChain) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedChain.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedChain proto.InternalMessageInfo

func (m *SupportedChain) GetChain() github_com_scalarorg_scalar_core_x_nexus_exported.ChainName {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *SupportedChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SupportedChain) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Protocol struct {
	BitcoinPubkey  []byte                `protobuf:"bytes,1,opt,name=bitcoin_pubkey,json=bitcoinPubkey,proto3" json:"bitcoin_pubkey,omitempty"`
	ScalarPubkey   []byte                `protobuf:"bytes,2,opt,name=scalar_pubkey,json=scalarPubkey,proto3" json:"scalar_pubkey,omitempty"`
	ScalarAddress  []byte                `protobuf:"bytes,3,opt,name=scalar_address,json=scalarAddress,proto3" json:"scalar_address,omitempty"`
	Name           string                `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Tag            []byte                `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
	Attribute      *ProtocolAttribute    `protobuf:"bytes,6,opt,name=attribute,proto3" json:"attribute,omitempty"`
	Status         Status                `protobuf:"varint,7,opt,name=status,proto3,enum=scalar.protocol.v1beta1.Status" json:"status,omitempty"`
	CustodianGroup *types.CustodianGroup `protobuf:"bytes,8,opt,name=custodian_group,json=custodianGroup,proto3" json:"custodian_group,omitempty"`
	Asset          *types1.Asset         `protobuf:"bytes,9,opt,name=asset,proto3" json:"asset,omitempty"`
	Chains         []*SupportedChain     `protobuf:"bytes,10,rep,name=chains,proto3" json:"chains,omitempty"`
}

func (m *Protocol) Reset()         { *m = Protocol{} }
func (m *Protocol) String() string { return proto.CompactTextString(m) }
func (*Protocol) ProtoMessage()    {}
func (*Protocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d53a37c7b7ae195, []int{2}
}
func (m *Protocol) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Protocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Protocol.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Protocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Protocol.Merge(m, src)
}
func (m *Protocol) XXX_Size() int {
	return m.Size()
}
func (m *Protocol) XXX_DiscardUnknown() {
	xxx_messageInfo_Protocol.DiscardUnknown(m)
}

var xxx_messageInfo_Protocol proto.InternalMessageInfo

func (m *Protocol) GetBitcoinPubkey() []byte {
	if m != nil {
		return m.BitcoinPubkey
	}
	return nil
}

func (m *Protocol) GetScalarPubkey() []byte {
	if m != nil {
		return m.ScalarPubkey
	}
	return nil
}

func (m *Protocol) GetScalarAddress() []byte {
	if m != nil {
		return m.ScalarAddress
	}
	return nil
}

func (m *Protocol) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Protocol) GetTag() []byte {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *Protocol) GetAttribute() *ProtocolAttribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *Protocol) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Unspecified
}

func (m *Protocol) GetCustodianGroup() *types.CustodianGroup {
	if m != nil {
		return m.CustodianGroup
	}
	return nil
}

func (m *Protocol) GetAsset() *types1.Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *Protocol) GetChains() []*SupportedChain {
	if m != nil {
		return m.Chains
	}
	return nil
}

func init() {
	proto.RegisterEnum("scalar.protocol.v1beta1.Status", Status_name, Status_value)
	proto.RegisterType((*ProtocolAttribute)(nil), "scalar.protocol.v1beta1.ProtocolAttribute")
	proto.RegisterType((*SupportedChain)(nil), "scalar.protocol.v1beta1.SupportedChain")
	proto.RegisterType((*Protocol)(nil), "scalar.protocol.v1beta1.Protocol")
}

func init() {
	proto.RegisterFile("scalar/protocol/v1beta1/types.proto", fileDescriptor_1d53a37c7b7ae195)
}

var fileDescriptor_1d53a37c7b7ae195 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x4f, 0xdb, 0x3e,
	0x14, 0xc0, 0x1b, 0x4a, 0x0b, 0x35, 0x50, 0xfa, 0xb5, 0xbe, 0xd2, 0xa2, 0x6a, 0x0a, 0x5d, 0x11,
	0x02, 0xa1, 0x2d, 0x19, 0xec, 0xb0, 0xc3, 0x0e, 0x28, 0xb4, 0x65, 0x43, 0xdb, 0x50, 0x95, 0xb6,
	0x3b, 0x6c, 0x07, 0xe4, 0x38, 0x5e, 0xb1, 0xa0, 0x71, 0x16, 0x3b, 0x08, 0xfe, 0x80, 0x49, 0x13,
	0xa7, 0x5d, 0x76, 0xe4, 0xb4, 0x7f, 0x66, 0x47, 0x8e, 0x3b, 0x4d, 0x13, 0x9c, 0xf7, 0x0f, 0xec,
	0x34, 0xc5, 0x8e, 0x43, 0xf9, 0xa5, 0xed, 0xe6, 0xbc, 0x7c, 0xde, 0xc7, 0xef, 0xd9, 0xcf, 0x60,
	0x91, 0x63, 0x74, 0x80, 0x62, 0x27, 0x8a, 0x99, 0x60, 0x98, 0x1d, 0x38, 0x87, 0x6b, 0x3e, 0x11,
	0x68, 0xcd, 0x11, 0xc7, 0x11, 0xe1, 0xb6, 0x0c, 0xc3, 0x7b, 0x0a, 0xb2, 0x35, 0x64, 0x67, 0x50,
	0xfd, 0xff, 0x21, 0x1b, 0x32, 0x19, 0x75, 0xd2, 0x95, 0x02, 0xea, 0x4b, 0x99, 0x13, 0xef, 0x21,
	0x1a, 0x72, 0xc7, 0x17, 0xf8, 0x36, 0x6b, 0xbd, 0x79, 0x15, 0xd3, 0x48, 0x84, 0x62, 0x34, 0xfa,
	0x0b, 0x23, 0xd8, 0x3e, 0x09, 0x35, 0xf3, 0xe0, 0x0e, 0x66, 0x6c, 0x2b, 0xdd, 0x25, 0x66, 0x87,
	0x24, 0x44, 0xa1, 0xb8, 0x15, 0x7a, 0x78, 0xfd, 0x28, 0xc8, 0x51, 0xc4, 0x62, 0x41, 0x82, 0xdb,
	0xe8, 0xe6, 0x3b, 0xf0, 0x5f, 0x37, 0x03, 0x5d, 0x21, 0x62, 0xea, 0x27, 0x82, 0xc0, 0x2d, 0x50,
	0x1a, 0xb1, 0x80, 0x1c, 0x98, 0x46, 0xc3, 0x58, 0xa9, 0xae, 0x3f, 0xb6, 0xaf, 0x1f, 0x9c, 0x56,
	0xea, 0x13, 0xb4, 0x5f, 0xd1, 0x0f, 0x09, 0x0d, 0xa8, 0x38, 0x7e, 0x9d, 0xe6, 0x79, 0x2a, 0xbd,
	0xf9, 0xc5, 0x00, 0xd5, 0x5e, 0x12, 0x29, 0xb6, 0x95, 0xf6, 0x05, 0x07, 0xa0, 0x24, 0x1b, 0x94,
	0xea, 0xca, 0xe6, 0xc6, 0xef, 0x1f, 0x0b, 0xcf, 0x86, 0x54, 0xec, 0x25, 0xbe, 0x8d, 0xd9, 0xc8,
	0x51, 0x1b, 0xb1, 0x78, 0x98, 0xad, 0x1e, 0x61, 0x16, 0x13, 0xe7, 0xc8, 0x09, 0xc9, 0x51, 0xc2,
	0xf3, 0x4e, 0x6c, 0xe9, 0xda, 0x41, 0x23, 0xe2, 0x29, 0x1b, 0x84, 0x60, 0x32, 0x44, 0x23, 0x62,
	0x4e, 0xa4, 0x56, 0x4f, 0xae, 0xa1, 0x09, 0xa6, 0x50, 0x10, 0xc4, 0x84, 0x73, 0xb3, 0x28, 0xc3,
	0xfa, 0xb3, 0xf9, 0xab, 0x08, 0xa6, 0x75, 0xd7, 0x70, 0x09, 0x54, 0x7d, 0x2a, 0x30, 0xa3, 0xe1,
	0x6e, 0x94, 0xf8, 0xfb, 0xe4, 0x58, 0x96, 0x36, 0xeb, 0xcd, 0x65, 0xd1, 0xae, 0x0c, 0xc2, 0x45,
	0x30, 0xa7, 0x4a, 0xd2, 0xd4, 0x84, 0xa4, 0x66, 0x55, 0x30, 0x83, 0x96, 0x40, 0x35, 0x83, 0xc6,
	0x77, 0x9e, 0xf5, 0xb2, 0x54, 0x57, 0x05, 0xf3, 0x6a, 0x27, 0xc7, 0xaa, 0xad, 0x81, 0xa2, 0x40,
	0x43, 0xb3, 0x24, 0xf9, 0x74, 0x09, 0x5f, 0x80, 0x0a, 0xd2, 0x57, 0x62, 0x96, 0x1b, 0xc6, 0xca,
	0xcc, 0xfa, 0xaa, 0x7d, 0xc7, 0x08, 0xdb, 0x37, 0x2e, 0xd1, 0xbb, 0x4c, 0x86, 0x4f, 0x41, 0x99,
	0x0b, 0x24, 0x12, 0x6e, 0x4e, 0xc9, 0x0b, 0x5d, 0xb8, 0x53, 0xd3, 0x93, 0x98, 0x97, 0xe1, 0xb0,
	0x0b, 0xe6, 0x71, 0xc2, 0x05, 0x0b, 0x28, 0x0a, 0x77, 0x87, 0x31, 0x4b, 0x22, 0x73, 0x5a, 0x16,
	0xb2, 0xac, 0x0d, 0x7a, 0x14, 0x73, 0x43, 0x4b, 0xf3, 0xcf, 0x53, 0xdc, 0xab, 0xe2, 0x2b, 0xdf,
	0x70, 0x1d, 0x94, 0x10, 0xe7, 0x44, 0x98, 0x15, 0xe9, 0xb9, 0x9f, 0x7b, 0xe4, 0xd4, 0xe7, 0x16,
	0x37, 0x65, 0x3c, 0x85, 0xc2, 0x0d, 0x50, 0x56, 0xbf, 0x4d, 0xd0, 0x28, 0x8e, 0x6f, 0x7e, 0xb3,
	0xfc, 0x2b, 0xc3, 0xe6, 0x65, 0x69, 0xab, 0x1f, 0x0d, 0x50, 0x56, 0x9d, 0xc1, 0x65, 0x00, 0x7b,
	0x7d, 0xb7, 0x3f, 0xe8, 0xed, 0x0e, 0x76, 0x7a, 0xdd, 0x4e, 0x6b, 0x7b, 0x6b, 0xbb, 0xd3, 0xae,
	0x15, 0xea, 0xf3, 0x27, 0xa7, 0x8d, 0x99, 0x41, 0xc8, 0x23, 0x82, 0xe9, 0x7b, 0x4a, 0x02, 0xb8,
	0x08, 0x6a, 0x19, 0xe8, 0xb6, 0xfa, 0xdb, 0x6f, 0xdc, 0x7e, 0xa7, 0x5d, 0x33, 0xea, 0x73, 0x27,
	0xa7, 0x8d, 0x8a, 0x8b, 0x05, 0x3d, 0x44, 0x82, 0x04, 0x63, 0xb6, 0x76, 0xe7, 0x12, 0x9b, 0x50,
	0xb6, 0x36, 0x41, 0x1a, 0xac, 0x4f, 0x7e, 0xfa, 0x6a, 0x15, 0x36, 0x5f, 0x7e, 0x3b, 0xb7, 0x8c,
	0xb3, 0x73, 0xcb, 0xf8, 0x79, 0x6e, 0x19, 0x9f, 0x2f, 0xac, 0xc2, 0xd9, 0x85, 0x55, 0xf8, 0x7e,
	0x61, 0x15, 0xde, 0xae, 0xfd, 0xc3, 0x1b, 0xc8, 0x1f, 0xb4, 0x7c, 0xbf, 0x7e, 0x59, 0x7e, 0x3f,
	0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x13, 0x9a, 0xd5, 0xfb, 0x04, 0x00, 0x00,
}

func (m *ProtocolAttribute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolAttribute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolAttribute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Model != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Model))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SupportedChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupportedChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupportedChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Protocol) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Protocol) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Protocol) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.Asset != nil {
		{
			size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.CustodianGroup != nil {
		{
			size, err := m.CustodianGroup.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if m.Attribute != nil {
		{
			size, err := m.Attribute.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.Tag) > 0 {
		i -= len(m.Tag)
		copy(dAtA[i:], m.Tag)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tag)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ScalarAddress) > 0 {
		i -= len(m.ScalarAddress)
		copy(dAtA[i:], m.ScalarAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ScalarAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ScalarPubkey) > 0 {
		i -= len(m.ScalarPubkey)
		copy(dAtA[i:], m.ScalarPubkey)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ScalarPubkey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BitcoinPubkey) > 0 {
		i -= len(m.BitcoinPubkey)
		copy(dAtA[i:], m.BitcoinPubkey)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BitcoinPubkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProtocolAttribute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Model != 0 {
		n += 1 + sovTypes(uint64(m.Model))
	}
	return n
}

func (m *SupportedChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Protocol) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BitcoinPubkey)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ScalarPubkey)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ScalarAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Tag)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Attribute != nil {
		l = m.Attribute.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovTypes(uint64(m.Status))
	}
	if m.CustodianGroup != nil {
		l = m.CustodianGroup.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Asset != nil {
		l = m.Asset.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Chains) > 0 {
		for _, e := range m.Chains {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProtocolAttribute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ProtocolAttribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolAttribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Model", wireType)
			}
			m.Model = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Model |= exported.LiquidityModel(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *SupportedChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SupportedChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupportedChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Protocol) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Protocol: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Protocol: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitcoinPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BitcoinPubkey = append(m.BitcoinPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.BitcoinPubkey == nil {
				m.BitcoinPubkey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScalarPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ScalarPubkey = append(m.ScalarPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.ScalarPubkey == nil {
				m.ScalarPubkey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScalarAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ScalarAddress = append(m.ScalarAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ScalarAddress == nil {
				m.ScalarAddress = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = append(m.Tag[:0], dAtA[iNdEx:postIndex]...)
			if m.Tag == nil {
				m.Tag = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attribute", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attribute == nil {
				m.Attribute = &ProtocolAttribute{}
			}
			if err := m.Attribute.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustodianGroup", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CustodianGroup == nil {
				m.CustodianGroup = &types.CustodianGroup{}
			}
			if err := m.CustodianGroup.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Asset == nil {
				m.Asset = &types1.Asset{}
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, &SupportedChain{})
			if err := m.Chains[len(m.Chains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
