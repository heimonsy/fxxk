// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fxxk.proto

package fxxk

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

type Command_Type int32

const (
	Command_PING      Command_Type = 0
	Command_NEW_TUNEL Command_Type = 1
	Command_CLOSE     Command_Type = 99
)

var Command_Type_name = map[int32]string{
	0:  "PING",
	1:  "NEW_TUNEL",
	99: "CLOSE",
}

var Command_Type_value = map[string]int32{
	"PING":      0,
	"NEW_TUNEL": 1,
	"CLOSE":     99,
}

func (x Command_Type) String() string {
	return proto.EnumName(Command_Type_name, int32(x))
}

func (Command_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_900b8d357a079f0a, []int{0, 0}
}

type Command struct {
	Type                 Command_Type `protobuf:"varint,1,opt,name=type,proto3,enum=heimonsy.fxxk.Command_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_900b8d357a079f0a, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetType() Command_Type {
	if m != nil {
		return m.Type
	}
	return Command_PING
}

type ConnectRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_900b8d357a079f0a, []int{1}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type TunnelPackage struct {
	ConnId               string   `protobuf:"bytes,1,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TunnelPackage) Reset()         { *m = TunnelPackage{} }
func (m *TunnelPackage) String() string { return proto.CompactTextString(m) }
func (*TunnelPackage) ProtoMessage()    {}
func (*TunnelPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_900b8d357a079f0a, []int{2}
}

func (m *TunnelPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TunnelPackage.Unmarshal(m, b)
}
func (m *TunnelPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TunnelPackage.Marshal(b, m, deterministic)
}
func (m *TunnelPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TunnelPackage.Merge(m, src)
}
func (m *TunnelPackage) XXX_Size() int {
	return xxx_messageInfo_TunnelPackage.Size(m)
}
func (m *TunnelPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_TunnelPackage.DiscardUnknown(m)
}

var xxx_messageInfo_TunnelPackage proto.InternalMessageInfo

func (m *TunnelPackage) GetConnId() string {
	if m != nil {
		return m.ConnId
	}
	return ""
}

func (m *TunnelPackage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type TunnelRequest struct {
	// Types that are valid to be assigned to Req:
	//	*TunnelRequest_Data
	//	*TunnelRequest_ClientId
	Req                  isTunnelRequest_Req `protobuf_oneof:"req"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TunnelRequest) Reset()         { *m = TunnelRequest{} }
func (m *TunnelRequest) String() string { return proto.CompactTextString(m) }
func (*TunnelRequest) ProtoMessage()    {}
func (*TunnelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_900b8d357a079f0a, []int{3}
}

func (m *TunnelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TunnelRequest.Unmarshal(m, b)
}
func (m *TunnelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TunnelRequest.Marshal(b, m, deterministic)
}
func (m *TunnelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TunnelRequest.Merge(m, src)
}
func (m *TunnelRequest) XXX_Size() int {
	return xxx_messageInfo_TunnelRequest.Size(m)
}
func (m *TunnelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TunnelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TunnelRequest proto.InternalMessageInfo

type isTunnelRequest_Req interface {
	isTunnelRequest_Req()
}

type TunnelRequest_Data struct {
	Data *TunnelPackage `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type TunnelRequest_ClientId struct {
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3,oneof"`
}

func (*TunnelRequest_Data) isTunnelRequest_Req() {}

func (*TunnelRequest_ClientId) isTunnelRequest_Req() {}

func (m *TunnelRequest) GetReq() isTunnelRequest_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *TunnelRequest) GetData() *TunnelPackage {
	if x, ok := m.GetReq().(*TunnelRequest_Data); ok {
		return x.Data
	}
	return nil
}

func (m *TunnelRequest) GetClientId() string {
	if x, ok := m.GetReq().(*TunnelRequest_ClientId); ok {
		return x.ClientId
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TunnelRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TunnelRequest_Data)(nil),
		(*TunnelRequest_ClientId)(nil),
	}
}

func init() {
	proto.RegisterEnum("heimonsy.fxxk.Command_Type", Command_Type_name, Command_Type_value)
	proto.RegisterType((*Command)(nil), "heimonsy.fxxk.Command")
	proto.RegisterType((*ConnectRequest)(nil), "heimonsy.fxxk.ConnectRequest")
	proto.RegisterType((*TunnelPackage)(nil), "heimonsy.fxxk.TunnelPackage")
	proto.RegisterType((*TunnelRequest)(nil), "heimonsy.fxxk.TunnelRequest")
}

func init() { proto.RegisterFile("fxxk.proto", fileDescriptor_900b8d357a079f0a) }

var fileDescriptor_900b8d357a079f0a = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6a, 0xea, 0x40,
	0x10, 0xc6, 0x5d, 0x8d, 0x7f, 0x32, 0xe7, 0x28, 0x32, 0x17, 0xe7, 0x48, 0xad, 0x20, 0xb9, 0x92,
	0x42, 0xa3, 0xa4, 0xb7, 0xbd, 0x32, 0x58, 0x15, 0xc4, 0x86, 0x34, 0xa5, 0xd0, 0x1b, 0x89, 0xc9,
	0xaa, 0x41, 0xb3, 0xeb, 0x9f, 0x0d, 0xc4, 0xe7, 0xe8, 0x1b, 0xf4, 0x49, 0xcb, 0x46, 0x83, 0xda,
	0x4a, 0xef, 0x76, 0x86, 0xdf, 0xcc, 0x7c, 0xdf, 0xc7, 0x02, 0xcc, 0xe2, 0x78, 0xa9, 0xaf, 0xb7,
	0x5c, 0x70, 0x2c, 0x2f, 0x68, 0x10, 0x72, 0xb6, 0xdb, 0xeb, 0xb2, 0xa9, 0xcd, 0xa0, 0x68, 0xf2,
	0x30, 0x74, 0x99, 0x8f, 0x6d, 0x50, 0xc4, 0x7e, 0x4d, 0x6b, 0xa4, 0x49, 0x5a, 0x15, 0xa3, 0xae,
	0x5f, 0x80, 0xfa, 0x91, 0xd2, 0x9d, 0xfd, 0x9a, 0xda, 0x09, 0xa8, 0xdd, 0x81, 0x22, 0x2b, 0x2c,
	0x81, 0x62, 0x0d, 0xc7, 0xfd, 0x6a, 0x06, 0xcb, 0xa0, 0x8e, 0x7b, 0x6f, 0x13, 0xe7, 0x75, 0xdc,
	0x1b, 0x55, 0x09, 0xaa, 0x90, 0x37, 0x47, 0xcf, 0x2f, 0xbd, 0xaa, 0xa7, 0xdd, 0x43, 0xc5, 0xe4,
	0x8c, 0x51, 0x4f, 0xd8, 0x74, 0x13, 0xd1, 0x9d, 0xc0, 0x3a, 0xa8, 0xde, 0x2a, 0xa0, 0x4c, 0x4c,
	0x02, 0x3f, 0xb9, 0xa9, 0xda, 0xa5, 0x43, 0x63, 0xe8, 0x6b, 0x8f, 0x50, 0x76, 0x22, 0xc6, 0xe8,
	0xca, 0x72, 0xbd, 0xa5, 0x3b, 0xa7, 0xf8, 0x1f, 0x8a, 0x1e, 0x67, 0xec, 0xc4, 0x16, 0x64, 0x39,
	0xf4, 0x11, 0x41, 0xf1, 0x5d, 0xe1, 0xd6, 0xb2, 0x4d, 0xd2, 0xfa, 0x6b, 0x27, 0x6f, 0x2d, 0x48,
	0xa7, 0xd3, 0x5b, 0xc6, 0x11, 0x92, 0xa3, 0x7f, 0x8c, 0xdb, 0x6f, 0xd6, 0x2e, 0x2e, 0x0d, 0x32,
	0x87, 0x25, 0xd8, 0x38, 0xd7, 0x27, 0xb7, 0xab, 0x83, 0xcc, 0x49, 0x61, 0x37, 0x0f, 0xb9, 0x2d,
	0xdd, 0x18, 0x1f, 0x04, 0x94, 0xa7, 0x38, 0x5e, 0x62, 0x57, 0x06, 0x99, 0x18, 0xc4, 0xc6, 0x8f,
	0xe8, 0xce, 0x8d, 0xdf, 0xfc, 0xbb, 0x9e, 0x6c, 0x87, 0x60, 0x1f, 0xf2, 0x4e, 0xc4, 0xe8, 0x0a,
	0xaf, 0x2b, 0x4c, 0x17, 0xfc, 0xaa, 0xbf, 0x45, 0x3a, 0xa4, 0xab, 0x59, 0xe4, 0xbd, 0x36, 0x0f,
	0xc4, 0x22, 0x9a, 0xea, 0x1e, 0x0f, 0xdb, 0x29, 0xdf, 0x96, 0xfc, 0x67, 0x36, 0x67, 0x5b, 0xe6,
	0xb4, 0x90, 0xfc, 0x87, 0x87, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x67, 0x3d, 0x50, 0x1d,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FxxkClient is the client API for Fxxk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FxxkClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (Fxxk_ConnectClient, error)
	Tunel(ctx context.Context, opts ...grpc.CallOption) (Fxxk_TunelClient, error)
}

