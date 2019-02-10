// Code generated by protoc-gen-go. DO NOT EDIT.
// source: farmer/farmer.proto

package go_micro_srv_farmer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

type Farmer struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	IsEnabled            bool     `protobuf:"varint,5,opt,name=isEnabled,proto3" json:"isEnabled,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Image                []*Image `protobuf:"bytes,8,rep,name=image,proto3" json:"image,omitempty"`
	IsConfirmed          bool     `protobuf:"varint,9,opt,name=isConfirmed,proto3" json:"isConfirmed,omitempty"`
	IsVerified           bool     `protobuf:"varint,10,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
	Quantity             int32    `protobuf:"varint,11,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Farmer) Reset()         { *m = Farmer{} }
func (m *Farmer) String() string { return proto.CompactTextString(m) }
func (*Farmer) ProtoMessage()    {}
func (*Farmer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceb4f917f19c5880, []int{0}
}

func (m *Farmer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Farmer.Unmarshal(m, b)
}
func (m *Farmer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Farmer.Marshal(b, m, deterministic)
}
func (m *Farmer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Farmer.Merge(m, src)
}
func (m *Farmer) XXX_Size() int {
	return xxx_messageInfo_Farmer.Size(m)
}
func (m *Farmer) XXX_DiscardUnknown() {
	xxx_messageInfo_Farmer.DiscardUnknown(m)
}

var xxx_messageInfo_Farmer proto.InternalMessageInfo

func (m *Farmer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Farmer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Farmer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Farmer) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Farmer) GetIsEnabled() bool {
	if m != nil {
		return m.IsEnabled
	}
	return false
}

func (m *Farmer) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Farmer) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Farmer) GetImage() []*Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Farmer) GetIsConfirmed() bool {
	if m != nil {
		return m.IsConfirmed
	}
	return false
}

func (m *Farmer) GetIsVerified() bool {
	if m != nil {
		return m.IsVerified
	}
	return false
}

func (m *Farmer) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type Image struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Order                int32    `protobuf:"varint,4,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceb4f917f19c5880, []int{1}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Image) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Image) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceb4f917f19c5880, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Farmer               *Farmer   `protobuf:"bytes,2,opt,name=farmer,proto3" json:"farmer,omitempty"`
	Farmers              []*Farmer `protobuf:"bytes,3,rep,name=farmers,proto3" json:"farmers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceb4f917f19c5880, []int{3}
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

func (m *Response) GetFarmer() *Farmer {
	if m != nil {
		return m.Farmer
	}
	return nil
}

func (m *Response) GetFarmers() []*Farmer {
	if m != nil {
		return m.Farmers
	}
	return nil
}

func init() {
	proto.RegisterType((*Farmer)(nil), "go.micro.srv.farmer.Farmer")
	proto.RegisterType((*Image)(nil), "go.micro.srv.farmer.Image")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.farmer.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.farmer.Response")
}

func init() { proto.RegisterFile("farmer/farmer.proto", fileDescriptor_ceb4f917f19c5880) }

var fileDescriptor_ceb4f917f19c5880 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x26, 0xd9, 0x26, 0x9b, 0x4c, 0x0a, 0x07, 0x97, 0x83, 0xb5, 0xfc, 0x45, 0x39, 0xe5, 0x14,
	0xd0, 0x56, 0x3c, 0x00, 0xaa, 0xa0, 0x42, 0x82, 0x8b, 0x11, 0xdc, 0xdd, 0xf5, 0x6c, 0xb1, 0xd4,
	0xc4, 0xa9, 0xed, 0x2d, 0xea, 0x5b, 0xf0, 0x20, 0x3c, 0x08, 0x8f, 0x85, 0x32, 0x4e, 0x76, 0xf7,
	0x10, 0x2d, 0xa7, 0xcc, 0xf7, 0xe3, 0x19, 0xcf, 0x97, 0x04, 0x2e, 0xb6, 0xd2, 0xb6, 0x68, 0xdf,
	0x86, 0x47, 0xd3, 0x5b, 0xe3, 0x0d, 0xbb, 0xb8, 0x35, 0x4d, 0xab, 0x37, 0xd6, 0x34, 0xce, 0x3e,
	0x34, 0x41, 0xaa, 0xfe, 0xc6, 0x90, 0x7e, 0xa2, 0x92, 0x3d, 0x83, 0x58, 0x2b, 0x1e, 0x95, 0x51,
	0x9d, 0x8b, 0x58, 0x2b, 0xc6, 0xe0, 0xac, 0x93, 0x2d, 0xf2, 0x98, 0x18, 0xaa, 0xd9, 0x73, 0x48,
	0xb0, 0x95, 0xfa, 0x8e, 0x2f, 0x88, 0x0c, 0x80, 0xad, 0x20, 0xeb, 0xa5, 0x73, 0xbf, 0x8c, 0x55,
	0xfc, 0x8c, 0x84, 0x3d, 0x66, 0x2f, 0x21, 0xd7, 0xee, 0x63, 0x27, 0x6f, 0xee, 0x50, 0xf1, 0xa4,
	0x8c, 0xea, 0x4c, 0x1c, 0x88, 0x41, 0xdd, 0x58, 0x94, 0x1e, 0xd5, 0x07, 0xcf, 0x53, 0x3a, 0x7a,
	0x20, 0x06, 0x75, 0xd7, 0xab, 0x51, 0x5d, 0x06, 0x75, 0x4f, 0xb0, 0x77, 0x90, 0xe8, 0x56, 0xde,
	0x22, 0xcf, 0xca, 0x45, 0x5d, 0xac, 0x57, 0xcd, 0xcc, 0x7e, 0xcd, 0xe7, 0xc1, 0x21, 0x82, 0x91,
	0x95, 0x50, 0x68, 0x77, 0x65, 0xba, 0xad, 0xb6, 0x2d, 0x2a, 0x9e, 0xd3, 0x6d, 0x8e, 0x29, 0xf6,
	0x1a, 0x40, 0xbb, 0x1f, 0x68, 0xf5, 0x56, 0xa3, 0xe2, 0x40, 0x86, 0x23, 0x66, 0xd8, 0xf4, 0x7e,
	0x27, 0x3b, 0xaf, 0xfd, 0x23, 0x2f, 0xca, 0xa8, 0x4e, 0xc4, 0x1e, 0x57, 0xdf, 0x21, 0xa1, 0x69,
	0x73, 0x41, 0xf6, 0xd2, 0xff, 0x9c, 0x82, 0x1c, 0xea, 0x81, 0xf3, 0x8f, 0x3d, 0x8e, 0x39, 0x52,
	0x3d, 0x84, 0x6b, 0xac, 0x42, 0x4b, 0x19, 0x26, 0x22, 0x80, 0xea, 0x1c, 0xe0, 0x1a, 0xbd, 0xc0,
	0xfb, 0x1d, 0x3a, 0x5f, 0xfd, 0x8e, 0x20, 0x13, 0xe8, 0x7a, 0xd3, 0x39, 0x64, 0x1c, 0x96, 0x63,
	0x58, 0x34, 0x2d, 0x13, 0x13, 0x64, 0x97, 0x90, 0x86, 0x00, 0x68, 0x68, 0xb1, 0x7e, 0x31, 0x1b,
	0x4e, 0x78, 0xf1, 0x62, 0xb4, 0xb2, 0xf7, 0xb0, 0x0c, 0x95, 0xe3, 0x0b, 0x8a, 0xf4, 0xe4, 0xa9,
	0xc9, 0xbb, 0xfe, 0x13, 0xc1, 0xd3, 0xc0, 0x7d, 0x43, 0xfb, 0xa0, 0x37, 0xc8, 0xbe, 0xc0, 0xf9,
	0x15, 0x5d, 0x64, 0xfc, 0xb2, 0x4e, 0xf5, 0x59, 0xbd, 0x9a, 0x15, 0xa7, 0x1d, 0xab, 0x27, 0xec,
	0x2b, 0xe4, 0xd7, 0xe8, 0xc7, 0x56, 0x6f, 0x66, 0xdd, 0x87, 0x80, 0xfe, 0xdb, 0xee, 0x26, 0xa5,
	0xbf, 0xe1, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x42, 0x54, 0xd0, 0x24, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for FarmerService service

type FarmerServiceClient interface {
	CreateFarmer(ctx context.Context, in *Farmer, opts ...client.CallOption) (*Response, error)
	GetFarmer(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type farmerServiceClient struct {
	c           client.Client
	serviceName string
}

func NewFarmerServiceClient(serviceName string, c client.Client) FarmerServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.farmer"
	}
	return &farmerServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *farmerServiceClient) CreateFarmer(ctx context.Context, in *Farmer, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "FarmerService.CreateFarmer", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *farmerServiceClient) GetFarmer(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "FarmerService.GetFarmer", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FarmerService service

type FarmerServiceHandler interface {
	CreateFarmer(context.Context, *Farmer, *Response) error
	GetFarmer(context.Context, *GetRequest, *Response) error
}

func RegisterFarmerServiceHandler(s server.Server, hdlr FarmerServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&FarmerService{hdlr}, opts...))
}

type FarmerService struct {
	FarmerServiceHandler
}

func (h *FarmerService) CreateFarmer(ctx context.Context, in *Farmer, out *Response) error {
	return h.FarmerServiceHandler.CreateFarmer(ctx, in, out)
}

func (h *FarmerService) GetFarmer(ctx context.Context, in *GetRequest, out *Response) error {
	return h.FarmerServiceHandler.GetFarmer(ctx, in, out)
}
