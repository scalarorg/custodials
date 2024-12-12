// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/vote/exported/v1beta1/types.proto

package exported

import (
	fmt "fmt"
	utils "github.com/axelarnetwork/axelar-core/utils"
	exported "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

type PollState int32

const (
	NonExistent PollState = 0
	Pending     PollState = 1
	Completed   PollState = 2
	Failed      PollState = 3
)

var PollState_name = map[int32]string{
	0: "POLL_STATE_UNSPECIFIED",
	1: "POLL_STATE_PENDING",
	2: "POLL_STATE_COMPLETED",
	3: "POLL_STATE_FAILED",
}

var PollState_value = map[string]int32{
	"POLL_STATE_UNSPECIFIED": 0,
	"POLL_STATE_PENDING":     1,
	"POLL_STATE_COMPLETED":   2,
	"POLL_STATE_FAILED":      3,
}

func (x PollState) String() string {
	return proto.EnumName(PollState_name, int32(x))
}

func (PollState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bcf266de9f0ebe6d, []int{0}
}

// PollMetadata represents a poll with write-in voting, i.e. the result of the
// vote can have any data type
type PollMetadata struct {
	ExpiresAt       int64             `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Result          *types.Any        `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	VotingThreshold utils.Threshold   `protobuf:"bytes,5,opt,name=voting_threshold,json=votingThreshold,proto3" json:"voting_threshold"`
	State           PollState         `protobuf:"varint,6,opt,name=state,proto3,enum=scalar.vote.exported.v1beta1.PollState" json:"state,omitempty"`
	MinVoterCount   int64             `protobuf:"varint,7,opt,name=min_voter_count,json=minVoterCount,proto3" json:"min_voter_count,omitempty"`
	RewardPoolName  string            `protobuf:"bytes,10,opt,name=reward_pool_name,json=rewardPoolName,proto3" json:"reward_pool_name,omitempty"`
	GracePeriod     int64             `protobuf:"varint,11,opt,name=grace_period,json=gracePeriod,proto3" json:"grace_period,omitempty"`
	CompletedAt     int64             `protobuf:"varint,12,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	ID              PollID            `protobuf:"varint,13,opt,name=id,proto3,customtype=PollID" json:"id"`
	Snapshot        exported.Snapshot `protobuf:"bytes,15,opt,name=snapshot,proto3" json:"snapshot"`
	Module          string            `protobuf:"bytes,16,opt,name=module,proto3" json:"module,omitempty"`
	ModuleMetadata  *types.Any        `protobuf:"bytes,17,opt,name=module_metadata,json=moduleMetadata,proto3" json:"module_metadata,omitempty"`
}

func (m *PollMetadata) Reset()         { *m = PollMetadata{} }
func (m *PollMetadata) String() string { return proto.CompactTextString(m) }
func (*PollMetadata) ProtoMessage()    {}
func (*PollMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcf266de9f0ebe6d, []int{0}
}
func (m *PollMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PollMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PollMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PollMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollMetadata.Merge(m, src)
}
func (m *PollMetadata) XXX_Size() int {
	return m.Size()
}
func (m *PollMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PollMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_PollMetadata proto.InternalMessageInfo

// PollKey represents the key data for a poll
//
// Deprecated: Do not use.
type PollKey struct {
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	ID     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *PollKey) Reset()      { *m = PollKey{} }
func (*PollKey) ProtoMessage() {}
func (*PollKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcf266de9f0ebe6d, []int{1}
}
func (m *PollKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PollKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PollKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PollKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollKey.Merge(m, src)
}
func (m *PollKey) XXX_Size() int {
	return m.Size()
}
func (m *PollKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PollKey.DiscardUnknown(m)
}

var xxx_messageInfo_PollKey proto.InternalMessageInfo

// PollParticipants should be embedded in poll events in other modules
type PollParticipants struct {
	PollID       PollID                                          `protobuf:"varint,1,opt,name=poll_id,json=pollId,proto3,customtype=PollID" json:"poll_id"`
	Participants []github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,2,rep,name=participants,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"participants,omitempty"`
}

func (m *PollParticipants) Reset()         { *m = PollParticipants{} }
func (m *PollParticipants) String() string { return proto.CompactTextString(m) }
func (*PollParticipants) ProtoMessage()    {}
func (*PollParticipants) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcf266de9f0ebe6d, []int{2}
}
func (m *PollParticipants) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PollParticipants) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PollParticipants.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PollParticipants) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollParticipants.Merge(m, src)
}
func (m *PollParticipants) XXX_Size() int {
	return m.Size()
}
func (m *PollParticipants) XXX_DiscardUnknown() {
	xxx_messageInfo_PollParticipants.DiscardUnknown(m)
}

