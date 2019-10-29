// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go-internal/log_entry_message.proto

package group

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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

type MessageEntryPayload_PayloadType int32

const (
	MessageEntryPayload_PayloadTypeUnknown    MessageEntryPayload_PayloadType = 0
	MessageEntryPayload_PayloadTypeMessage    MessageEntryPayload_PayloadType = 1
	MessageEntryPayload_PayloadTypeInvitation MessageEntryPayload_PayloadType = 2
)

var MessageEntryPayload_PayloadType_name = map[int32]string{
	0: "PayloadTypeUnknown",
	1: "PayloadTypeMessage",
	2: "PayloadTypeInvitation",
}

var MessageEntryPayload_PayloadType_value = map[string]int32{
	"PayloadTypeUnknown":    0,
	"PayloadTypeMessage":    1,
	"PayloadTypeInvitation": 2,
}

func (x MessageEntryPayload_PayloadType) String() string {
	return proto.EnumName(MessageEntryPayload_PayloadType_name, int32(x))
}

func (MessageEntryPayload_PayloadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb30c2d83a30b073, []int{1, 0}
}

type MessageEntryEnvelope struct {
	Counter          uint64 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
	EncryptedPayload []byte `protobuf:"bytes,2,opt,name=encrypted_payload,json=encryptedPayload,proto3" json:"encrypted_payload,omitempty"`
	Signature        []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *MessageEntryEnvelope) Reset()         { *m = MessageEntryEnvelope{} }
func (m *MessageEntryEnvelope) String() string { return proto.CompactTextString(m) }
func (*MessageEntryEnvelope) ProtoMessage()    {}
func (*MessageEntryEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb30c2d83a30b073, []int{0}
}
func (m *MessageEntryEnvelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageEntryEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageEntryEnvelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageEntryEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageEntryEnvelope.Merge(m, src)
}
func (m *MessageEntryEnvelope) XXX_Size() int {
	return m.Size()
}
func (m *MessageEntryEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageEntryEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_MessageEntryEnvelope proto.InternalMessageInfo

func (m *MessageEntryEnvelope) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *MessageEntryEnvelope) GetEncryptedPayload() []byte {
	if m != nil {
		return m.EncryptedPayload
	}
	return nil
}

func (m *MessageEntryEnvelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type MessageEntryPayload struct {
	Type        MessageEntryPayload_PayloadType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageEntryPayload_PayloadType" json:"type,omitempty"`
	MessageBody []byte                          `protobuf:"bytes,2,opt,name=message_body,json=messageBody,proto3" json:"message_body,omitempty"`
	Invitation  *Invitation                     `protobuf:"bytes,3,opt,name=invitation,proto3" json:"invitation,omitempty"`
}

func (m *MessageEntryPayload) Reset()         { *m = MessageEntryPayload{} }
func (m *MessageEntryPayload) String() string { return proto.CompactTextString(m) }
func (*MessageEntryPayload) ProtoMessage()    {}
func (*MessageEntryPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb30c2d83a30b073, []int{1}
}
func (m *MessageEntryPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageEntryPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageEntryPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageEntryPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageEntryPayload.Merge(m, src)
}
func (m *MessageEntryPayload) XXX_Size() int {
	return m.Size()
}
func (m *MessageEntryPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageEntryPayload.DiscardUnknown(m)
}

var xxx_messageInfo_MessageEntryPayload proto.InternalMessageInfo

func (m *MessageEntryPayload) GetType() MessageEntryPayload_PayloadType {
	if m != nil {
		return m.Type
	}
	return MessageEntryPayload_PayloadTypeUnknown
}

func (m *MessageEntryPayload) GetMessageBody() []byte {
	if m != nil {
		return m.MessageBody
	}
	return nil
}

func (m *MessageEntryPayload) GetInvitation() *Invitation {
	if m != nil {
		return m.Invitation
	}
	return nil
}

type Invitation struct {
	InviterMemberPubKey       []byte `protobuf:"bytes,1,opt,name=inviter_member_pub_key,json=inviterMemberPubKey,proto3" json:"inviter_member_pub_key,omitempty"`
	InvitationPrivKey         []byte `protobuf:"bytes,2,opt,name=invitation_priv_key,json=invitationPrivKey,proto3" json:"invitation_priv_key,omitempty"`
	InvitationPubKeySignature []byte `protobuf:"bytes,3,opt,name=invitation_pub_key_signature,json=invitationPubKeySignature,proto3" json:"invitation_pub_key_signature,omitempty"`
	GroupVersion              uint32 `protobuf:"varint,4,opt,name=group_version,json=groupVersion,proto3" json:"group_version,omitempty"`
	GroupIdPubKey             []byte `protobuf:"bytes,5,opt,name=group_id_pub_key,json=groupIdPubKey,proto3" json:"group_id_pub_key,omitempty"`
	SharedSecret              []byte `protobuf:"bytes,6,opt,name=shared_secret,json=sharedSecret,proto3" json:"shared_secret,omitempty"`
}

func (m *Invitation) Reset()         { *m = Invitation{} }
func (m *Invitation) String() string { return proto.CompactTextString(m) }
func (*Invitation) ProtoMessage()    {}
func (*Invitation) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb30c2d83a30b073, []int{2}
}
func (m *Invitation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Invitation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Invitation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Invitation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invitation.Merge(m, src)
}
func (m *Invitation) XXX_Size() int {
	return m.Size()
}
func (m *Invitation) XXX_DiscardUnknown() {
	xxx_messageInfo_Invitation.DiscardUnknown(m)
}

