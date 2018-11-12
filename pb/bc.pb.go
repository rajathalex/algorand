// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bc.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represent a void message indicating success
type Success struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Success) Reset()         { *m = Success{} }
func (m *Success) String() string { return proto.CompactTextString(m) }
func (*Success) ProtoMessage()    {}
func (*Success) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e2a20f8b284799, []int{0}
}

func (m *Success) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Success.Unmarshal(m, b)
}
func (m *Success) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Success.Marshal(b, m, deterministic)
}
func (m *Success) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Success.Merge(m, src)
}
func (m *Success) XXX_Size() int {
	return xxx_messageInfo_Success.Size(m)
}
func (m *Success) XXX_DiscardUnknown() {
	xxx_messageInfo_Success.DiscardUnknown(m)
}

var xxx_messageInfo_Success proto.InternalMessageInfo

// Represents an error.
type Error struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e2a20f8b284799, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// A single Block on a Blockchain
type Block struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PrevHash             string   `protobuf:"bytes,3,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	Hash                 string   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e2a20f8b284799, []int{2}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Block) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Block) GetPrevHash() string {
	if m != nil {
		return m.PrevHash
	}
	return ""
}

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Block) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Input to AppendBlock
type AppendBlockArgs struct {
	// Block block = 1;
	V                    string   `protobuf:"bytes,1,opt,name=v,proto3" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendBlockArgs) Reset()         { *m = AppendBlockArgs{} }
func (m *AppendBlockArgs) String() string { return proto.CompactTextString(m) }
func (*AppendBlockArgs) ProtoMessage()    {}
func (*AppendBlockArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e2a20f8b284799, []int{3}
}

func (m *AppendBlockArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendBlockArgs.Unmarshal(m, b)
}
func (m *AppendBlockArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendBlockArgs.Marshal(b, m, deterministic)
}
func (m *AppendBlockArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendBlockArgs.Merge(m, src)
}
func (m *AppendBlockArgs) XXX_Size() int {
	return xxx_messageInfo_AppendBlockArgs.Size(m)
}
func (m *AppendBlockArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendBlockArgs.DiscardUnknown(m)
}

var xxx_messageInfo_AppendBlockArgs proto.InternalMessageInfo

func (m *AppendBlockArgs) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

type AppendBlockRet struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendBlockRet) Reset()         { *m = AppendBlockRet{} }
func (m *AppendBlockRet) String() string { return proto.CompactTextString(m) }
func (*AppendBlockRet) ProtoMessage()    {}
func (*AppendBlockRet) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e2a20f8b284799, []int{4}
}

func (m *AppendBlockRet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendBlockRet.Unmarshal(m, b)
}
func (m *AppendBlockRet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendBlockRet.Marshal(b, m, deterministic)
}
func (m *AppendBlockRet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendBlockRet.Merge(m, src)
}
func (m *AppendBlockRet) XXX_Size() int {
	return xxx_messageInfo_AppendBlockRet.Size(m)
}
func (m *AppendBlockRet) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendBlockRet.DiscardUnknown(m)
}

var xxx_messageInfo_AppendBlockRet proto.InternalMessageInfo

func (m *AppendBlockRet) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SIGArgs struct {
	V                    string   `protobuf:"bytes,1,opt,name=v,proto3" json:"v,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	P                    int64    `protobuf:"varint,3,opt,name=p,proto3" json:"p,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SIGArgs) Reset()         { *m = SIGArgs{} }
func (m *SIGArgs) String() string { return proto.CompactTextString(m) }
func (*SIGArgs) ProtoMessage()    {}
func (*SIGArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e2a20f8b284799, []int{5}
}

func (m *SIGArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SIGArgs.Unmarshal(m, b)
}
func (m *SIGArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SIGArgs.Marshal(b, m, deterministic)
}
func (m *SIGArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SIGArgs.Merge(m, src)
}
func (m *SIGArgs) XXX_Size() int {
	return xxx_messageInfo_SIGArgs.Size(m)
}
func (m *SIGArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_SIGArgs.DiscardUnknown(m)
}

var xxx_messageInfo_SIGArgs proto.InternalMessageInfo

func (m *SIGArgs) GetV() string {
	if m != nil {
		return m.V
	}
	return ""
}

func (m *SIGArgs) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SIGArgs) GetP() int64 {
	if m != nil {
		return m.P
	}
	return 0
}

type SIGRet struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SIGRet) Reset()         { *m = SIGRet{} }
func (m *SIGRet) String() string { return proto.CompactTextString(m) }
func (*SIGRet) ProtoMessage()    {}
func (*SIGRet) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e2a20f8b284799, []int{6}
}

func (m *SIGRet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SIGRet.Unmarshal(m, b)
}
func (m *SIGRet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SIGRet.Marshal(b, m, deterministic)
}
func (m *SIGRet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SIGRet.Merge(m, src)
}
func (m *SIGRet) XXX_Size() int {
	return xxx_messageInfo_SIGRet.Size(m)
}
func (m *SIGRet) XXX_DiscardUnknown() {
	xxx_messageInfo_SIGRet.DiscardUnknown(m)
}

var xxx_messageInfo_SIGRet proto.InternalMessageInfo

func (m *SIGRet) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Success)(nil), "pb.Success")
	proto.RegisterType((*Error)(nil), "pb.Error")
	proto.RegisterType((*Block)(nil), "pb.Block")
	proto.RegisterType((*AppendBlockArgs)(nil), "pb.AppendBlockArgs")
	proto.RegisterType((*AppendBlockRet)(nil), "pb.AppendBlockRet")
	proto.RegisterType((*SIGArgs)(nil), "pb.SIGArgs")
	proto.RegisterType((*SIGRet)(nil), "pb.SIGRet")
}

