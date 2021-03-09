// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vess/vess.proto

package go_micro_srv_vessel

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

// 每条货轮的熟悉
type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwerId               string   `protobuf:"bytes,6,opt,name=ower_id,json=owerId,proto3" json:"ower_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_4854a706581f5ee8, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwerId() string {
	if m != nil {
		return m.OwerId
	}
	return ""
}

// 等待运送的货物
type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_4854a706581f5ee8, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

// 货轮装得下的话
// 返回的多条货轮信息
type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4854a706581f5ee8, []int{2}
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

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.srv.vessel.Response")
}

func init() { proto.RegisterFile("proto/vess/vess.proto", fileDescriptor_4854a706581f5ee8) }

var fileDescriptor_4854a706581f5ee8 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xfd, 0x26, 0x6d, 0xd3, 0xe6, 0x7e, 0xc4, 0xc5, 0x88, 0x38, 0x54, 0x0b, 0x21, 0xab, 0xac,
	0x46, 0x68, 0xf1, 0x01, 0xdc, 0x08, 0xba, 0x4c, 0x45, 0x97, 0x65, 0x3a, 0xb9, 0xd6, 0x0b, 0x49,
	0x26, 0x64, 0x42, 0x5a, 0x5f, 0xc6, 0x67, 0x95, 0x4e, 0x6c, 0xa5, 0x12, 0x74, 0x33, 0xdc, 0x9f,
	0x73, 0x0f, 0xe7, 0x9c, 0x81, 0x8b, 0xaa, 0x36, 0x8d, 0xb9, 0x69, 0xd1, 0x5a, 0xf7, 0x48, 0xd7,
	0xf3, 0xf3, 0x8d, 0x91, 0x05, 0xe9, 0xda, 0x48, 0x5b, 0xb7, 0x72, 0xbf, 0xc0, 0x3c, 0xfe, 0x60,
	0xe0, 0x3f, 0xbb, 0x92, 0x9f, 0x81, 0x47, 0x99, 0x60, 0x11, 0x4b, 0x82, 0xd4, 0xa3, 0x8c, 0x4f,
	0x61, 0xa2, 0x55, 0xa5, 0x34, 0x35, 0xef, 0xc2, 0x8b, 0x58, 0x32, 0x4a, 0x8f, 0x3d, 0x9f, 0x01,
	0x14, 0x6a, 0xb7, 0xda, 0x22, 0x6d, 0xde, 0x1a, 0x31, 0x70, 0xdb, 0xa0, 0x50, 0xbb, 0x17, 0x37,
	0xe0, 0x1c, 0x86, 0xa5, 0x2a, 0x50, 0x0c, 0x1d, 0x99, 0xab, 0xf9, 0x35, 0x04, 0xaa, 0x55, 0x94,
	0xab, 0x75, 0x8e, 0x62, 0x14, 0xb1, 0x64, 0x92, 0x7e, 0x0f, 0xf8, 0x25, 0x8c, 0xcd, 0x16, 0xeb,
	0x15, 0x65, 0xc2, 0x77, 0x47, 0xfe, 0xbe, 0x7d, 0xc8, 0xe2, 0x47, 0x08, 0x97, 0x15, 0x6a, 0x7a,
	0x25, 0xad, 0x1a, 0x32, 0xe5, 0x89, 0x2c, 0xf6, 0xab, 0x2c, 0xef, 0x87, 0xac, 0xb8, 0x85, 0x49,
	0x8a, 0xb6, 0x32, 0xa5, 0x45, 0xbe, 0x00, 0xbf, 0x8b, 0xc0, 0x91, 0xfc, 0x9f, 0x5f, 0xc9, 0x9e,
	0x78, 0x64, 0x17, 0x4d, 0xfa, 0x05, 0xe5, 0xb7, 0x30, 0xee, 0x2a, 0x2b, 0xbc, 0x68, 0xf0, 0xd7,
	0xd5, 0x01, 0x3b, 0x47, 0x08, 0xbb, 0xd1, 0x12, 0xeb, 0x96, 0x34, 0xf2, 0x27, 0x08, 0xef, 0xa9,
	0xcc, 0xee, 0x8e, 0xf6, 0xe3, 0x5e, 0x9e, 0x13, 0xe3, 0xd3, 0x59, 0x2f, 0xe6, 0x60, 0x28, 0xfe,
	0xb7, 0xf6, 0xdd, 0x3f, 0x2f, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x08, 0x24, 0x2f, 0xd6, 0x00,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	// 检查是否有能运送货物的轮船
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	// 检查是否有能运送货物的轮船
	FindAvailable(context.Context, *Specification, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}
