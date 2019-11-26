// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet.proto

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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type ServerStreamRequest struct {
	MaxPrime             int64    `protobuf:"varint,1,opt,name=MaxPrime,proto3" json:"MaxPrime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStreamRequest) Reset()         { *m = ServerStreamRequest{} }
func (m *ServerStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ServerStreamRequest) ProtoMessage()    {}
func (*ServerStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{3}
}

func (m *ServerStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStreamRequest.Unmarshal(m, b)
}
func (m *ServerStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStreamRequest.Marshal(b, m, deterministic)
}
func (m *ServerStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStreamRequest.Merge(m, src)
}
func (m *ServerStreamRequest) XXX_Size() int {
	return xxx_messageInfo_ServerStreamRequest.Size(m)
}
func (m *ServerStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStreamRequest proto.InternalMessageInfo

func (m *ServerStreamRequest) GetMaxPrime() int64 {
	if m != nil {
		return m.MaxPrime
	}
	return 0
}

type ServerStreamResponse struct {
	Prime                int64    `protobuf:"varint,1,opt,name=Prime,proto3" json:"Prime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStreamResponse) Reset()         { *m = ServerStreamResponse{} }
func (m *ServerStreamResponse) String() string { return proto.CompactTextString(m) }
func (*ServerStreamResponse) ProtoMessage()    {}
func (*ServerStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{4}
}

func (m *ServerStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStreamResponse.Unmarshal(m, b)
}
func (m *ServerStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStreamResponse.Marshal(b, m, deterministic)
}
func (m *ServerStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStreamResponse.Merge(m, src)
}
func (m *ServerStreamResponse) XXX_Size() int {
	return xxx_messageInfo_ServerStreamResponse.Size(m)
}
func (m *ServerStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStreamResponse proto.InternalMessageInfo

func (m *ServerStreamResponse) GetPrime() int64 {
	if m != nil {
		return m.Prime
	}
	return 0
}

type ClientStreamRequest struct {
	Index                int64    `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Number               int64    `protobuf:"varint,2,opt,name=Number,proto3" json:"Number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStreamRequest) Reset()         { *m = ClientStreamRequest{} }
func (m *ClientStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ClientStreamRequest) ProtoMessage()    {}
func (*ClientStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{5}
}

func (m *ClientStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStreamRequest.Unmarshal(m, b)
}
func (m *ClientStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStreamRequest.Marshal(b, m, deterministic)
}
func (m *ClientStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStreamRequest.Merge(m, src)
}
func (m *ClientStreamRequest) XXX_Size() int {
	return xxx_messageInfo_ClientStreamRequest.Size(m)
}
func (m *ClientStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStreamRequest proto.InternalMessageInfo

func (m *ClientStreamRequest) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ClientStreamRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ClientStreamResponse struct {
	StartIndex           int64    `protobuf:"varint,1,opt,name=StartIndex,proto3" json:"StartIndex,omitempty"`
	EndIndex             int64    `protobuf:"varint,2,opt,name=EndIndex,proto3" json:"EndIndex,omitempty"`
	Sum                  int64    `protobuf:"varint,3,opt,name=Sum,proto3" json:"Sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStreamResponse) Reset()         { *m = ClientStreamResponse{} }
func (m *ClientStreamResponse) String() string { return proto.CompactTextString(m) }
func (*ClientStreamResponse) ProtoMessage()    {}
func (*ClientStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{6}
}

func (m *ClientStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStreamResponse.Unmarshal(m, b)
}
func (m *ClientStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStreamResponse.Marshal(b, m, deterministic)
}
func (m *ClientStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStreamResponse.Merge(m, src)
}
func (m *ClientStreamResponse) XXX_Size() int {
	return xxx_messageInfo_ClientStreamResponse.Size(m)
}
func (m *ClientStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStreamResponse proto.InternalMessageInfo

func (m *ClientStreamResponse) GetStartIndex() int64 {
	if m != nil {
		return m.StartIndex
	}
	return 0
}

func (m *ClientStreamResponse) GetEndIndex() int64 {
	if m != nil {
		return m.EndIndex
	}
	return 0
}

func (m *ClientStreamResponse) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

type BiDirStreamRequest struct {
	Index                int64    `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Number               int64    `protobuf:"varint,2,opt,name=Number,proto3" json:"Number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiDirStreamRequest) Reset()         { *m = BiDirStreamRequest{} }
func (m *BiDirStreamRequest) String() string { return proto.CompactTextString(m) }
func (*BiDirStreamRequest) ProtoMessage()    {}
func (*BiDirStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{7}
}

func (m *BiDirStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiDirStreamRequest.Unmarshal(m, b)
}
func (m *BiDirStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiDirStreamRequest.Marshal(b, m, deterministic)
}
func (m *BiDirStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiDirStreamRequest.Merge(m, src)
}
func (m *BiDirStreamRequest) XXX_Size() int {
	return xxx_messageInfo_BiDirStreamRequest.Size(m)
}
func (m *BiDirStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BiDirStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BiDirStreamRequest proto.InternalMessageInfo

func (m *BiDirStreamRequest) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BiDirStreamRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type BiDirStreamResponse struct {
	StartIndex           int64    `protobuf:"varint,1,opt,name=StartIndex,proto3" json:"StartIndex,omitempty"`
	EndIndex             int64    `protobuf:"varint,2,opt,name=EndIndex,proto3" json:"EndIndex,omitempty"`
	Sum                  int64    `protobuf:"varint,3,opt,name=Sum,proto3" json:"Sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiDirStreamResponse) Reset()         { *m = BiDirStreamResponse{} }
func (m *BiDirStreamResponse) String() string { return proto.CompactTextString(m) }
func (*BiDirStreamResponse) ProtoMessage()    {}
func (*BiDirStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{8}
}

func (m *BiDirStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiDirStreamResponse.Unmarshal(m, b)
}
func (m *BiDirStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiDirStreamResponse.Marshal(b, m, deterministic)
}
func (m *BiDirStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiDirStreamResponse.Merge(m, src)
}
func (m *BiDirStreamResponse) XXX_Size() int {
	return xxx_messageInfo_BiDirStreamResponse.Size(m)
}
func (m *BiDirStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BiDirStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BiDirStreamResponse proto.InternalMessageInfo

func (m *BiDirStreamResponse) GetStartIndex() int64 {
	if m != nil {
		return m.StartIndex
	}
	return 0
}

func (m *BiDirStreamResponse) GetEndIndex() int64 {
	if m != nil {
		return m.EndIndex
	}
	return 0
}

func (m *BiDirStreamResponse) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func init() {
	proto.RegisterType((*Greeting)(nil), "Greeting")
	proto.RegisterType((*GreetRequest)(nil), "GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "GreetResponse")
	proto.RegisterType((*ServerStreamRequest)(nil), "ServerStreamRequest")
	proto.RegisterType((*ServerStreamResponse)(nil), "ServerStreamResponse")
	proto.RegisterType((*ClientStreamRequest)(nil), "ClientStreamRequest")
	proto.RegisterType((*ClientStreamResponse)(nil), "ClientStreamResponse")
	proto.RegisterType((*BiDirStreamRequest)(nil), "BiDirStreamRequest")
	proto.RegisterType((*BiDirStreamResponse)(nil), "BiDirStreamResponse")
}

func init() { proto.RegisterFile("greet.proto", fileDescriptor_32c0044392f32579) }

var fileDescriptor_32c0044392f32579 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcf, 0x6b, 0xea, 0x40,
	0x10, 0x36, 0x2f, 0x28, 0xc9, 0xa8, 0x8f, 0xc7, 0x26, 0x3e, 0x24, 0x8f, 0x57, 0xca, 0x42, 0xa9,
	0x87, 0xb2, 0x58, 0x4b, 0xaf, 0x2d, 0x68, 0x7f, 0xd0, 0x43, 0xa5, 0x24, 0xb7, 0x5e, 0x4a, 0xd4,
	0xa9, 0x04, 0x4c, 0xb4, 0x9b, 0x4d, 0xf1, 0x6f, 0xef, 0xa9, 0xec, 0x66, 0x23, 0x1b, 0x9a, 0x5b,
	0xe9, 0x49, 0xbf, 0xf9, 0x32, 0xdf, 0x37, 0x99, 0x6f, 0x02, 0xdd, 0x35, 0x47, 0x14, 0x6c, 0xc7,
	0xb7, 0x62, 0x4b, 0xef, 0xc0, 0xb9, 0x97, 0x30, 0xc9, 0xd6, 0xe4, 0x3f, 0xc0, 0x6b, 0xc2, 0x73,
	0xf1, 0x92, 0xc5, 0x29, 0x0e, 0xad, 0x63, 0x6b, 0xe4, 0x86, 0xae, 0xaa, 0xcc, 0xe3, 0x14, 0xc9,
	0x3f, 0x70, 0x37, 0x71, 0xc5, 0xfe, 0x52, 0xac, 0x23, 0x0b, 0x92, 0xa4, 0x97, 0xd0, 0x53, 0x3a,
	0x21, 0xbe, 0x15, 0x98, 0x0b, 0x72, 0x02, 0xce, 0x5a, 0xeb, 0x2a, 0xa5, 0xee, 0xc4, 0x65, 0x95,
	0x51, 0x78, 0xa0, 0xe8, 0x29, 0xf4, 0x75, 0x5b, 0xbe, 0xdb, 0x66, 0x39, 0x92, 0xbf, 0xd0, 0xe1,
	0x98, 0x17, 0x1b, 0xa1, 0xfd, 0x35, 0xa2, 0xe7, 0xe0, 0x45, 0xc8, 0xdf, 0x91, 0x47, 0x82, 0x63,
	0x9c, 0x56, 0x36, 0x01, 0x38, 0x8f, 0xf1, 0xfe, 0x89, 0x27, 0x7a, 0x60, 0x3b, 0x3c, 0x60, 0x7a,
	0x06, 0x7e, 0xbd, 0x45, 0x5b, 0xf8, 0xd0, 0x36, 0x1b, 0x4a, 0x40, 0x67, 0xe0, 0xcd, 0x36, 0x09,
	0x66, 0xa2, 0x6e, 0xe0, 0x43, 0xfb, 0x21, 0x5b, 0xe1, 0xbe, 0x7a, 0x58, 0x01, 0x39, 0xe5, 0xbc,
	0x48, 0x17, 0xc8, 0xd5, 0x1e, 0xec, 0x50, 0x23, 0xba, 0x02, 0xbf, 0x2e, 0xa2, 0x2d, 0x8f, 0x00,
	0x22, 0x11, 0x73, 0x61, 0x4a, 0x19, 0x15, 0xf9, 0x1a, 0xb7, 0xd9, 0xaa, 0x64, 0x4b, 0xc5, 0x03,
	0x26, 0x7f, 0xc0, 0x8e, 0x8a, 0x74, 0x68, 0xab, 0xb2, 0xfc, 0x4b, 0xa7, 0x40, 0xa6, 0xc9, 0x4d,
	0xc2, 0xbf, 0x33, 0xe9, 0x12, 0xbc, 0x9a, 0xc6, 0x4f, 0x0c, 0x3a, 0xf9, 0xb0, 0xf4, 0x55, 0xc8,
	0x1c, 0x92, 0x25, 0x92, 0x11, 0xb4, 0x15, 0x26, 0x7d, 0x66, 0x5e, 0x4b, 0xf0, 0x9b, 0xd5, 0xae,
	0x80, 0xb6, 0xc8, 0x35, 0xf4, 0xcc, 0xf0, 0x88, 0xcf, 0x1a, 0xe2, 0x0f, 0x06, 0xac, 0x29, 0x61,
	0xda, 0x1a, 0x5b, 0x52, 0xc0, 0x8c, 0x82, 0xf8, 0xac, 0x21, 0xde, 0x60, 0xc0, 0x9a, 0xf2, 0xa2,
	0xad, 0x91, 0x45, 0xae, 0xa0, 0x6b, 0x6c, 0x88, 0x78, 0xec, 0xeb, 0xce, 0x03, 0x9f, 0x35, 0x2c,
	0x51, 0x76, 0x8f, 0xad, 0xa9, 0xf3, 0xdc, 0x51, 0x9f, 0x58, 0xbe, 0x28, 0x7f, 0x2f, 0x3e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x00, 0x38, 0xef, 0xc7, 0x79, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	// server streaming
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...grpc.CallOption) (GreetService_ServerStreamClient, error)
	// client streaming
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_ClientStreamClient, error)
	// bi-directional streaming
	BiDirStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_BiDirStreamClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...grpc.CallOption) (GreetService_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/GreetService/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_ServerStreamClient interface {
	Recv() (*ServerStreamResponse, error)
	grpc.ClientStream
}

type greetServiceServerStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceServerStreamClient) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[1], "/GreetService/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceClientStreamClient{stream}
	return x, nil
}

type GreetService_ClientStreamClient interface {
	Send(*ClientStreamRequest) error
	CloseAndRecv() (*ClientStreamResponse, error)
	grpc.ClientStream
}

type greetServiceClientStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceClientStreamClient) Send(m *ClientStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceClientStreamClient) CloseAndRecv() (*ClientStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ClientStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) BiDirStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_BiDirStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[2], "/GreetService/BiDirStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceBiDirStreamClient{stream}
	return x, nil
}

type GreetService_BiDirStreamClient interface {
	Send(*BiDirStreamRequest) error
	Recv() (*BiDirStreamResponse, error)
	grpc.ClientStream
}

type greetServiceBiDirStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceBiDirStreamClient) Send(m *BiDirStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceBiDirStreamClient) Recv() (*BiDirStreamResponse, error) {
	m := new(BiDirStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// unary
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	// server streaming
	ServerStream(*ServerStreamRequest, GreetService_ServerStreamServer) error
	// client streaming
	ClientStream(GreetService_ClientStreamServer) error
	// bi-directional streaming
	BiDirStream(GreetService_BiDirStreamServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreetServiceServer) ServerStream(req *ServerStreamRequest, srv GreetService_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (*UnimplementedGreetServiceServer) ClientStream(srv GreetService_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (*UnimplementedGreetServiceServer) BiDirStream(srv GreetService_BiDirStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirStream not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).ServerStream(m, &greetServiceServerStreamServer{stream})
}

type GreetService_ServerStreamServer interface {
	Send(*ServerStreamResponse) error
	grpc.ServerStream
}

type greetServiceServerStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceServerStreamServer) Send(m *ServerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).ClientStream(&greetServiceClientStreamServer{stream})
}

type GreetService_ClientStreamServer interface {
	SendAndClose(*ClientStreamResponse) error
	Recv() (*ClientStreamRequest, error)
	grpc.ServerStream
}

type greetServiceClientStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceClientStreamServer) SendAndClose(m *ClientStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceClientStreamServer) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_BiDirStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).BiDirStream(&greetServiceBiDirStreamServer{stream})
}

type GreetService_BiDirStreamServer interface {
	Send(*BiDirStreamResponse) error
	Recv() (*BiDirStreamRequest, error)
	grpc.ServerStream
}

type greetServiceBiDirStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceBiDirStreamServer) Send(m *BiDirStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceBiDirStreamServer) Recv() (*BiDirStreamRequest, error) {
	m := new(BiDirStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _GreetService_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStream",
			Handler:       _GreetService_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BiDirStream",
			Handler:       _GreetService_BiDirStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet.proto",
}
