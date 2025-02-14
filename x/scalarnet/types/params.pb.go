// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/scalarnet/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params represent the genesis parameters for the module
type Params struct {
	// IBC packet route timeout window
	RouteTimeoutWindow               uint64                          `protobuf:"varint,1,opt,name=route_timeout_window,json=routeTimeoutWindow,proto3" json:"route_timeout_window,omitempty"`
	TransferLimit                    uint64                          `protobuf:"varint,2,opt,name=transfer_limit,json=transferLimit,proto3" json:"transfer_limit,omitempty"`
	EndBlockerLimit                  uint64                          `protobuf:"varint,3,opt,name=end_blocker_limit,json=endBlockerLimit,proto3" json:"end_blocker_limit,omitempty"`
	Version                          uint32                          `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Tag                              []byte                          `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
	CallContractsProposalMinDeposits CallContractProposalMinDeposits `protobuf:"bytes,6,rep,name=call_contracts_proposal_min_deposits,json=callContractsProposalMinDeposits,proto3,castrepeated=CallContractProposalMinDeposits" json:"call_contracts_proposal_min_deposits"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_68386da21503c6a6, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

type CallContractProposalMinDeposit struct {
	Chain           github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	ContractAddress string                                                      `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	MinDeposits     github_com_cosmos_cosmos_sdk_types.Coins                    `protobuf:"bytes,3,rep,name=min_deposits,json=minDeposits,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_deposits"`
}

func (m *CallContractProposalMinDeposit) Reset()         { *m = CallContractProposalMinDeposit{} }
func (m *CallContractProposalMinDeposit) String() string { return proto.CompactTextString(m) }
func (*CallContractProposalMinDeposit) ProtoMessage()    {}
func (*CallContractProposalMinDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_68386da21503c6a6, []int{1}
}
func (m *CallContractProposalMinDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CallContractProposalMinDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CallContractProposalMinDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CallContractProposalMinDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallContractProposalMinDeposit.Merge(m, src)
}
func (m *CallContractProposalMinDeposit) XXX_Size() int {
	return m.Size()
}
func (m *CallContractProposalMinDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_CallContractProposalMinDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_CallContractProposalMinDeposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "scalar.scalarnet.v1beta1.Params")
	proto.RegisterType((*CallContractProposalMinDeposit)(nil), "scalar.scalarnet.v1beta1.CallContractProposalMinDeposit")
}

func init() {
	proto.RegisterFile("scalar/scalarnet/v1beta1/params.proto", fileDescriptor_68386da21503c6a6)
}