type fxxkClient struct {
	cc *grpc.ClientConn
}

func NewFxxkClient(cc *grpc.ClientConn) FxxkClient {
	return &fxxkClient{cc}
}

func (c *fxxkClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (Fxxk_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fxxk_serviceDesc.Streams[0], "/heimonsy.fxxk.Fxxk/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &fxxkConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fxxk_ConnectClient interface {
	Recv() (*Command, error)
	grpc.ClientStream
}

type fxxkConnectClient struct {
	grpc.ClientStream
}

func (x *fxxkConnectClient) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fxxkClient) Tunel(ctx context.Context, opts ...grpc.CallOption) (Fxxk_TunelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fxxk_serviceDesc.Streams[1], "/heimonsy.fxxk.Fxxk/Tunel", opts...)
	if err != nil {
		return nil, err
	}
	x := &fxxkTunelClient{stream}
	return x, nil
}

type Fxxk_TunelClient interface {
	Send(*TunnelRequest) error
	Recv() (*TunnelPackage, error)
	grpc.ClientStream
}

type fxxkTunelClient struct {
	grpc.ClientStream
}

func (x *fxxkTunelClient) Send(m *TunnelRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fxxkTunelClient) Recv() (*TunnelPackage, error) {
	m := new(TunnelPackage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FxxkServer is the server API for Fxxk service.
type FxxkServer interface {
	Connect(*ConnectRequest, Fxxk_ConnectServer) error
	Tunel(Fxxk_TunelServer) error
}

// UnimplementedFxxkServer can be embedded to have forward compatible implementations.
type UnimplementedFxxkServer struct {
}

func (*UnimplementedFxxkServer) Connect(req *ConnectRequest, srv Fxxk_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedFxxkServer) Tunel(srv Fxxk_TunelServer) error {
	return status.Errorf(codes.Unimplemented, "method Tunel not implemented")
}

func RegisterFxxkServer(s *grpc.Server, srv FxxkServer) {
	s.RegisterService(&_Fxxk_serviceDesc, srv)
}

func _Fxxk_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FxxkServer).Connect(m, &fxxkConnectServer{stream})
}

type Fxxk_ConnectServer interface {
	Send(*Command) error
	grpc.ServerStream
}

type fxxkConnectServer struct {
	grpc.ServerStream
}

func (x *fxxkConnectServer) Send(m *Command) error {
	return x.ServerStream.SendMsg(m)
}

func _Fxxk_Tunel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FxxkServer).Tunel(&fxxkTunelServer{stream})
}

type Fxxk_TunelServer interface {
	Send(*TunnelPackage) error
	Recv() (*TunnelRequest, error)
	grpc.ServerStream
}

type fxxkTunelServer struct {
	grpc.ServerStream
}

func (x *fxxkTunelServer) Send(m *TunnelPackage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fxxkTunelServer) Recv() (*TunnelRequest, error) {
	m := new(TunnelRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Fxxk_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heimonsy.fxxk.Fxxk",
	HandlerType: (*FxxkServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Fxxk_Connect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Tunel",
			Handler:       _Fxxk_Tunel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "fxxk.proto",
}