var xxx_messageInfo_Invitation proto.InternalMessageInfo

func (m *Invitation) GetInviterMemberPubKey() []byte {
	if m != nil {
		return m.InviterMemberPubKey
	}
	return nil
}

func (m *Invitation) GetInvitationPrivKey() []byte {
	if m != nil {
		return m.InvitationPrivKey
	}
	return nil
}

func (m *Invitation) GetInvitationPubKeySignature() []byte {
	if m != nil {
		return m.InvitationPubKeySignature
	}
	return nil
}

func (m *Invitation) GetGroupVersion() uint32 {
	if m != nil {
		return m.GroupVersion
	}
	return 0
}

func (m *Invitation) GetGroupIdPubKey() []byte {
	if m != nil {
		return m.GroupIdPubKey
	}
	return nil
}

func (m *Invitation) GetSharedSecret() []byte {
	if m != nil {
		return m.SharedSecret
	}
	return nil
}

func init() {
	proto.RegisterEnum("MessageEntryPayload_PayloadType", MessageEntryPayload_PayloadType_name, MessageEntryPayload_PayloadType_value)
	proto.RegisterType((*MessageEntryEnvelope)(nil), "MessageEntryEnvelope")
	proto.RegisterType((*MessageEntryPayload)(nil), "MessageEntryPayload")
	proto.RegisterType((*Invitation)(nil), "Invitation")
}

func init() {
	proto.RegisterFile("go-internal/log_entry_message.proto", fileDescriptor_eb30c2d83a30b073)
}