var xxx_messageInfo_PollParticipants proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("scalar.vote.exported.v1beta1.PollState", PollState_name, PollState_value)
	proto.RegisterType((*PollMetadata)(nil), "scalar.vote.exported.v1beta1.PollMetadata")
	proto.RegisterType((*PollKey)(nil), "scalar.vote.exported.v1beta1.PollKey")
	proto.RegisterType((*PollParticipants)(nil), "scalar.vote.exported.v1beta1.PollParticipants")
}

func init() {
	proto.RegisterFile("scalar/vote/exported/v1beta1/types.proto", fileDescriptor_bcf266de9f0ebe6d)
}

var fileDescriptor_bcf266de9f0ebe6d = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x6f, 0xe3, 0x44,
	0x18, 0xc6, 0xe3, 0x34, 0x75, 0x93, 0x49, 0xda, 0x78, 0x47, 0x55, 0x65, 0x22, 0x70, 0xbc, 0x65,
	0xb5, 0x1b, 0x15, 0x6a, 0xd3, 0xe5, 0x86, 0xc4, 0x21, 0x69, 0x52, 0x94, 0x92, 0x66, 0xad, 0xb4,
	0xbb, 0x42, 0x5c, 0xac, 0xa9, 0x67, 0x70, 0x2c, 0x6c, 0x8f, 0x35, 0x33, 0x29, 0xe9, 0x37, 0x40,
	0x3d, 0x71, 0xe4, 0x40, 0x25, 0x24, 0x38, 0x70, 0x87, 0x3b, 0xd7, 0x8a, 0xd3, 0x1e, 0x11, 0x87,
	0x0a, 0xd2, 0x6f, 0xc1, 0x09, 0x8d, 0xff, 0x64, 0x53, 0x2d, 0xda, 0xd3, 0x9e, 0x32, 0xf3, 0xcc,
	0x6f, 0xde, 0xbc, 0xcf, 0xe3, 0xd7, 0x06, 0x1d, 0xee, 0xa1, 0x10, 0x31, 0xfb, 0x82, 0x0a, 0x62,
	0x93, 0x79, 0x42, 0x99, 0x20, 0xd8, 0xbe, 0x38, 0x38, 0x27, 0x02, 0x1d, 0xd8, 0xe2, 0x32, 0x21,
	0xdc, 0x4a, 0x18, 0x15, 0x14, 0xbe, 0x9b, 0x91, 0x96, 0x24, 0xad, 0x82, 0xb4, 0x72, 0xb2, 0xb5,
	0xed, 0x53, 0x9f, 0xa6, 0xa0, 0x2d, 0x57, 0xd9, 0x9d, 0xd6, 0x3b, 0x3e, 0xa5, 0x7e, 0x48, 0xec,
	0x74, 0x77, 0x3e, 0xfb, 0xca, 0x46, 0xf1, 0x65, 0x71, 0xe4, 0x51, 0x1e, 0x51, 0xee, 0x66, 0x77,
	0xb2, 0x4d, 0x7e, 0xf4, 0x21, 0x9a, 0x13, 0xd9, 0x13, 0x8f, 0x51, 0xc2, 0xa7, 0x54, 0xbc, 0xb1,
	0xaf, 0xd6, 0xa3, 0x9c, 0x9e, 0x89, 0x20, 0xe4, 0xaf, 0x88, 0x29, 0x23, 0x7c, 0x4a, 0x43, 0x9c,
	0x51, 0xbb, 0xbf, 0xaf, 0x83, 0x86, 0x43, 0xc3, 0xf0, 0x84, 0x08, 0x84, 0x91, 0x40, 0xf0, 0x3d,
	0x00, 0xc8, 0x3c, 0x09, 0x18, 0xe1, 0x2e, 0x12, 0xfa, 0x9a, 0xa9, 0x74, 0xd6, 0x26, 0xb5, 0x5c,
	0xe9, 0x0a, 0xf8, 0x05, 0x50, 0x19, 0xe1, 0xb3, 0x50, 0xe8, 0x15, 0x53, 0xe9, 0xd4, 0x9f, 0x6e,
	0x5b, 0x99, 0x15, 0xab, 0xb0, 0x62, 0x75, 0xe3, 0xcb, 0xde, 0xde, 0x1f, 0xbf, 0xed, 0x3f, 0xf6,
	0x03, 0x31, 0x9d, 0x9d, 0x5b, 0x1e, 0x8d, 0x72, 0x1b, 0xb6, 0x47, 0x31, 0xf1, 0x6c, 0x47, 0x92,
	0x27, 0x88, 0xf1, 0x29, 0x0a, 0x09, 0x9b, 0xe4, 0xf5, 0xa0, 0x03, 0xb4, 0x0b, 0x2a, 0x82, 0xd8,
	0x77, 0x97, 0x3d, 0xea, 0xeb, 0xe9, 0x7f, 0xb4, 0xad, 0xcc, 0x8a, 0x95, 0x5a, 0x29, 0xa2, 0xb5,
	0xce, 0x0a, 0xac, 0x57, 0xb9, 0xb9, 0x6d, 0x97, 0x26, 0xcd, 0xec, 0xfa, 0x52, 0x86, 0x9f, 0x82,
	0x75, 0x2e, 0x90, 0x20, 0xba, 0x6a, 0x2a, 0x9d, 0xad, 0xa7, 0x4f, 0xac, 0x37, 0x3d, 0x29, 0x4b,
	0xa6, 0x70, 0x2a, 0xf1, 0x49, 0x76, 0x0b, 0x3e, 0x06, 0xcd, 0x28, 0x88, 0x5d, 0x49, 0x33, 0xd7,
	0xa3, 0xb3, 0x58, 0xe8, 0x1b, 0x69, 0x1c, 0x9b, 0x51, 0x10, 0xbf, 0x90, 0xea, 0xa1, 0x14, 0x61,
	0x07, 0x68, 0x8c, 0x7c, 0x83, 0x18, 0x76, 0x13, 0x4a, 0x43, 0x37, 0x46, 0x11, 0xd1, 0x81, 0xa9,
	0x74, 0x6a, 0x93, 0xad, 0x4c, 0x77, 0x28, 0x0d, 0xc7, 0x28, 0x22, 0xf0, 0x21, 0x68, 0xf8, 0x0c,
	0x79, 0xc4, 0x4d, 0x08, 0x0b, 0x28, 0xd6, 0xeb, 0x69, 0xb9, 0x7a, 0xaa, 0x39, 0xa9, 0x24, 0x11,
	0x8f, 0x46, 0x49, 0x48, 0x04, 0xc1, 0xf2, 0x01, 0x34, 0x32, 0x64, 0xa9, 0x75, 0x05, 0x7c, 0x04,
	0xca, 0x01, 0xd6, 0x37, 0x4d, 0xa5, 0x53, 0xe9, 0x6d, 0x4b, 0xe7, 0x7f, 0xdd, 0xb6, 0x55, 0xd9,
	0xfd, 0xb0, 0xbf, 0xb8, 0x6d, 0x97, 0x87, 0xfd, 0x49, 0x39, 0xc0, 0x70, 0x04, 0xaa, 0xc5, 0x9c,
	0xe8, 0xcd, 0x34, 0xc6, 0xbd, 0x22, 0xc6, 0x42, 0x7f, 0x3d, 0x83, 0xd3, 0xfc, 0x24, 0x4f, 0x74,
	0x59, 0x01, 0xee, 0x00, 0x35, 0xa2, 0x78, 0x16, 0x12, 0x5d, 0x4b, 0x9d, 0xe5, 0x3b, 0x18, 0x80,
	0x66, 0xb6, 0x72, 0xa3, 0x7c, 0x80, 0xf4, 0x07, 0x6f, 0x69, 0x2e, 0xb6, 0xb2, 0xc2, 0xc5, 0x60,
	0x1e, 0x57, 0xaa, 0x8a, 0x56, 0x3e, 0xae, 0x54, 0xab, 0x5a, 0xed, 0xb8, 0x52, 0xad, 0x69, 0xe0,
	0xb8, 0x52, 0xdd, 0xd2, 0x9a, 0xbb, 0x5d, 0xb0, 0x21, 0xcd, 0x7f, 0x4e, 0x2e, 0x57, 0xba, 0x54,
	0xee, 0x75, 0xb9, 0x93, 0x26, 0x56, 0x96, 0x5a, 0x4f, 0x7d, 0x95, 0xd1, 0x27, 0xea, 0xf7, 0x3f,
	0xb6, 0x4b, 0xba, 0xb2, 0xfb, 0x83, 0x02, 0x34, 0x59, 0xc3, 0x41, 0x4c, 0x04, 0x5e, 0x90, 0xa0,
	0x58, 0x70, 0x78, 0x00, 0x36, 0x12, 0x1a, 0x86, 0x6e, 0x80, 0xd3, 0x6a, 0x95, 0x9e, 0xfe, 0x5a,
	0xd6, 0xf9, 0x6a, 0xa2, 0x4a, 0x70, 0x88, 0xe1, 0x73, 0xd0, 0x48, 0x56, 0x4a, 0xe8, 0x65, 0x73,
	0xad, 0xd3, 0xe8, 0x1d, 0xfc, 0x7b, 0xdb, 0xde, 0xff, 0x3f, 0xd3, 0xf2, 0x67, 0x9f, 0xe3, 0xaf,
	0xf3, 0xd7, 0xf6, 0x05, 0x0a, 0xbb, 0x18, 0x33, 0xc2, 0xf9, 0xe4, 0x5e, 0x99, 0xbd, 0x5f, 0x15,
	0x50, 0x5b, 0x4e, 0x27, 0xfc, 0x00, 0xec, 0x38, 0xcf, 0x46, 0x23, 0xf7, 0xf4, 0xac, 0x7b, 0x36,
	0x70, 0x9f, 0x8f, 0x4f, 0x9d, 0xc1, 0xe1, 0xf0, 0x68, 0x38, 0xe8, 0x6b, 0xa5, 0x56, 0xf3, 0xea,
	0xda, 0xac, 0x8f, 0x69, 0x3c, 0x98, 0x07, 0x5c, 0x90, 0x58, 0xc0, 0xf7, 0x01, 0x5c, 0x81, 0x9d,
	0xc1, 0xb8, 0x3f, 0x1c, 0x7f, 0xa6, 0x29, 0xad, 0xfa, 0xd5, 0xb5, 0xb9, 0xe1, 0x90, 0x18, 0x07,
	0xb1, 0x0f, 0x9f, 0x80, 0xed, 0x15, 0xe8, 0xf0, 0xd9, 0x89, 0x33, 0x1a, 0x9c, 0x0d, 0xfa, 0x5a,
	0xb9, 0xb5, 0x79, 0x75, 0x6d, 0xd6, 0x0e, 0x8b, 0xd9, 0x83, 0x0f, 0xc1, 0x83, 0x15, 0xf0, 0xa8,
	0x3b, 0x1c, 0x0d, 0xfa, 0xda, 0x5a, 0x0b, 0x5c, 0x5d, 0x9b, 0xea, 0x11, 0x0a, 0x42, 0x82, 0x5b,
	0xd5, 0x6f, 0x7f, 0x32, 0x4a, 0xbf, 0xfc, 0x6c, 0x28, 0xbd, 0xf1, 0xcd, 0x3f, 0x46, 0xe9, 0x66,
	0x61, 0x28, 0x2f, 0x17, 0x86, 0xf2, 0xf7, 0xc2, 0x50, 0xbe, 0xbb, 0x33, 0x4a, 0x2f, 0xef, 0x8c,
	0xd2, 0x9f, 0x77, 0x46, 0xe9, 0xcb, 0x8f, 0x56, 0x02, 0xc9, 0x5e, 0x4b, 0xca, 0xfc, 0x7c, 0xb5,
	0xef, 0x51, 0x46, 0xec, 0xf9, 0xfd, 0x6f, 0xef, 0xb9, 0x9a, 0x4e, 0xd2, 0xc7, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x76, 0xc9, 0x9a, 0x29, 0x9a, 0x05, 0x00, 0x00,
}

