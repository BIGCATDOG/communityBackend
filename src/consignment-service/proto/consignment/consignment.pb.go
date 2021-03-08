// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consignment.proto

package go_micro_srv_consignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
	context "golang.org/x/net/context"
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

// 货轮承运的一批货物
type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

// 单个集装箱
type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// 托运结果
type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{3}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
	proto.RegisterType((*Request)(nil), "go.micro.srv.consignment.Request")
}

func init() { proto.RegisterFile("consignment.proto", fileDescriptor_3804bf87090b51a9) }

var fileDescriptor_3804bf87090b51a9 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x49, 0xff, 0x33, 0x41, 0xa0, 0x7a, 0x01, 0x16, 0x2c, 0x08, 0x41, 0x48, 0x5d, 0x65,
	0x51, 0x4e, 0x80, 0xb2, 0xa8, 0xb2, 0x75, 0xb7, 0x48, 0x50, 0xe2, 0x51, 0x6a, 0xa9, 0xb1, 0x83,
	0xed, 0x96, 0xab, 0xc1, 0x31, 0xb8, 0x11, 0xaa, 0xd3, 0x80, 0x11, 0x6a, 0xd5, 0x9d, 0xdf, 0x8c,
	0x9f, 0xdf, 0xe7, 0xd1, 0xc0, 0xb8, 0x50, 0xd2, 0x88, 0x52, 0x56, 0x28, 0x6d, 0x5a, 0x6b, 0x65,
	0x15, 0xa1, 0xa5, 0x4a, 0x2b, 0x51, 0x68, 0x95, 0x1a, 0xbd, 0x49, 0xbd, 0x7e, 0xf2, 0x19, 0x40,
	0x94, 0xfd, 0x6a, 0x72, 0x06, 0x1d, 0xc1, 0x69, 0x10, 0x07, 0x93, 0x90, 0x75, 0x04, 0x27, 0x31,
	0x44, 0x1c, 0x4d, 0xa1, 0x45, 0x6d, 0x85, 0x92, 0xb4, 0xe3, 0x1a, 0x7e, 0x89, 0x5c, 0xc0, 0xe0,
	0x1d, 0x45, 0xb9, 0xb4, 0xb4, 0x1b, 0x07, 0x93, 0x3e, 0xdb, 0x29, 0x92, 0x01, 0x14, 0x4a, 0xda,
	0x85, 0x90, 0xa8, 0x0d, 0xed, 0xc5, 0xdd, 0x49, 0x34, 0xbd, 0x4b, 0xf7, 0x81, 0xa4, 0x59, 0x7b,
	0x97, 0x79, 0x36, 0x72, 0x0d, 0xe1, 0x06, 0x8d, 0xc1, 0xd5, 0xb3, 0xe0, 0xb4, 0xef, 0xc2, 0x47,
	0x4d, 0x21, 0xe7, 0x49, 0x05, 0xe1, 0x8f, 0xeb, 0x1f, 0xf8, 0x0d, 0x44, 0xc5, 0xda, 0x58, 0x55,
	0xa1, 0xde, 0x7a, 0x1b, 0x70, 0x68, 0x4b, 0x39, 0xdf, 0x72, 0x2b, 0x2d, 0x4a, 0x21, 0x1d, 0x77,
	0xc8, 0x76, 0x8a, 0x5c, 0xc2, 0x70, 0x6d, 0x1a, 0x53, 0xaf, 0x69, 0x6c, 0x65, 0xce, 0x93, 0x8f,
	0x00, 0x46, 0x0c, 0x4d, 0xad, 0xa4, 0x41, 0x42, 0x61, 0x58, 0x68, 0x5c, 0x58, 0x6c, 0x32, 0x47,
	0xac, 0x95, 0x64, 0x06, 0x91, 0xf7, 0x2f, 0x17, 0x1c, 0x4d, 0xef, 0x0f, 0x7e, 0xbc, 0x3d, 0x33,
	0xdf, 0x49, 0x72, 0x38, 0xf5, 0xa4, 0xa1, 0x5d, 0x37, 0xc2, 0x23, 0x5f, 0xfa, 0x63, 0x4d, 0x42,
	0x18, 0x32, 0x7c, 0x5b, 0xa3, 0xb1, 0xd3, 0xaf, 0x00, 0xce, 0xe7, 0x4b, 0x51, 0xd7, 0x42, 0x96,
	0x73, 0xd4, 0x1b, 0x51, 0x20, 0x79, 0x81, 0x71, 0xe6, 0xe8, 0xfd, 0x4d, 0x38, 0x2e, 0xe8, 0x2a,
	0xd9, 0x7f, 0xad, 0x1d, 0x56, 0x72, 0x42, 0x9e, 0x60, 0x3c, 0x43, 0xfb, 0xb8, 0x5a, 0xf9, 0x09,
	0xb7, 0x87, 0xac, 0x8e, 0xf6, 0xb8, 0xd7, 0x5f, 0x07, 0x6e, 0xcb, 0x1f, 0xbe, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x83, 0x6f, 0x95, 0xd5, 0xfa, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	// 托运一批货物
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetAllConsignment(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.consignment"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetAllConsignment(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetAllConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	// 托运一批货物
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetAllConsignment(context.Context, *Request, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetAllConsignment(ctx context.Context, in *Request, out *Response) error {
	return h.ShippingServiceHandler.GetAllConsignment(ctx, in, out)
}
