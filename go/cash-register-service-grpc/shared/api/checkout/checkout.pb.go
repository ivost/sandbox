// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkout/checkout.proto

// register checkout service

package checkout

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Item struct {
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Price                float32  `protobuf:"fixed32,8,opt,name=price,proto3" json:"price,omitempty"`
	Valid                bool     `protobuf:"varint,10,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89273fd6fb3348, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type AddItemRequest struct {
	Item                 *Item    `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddItemRequest) Reset()         { *m = AddItemRequest{} }
func (m *AddItemRequest) String() string { return proto.CompactTextString(m) }
func (*AddItemRequest) ProtoMessage()    {}
func (*AddItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89273fd6fb3348, []int{1}
}

func (m *AddItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddItemRequest.Unmarshal(m, b)
}
func (m *AddItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddItemRequest.Marshal(b, m, deterministic)
}
func (m *AddItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddItemRequest.Merge(m, src)
}
func (m *AddItemRequest) XXX_Size() int {
	return xxx_messageInfo_AddItemRequest.Size(m)
}
func (m *AddItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddItemRequest proto.InternalMessageInfo

func (m *AddItemRequest) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type AddItemResponse struct {
	Item                 *Item    `protobuf:"bytes,4,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddItemResponse) Reset()         { *m = AddItemResponse{} }
func (m *AddItemResponse) String() string { return proto.CompactTextString(m) }
func (*AddItemResponse) ProtoMessage()    {}
func (*AddItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89273fd6fb3348, []int{2}
}

func (m *AddItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddItemResponse.Unmarshal(m, b)
}
func (m *AddItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddItemResponse.Marshal(b, m, deterministic)
}
func (m *AddItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddItemResponse.Merge(m, src)
}
func (m *AddItemResponse) XXX_Size() int {
	return xxx_messageInfo_AddItemResponse.Size(m)
}
func (m *AddItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddItemResponse proto.InternalMessageInfo

func (m *AddItemResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type LineItem struct {
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Qty                  int32    `protobuf:"varint,4,opt,name=qty,proto3" json:"qty,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Price                float32  `protobuf:"fixed32,8,opt,name=price,proto3" json:"price,omitempty"`
	Extprice             float32  `protobuf:"fixed32,10,opt,name=extprice,proto3" json:"extprice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LineItem) Reset()         { *m = LineItem{} }
func (m *LineItem) String() string { return proto.CompactTextString(m) }
func (*LineItem) ProtoMessage()    {}
func (*LineItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89273fd6fb3348, []int{3}
}

func (m *LineItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineItem.Unmarshal(m, b)
}
func (m *LineItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineItem.Marshal(b, m, deterministic)
}
func (m *LineItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineItem.Merge(m, src)
}
func (m *LineItem) XXX_Size() int {
	return xxx_messageInfo_LineItem.Size(m)
}
func (m *LineItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LineItem.DiscardUnknown(m)
}

var xxx_messageInfo_LineItem proto.InternalMessageInfo

func (m *LineItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LineItem) GetQty() int32 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *LineItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LineItem) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *LineItem) GetExtprice() float32 {
	if m != nil {
		return m.Extprice
	}
	return 0
}

type Receipt struct {
	Register             string      `protobuf:"bytes,2,opt,name=register,proto3" json:"register,omitempty"`
	Items                []*LineItem `protobuf:"bytes,6,rep,name=items,proto3" json:"items,omitempty"`
	Time                 int64       `protobuf:"varint,10,opt,name=time,proto3" json:"time,omitempty"`
	Total                float32     `protobuf:"fixed32,12,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89273fd6fb3348, []int{4}
}

func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receipt.Unmarshal(m, b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return xxx_messageInfo_Receipt.Size(m)
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetRegister() string {
	if m != nil {
		return m.Register
	}
	return ""
}

func (m *Receipt) GetItems() []*LineItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Receipt) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Receipt) GetTotal() float32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type ScanItem struct {
	Upc                  string   `protobuf:"bytes,6,opt,name=upc,proto3" json:"upc,omitempty"`
	Time                 int64    `protobuf:"varint,10,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanItem) Reset()         { *m = ScanItem{} }
func (m *ScanItem) String() string { return proto.CompactTextString(m) }
func (*ScanItem) ProtoMessage()    {}
func (*ScanItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89273fd6fb3348, []int{5}
}

func (m *ScanItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanItem.Unmarshal(m, b)
}
func (m *ScanItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanItem.Marshal(b, m, deterministic)
}
func (m *ScanItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanItem.Merge(m, src)
}
func (m *ScanItem) XXX_Size() int {
	return xxx_messageInfo_ScanItem.Size(m)
}
func (m *ScanItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanItem.DiscardUnknown(m)
}

var xxx_messageInfo_ScanItem proto.InternalMessageInfo

func (m *ScanItem) GetUpc() string {
	if m != nil {
		return m.Upc
	}
	return ""
}

func (m *ScanItem) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type CheckoutRequest struct {
	Register             string      `protobuf:"bytes,2,opt,name=register,proto3" json:"register,omitempty"`
	Items                []*ScanItem `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CheckoutRequest) Reset()         { *m = CheckoutRequest{} }
