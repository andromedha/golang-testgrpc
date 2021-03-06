// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pushmetadata.proto

package pushmetadata

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegistrationRequest struct {
	Clientnumber         int32    `protobuf:"varint,1,opt,name=Clientnumber,proto3" json:"Clientnumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationRequest) Reset()         { *m = RegistrationRequest{} }
func (m *RegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*RegistrationRequest) ProtoMessage()    {}
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49959e1f3ad47d4b, []int{0}
}

func (m *RegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationRequest.Unmarshal(m, b)
}
func (m *RegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationRequest.Marshal(b, m, deterministic)
}
func (m *RegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationRequest.Merge(m, src)
}
func (m *RegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_RegistrationRequest.Size(m)
}
func (m *RegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationRequest proto.InternalMessageInfo

func (m *RegistrationRequest) GetClientnumber() int32 {
	if m != nil {
		return m.Clientnumber
	}
	return 0
}

type RegistrationResponse struct {
	Clientnumber         int32    `protobuf:"varint,1,opt,name=Clientnumber,proto3" json:"Clientnumber,omitempty"`
	Servername           string   `protobuf:"bytes,2,opt,name=Servername,proto3" json:"Servername,omitempty"`
	ConnectedClients     int32    `protobuf:"varint,3,opt,name=ConnectedClients,proto3" json:"ConnectedClients,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrationResponse) Reset()         { *m = RegistrationResponse{} }
func (m *RegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*RegistrationResponse) ProtoMessage()    {}
func (*RegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49959e1f3ad47d4b, []int{1}
}

func (m *RegistrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationResponse.Unmarshal(m, b)
}
func (m *RegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationResponse.Marshal(b, m, deterministic)
}
func (m *RegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationResponse.Merge(m, src)
}
func (m *RegistrationResponse) XXX_Size() int {
	return xxx_messageInfo_RegistrationResponse.Size(m)
}
func (m *RegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationResponse proto.InternalMessageInfo

func (m *RegistrationResponse) GetClientnumber() int32 {
	if m != nil {
		return m.Clientnumber
	}
	return 0
}

func (m *RegistrationResponse) GetServername() string {
	if m != nil {
		return m.Servername
	}
	return ""
}

func (m *RegistrationResponse) GetConnectedClients() int32 {
	if m != nil {
		return m.ConnectedClients
	}
	return 0
}

type Clientdata struct {
	Clientnumber         int32    `protobuf:"varint,1,opt,name=Clientnumber,proto3" json:"Clientnumber,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Clientdata) Reset()         { *m = Clientdata{} }
func (m *Clientdata) String() string { return proto.CompactTextString(m) }
func (*Clientdata) ProtoMessage()    {}
func (*Clientdata) Descriptor() ([]byte, []int) {
	return fileDescriptor_49959e1f3ad47d4b, []int{2}
}

func (m *Clientdata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clientdata.Unmarshal(m, b)
}
func (m *Clientdata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clientdata.Marshal(b, m, deterministic)
}
func (m *Clientdata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clientdata.Merge(m, src)
}
func (m *Clientdata) XXX_Size() int {
	return xxx_messageInfo_Clientdata.Size(m)
}
func (m *Clientdata) XXX_DiscardUnknown() {
	xxx_messageInfo_Clientdata.DiscardUnknown(m)
}

var xxx_messageInfo_Clientdata proto.InternalMessageInfo

func (m *Clientdata) GetClientnumber() int32 {
	if m != nil {
		return m.Clientnumber
	}
	return 0
}

func (m *Clientdata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Notification struct {
	Streamers            map[int32]string `protobuf:"bytes,1,rep,name=Streamers,proto3" json:"Streamers,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_49959e1f3ad47d4b, []int{3}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetStreamers() map[int32]string {
	if m != nil {
		return m.Streamers
	}
	return nil
}

func init() {
	proto.RegisterType((*RegistrationRequest)(nil), "pushmetadata.RegistrationRequest")
	proto.RegisterType((*RegistrationResponse)(nil), "pushmetadata.RegistrationResponse")
	proto.RegisterType((*Clientdata)(nil), "pushmetadata.Clientdata")
	proto.RegisterType((*Notification)(nil), "pushmetadata.Notification")
	proto.RegisterMapType((map[int32]string)(nil), "pushmetadata.Notification.StreamersEntry")
}

func init() { proto.RegisterFile("pushmetadata.proto", fileDescriptor_49959e1f3ad47d4b) }

var fileDescriptor_49959e1f3ad47d4b = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4e, 0x32, 0x31,
	0x14, 0xa5, 0xf0, 0xf1, 0x45, 0xae, 0x13, 0x43, 0x2a, 0x8b, 0xc9, 0x2c, 0x0c, 0x76, 0x35, 0xba,
	0x20, 0x04, 0x37, 0x6a, 0xdc, 0x81, 0x71, 0x25, 0x31, 0x43, 0x7c, 0x80, 0x0e, 0x5c, 0x65, 0x22,
	0xd3, 0x62, 0x7b, 0x4b, 0xc2, 0x0b, 0xf8, 0x06, 0xbe, 0x86, 0xcf, 0x68, 0x98, 0x51, 0x99, 0xc6,
	0x9f, 0xb0, 0xbb, 0x3d, 0xbd, 0xa7, 0xf7, 0x9c, 0x73, 0x0b, 0x7c, 0xe9, 0xec, 0x3c, 0x47, 0x92,
	0x33, 0x49, 0xb2, 0xb7, 0x34, 0x9a, 0x34, 0x0f, 0xaa, 0x98, 0xb8, 0x80, 0xc3, 0x04, 0x1f, 0x33,
	0x4b, 0x46, 0x52, 0xa6, 0x55, 0x82, 0xcf, 0x0e, 0x2d, 0x71, 0x01, 0xc1, 0x70, 0x91, 0xa1, 0x22,
	0xe5, 0xf2, 0x14, 0x4d, 0xc8, 0xba, 0x2c, 0x6e, 0x26, 0x1e, 0x26, 0x5e, 0x18, 0x74, 0x7c, 0xae,
	0x5d, 0x6a, 0x65, 0x71, 0x17, 0x32, 0x3f, 0x02, 0x98, 0xa0, 0x59, 0xa1, 0x51, 0x32, 0xc7, 0xb0,
	0xde, 0x65, 0x71, 0x2b, 0xa9, 0x20, 0xfc, 0x14, 0xda, 0x43, 0xad, 0x14, 0x4e, 0x09, 0x67, 0x25,
	0xd1, 0x86, 0x8d, 0xe2, 0x9d, 0x6f, 0xb8, 0x18, 0x01, 0x94, 0xe5, 0xc6, 0xd1, 0x4e, 0xd3, 0x39,
	0xfc, 0x1b, 0x6f, 0xe7, 0x16, 0xb5, 0x78, 0x65, 0x10, 0x8c, 0x35, 0x65, 0x0f, 0xd9, 0xb4, 0xb0,
	0xc3, 0x6f, 0xa0, 0x35, 0x21, 0x83, 0x32, 0x47, 0x63, 0x43, 0xd6, 0x6d, 0xc4, 0xfb, 0x83, 0x93,
	0x9e, 0x17, 0x68, 0xb5, 0xbd, 0xf7, 0xd5, 0x7b, 0xad, 0xc8, 0xac, 0x93, 0x2d, 0x37, 0xba, 0x82,
	0x03, 0xff, 0x92, 0xb7, 0xa1, 0xf1, 0x84, 0xeb, 0x0f, 0x69, 0x9b, 0x92, 0x77, 0xa0, 0xb9, 0x92,
	0x0b, 0xf7, 0x29, 0xa9, 0x3c, 0x5c, 0xd6, 0xcf, 0xd9, 0xe0, 0x8d, 0x41, 0x70, 0xe7, 0xec, 0xfc,
	0x16, 0x49, 0x8e, 0x36, 0x06, 0xef, 0x61, 0xaf, 0x8c, 0x1d, 0x0d, 0x3f, 0xf6, 0x05, 0xfd, 0xb0,
	0xca, 0x48, 0xfc, 0xd5, 0x52, 0x6e, 0x4c, 0xd4, 0xfa, 0xac, 0xb0, 0xeb, 0x52, 0x3b, 0x35, 0x59,
	0x8a, 0x3c, 0xf4, 0x49, 0xdb, 0x78, 0xa3, 0xe8, 0xf7, 0x08, 0x44, 0x2d, 0x66, 0x7d, 0x96, 0xfe,
	0x2f, 0xfe, 0xd9, 0xd9, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x79, 0xc4, 0x6e, 0x7d, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PushMetaDataClient is the client API for PushMetaData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PushMetaDataClient interface {
	Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (PushMetaData_RegisterClient, error)
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (PushMetaData_SubscribeClient, error)
}

type pushMetaDataClient struct {
	cc *grpc.ClientConn
}

func NewPushMetaDataClient(cc *grpc.ClientConn) PushMetaDataClient {
	return &pushMetaDataClient{cc}
}

func (c *pushMetaDataClient) Register(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (PushMetaData_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PushMetaData_serviceDesc.Streams[0], "/pushmetadata.PushMetaData/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushMetaDataRegisterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushMetaData_RegisterClient interface {
	Recv() (*RegistrationResponse, error)
	grpc.ClientStream
}

type pushMetaDataRegisterClient struct {
	grpc.ClientStream
}

func (x *pushMetaDataRegisterClient) Recv() (*RegistrationResponse, error) {
	m := new(RegistrationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pushMetaDataClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (PushMetaData_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PushMetaData_serviceDesc.Streams[1], "/pushmetadata.PushMetaData/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushMetaDataSubscribeClient{stream}
	return x, nil
}

type PushMetaData_SubscribeClient interface {
	Send(*Clientdata) error
	Recv() (*Notification, error)
	grpc.ClientStream
}

type pushMetaDataSubscribeClient struct {
	grpc.ClientStream
}

func (x *pushMetaDataSubscribeClient) Send(m *Clientdata) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pushMetaDataSubscribeClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PushMetaDataServer is the server API for PushMetaData service.
type PushMetaDataServer interface {
	Register(*RegistrationRequest, PushMetaData_RegisterServer) error
	Subscribe(PushMetaData_SubscribeServer) error
}

func RegisterPushMetaDataServer(s *grpc.Server, srv PushMetaDataServer) {
	s.RegisterService(&_PushMetaData_serviceDesc, srv)
}

func _PushMetaData_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RegistrationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushMetaDataServer).Register(m, &pushMetaDataRegisterServer{stream})
}

type PushMetaData_RegisterServer interface {
	Send(*RegistrationResponse) error
	grpc.ServerStream
}

type pushMetaDataRegisterServer struct {
	grpc.ServerStream
}

func (x *pushMetaDataRegisterServer) Send(m *RegistrationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PushMetaData_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PushMetaDataServer).Subscribe(&pushMetaDataSubscribeServer{stream})
}

type PushMetaData_SubscribeServer interface {
	Send(*Notification) error
	Recv() (*Clientdata, error)
	grpc.ServerStream
}

type pushMetaDataSubscribeServer struct {
	grpc.ServerStream
}

func (x *pushMetaDataSubscribeServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pushMetaDataSubscribeServer) Recv() (*Clientdata, error) {
	m := new(Clientdata)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PushMetaData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pushmetadata.PushMetaData",
	HandlerType: (*PushMetaDataServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _PushMetaData_Register_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _PushMetaData_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pushmetadata.proto",
}