var fileDescriptor_eb30c2d83a30b073 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6a, 0xdb, 0x4c,
	0x14, 0x85, 0x3d, 0xfe, 0xfd, 0xa7, 0xf4, 0x5a, 0x2e, 0xca, 0xb8, 0x0d, 0x0a, 0x18, 0xe1, 0xda,
	0x8b, 0x1a, 0x42, 0x65, 0x48, 0x4a, 0xb7, 0x85, 0x40, 0x16, 0xa1, 0x04, 0x8c, 0xd2, 0x96, 0xd2,
	0xcd, 0x20, 0x59, 0x17, 0x45, 0xc4, 0x9e, 0x19, 0x46, 0x23, 0x95, 0xa1, 0x2f, 0xd1, 0x97, 0xe8,
	0xbb, 0x74, 0x99, 0x65, 0x97, 0xc5, 0x7e, 0x86, 0xee, 0x8b, 0x47, 0xb2, 0x25, 0x42, 0x57, 0x66,
	0xbe, 0x73, 0xae, 0xcf, 0xb9, 0xa3, 0x81, 0x69, 0x2a, 0x5e, 0x67, 0x5c, 0xa3, 0xe2, 0xd1, 0x6a,
	0xbe, 0x12, 0x29, 0x43, 0xae, 0x95, 0x61, 0x6b, 0xcc, 0xf3, 0x28, 0xc5, 0x40, 0x2a, 0xa1, 0xc5,
	0xe4, 0x1b, 0x3c, 0xbf, 0xa9, 0xc0, 0xd5, 0x4e, 0xbd, 0xe2, 0x25, 0xae, 0x84, 0x44, 0xea, 0xc1,
	0x93, 0xa5, 0x28, 0x76, 0xd3, 0x1e, 0x19, 0x93, 0x59, 0x2f, 0xdc, 0x1f, 0xe9, 0x19, 0x1c, 0x23,
	0x5f, 0x2a, 0x23, 0x35, 0x26, 0x4c, 0x46, 0x66, 0x25, 0xa2, 0xc4, 0xeb, 0x8e, 0xc9, 0xcc, 0x09,
	0xdd, 0x83, 0xb0, 0xa8, 0x38, 0x1d, 0xc1, 0xd3, 0x3c, 0x4b, 0x79, 0xa4, 0x0b, 0x85, 0xde, 0x7f,
	0xd6, 0xd4, 0x80, 0xc9, 0x1f, 0x02, 0xc3, 0x76, 0xfa, 0x7e, 0xea, 0x0d, 0xf4, 0xb4, 0x91, 0x68,
	0x93, 0x9f, 0x9d, 0x8f, 0x83, 0x7f, 0x78, 0x82, 0xfa, 0xf7, 0x83, 0x91, 0x18, 0x5a, 0x37, 0x7d,
	0x09, 0x4e, 0xbd, 0x1b, 0x8b, 0x45, 0x62, 0xea, 0x4e, 0xfd, 0x9a, 0x5d, 0x8a, 0xc4, 0xd0, 0x33,
	0x80, 0x8c, 0x97, 0x99, 0x8e, 0x74, 0x26, 0xb8, 0xed, 0xd3, 0x3f, 0xef, 0x07, 0xd7, 0x07, 0x14,
	0xb6, 0xe4, 0xc9, 0x67, 0xe8, 0xb7, 0x42, 0xe8, 0x09, 0xd0, 0xd6, 0xf1, 0x23, 0xbf, 0xe7, 0xe2,
	0x2b, 0x77, 0x3b, 0x8f, 0x78, 0x5d, 0xd5, 0x25, 0xf4, 0x14, 0x5e, 0xb4, 0x78, 0x93, 0xe1, 0x76,
	0x27, 0x3f, 0xba, 0x00, 0x0d, 0xa0, 0x17, 0x70, 0x62, 0x63, 0x51, 0xb1, 0x35, 0xae, 0x63, 0x54,
	0x4c, 0x16, 0x31, 0xbb, 0x47, 0x63, 0x2f, 0xc0, 0x09, 0x87, 0xb5, 0x7a, 0x63, 0xc5, 0x45, 0x11,
	0xbf, 0x47, 0x43, 0x03, 0x18, 0x36, 0x5d, 0x99, 0x54, 0x59, 0x69, 0x27, 0xaa, 0xa5, 0x8f, 0x1b,
	0x69, 0xa1, 0xb2, 0x72, 0xe7, 0x7f, 0x07, 0xa3, 0xb6, 0xbf, 0x0a, 0x60, 0x8f, 0x3f, 0xce, 0x69,
	0x6b, 0xd0, 0xe6, 0xdc, 0xee, 0x0d, 0x74, 0x0a, 0x83, 0x54, 0x89, 0x42, 0xb2, 0x12, 0x55, 0xbe,
	0xbb, 0xbe, 0xde, 0x98, 0xcc, 0x06, 0xa1, 0x63, 0xe1, 0xa7, 0x8a, 0xd1, 0x57, 0xe0, 0x56, 0xa6,
	0x2c, 0x39, 0x2c, 0xf1, 0xbf, 0xfd, 0xe7, 0x6a, 0xf8, 0x3a, 0xa9, 0xeb, 0x4f, 0x61, 0x90, 0xdf,
	0x45, 0x0a, 0x13, 0x96, 0xe3, 0x52, 0xa1, 0xf6, 0x8e, 0xac, 0xcb, 0xa9, 0xe0, 0xad, 0x65, 0x97,
	0x6f, 0x7f, 0x6e, 0x7c, 0xf2, 0xb0, 0xf1, 0xc9, 0xef, 0x8d, 0x4f, 0xbe, 0x6f, 0xfd, 0xce, 0xc3,
	0xd6, 0xef, 0xfc, 0xda, 0xfa, 0x9d, 0x2f, 0xa3, 0x18, 0x95, 0x36, 0x81, 0xc6, 0xe5, 0xdd, 0x3c,
	0x15, 0xf3, 0xc3, 0x33, 0xb7, 0x21, 0xf1, 0x91, 0x7d, 0xdb, 0x17, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0x1f, 0xe0, 0xbe, 0x02, 0x03, 0x00, 0x00,
}