func (m *CheckoutRequest) String() string { return proto.CompactTextString(m) }
func (*CheckoutRequest) ProtoMessage()    {}
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89273fd6fb3348, []int{6}
}

func (m *CheckoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckoutRequest.Unmarshal(m, b)
}
func (m *CheckoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckoutRequest.Marshal(b, m, deterministic)
}
func (m *CheckoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckoutRequest.Merge(m, src)
}
func (m *CheckoutRequest) XXX_Size() int {
	return xxx_messageInfo_CheckoutRequest.Size(m)
}
func (m *CheckoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckoutRequest proto.InternalMessageInfo

func (m *CheckoutRequest) GetRegister() string {
	if m != nil {
		return m.Register
	}
	return ""
}

func (m *CheckoutRequest) GetItems() []*ScanItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// end checkout  will return the receipt
type CheckoutResponse struct {
	Receipt              *Receipt `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckoutResponse) Reset()         { *m = CheckoutResponse{} }
func (m *CheckoutResponse) String() string { return proto.CompactTextString(m) }
func (*CheckoutResponse) ProtoMessage()    {}
func (*CheckoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89273fd6fb3348, []int{7}
}

func (m *CheckoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckoutResponse.Unmarshal(m, b)
}
func (m *CheckoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckoutResponse.Marshal(b, m, deterministic)
}
func (m *CheckoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckoutResponse.Merge(m, src)
}
func (m *CheckoutResponse) XXX_Size() int {
	return xxx_messageInfo_CheckoutResponse.Size(m)
}
func (m *CheckoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckoutResponse proto.InternalMessageInfo

func (m *CheckoutResponse) GetReceipt() *Receipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

// field numbers 1 through 15 take one byte to encode
type HealthResponse struct {
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Time                 string   `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,8,opt,name=Address,proto3" json:"Address,omitempty"`
	Version              string   `protobuf:"bytes,10,opt,name=version,proto3" json:"version,omitempty"`
	Build                string   `protobuf:"bytes,12,opt,name=build,proto3" json:"build,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89273fd6fb3348, []int{8}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *HealthResponse) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *HealthResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HealthResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *HealthResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *HealthResponse) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func init() {
	proto.RegisterType((*Item)(nil), "checkout.Item")
	proto.RegisterType((*AddItemRequest)(nil), "checkout.AddItemRequest")
	proto.RegisterType((*AddItemResponse)(nil), "checkout.AddItemResponse")
	proto.RegisterType((*LineItem)(nil), "checkout.LineItem")
	proto.RegisterType((*Receipt)(nil), "checkout.Receipt")
	proto.RegisterType((*ScanItem)(nil), "checkout.ScanItem")
	proto.RegisterType((*CheckoutRequest)(nil), "checkout.CheckoutRequest")
	proto.RegisterType((*CheckoutResponse)(nil), "checkout.CheckoutResponse")
	proto.RegisterType((*HealthResponse)(nil), "checkout.HealthResponse")
}

func init() { proto.RegisterFile("checkout/checkout.proto", fileDescriptor_ed89273fd6fb3348) }

var fileDescriptor_ed89273fd6fb3348 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x96, 0x1d, 0x27, 0x71, 0x26, 0x3f, 0x25, 0xf9, 0xad, 0xaa, 0xd6, 0x35, 0x20, 0xa2, 0x3d,
	0x59, 0x45, 0x8a, 0x21, 0xc0, 0xa5, 0x17, 0x54, 0x21, 0x24, 0x10, 0x9c, 0xb6, 0x12, 0x9c, 0x38,
	0x38, 0xf6, 0x92, 0xac, 0x48, 0x6c, 0xd7, 0xbb, 0x8e, 0xc8, 0x95, 0x57, 0xe0, 0xc6, 0x8d, 0x67,
	0xe2, 0x15, 0x78, 0x10, 0xb4, 0xb3, 0x6b, 0xa7, 0xa9, 0x72, 0x28, 0xb7, 0xf9, 0x3c, 0xff, 0xbe,
	0x6f, 0x66, 0xc7, 0x70, 0x96, 0xae, 0x78, 0xfa, 0xb5, 0xa8, 0x55, 0xdc, 0x18, 0xb3, 0xb2, 0x2a,
	0x54, 0x41, 0xfc, 0x06, 0x87, 0x0f, 0x97, 0x45, 0xb1, 0x5c, 0xf3, 0x38, 0x29, 0x45, 0x9c, 0xe4,
	0x79, 0xa1, 0x12, 0x25, 0x8a, 0x5c, 0x9a, 0xb8, 0xf0, 0x81, 0xf5, 0x22, 0x5a, 0xd4, 0x5f, 0x62,
	0xbe, 0x29, 0xd5, 0xce, 0x38, 0xe9, 0x47, 0xf0, 0xde, 0x29, 0xbe, 0x21, 0x23, 0x70, 0x45, 0x16,
	0xb8, 0x53, 0x27, 0x1a, 0x30, 0x57, 0x64, 0x84, 0x80, 0x97, 0x27, 0x1b, 0x1e, 0xf4, 0xf0, 0x0b,
	0xda, 0xe4, 0x04, 0xba, 0x65, 0x25, 0x52, 0x1e, 0xf8, 0x53, 0x27, 0x72, 0x99, 0x01, 0xfa, 0xeb,
	0x36, 0x59, 0x8b, 0x2c, 0x80, 0xa9, 0x13, 0xf9, 0xcc, 0x00, 0xfa, 0x02, 0x46, 0x57, 0x59, 0xa6,
	0x4b, 0x33, 0x7e, 0x53, 0x73, 0xa9, 0x08, 0x05, 0x4f, 0x28, 0xbe, 0x09, 0xbc, 0xa9, 0x13, 0x0d,
	0xe7, 0xa3, 0x59, 0xab, 0x06, 0x83, 0xd0, 0x47, 0x5f, 0xc2, 0xb8, 0xcd, 0x92, 0x65, 0x91, 0x4b,
	0x7e, 0xaf, 0xb4, 0x0a, 0xfc, 0x0f, 0x22, 0xe7, 0x47, 0x85, 0x4c, 0xa0, 0x73, 0xa3, 0x76, 0x98,
	0xde, 0x65, 0xda, 0xfc, 0x07, 0x69, 0x21, 0xf8, 0xfc, 0x9b, 0x32, 0x0e, 0x40, 0x47, 0x8b, 0xe9,
	0x0e, 0xfa, 0x8c, 0xa7, 0x5c, 0x94, 0x4a, 0x87, 0x55, 0x7c, 0x29, 0xa4, 0xe2, 0x95, 0x6d, 0xdc,
	0x62, 0x12, 0x41, 0x57, 0x53, 0x94, 0x41, 0x6f, 0xda, 0x89, 0x86, 0x73, 0xb2, 0xe7, 0xdf, 0x30,
	0x66, 0x26, 0x40, 0xd3, 0x52, 0x62, 0x63, 0x1a, 0x75, 0x18, 0xda, 0x9a, 0x96, 0x2a, 0x54, 0xb2,
	0x0e, 0xfe, 0x33, 0xb4, 0x10, 0xd0, 0xa7, 0xe0, 0x5f, 0xa7, 0x49, 0x8e, 0x72, 0x27, 0xd0, 0xa9,
	0xcb, 0xd4, 0x6a, 0xd1, 0xe6, 0xb1, 0x3a, 0xf4, 0x13, 0x8c, 0x5f, 0xdb, 0xbe, 0xcd, 0x3a, 0xee,
	0x45, 0xda, 0xbb, 0x4b, 0xba, 0xe9, 0x6b, 0x49, 0xd3, 0x57, 0x30, 0xd9, 0x17, 0xb6, 0x1b, 0x7b,
	0x02, 0xfd, 0xca, 0x4c, 0x06, 0x0b, 0x0f, 0xe7, 0xff, 0xef, 0xf3, 0xed, 0xc8, 0x58, 0x13, 0x41,
	0x7f, 0x3a, 0x30, 0x7a, 0xcb, 0x93, 0xb5, 0x5a, 0xb5, 0xf9, 0xa7, 0xd0, 0x93, 0x2a, 0x51, 0xb5,
	0xb4, 0xbc, 0x2c, 0x6a, 0x85, 0x79, 0x66, 0x6f, 0x38, 0xa0, 0x63, 0xbb, 0x0c, 0xa0, 0x7f, 0x95,
	0x65, 0x15, 0x97, 0x12, 0xb7, 0x39, 0x60, 0x0d, 0xd4, 0x9e, 0x2d, 0xaf, 0xa4, 0x28, 0x72, 0x9c,
	0xce, 0x80, 0x35, 0x50, 0x0f, 0x7a, 0x51, 0x8b, 0x75, 0x86, 0x83, 0x1e, 0x30, 0x03, 0xe6, 0xbf,
	0xdc, 0xfd, 0xdc, 0xae, 0x79, 0xb5, 0xd5, 0x6f, 0xe2, 0x3d, 0xf4, 0x0c, 0x5f, 0x72, 0x3a, 0x33,
	0x87, 0x35, 0x6b, 0x0e, 0x6b, 0xf6, 0x46, 0x1f, 0x56, 0x18, 0xec, 0xe5, 0x1e, 0x2a, 0xa3, 0xe3,
	0xef, 0xbf, 0xff, 0xfc, 0x70, 0x07, 0xa4, 0x1f, 0xaf, 0x4c, 0x89, 0xcf, 0x48, 0x15, 0x17, 0x79,
	0x2b, 0xeb, 0xf0, 0x70, 0xc2, 0xf3, 0x23, 0x1e, 0x5b, 0xf0, 0x11, 0x16, 0x3c, 0xa3, 0x24, 0xde,
	0x3e, 0x6b, 0x7f, 0x0f, 0x31, 0xae, 0xe6, 0xd2, 0xb9, 0x20, 0x0b, 0xf0, 0x1b, 0xfa, 0xe4, 0x56,
	0x95, 0x3b, 0x4f, 0x21, 0x0c, 0x8f, 0xb9, 0x6c, 0x87, 0xc7, 0xd8, 0xe1, 0x9c, 0x9e, 0x1c, 0x74,
	0xb0, 0xdb, 0xbb, 0x74, 0x2e, 0x16, 0x3d, 0x54, 0xff, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x14, 0x78, 0xac, 0x14, 0xa7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CheckoutServiceClient is the client API for CheckoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckoutServiceClient interface {
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
	// catalog logically belongs to a separate service
	// moved it here due to limited time
	AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error)
	Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error)
}

type checkoutServiceClient struct {
	cc *grpc.ClientConn
}

func NewCheckoutServiceClient(cc *grpc.ClientConn) CheckoutServiceClient {
	return &checkoutServiceClient{cc}
}

func (c *checkoutServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/checkout.CheckoutService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkoutServiceClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*AddItemResponse, error) {
	out := new(AddItemResponse)
	err := c.cc.Invoke(ctx, "/checkout.CheckoutService/AddItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkoutServiceClient) Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error) {
	out := new(CheckoutResponse)
	err := c.cc.Invoke(ctx, "/checkout.CheckoutService/Checkout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckoutServiceServer is the server API for CheckoutService service.
type CheckoutServiceServer interface {
	Health(context.Context, *empty.Empty) (*HealthResponse, error)
	// catalog logically belongs to a separate service
	// moved it here due to limited time
	AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error)
	Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error)
}

// UnimplementedCheckoutServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCheckoutServiceServer struct {
}

func (*UnimplementedCheckoutServiceServer) Health(ctx context.Context, req *empty.Empty) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedCheckoutServiceServer) AddItem(ctx context.Context, req *AddItemRequest) (*AddItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItem not implemented")
}
func (*UnimplementedCheckoutServiceServer) Checkout(ctx context.Context, req *CheckoutRequest) (*CheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}

func RegisterCheckoutServiceServer(s *grpc.Server, srv CheckoutServiceServer) {
	s.RegisterService(&_CheckoutService_serviceDesc, srv)
}

func _CheckoutService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkout.CheckoutService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckoutService_AddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).AddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkout.CheckoutService/AddItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).AddItem(ctx, req.(*AddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckoutService_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckoutServiceServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkout.CheckoutService/Checkout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckoutServiceServer).Checkout(ctx, req.(*CheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CheckoutService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "checkout.CheckoutService",
	HandlerType: (*CheckoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _CheckoutService_Health_Handler,
		},
		{
			MethodName: "AddItem",
			Handler:    _CheckoutService_AddItem_Handler,
		},
		{
			MethodName: "Checkout",
			Handler:    _CheckoutService_Checkout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkout/checkout.proto",
}