var fileDescriptor_68386da21503c6a6 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xd6, 0xad, 0x68, 0xde, 0xc6, 0x86, 0xb5, 0x43, 0xd8, 0x21, 0x8d, 0x26, 0x26, 0x02,
	0xd2, 0x92, 0x6d, 0x5c, 0x90, 0x38, 0x20, 0x52, 0x8e, 0x80, 0xaa, 0x08, 0x84, 0xc4, 0x25, 0x72,
	0x1c, 0x93, 0x59, 0x4b, 0xfc, 0x22, 0xdb, 0xdd, 0xca, 0xbf, 0x40, 0xfc, 0x0c, 0x6e, 0xfc, 0x8b,
	0x1e, 0x77, 0xe4, 0x34, 0xa0, 0x3d, 0xf3, 0x07, 0x38, 0xa1, 0xd8, 0x69, 0x3b, 0x24, 0x98, 0x38,
	0xf9, 0xe5, 0x7b, 0xdf, 0xfb, 0x5e, 0xf2, 0xe5, 0x33, 0x3a, 0x50, 0x94, 0x94, 0x44, 0x46, 0xf6,
	0x10, 0x4c, 0x47, 0xe7, 0xc7, 0x19, 0xd3, 0xe4, 0x38, 0xaa, 0x89, 0x24, 0x95, 0x0a, 0x6b, 0x09,
	0x1a, 0xb0, 0x6b, 0xfb, 0xe1, 0x82, 0x16, 0xb6, 0xb4, 0x3d, 0x8f, 0x82, 0xaa, 0x40, 0x45, 0x19,
	0x51, 0x6c, 0x31, 0x4b, 0x81, 0x0b, 0x3b, 0xb9, 0xb7, 0x5b, 0x40, 0x01, 0xa6, 0x8c, 0x9a, 0xca,
	0xa2, 0xfb, 0x3f, 0x57, 0x50, 0x6f, 0x68, 0x16, 0xe0, 0x23, 0xb4, 0x2b, 0x61, 0xa4, 0x59, 0xaa,
	0x79, 0xc5, 0x60, 0xa4, 0xd3, 0x0b, 0x2e, 0x72, 0xb8, 0x70, 0x1d, 0xdf, 0x09, 0x56, 0x13, 0x6c,
	0x7a, 0xaf, 0x6d, 0xeb, 0xad, 0xe9, 0xe0, 0x03, 0x74, 0x5b, 0x4b, 0x22, 0xd4, 0x7b, 0x26, 0xd3,
	0x92, 0x57, 0x5c, 0xbb, 0x2b, 0x86, 0xbb, 0x35, 0x47, 0x5f, 0x34, 0x20, 0x7e, 0x88, 0xee, 0x30,
	0x91, 0xa7, 0x59, 0x09, 0xf4, 0x6c, 0xc1, 0xec, 0x1a, 0xe6, 0x36, 0x13, 0x79, 0x6c, 0x71, 0xcb,
	0x75, 0xd1, 0xad, 0x73, 0x26, 0x15, 0x07, 0xe1, 0xae, 0xfa, 0x4e, 0xb0, 0x95, 0xcc, 0x1f, 0xf1,
	0x0e, 0xea, 0x6a, 0x52, 0xb8, 0x6b, 0xbe, 0x13, 0x6c, 0x26, 0x4d, 0x89, 0xbf, 0x38, 0xe8, 0x1e,
	0x25, 0x65, 0x99, 0x52, 0x10, 0x5a, 0x12, 0xaa, 0x55, 0x5a, 0x4b, 0xa8, 0x41, 0x91, 0x32, 0xad,
	0xb8, 0x48, 0x73, 0x56, 0x83, 0xe2, 0x5a, 0xb9, 0x3d, 0xbf, 0x1b, 0x6c, 0x9c, 0x3c, 0x0e, 0xff,
	0xe5, 0x5d, 0x38, 0x20, 0x65, 0x39, 0x68, 0x45, 0x86, 0xad, 0xc4, 0x4b, 0x2e, 0x9e, 0x5b, 0x81,
	0xf8, 0xfe, 0xe4, 0xaa, 0xdf, 0xf9, 0xfc, 0xad, 0xdf, 0xbf, 0x99, 0xa7, 0x12, 0x9f, 0x5e, 0x23,
	0xa8, 0xbf, 0x30, 0xf6, 0x3f, 0xad, 0x20, 0xef, 0x66, 0x15, 0xfc, 0x06, 0xad, 0xd1, 0x53, 0xc2,
	0x85, 0x31, 0x7e, 0x3d, 0x7e, 0xfa, 0xeb, 0xaa, 0xff, 0xa4, 0xe0, 0xfa, 0x74, 0x94, 0x85, 0x14,
	0xaa, 0x36, 0x20, 0x20, 0x8b, 0xb6, 0x3a, 0xa4, 0x20, 0x59, 0x34, 0x8e, 0x04, 0x1b, 0x8f, 0x54,
	0xc4, 0xc6, 0x35, 0x48, 0xcd, 0xf2, 0x70, 0xd0, 0x48, 0xbc, 0x22, 0x15, 0x4b, 0xac, 0x1a, 0x7e,
	0x80, 0x76, 0xe6, 0x3e, 0xa5, 0x24, 0xcf, 0x25, 0x53, 0xca, 0xfc, 0xae, 0xf5, 0x64, 0x7b, 0x8e,
	0x3f, 0xb3, 0x30, 0x16, 0x68, 0xf3, 0x0f, 0xff, 0xba, 0xc6, 0xbf, 0xbb, 0xa1, 0x4d, 0x58, 0xd8,
	0x24, 0x6c, 0x69, 0x1d, 0x70, 0x11, 0x1f, 0xb5, 0x06, 0x05, 0xd7, 0xde, 0xb3, 0x8d, 0xa3, 0x3d,
	0x0e, 0x55, 0x7e, 0x16, 0xe9, 0x0f, 0x35, 0x53, 0x66, 0x40, 0x25, 0x1b, 0xd5, 0xd2, 0x94, 0x78,
	0x38, 0xf9, 0xe1, 0x75, 0x26, 0x53, 0xcf, 0xb9, 0x9c, 0x7a, 0xce, 0xf7, 0xa9, 0xe7, 0x7c, 0x9c,
	0x79, 0x9d, 0xcb, 0x99, 0xd7, 0xf9, 0x3a, 0xf3, 0x3a, 0xef, 0x4e, 0xfe, 0xe3, 0xe3, 0x97, 0xb7,
	0xc6, 0x2c, 0xc9, 0x7a, 0x26, 0xdd, 0x8f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x19, 0x94, 0x33,
	0xc8, 0x56, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CallContractsProposalMinDeposits) > 0 {
		for iNdEx := len(m.CallContractsProposalMinDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CallContractsProposalMinDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Tag) > 0 {
		i -= len(m.Tag)
		copy(dAtA[i:], m.Tag)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Tag)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Version != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x20
	}
	if m.EndBlockerLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EndBlockerLimit))
		i--
		dAtA[i] = 0x18
	}
	if m.TransferLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TransferLimit))
		i--
		dAtA[i] = 0x10
	}
	if m.RouteTimeoutWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RouteTimeoutWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CallContractProposalMinDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallContractProposalMinDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CallContractProposalMinDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinDeposits) > 0 {
		for iNdEx := len(m.MinDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RouteTimeoutWindow != 0 {
		n += 1 + sovParams(uint64(m.RouteTimeoutWindow))
	}
	if m.TransferLimit != 0 {
		n += 1 + sovParams(uint64(m.TransferLimit))
	}
	if m.EndBlockerLimit != 0 {
		n += 1 + sovParams(uint64(m.EndBlockerLimit))
	}
	if m.Version != 0 {
		n += 1 + sovParams(uint64(m.Version))
	}
	l = len(m.Tag)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.CallContractsProposalMinDeposits) > 0 {
		for _, e := range m.CallContractsProposalMinDeposits {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *CallContractProposalMinDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.MinDeposits) > 0 {
		for _, e := range m.MinDeposits {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteTimeoutWindow", wireType)
			}
			m.RouteTimeoutWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RouteTimeoutWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferLimit", wireType)
			}
			m.TransferLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransferLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockerLimit", wireType)
			}
			m.EndBlockerLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockerLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field CallContractsProposalMinDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallContractsProposalMinDeposits = append(m.CallContractsProposalMinDeposits, CallContractProposalMinDeposit{})
			if err := m.CallContractsProposalMinDeposits[len(m.CallContractsProposalMinDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *CallContractProposalMinDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: CallContractProposalMinDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallContractProposalMinDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinDeposits = append(m.MinDeposits, types.Coin{})
			if err := m.MinDeposits[len(m.MinDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