func init() { proto.RegisterFile("bc.proto", fileDescriptor_99e2a20f8b284799) }

var fileDescriptor_99e2a20f8b284799 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0x05, 0xda, 0x02, 0x53, 0x53, 0xcd, 0x78, 0x41, 0x62, 0x62, 0xb3, 0x27, 0xe3, 0x81,
	0x83, 0x5e, 0xf4, 0x88, 0x89, 0x41, 0xae, 0xcb, 0x27, 0xe0, 0xcf, 0x4a, 0x89, 0xa5, 0x6c, 0x76,
	0x57, 0x12, 0x4f, 0x7e, 0x75, 0xb3, 0x23, 0xd5, 0xa6, 0x26, 0xde, 0xe6, 0xbd, 0x79, 0x3b, 0xf9,
	0xbd, 0x2c, 0x04, 0x55, 0x9d, 0x48, 0x35, 0x98, 0x01, 0x5d, 0x59, 0xb1, 0x10, 0xfc, 0xe2, 0xbd,
	0xae, 0x85, 0xd6, 0xec, 0x12, 0xe6, 0xcf, 0x4a, 0x0d, 0x0a, 0xcf, 0xc1, 0xeb, 0x75, 0x1b, 0x39,
	0x6b, 0xe7, 0x26, 0xe4, 0x76, 0x64, 0x9f, 0x30, 0x7f, 0xda, 0x0e, 0xf5, 0x1b, 0xae, 0xc0, 0xed,
	0x1a, 0xda, 0x78, 0xdc, 0xed, 0x1a, 0xbc, 0x82, 0xd0, 0x74, 0xbd, 0xd0, 0xa6, 0xec, 0x65, 0xe4,
	0xd2, 0x83, 0x5f, 0x03, 0x63, 0x08, 0xa4, 0x12, 0xe3, 0x4b, 0xa9, 0x37, 0x91, 0x47, 0xcb, 0x1f,
	0x8d, 0x08, 0xb3, 0x8d, 0xf5, 0x67, 0xe4, 0xd3, 0x8c, 0x11, 0xf8, 0xbd, 0xd0, 0xba, 0x6c, 0x45,
	0x34, 0x27, 0x7b, 0x2f, 0xd9, 0x35, 0x9c, 0xa5, 0x52, 0x8a, 0x5d, 0x43, 0x18, 0xa9, 0x6a, 0x35,
	0x9e, 0x82, 0x33, 0x4e, 0x8c, 0xce, 0xc8, 0x6e, 0x61, 0x75, 0x10, 0xe0, 0xc2, 0xd8, 0x63, 0xfa,
	0xbb, 0x19, 0xa5, 0x02, 0xbe, 0x97, 0xec, 0x11, 0xfc, 0x22, 0xcf, 0xfe, 0x1e, 0xb1, 0x4c, 0xe6,
	0x43, 0x8a, 0xa9, 0x08, 0xcd, 0x36, 0x21, 0x09, 0xde, 0xe3, 0x8e, 0x64, 0x0c, 0x16, 0x45, 0x9e,
	0xfd, 0x7b, 0xfe, 0xee, 0x15, 0x82, 0x74, 0xdb, 0x0e, 0xaa, 0xdc, 0x35, 0xf8, 0x00, 0xcb, 0x03,
	0x2c, 0xbc, 0x48, 0x64, 0x95, 0x1c, 0x15, 0x89, 0xf1, 0xc8, 0xe4, 0xc2, 0xb0, 0x13, 0x5c, 0x83,
	0x57, 0xe4, 0x19, 0x2e, 0xed, 0x72, 0xa2, 0x8d, 0x61, 0x12, 0x94, 0xa8, 0x16, 0xf4, 0x8b, 0xf7,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xb1, 0x5a, 0x42, 0xd1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlgorandClient is the client API for Algorand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlgorandClient interface {
	AppendBlock(ctx context.Context, in *AppendBlockArgs, opts ...grpc.CallOption) (*AppendBlockRet, error)
	SIG(ctx context.Context, in *SIGArgs, opts ...grpc.CallOption) (*SIGRet, error)
}

type algorandClient struct {
	cc *grpc.ClientConn
}

func NewAlgorandClient(cc *grpc.ClientConn) AlgorandClient {
	return &algorandClient{cc}
}

func (c *algorandClient) AppendBlock(ctx context.Context, in *AppendBlockArgs, opts ...grpc.CallOption) (*AppendBlockRet, error) {
	out := new(AppendBlockRet)
	err := c.cc.Invoke(ctx, "/pb.Algorand/AppendBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algorandClient) SIG(ctx context.Context, in *SIGArgs, opts ...grpc.CallOption) (*SIGRet, error) {
	out := new(SIGRet)
	err := c.cc.Invoke(ctx, "/pb.Algorand/SIG", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlgorandServer is the server API for Algorand service.
type AlgorandServer interface {
	AppendBlock(context.Context, *AppendBlockArgs) (*AppendBlockRet, error)
	SIG(context.Context, *SIGArgs) (*SIGRet, error)
}

func RegisterAlgorandServer(s *grpc.Server, srv AlgorandServer) {
	s.RegisterService(&_Algorand_serviceDesc, srv)
}

func _Algorand_AppendBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendBlockArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorandServer).AppendBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Algorand/AppendBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorandServer).AppendBlock(ctx, req.(*AppendBlockArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Algorand_SIG_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SIGArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorandServer).SIG(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Algorand/SIG",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorandServer).SIG(ctx, req.(*SIGArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Algorand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Algorand",
	HandlerType: (*AlgorandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendBlock",
			Handler:    _Algorand_AppendBlock_Handler,
		},
		{
			MethodName: "SIG",
			Handler:    _Algorand_SIG_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bc.proto",
}