func (m *MessageEntryEnvelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageEntryEnvelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MessageEntryEnvelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EncryptedPayload) > 0 {
		i -= len(m.EncryptedPayload)
		copy(dAtA[i:], m.EncryptedPayload)
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(len(m.EncryptedPayload)))
		i--
		dAtA[i] = 0x12
	}
	if m.Counter != 0 {
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MessageEntryPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageEntryPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MessageEntryPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Invitation != nil {
		{
			size, err := m.Invitation.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLogEntryMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MessageBody) > 0 {
		i -= len(m.MessageBody)
		copy(dAtA[i:], m.MessageBody)
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(len(m.MessageBody)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Invitation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Invitation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Invitation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SharedSecret) > 0 {
		i -= len(m.SharedSecret)
		copy(dAtA[i:], m.SharedSecret)
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(len(m.SharedSecret)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.GroupIdPubKey) > 0 {
		i -= len(m.GroupIdPubKey)
		copy(dAtA[i:], m.GroupIdPubKey)
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(len(m.GroupIdPubKey)))
		i--
		dAtA[i] = 0x2a
	}
	if m.GroupVersion != 0 {
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(m.GroupVersion))
		i--
		dAtA[i] = 0x20
	}
	if len(m.InvitationPubKeySignature) > 0 {
		i -= len(m.InvitationPubKeySignature)
		copy(dAtA[i:], m.InvitationPubKeySignature)
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(len(m.InvitationPubKeySignature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InvitationPrivKey) > 0 {
		i -= len(m.InvitationPrivKey)
		copy(dAtA[i:], m.InvitationPrivKey)
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(len(m.InvitationPrivKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.InviterMemberPubKey) > 0 {
		i -= len(m.InviterMemberPubKey)
		copy(dAtA[i:], m.InviterMemberPubKey)
		i = encodeVarintLogEntryMessage(dAtA, i, uint64(len(m.InviterMemberPubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLogEntryMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovLogEntryMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MessageEntryEnvelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Counter != 0 {
		n += 1 + sovLogEntryMessage(uint64(m.Counter))
	}
	l = len(m.EncryptedPayload)
	if l > 0 {
		n += 1 + l + sovLogEntryMessage(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovLogEntryMessage(uint64(l))
	}
	return n
}

func (m *MessageEntryPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovLogEntryMessage(uint64(m.Type))
	}
	l = len(m.MessageBody)
	if l > 0 {
		n += 1 + l + sovLogEntryMessage(uint64(l))
	}
	if m.Invitation != nil {
		l = m.Invitation.Size()
		n += 1 + l + sovLogEntryMessage(uint64(l))
	}
	return n
}

func (m *Invitation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InviterMemberPubKey)
	if l > 0 {
		n += 1 + l + sovLogEntryMessage(uint64(l))
	}
	l = len(m.InvitationPrivKey)
	if l > 0 {
		n += 1 + l + sovLogEntryMessage(uint64(l))
	}
	l = len(m.InvitationPubKeySignature)
	if l > 0 {
		n += 1 + l + sovLogEntryMessage(uint64(l))
	}
	if m.GroupVersion != 0 {
		n += 1 + sovLogEntryMessage(uint64(m.GroupVersion))
	}
	l = len(m.GroupIdPubKey)
	if l > 0 {
		n += 1 + l + sovLogEntryMessage(uint64(l))
	}
	l = len(m.SharedSecret)
	if l > 0 {
		n += 1 + l + sovLogEntryMessage(uint64(l))
	}
	return n
}

func sovLogEntryMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLogEntryMessage(x uint64) (n int) {
	return sovLogEntryMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MessageEntryEnvelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogEntryMessage
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
			return fmt.Errorf("proto: MessageEntryEnvelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageEntryEnvelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
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
				return ErrInvalidLengthLogEntryMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedPayload = append(m.EncryptedPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptedPayload == nil {
				m.EncryptedPayload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
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
				return ErrInvalidLengthLogEntryMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogEntryMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogEntryMessage
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
func (m *MessageEntryPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogEntryMessage
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
			return fmt.Errorf("proto: MessageEntryPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageEntryPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= MessageEntryPayload_PayloadType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageBody", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
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
				return ErrInvalidLengthLogEntryMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageBody = append(m.MessageBody[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageBody == nil {
				m.MessageBody = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Invitation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
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
				return ErrInvalidLengthLogEntryMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Invitation == nil {
				m.Invitation = &Invitation{}
			}
			if err := m.Invitation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogEntryMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogEntryMessage
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
func (m *Invitation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogEntryMessage
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
			return fmt.Errorf("proto: Invitation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Invitation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InviterMemberPubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
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
				return ErrInvalidLengthLogEntryMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InviterMemberPubKey = append(m.InviterMemberPubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.InviterMemberPubKey == nil {
				m.InviterMemberPubKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvitationPrivKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
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
				return ErrInvalidLengthLogEntryMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InvitationPrivKey = append(m.InvitationPrivKey[:0], dAtA[iNdEx:postIndex]...)
			if m.InvitationPrivKey == nil {
				m.InvitationPrivKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvitationPubKeySignature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
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
				return ErrInvalidLengthLogEntryMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InvitationPubKeySignature = append(m.InvitationPubKeySignature[:0], dAtA[iNdEx:postIndex]...)
			if m.InvitationPubKeySignature == nil {
				m.InvitationPubKeySignature = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupVersion", wireType)
			}
			m.GroupVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupIdPubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
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
				return ErrInvalidLengthLogEntryMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupIdPubKey = append(m.GroupIdPubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.GroupIdPubKey == nil {
				m.GroupIdPubKey = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharedSecret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntryMessage
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
				return ErrInvalidLengthLogEntryMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SharedSecret = append(m.SharedSecret[:0], dAtA[iNdEx:postIndex]...)
			if m.SharedSecret == nil {
				m.SharedSecret = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogEntryMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogEntryMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogEntryMessage
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
func skipLogEntryMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogEntryMessage
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
					return 0, ErrIntOverflowLogEntryMessage
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
					return 0, ErrIntOverflowLogEntryMessage
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
				return 0, ErrInvalidLengthLogEntryMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLogEntryMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLogEntryMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLogEntryMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogEntryMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLogEntryMessage = fmt.Errorf("proto: unexpected end of group")
)
