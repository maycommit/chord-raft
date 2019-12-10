// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chord.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ID struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_541dae51990542ec, []int{0}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Node struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_541dae51990542ec, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Key struct {
	Key                  int64    `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_541dae51990542ec, []int{2}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

type Value struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_541dae51990542ec, []int{3}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Data struct {
	Key                  int64    `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_541dae51990542ec, []int{4}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Data) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Datas struct {
	Datas                []*Data  `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Datas) Reset()         { *m = Datas{} }
func (m *Datas) String() string { return proto.CompactTextString(m) }
func (*Datas) ProtoMessage()    {}
func (*Datas) Descriptor() ([]byte, []int) {
	return fileDescriptor_541dae51990542ec, []int{5}
}

func (m *Datas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Datas.Unmarshal(m, b)
}
func (m *Datas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Datas.Marshal(b, m, deterministic)
}
func (m *Datas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Datas.Merge(m, src)
}
func (m *Datas) XXX_Size() int {
	return xxx_messageInfo_Datas.Size(m)
}
func (m *Datas) XXX_DiscardUnknown() {
	xxx_messageInfo_Datas.DiscardUnknown(m)
}

var xxx_messageInfo_Datas proto.InternalMessageInfo

func (m *Datas) GetDatas() []*Data {
	if m != nil {
		return m.Datas
	}
	return nil
}

type Any struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Any) Reset()         { *m = Any{} }
func (m *Any) String() string { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()    {}
func (*Any) Descriptor() ([]byte, []int) {
	return fileDescriptor_541dae51990542ec, []int{6}
}

func (m *Any) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Any.Unmarshal(m, b)
}
func (m *Any) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Any.Marshal(b, m, deterministic)
}
func (m *Any) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Any.Merge(m, src)
}
func (m *Any) XXX_Size() int {
	return xxx_messageInfo_Any.Size(m)
}
func (m *Any) XXX_DiscardUnknown() {
	xxx_messageInfo_Any.DiscardUnknown(m)
}

var xxx_messageInfo_Any proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ID)(nil), "protos.ID")
	proto.RegisterType((*Node)(nil), "protos.Node")
	proto.RegisterType((*Key)(nil), "protos.Key")
	proto.RegisterType((*Value)(nil), "protos.Value")
	proto.RegisterType((*Data)(nil), "protos.Data")
	proto.RegisterType((*Datas)(nil), "protos.Datas")
	proto.RegisterType((*Any)(nil), "protos.Any")
}

func init() { proto.RegisterFile("chord.proto", fileDescriptor_541dae51990542ec) }

var fileDescriptor_541dae51990542ec = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x6b, 0xc2, 0x30,
	0x10, 0xc6, 0x69, 0x6b, 0x37, 0x3c, 0xe7, 0xe6, 0x82, 0x6c, 0x45, 0x18, 0x48, 0x1e, 0x86, 0x4c,
	0x29, 0xc3, 0xfd, 0x05, 0xa2, 0x4c, 0x44, 0x10, 0xa9, 0xb0, 0xf7, 0xae, 0xb9, 0xb9, 0xb2, 0xd2,
	0x8c, 0x24, 0x0e, 0xf2, 0x9f, 0xef, 0x71, 0xa4, 0x59, 0xb5, 0x55, 0x18, 0x7b, 0xea, 0xdd, 0x7d,
	0xbf, 0xef, 0x6b, 0xaf, 0x09, 0xb4, 0x92, 0x77, 0x2e, 0x58, 0xf8, 0x29, 0xb8, 0xe2, 0xe4, 0xac,
	0x78, 0x48, 0xda, 0x05, 0x77, 0x31, 0x23, 0x97, 0xe0, 0xa6, 0x2c, 0x70, 0xfa, 0xce, 0xc0, 0x8b,
	0xdc, 0x94, 0xd1, 0x47, 0x68, 0xac, 0x38, 0xc3, 0xe3, 0x39, 0x09, 0xe0, 0x3c, 0x66, 0x4c, 0xa0,
	0x94, 0x81, 0xdb, 0x77, 0x06, 0xcd, 0xa8, 0x6c, 0xe9, 0x2d, 0x78, 0x4b, 0xd4, 0xa4, 0x03, 0xde,
	0x07, 0xea, 0x5f, 0x87, 0x29, 0xe9, 0x1d, 0xf8, 0x2f, 0x71, 0xb6, 0x43, 0xd2, 0x05, 0xff, 0xcb,
	0x14, 0x85, 0xd8, 0x8c, 0x6c, 0x43, 0x43, 0x68, 0xcc, 0x62, 0x15, 0x9f, 0x1a, 0x0f, 0xbc, 0x5b,
	0xe5, 0x87, 0xe0, 0x1b, 0x5e, 0x12, 0x0a, 0x3e, 0x33, 0x45, 0xe0, 0xf4, 0xbd, 0x41, 0x6b, 0x7c,
	0x61, 0xf7, 0x92, 0xa1, 0x51, 0x23, 0x2b, 0x51, 0x1f, 0xbc, 0x49, 0xae, 0xc7, 0xdf, 0x2e, 0xf8,
	0x53, 0xb3, 0x3b, 0x19, 0x41, 0xe7, 0x39, 0xcd, 0xd9, 0x66, 0x97, 0x24, 0x28, 0x25, 0x17, 0xd1,
	0x7a, 0x4a, 0xa0, 0x74, 0x2e, 0x66, 0xbd, 0x7d, 0x4a, 0xb1, 0xfd, 0x18, 0x6e, 0xa6, 0x19, 0x97,
	0xa8, 0xd6, 0x02, 0x13, 0x64, 0x69, 0xbe, 0x35, 0xe3, 0xbf, 0x3d, 0x23, 0xb8, 0x9a, 0xa3, 0xaa,
	0xbd, 0xa0, 0x55, 0x02, 0x93, 0x5c, 0x1f, 0xd1, 0x21, 0x5c, 0xcf, 0x8b, 0x78, 0x86, 0xff, 0xe2,
	0xef, 0xa1, 0xb9, 0xe2, 0x2a, 0x7d, 0xd3, 0x86, 0xab, 0x49, 0xbd, 0xaa, 0x8b, 0x0c, 0xa1, 0xbd,
	0x51, 0x5c, 0xc4, 0x5b, 0x9c, 0xa3, 0xaa, 0x65, 0x2e, 0x51, 0xf7, 0xda, 0x65, 0x63, 0x0f, 0xe6,
	0x61, 0x0f, 0x6f, 0x2c, 0x5c, 0xfb, 0x97, 0xf5, 0xe0, 0x10, 0x3a, 0x87, 0xe0, 0x49, 0x96, 0x9d,
	0x7c, 0x6f, 0xbb, 0xea, 0x95, 0xaf, 0xf6, 0x9a, 0x3d, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6e,
	0x4f, 0x2e, 0x53, 0x7c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChordClient is the client API for Chord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChordClient interface {
	FindSuccessorRPC(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error)
	ClosetPrecedingNodeRPC(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error)
	GetSuccessorRPC(ctx context.Context, in *Any, opts ...grpc.CallOption) (*Node, error)
	GetPredecessorRPC(ctx context.Context, in *Any, opts ...grpc.CallOption) (*Node, error)
	NotifyRPC(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Any, error)
	StorageGetRPC(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	StorageSetRPC(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Any, error)
	StorageGetAllRPC(ctx context.Context, in *Any, opts ...grpc.CallOption) (*Datas, error)
}

type chordClient struct {
	cc *grpc.ClientConn
}

func NewChordClient(cc *grpc.ClientConn) ChordClient {
	return &chordClient{cc}
}

func (c *chordClient) FindSuccessorRPC(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/protos.Chord/FindSuccessorRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) ClosetPrecedingNodeRPC(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/protos.Chord/ClosetPrecedingNodeRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) GetSuccessorRPC(ctx context.Context, in *Any, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/protos.Chord/GetSuccessorRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) GetPredecessorRPC(ctx context.Context, in *Any, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/protos.Chord/GetPredecessorRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) NotifyRPC(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Any, error) {
	out := new(Any)
	err := c.cc.Invoke(ctx, "/protos.Chord/NotifyRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) StorageGetRPC(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/protos.Chord/StorageGetRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) StorageSetRPC(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Any, error) {
	out := new(Any)
	err := c.cc.Invoke(ctx, "/protos.Chord/StorageSetRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) StorageGetAllRPC(ctx context.Context, in *Any, opts ...grpc.CallOption) (*Datas, error) {
	out := new(Datas)
	err := c.cc.Invoke(ctx, "/protos.Chord/StorageGetAllRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChordServer is the server API for Chord service.
type ChordServer interface {
	FindSuccessorRPC(context.Context, *ID) (*Node, error)
	ClosetPrecedingNodeRPC(context.Context, *ID) (*Node, error)
	GetSuccessorRPC(context.Context, *Any) (*Node, error)
	GetPredecessorRPC(context.Context, *Any) (*Node, error)
	NotifyRPC(context.Context, *Node) (*Any, error)
	StorageGetRPC(context.Context, *Key) (*Value, error)
	StorageSetRPC(context.Context, *Data) (*Any, error)
	StorageGetAllRPC(context.Context, *Any) (*Datas, error)
}

// UnimplementedChordServer can be embedded to have forward compatible implementations.
type UnimplementedChordServer struct {
}

func (*UnimplementedChordServer) FindSuccessorRPC(ctx context.Context, req *ID) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessorRPC not implemented")
}
func (*UnimplementedChordServer) ClosetPrecedingNodeRPC(ctx context.Context, req *ID) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClosetPrecedingNodeRPC not implemented")
}
func (*UnimplementedChordServer) GetSuccessorRPC(ctx context.Context, req *Any) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuccessorRPC not implemented")
}
func (*UnimplementedChordServer) GetPredecessorRPC(ctx context.Context, req *Any) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredecessorRPC not implemented")
}
func (*UnimplementedChordServer) NotifyRPC(ctx context.Context, req *Node) (*Any, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyRPC not implemented")
}
func (*UnimplementedChordServer) StorageGetRPC(ctx context.Context, req *Key) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageGetRPC not implemented")
}
func (*UnimplementedChordServer) StorageSetRPC(ctx context.Context, req *Data) (*Any, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageSetRPC not implemented")
}
func (*UnimplementedChordServer) StorageGetAllRPC(ctx context.Context, req *Any) (*Datas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageGetAllRPC not implemented")
}

func RegisterChordServer(s *grpc.Server, srv ChordServer) {
	s.RegisterService(&_Chord_serviceDesc, srv)
}

func _Chord_FindSuccessorRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).FindSuccessorRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/FindSuccessorRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).FindSuccessorRPC(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_ClosetPrecedingNodeRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).ClosetPrecedingNodeRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/ClosetPrecedingNodeRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).ClosetPrecedingNodeRPC(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_GetSuccessorRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Any)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetSuccessorRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/GetSuccessorRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetSuccessorRPC(ctx, req.(*Any))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_GetPredecessorRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Any)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetPredecessorRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/GetPredecessorRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetPredecessorRPC(ctx, req.(*Any))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_NotifyRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).NotifyRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/NotifyRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).NotifyRPC(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_StorageGetRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).StorageGetRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/StorageGetRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).StorageGetRPC(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_StorageSetRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).StorageSetRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/StorageSetRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).StorageSetRPC(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_StorageGetAllRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Any)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).StorageGetAllRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/StorageGetAllRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).StorageGetAllRPC(ctx, req.(*Any))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chord_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Chord",
	HandlerType: (*ChordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindSuccessorRPC",
			Handler:    _Chord_FindSuccessorRPC_Handler,
		},
		{
			MethodName: "ClosetPrecedingNodeRPC",
			Handler:    _Chord_ClosetPrecedingNodeRPC_Handler,
		},
		{
			MethodName: "GetSuccessorRPC",
			Handler:    _Chord_GetSuccessorRPC_Handler,
		},
		{
			MethodName: "GetPredecessorRPC",
			Handler:    _Chord_GetPredecessorRPC_Handler,
		},
		{
			MethodName: "NotifyRPC",
			Handler:    _Chord_NotifyRPC_Handler,
		},
		{
			MethodName: "StorageGetRPC",
			Handler:    _Chord_StorageGetRPC_Handler,
		},
		{
			MethodName: "StorageSetRPC",
			Handler:    _Chord_StorageSetRPC_Handler,
		},
		{
			MethodName: "StorageGetAllRPC",
			Handler:    _Chord_StorageGetAllRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chord.proto",
}
