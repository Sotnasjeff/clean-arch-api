// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/infra/grpc/protofiles/order.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateOrderRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{0}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOrderRequest) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreateOrderRequest) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

type CreateOrderResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice           float32  `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{1}
}

func (m *CreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResponse.Unmarshal(m, b)
}
func (m *CreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResponse.Merge(m, src)
}
func (m *CreateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResponse.Size(m)
}
func (m *CreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResponse proto.InternalMessageInfo

func (m *CreateOrderResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOrderResponse) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreateOrderResponse) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

func (m *CreateOrderResponse) GetFinalPrice() float32 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

type GetOrderByIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderByIdRequest) Reset()         { *m = GetOrderByIdRequest{} }
func (m *GetOrderByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrderByIdRequest) ProtoMessage()    {}
func (*GetOrderByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{2}
}

func (m *GetOrderByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderByIdRequest.Unmarshal(m, b)
}
func (m *GetOrderByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetOrderByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderByIdRequest.Merge(m, src)
}
func (m *GetOrderByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrderByIdRequest.Size(m)
}
func (m *GetOrderByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderByIdRequest proto.InternalMessageInfo

func (m *GetOrderByIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Blank struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blank) Reset()         { *m = Blank{} }
func (m *Blank) String() string { return proto.CompactTextString(m) }
func (*Blank) ProtoMessage()    {}
func (*Blank) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{3}
}

func (m *Blank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blank.Unmarshal(m, b)
}
func (m *Blank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blank.Marshal(b, m, deterministic)
}
func (m *Blank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blank.Merge(m, src)
}
func (m *Blank) XXX_Size() int {
	return xxx_messageInfo_Blank.Size(m)
}
func (m *Blank) XXX_DiscardUnknown() {
	xxx_messageInfo_Blank.DiscardUnknown(m)
}

var xxx_messageInfo_Blank proto.InternalMessageInfo

type GetOrderResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice           float32  `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderResponse) Reset()         { *m = GetOrderResponse{} }
func (m *GetOrderResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrderResponse) ProtoMessage()    {}
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{4}
}

func (m *GetOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderResponse.Unmarshal(m, b)
}
func (m *GetOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderResponse.Marshal(b, m, deterministic)
}
func (m *GetOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderResponse.Merge(m, src)
}
func (m *GetOrderResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrderResponse.Size(m)
}
func (m *GetOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderResponse proto.InternalMessageInfo

func (m *GetOrderResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetOrderResponse) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *GetOrderResponse) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

func (m *GetOrderResponse) GetFinalPrice() float32 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