func (m *PollMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PollMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PollMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ModuleMetadata != nil {
		{
			size, err := m.ModuleMetadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	{
		size, err := m.Snapshot.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x7a
	if m.ID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x68
	}
	if m.CompletedAt != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.CompletedAt))
		i--
		dAtA[i] = 0x60
	}
	if m.GracePeriod != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.GracePeriod))
		i--
		dAtA[i] = 0x58
	}
	if len(m.RewardPoolName) > 0 {
		i -= len(m.RewardPoolName)
		copy(dAtA[i:], m.RewardPoolName)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.RewardPoolName)))
		i--
		dAtA[i] = 0x52
	}
	if m.MinVoterCount != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MinVoterCount))
		i--
		dAtA[i] = 0x38
	}
	if m.State != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.VotingThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ExpiresAt != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ExpiresAt))
		i--
		dAtA[i] = 0x18
	}
	return len(dAtA) - i, nil
}

func (m *PollKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PollKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PollKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PollParticipants) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PollParticipants) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PollParticipants) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Participants[iNdEx])
			copy(dAtA[i:], m.Participants[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Participants[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.PollID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.PollID))
		i--
		dAtA[i] = 0x8
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
func (m *PollMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExpiresAt != 0 {
		n += 1 + sovTypes(uint64(m.ExpiresAt))
	}
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.VotingThreshold.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.State != 0 {
		n += 1 + sovTypes(uint64(m.State))
	}
	if m.MinVoterCount != 0 {
		n += 1 + sovTypes(uint64(m.MinVoterCount))
	}
	l = len(m.RewardPoolName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.GracePeriod != 0 {
		n += 1 + sovTypes(uint64(m.GracePeriod))
	}
	if m.CompletedAt != 0 {
		n += 1 + sovTypes(uint64(m.CompletedAt))
	}
	if m.ID != 0 {
		n += 1 + sovTypes(uint64(m.ID))
	}
	l = m.Snapshot.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.Module)
	if l > 0 {
		n += 2 + l + sovTypes(uint64(l))
	}
	if m.ModuleMetadata != nil {
		l = m.ModuleMetadata.Size()
		n += 2 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *PollKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *PollParticipants) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PollID != 0 {
		n += 1 + sovTypes(uint64(m.PollID))
	}
	if len(m.Participants) > 0 {
		for _, b := range m.Participants {
			l = len(b)
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
func (m *PollMetadata) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PollMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PollMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
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
			if m.Result == nil {
				m.Result = &types.Any{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingThreshold", wireType)
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
			if err := m.VotingThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= PollState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinVoterCount", wireType)
			}
			m.MinVoterCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinVoterCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPoolName", wireType)
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
			m.RewardPoolName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GracePeriod", wireType)
			}
			m.GracePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GracePeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletedAt", wireType)
			}
			m.CompletedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompletedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= PollID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snapshot", wireType)
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
			if err := m.Snapshot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
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
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleMetadata", wireType)
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
			if m.ModuleMetadata == nil {
				m.ModuleMetadata = &types.Any{}
			}
			if err := m.ModuleMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *PollKey) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PollKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PollKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
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
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
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
			m.ID = string(dAtA[iNdEx:postIndex])
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
func (m *PollParticipants) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PollParticipants: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PollParticipants: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollID", wireType)
			}
			m.PollID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PollID |= PollID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
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
			m.Participants = append(m.Participants, make([]byte, postIndex-iNdEx))
			copy(m.Participants[len(m.Participants)-1], dAtA[iNdEx:postIndex])
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
