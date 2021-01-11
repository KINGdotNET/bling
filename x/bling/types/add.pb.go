// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bling/add.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Add struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Post    string `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
	Title   string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Body    string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Ipfs    string `protobuf:"bytes,6,opt,name=ipfs,proto3" json:"ipfs,omitempty"`
	Parent  string `protobuf:"bytes,7,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (m *Add) Reset()         { *m = Add{} }
func (m *Add) String() string { return proto.CompactTextString(m) }
func (*Add) ProtoMessage()    {}
func (*Add) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f26b1e7726a791, []int{0}
}
func (m *Add) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Add) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Add.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Add) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Add.Merge(m, src)
}
func (m *Add) XXX_Size() int {
	return m.Size()
}
func (m *Add) XXX_DiscardUnknown() {
	xxx_messageInfo_Add.DiscardUnknown(m)
}

var xxx_messageInfo_Add proto.InternalMessageInfo

func (m *Add) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Add) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Add) GetPost() string {
	if m != nil {
		return m.Post
	}
	return ""
}

func (m *Add) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Add) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Add) GetIpfs() string {
	if m != nil {
		return m.Ipfs
	}
	return ""
}

func (m *Add) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

type MsgCreateAdd struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Post    string `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body    string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Ipfs    string `protobuf:"bytes,5,opt,name=ipfs,proto3" json:"ipfs,omitempty"`
	Parent  string `protobuf:"bytes,6,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (m *MsgCreateAdd) Reset()         { *m = MsgCreateAdd{} }
func (m *MsgCreateAdd) String() string { return proto.CompactTextString(m) }
func (*MsgCreateAdd) ProtoMessage()    {}
func (*MsgCreateAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f26b1e7726a791, []int{1}
}
func (m *MsgCreateAdd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateAdd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateAdd.Merge(m, src)
}
func (m *MsgCreateAdd) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateAdd.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateAdd proto.InternalMessageInfo

func (m *MsgCreateAdd) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateAdd) GetPost() string {
	if m != nil {
		return m.Post
	}
	return ""
}

func (m *MsgCreateAdd) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgCreateAdd) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MsgCreateAdd) GetIpfs() string {
	if m != nil {
		return m.Ipfs
	}
	return ""
}

func (m *MsgCreateAdd) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

type MsgUpdateAdd struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Post    string `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
	Title   string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Body    string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Ipfs    string `protobuf:"bytes,6,opt,name=ipfs,proto3" json:"ipfs,omitempty"`
	Parent  string `protobuf:"bytes,7,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (m *MsgUpdateAdd) Reset()         { *m = MsgUpdateAdd{} }
func (m *MsgUpdateAdd) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateAdd) ProtoMessage()    {}
func (*MsgUpdateAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f26b1e7726a791, []int{2}
}
func (m *MsgUpdateAdd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateAdd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateAdd.Merge(m, src)
}
func (m *MsgUpdateAdd) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateAdd.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateAdd proto.InternalMessageInfo

func (m *MsgUpdateAdd) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateAdd) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgUpdateAdd) GetPost() string {
	if m != nil {
		return m.Post
	}
	return ""
}

func (m *MsgUpdateAdd) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgUpdateAdd) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MsgUpdateAdd) GetIpfs() string {
	if m != nil {
		return m.Ipfs
	}
	return ""
}

func (m *MsgUpdateAdd) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

type MsgDeleteAdd struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgDeleteAdd) Reset()         { *m = MsgDeleteAdd{} }
func (m *MsgDeleteAdd) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteAdd) ProtoMessage()    {}
func (*MsgDeleteAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f26b1e7726a791, []int{3}
}
func (m *MsgDeleteAdd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteAdd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteAdd.Merge(m, src)
}
func (m *MsgDeleteAdd) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteAdd.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteAdd proto.InternalMessageInfo

func (m *MsgDeleteAdd) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteAdd) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Add)(nil), "kingdotnet.bling.bling.Add")
	proto.RegisterType((*MsgCreateAdd)(nil), "kingdotnet.bling.bling.MsgCreateAdd")
	proto.RegisterType((*MsgUpdateAdd)(nil), "kingdotnet.bling.bling.MsgUpdateAdd")
	proto.RegisterType((*MsgDeleteAdd)(nil), "kingdotnet.bling.bling.MsgDeleteAdd")
}

func init() { proto.RegisterFile("bling/add.proto", fileDescriptor_02f26b1e7726a791) }

var fileDescriptor_02f26b1e7726a791 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x3f, 0x4b, 0xc4, 0x30,
	0x18, 0xc6, 0x9b, 0xfe, 0x3b, 0x0c, 0xa2, 0x10, 0x8e, 0x23, 0x38, 0x04, 0xb9, 0x49, 0x97, 0x76,
	0x70, 0x71, 0xd5, 0x73, 0x75, 0x11, 0x5c, 0xdc, 0xda, 0x4b, 0x8c, 0xc1, 0xda, 0x84, 0xf6, 0x15,
	0xbc, 0x6f, 0x21, 0x6e, 0xae, 0x7e, 0x1a, 0xc7, 0x1b, 0x1d, 0xa5, 0xfd, 0x22, 0x92, 0xe4, 0x8a,
	0x08, 0x45, 0x71, 0xba, 0x25, 0x79, 0xde, 0x5f, 0x9e, 0xe1, 0x47, 0x78, 0xf1, 0x7e, 0x59, 0xa9,
	0x5a, 0xe6, 0x05, 0xe7, 0x99, 0x69, 0x34, 0x68, 0x32, 0xbb, 0x57, 0xb5, 0xe4, 0x1a, 0x6a, 0x01,
	0x99, 0x7b, 0xf3, 0xe7, 0xc1, 0x54, 0x6a, 0xa9, 0x5d, 0x25, 0xb7, 0xc9, 0xb7, 0xe7, 0xaf, 0x08,
	0x47, 0x67, 0x9c, 0x13, 0x8a, 0x27, 0xcb, 0x46, 0x14, 0xa0, 0x1b, 0x8a, 0x0e, 0xd1, 0xd1, 0xce,
	0xd5, 0x30, 0x92, 0x3d, 0x1c, 0x2a, 0x4e, 0x43, 0x07, 0x43, 0xc5, 0x09, 0xc1, 0xb1, 0xd1, 0x2d,
	0xd0, 0xc8, 0x11, 0x97, 0xc9, 0x14, 0x27, 0xa0, 0xa0, 0x12, 0x34, 0x76, 0xd0, 0x0f, 0xb6, 0x59,
	0x6a, 0xbe, 0xa2, 0x89, 0x6f, 0xda, 0x6c, 0x99, 0x32, 0xb7, 0x2d, 0x4d, 0x3d, 0xb3, 0x99, 0xcc,
	0x70, 0x6a, 0x8a, 0x46, 0xd4, 0x40, 0x27, 0x8e, 0x6e, 0xa6, 0xf9, 0x0b, 0xc2, 0xbb, 0x97, 0xad,
	0x5c, 0x58, 0x11, 0xf1, 0xbb, 0xe4, 0x20, 0x15, 0x8e, 0x49, 0x45, 0x63, 0x52, 0xf1, 0x88, 0x54,
	0x32, 0x2a, 0x95, 0xfe, 0x90, 0x7a, 0xf3, 0x52, 0xd7, 0x86, 0xff, 0x29, 0xb5, 0xbd, 0x9f, 0x3b,
	0x75, 0x8e, 0x17, 0xa2, 0x12, 0xff, 0x74, 0x3c, 0x5f, 0xbc, 0x77, 0x0c, 0xad, 0x3b, 0x86, 0x3e,
	0x3b, 0x86, 0x9e, 0x7b, 0x16, 0xac, 0x7b, 0x16, 0x7c, 0xf4, 0x2c, 0xb8, 0x39, 0x96, 0x0a, 0xee,
	0x1e, 0xcb, 0x6c, 0xa9, 0x1f, 0xf2, 0xef, 0x15, 0xcb, 0xfd, 0xfa, 0x3d, 0x6d, 0x6e, 0x58, 0x19,
	0xd1, 0x96, 0xa9, 0xdb, 0xad, 0x93, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0xc2, 0xc9, 0x08,
	0x9c, 0x02, 0x00, 0x00,
}

func (m *Add) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Add) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Add) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Parent) > 0 {
		i -= len(m.Parent)
		copy(dAtA[i:], m.Parent)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Parent)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Ipfs) > 0 {
		i -= len(m.Ipfs)
		copy(dAtA[i:], m.Ipfs)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Ipfs)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Post) > 0 {
		i -= len(m.Post)
		copy(dAtA[i:], m.Post)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Post)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateAdd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateAdd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateAdd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Parent) > 0 {
		i -= len(m.Parent)
		copy(dAtA[i:], m.Parent)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Parent)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Ipfs) > 0 {
		i -= len(m.Ipfs)
		copy(dAtA[i:], m.Ipfs)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Ipfs)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Post) > 0 {
		i -= len(m.Post)
		copy(dAtA[i:], m.Post)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Post)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateAdd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateAdd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateAdd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Parent) > 0 {
		i -= len(m.Parent)
		copy(dAtA[i:], m.Parent)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Parent)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Ipfs) > 0 {
		i -= len(m.Ipfs)
		copy(dAtA[i:], m.Ipfs)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Ipfs)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Post) > 0 {
		i -= len(m.Post)
		copy(dAtA[i:], m.Post)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Post)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteAdd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteAdd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteAdd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAdd(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdd(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Add) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Post)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Ipfs)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Parent)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	return n
}

func (m *MsgCreateAdd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Post)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Ipfs)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Parent)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	return n
}

func (m *MsgUpdateAdd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Post)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Ipfs)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Parent)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	return n
}

func (m *MsgDeleteAdd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	return n
}

func sovAdd(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdd(x uint64) (n int) {
	return sovAdd(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Add) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdd
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
			return fmt.Errorf("proto: Add: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Add: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Post", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Post = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ipfs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ipfs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdd
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdd
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
func (m *MsgCreateAdd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdd
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
			return fmt.Errorf("proto: MsgCreateAdd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateAdd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Post", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Post = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ipfs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ipfs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdd
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdd
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
func (m *MsgUpdateAdd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdd
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
			return fmt.Errorf("proto: MsgUpdateAdd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateAdd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Post", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Post = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ipfs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ipfs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdd
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdd
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
func (m *MsgDeleteAdd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdd
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
			return fmt.Errorf("proto: MsgDeleteAdd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteAdd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
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
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAdd
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAdd
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
func skipAdd(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdd
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
					return 0, ErrIntOverflowAdd
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
					return 0, ErrIntOverflowAdd
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
				return 0, ErrInvalidLengthAdd
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdd
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdd
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdd        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdd          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdd = fmt.Errorf("proto: unexpected end of group")
)