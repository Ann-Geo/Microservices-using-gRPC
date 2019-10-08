// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Grpc_course/client_stream/greetpb/greet.proto

package greetpb

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
	return fileDescriptor_3109481dc6c123a3, []int{0}
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

type GreetStreamRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetStreamRequest) Reset()         { *m = GreetStreamRequest{} }
func (m *GreetStreamRequest) String() string { return proto.CompactTextString(m) }
func (*GreetStreamRequest) ProtoMessage()    {}
func (*GreetStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3109481dc6c123a3, []int{1}
}

func (m *GreetStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetStreamRequest.Unmarshal(m, b)
}
func (m *GreetStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetStreamRequest.Marshal(b, m, deterministic)
}
func (m *GreetStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetStreamRequest.Merge(m, src)
}
func (m *GreetStreamRequest) XXX_Size() int {
	return xxx_messageInfo_GreetStreamRequest.Size(m)
}
func (m *GreetStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetStreamRequest proto.InternalMessageInfo

func (m *GreetStreamRequest) GetGreeting() *Greeting {
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
	return fileDescriptor_3109481dc6c123a3, []int{2}
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

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetStreamRequest)(nil), "greet.GreetStreamRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
}

func init() {
	proto.RegisterFile("Grpc_course/client_stream/greetpb/greet.proto", fileDescriptor_3109481dc6c123a3)
}

var fileDescriptor_3109481dc6c123a3 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xad, 0xe0, 0xda, 0xce, 0x2a, 0xc2, 0x20, 0xe2, 0x1f, 0x04, 0xe9, 0xc5, 0x05, 0x71,
	0x17, 0xd6, 0x4f, 0xe0, 0x1e, 0xdc, 0x9b, 0x87, 0x78, 0xf3, 0x52, 0xba, 0x61, 0x2c, 0x81, 0x36,
	0x89, 0x93, 0xd4, 0xcf, 0x2f, 0x9d, 0xa6, 0x52, 0x4f, 0xc9, 0xbc, 0x37, 0xbc, 0xdf, 0x63, 0xe0,
	0x79, 0xcf, 0x5e, 0x57, 0xda, 0xf5, 0x1c, 0x68, 0xa3, 0x5b, 0x43, 0x36, 0x56, 0x21, 0x32, 0xd5,
	0xdd, 0xa6, 0x61, 0xa2, 0xe8, 0x0f, 0xe3, 0xbb, 0xf6, 0xec, 0xa2, 0xc3, 0x13, 0x19, 0xca, 0x37,
	0xc8, 0xf7, 0xc3, 0xc7, 0xd8, 0x06, 0xef, 0x01, 0xbe, 0x0c, 0x87, 0x58, 0xd9, 0xba, 0xa3, 0xeb,
	0xec, 0x21, 0x5b, 0x15, 0xaa, 0x10, 0xe5, 0xbd, 0xee, 0x08, 0xef, 0xa0, 0x68, 0xeb, 0xc9, 0x3d,
	0x16, 0x37, 0x1f, 0x84, 0xc1, 0x2c, 0x5f, 0x01, 0x25, 0xe7, 0x43, 0x88, 0x8a, 0xbe, 0x7b, 0x0a,
	0x11, 0x9f, 0x20, 0x6f, 0x52, 0xba, 0xe4, 0x2d, 0xb7, 0x17, 0xeb, 0xb1, 0xc4, 0x04, 0x55, 0x7f,
	0x0b, 0xe5, 0x23, 0x9c, 0x8b, 0xaa, 0x28, 0x78, 0x67, 0x03, 0xe1, 0x15, 0x2c, 0x98, 0x42, 0xdf,
	0xc6, 0xd4, 0x25, 0x4d, 0x5b, 0x05, 0x67, 0x23, 0x8b, 0xf8, 0xc7, 0x68, 0xc2, 0x1d, 0x2c, 0x67,
	0x6c, 0xbc, 0x99, 0x23, 0xfe, 0xf5, 0xb9, 0xbd, 0x9c, 0x5b, 0x13, 0xa7, 0x3c, 0x5a, 0x65, 0xbb,
	0xe2, 0xf3, 0x34, 0x5d, 0xe9, 0xb0, 0x90, 0x03, 0xbd, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc8,
	0xc5, 0xe1, 0xd8, 0x51, 0x01, 0x00, 0x00,
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
	GreetStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetStreamClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) GreetStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetStreamClient{stream}
	return x, nil
}

type GreetService_GreetStreamClient interface {
	Send(*GreetStreamRequest) error
	CloseAndRecv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreetStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetStreamClient) Send(m *GreetStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetStreamClient) CloseAndRecv() (*GreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	GreetStream(GreetService_GreetStreamServer) error
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_GreetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetStream(&greetServiceGreetStreamServer{stream})
}

type GreetService_GreetStreamServer interface {
	SendAndClose(*GreetResponse) error
	Recv() (*GreetStreamRequest, error)
	grpc.ServerStream
}

type greetServiceGreetStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetStreamServer) SendAndClose(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetStreamServer) Recv() (*GreetStreamRequest, error) {
	m := new(GreetStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetStream",
			Handler:       _GreetService_GreetStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Grpc_course/client_stream/greetpb/greet.proto",
}