type ListOrdersResponse struct {
	Orders               []*GetOrderResponse `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListOrdersResponse) Reset()         { *m = ListOrdersResponse{} }
func (m *ListOrdersResponse) String() string { return proto.CompactTextString(m) }
func (*ListOrdersResponse) ProtoMessage()    {}
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{5}
}

func (m *ListOrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOrdersResponse.Unmarshal(m, b)
}
func (m *ListOrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOrdersResponse.Marshal(b, m, deterministic)
}
func (m *ListOrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOrdersResponse.Merge(m, src)
}
func (m *ListOrdersResponse) XXX_Size() int {
	return xxx_messageInfo_ListOrdersResponse.Size(m)
}
func (m *ListOrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOrdersResponse proto.InternalMessageInfo

func (m *ListOrdersResponse) GetOrders() []*GetOrderResponse {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateOrderRequest)(nil), "pb.CreateOrderRequest")
	proto.RegisterType((*CreateOrderResponse)(nil), "pb.CreateOrderResponse")
	proto.RegisterType((*GetOrderByIdRequest)(nil), "pb.GetOrderByIdRequest")
	proto.RegisterType((*Blank)(nil), "pb.Blank")
	proto.RegisterType((*GetOrderResponse)(nil), "pb.GetOrderResponse")
	proto.RegisterType((*ListOrdersResponse)(nil), "pb.ListOrdersResponse")
}

func init() {
	proto.RegisterFile("internal/infra/grpc/protofiles/order.proto", fileDescriptor_a81083cb550ac9a5)
}

var fileDescriptor_a81083cb550ac9a5 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x69, 0xf7, 0xdf, 0xfe, 0xec, 0xed, 0x94, 0x91, 0x8d, 0xad, 0xf4, 0x62, 0x29, 0x08,
	0x45, 0xa4, 0x85, 0x7a, 0x56, 0xb0, 0x1e, 0x44, 0x18, 0x28, 0xf5, 0xe6, 0x45, 0xd2, 0x35, 0x95,
	0xb0, 0x90, 0xc6, 0x24, 0x8a, 0x7e, 0x3c, 0xbf, 0x99, 0x24, 0x9d, 0x52, 0x5d, 0xbd, 0x08, 0xde,
	0xda, 0xe7, 0x7d, 0xfa, 0x3e, 0xbf, 0x3e, 0x09, 0x1c, 0x51, 0xae, 0x89, 0xe4, 0x98, 0xa5, 0x94,
	0xd7, 0x12, 0xa7, 0x0f, 0x52, 0xac, 0x53, 0x21, 0x1b, 0xdd, 0xd4, 0x94, 0x11, 0x95, 0x36, 0xb2,
	0x22, 0x32, 0xb1, 0x02, 0x72, 0x45, 0x19, 0xad, 0x00, 0x5d, 0x48, 0x82, 0x35, 0xb9, 0x36, 0x83,
	0x82, 0x3c, 0x3e, 0x11, 0xa5, 0xd1, 0x3e, 0xb8, 0xb4, 0xf2, 0x9d, 0xd0, 0x89, 0xc7, 0x85, 0x4b,
	0x2b, 0x34, 0x87, 0xa1, 0x90, 0x74, 0x4d, 0x7c, 0x37, 0x74, 0x62, 0xb7, 0x68, 0x5f, 0xd0, 0x14,
	0x06, 0x1a, 0xbf, 0xf8, 0x03, 0xab, 0x99, 0xc7, 0x88, 0xc3, 0xec, 0xcb, 0x36, 0x25, 0x1a, 0xae,
	0xc8, 0x6f, 0xd7, 0xa1, 0x03, 0xf0, 0x6a, 0xca, 0x31, 0xbb, 0x6f, 0xdd, 0xff, 0xec, 0x04, 0xac,
	0x74, 0x63, 0x94, 0xe8, 0x10, 0x66, 0x97, 0x44, 0xdb, 0xb0, 0xfc, 0xf5, 0xaa, 0xfa, 0x01, 0x3f,
	0xfa, 0x0f, 0xc3, 0x9c, 0x61, 0xbe, 0x89, 0x36, 0x30, 0xfd, 0xf0, 0xff, 0x3d, 0x5c, 0x0e, 0x68,
	0x45, 0x55, 0x9b, 0xa6, 0x3e, 0xe3, 0x8e, 0x61, 0x64, 0xcf, 0x40, 0xf9, 0x4e, 0x38, 0x88, 0xbd,
	0x6c, 0x9e, 0x88, 0x32, 0xf9, 0x0e, 0x55, 0x6c, 0x3d, 0xd9, 0x9b, 0x03, 0x13, 0x3b, 0xb9, 0x25,
	0xf2, 0xd9, 0x70, 0x9c, 0x81, 0xd7, 0x69, 0x18, 0x2d, 0xcc, 0xd7, 0xbb, 0x07, 0x18, 0x2c, 0x77,
	0xf4, 0x6d, 0x7c, 0x06, 0x7b, 0x06, 0xea, 0x9c, 0xb1, 0x96, 0x0b, 0x8d, 0x8d, 0xd3, 0xb6, 0x13,
	0xd8, 0x65, 0x3d, 0xc8, 0xa7, 0x30, 0xe9, 0xb6, 0x8c, 0x96, 0x5d, 0xe4, 0x4e, 0xef, 0x41, 0xef,
	0xbf, 0xe4, 0xfe, 0xdd, 0xa2, 0xf7, 0x52, 0x96, 0xe5, 0xc8, 0xde, 0xc3, 0x93, 0xf7, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4c, 0x79, 0x01, 0x5f, 0xb5, 0x02, 0x00, 0x00,
}
