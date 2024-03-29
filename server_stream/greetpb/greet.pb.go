// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Grpc_course/server_stream/greetpb/greet.proto

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
	return fileDescriptor_eb2c88bdeebc5fd4, []int{0}
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
	return fileDescriptor_eb2c88bdeebc5fd4, []int{1}
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

type GreetStreamResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetStreamResponse) Reset()         { *m = GreetStreamResponse{} }
func (m *GreetStreamResponse) String() string { return proto.CompactTextString(m) }
func (*GreetStreamResponse) ProtoMessage()    {}
func (*GreetStreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb2c88bdeebc5fd4, []int{2}
}

func (m *GreetStreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetStreamResponse.Unmarshal(m, b)
}
func (m *GreetStreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetStreamResponse.Marshal(b, m, deterministic)
}
func (m *GreetStreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetStreamResponse.Merge(m, src)
}
func (m *GreetStreamResponse) XXX_Size() int {
	return xxx_messageInfo_GreetStreamResponse.Size(m)
}
func (m *GreetStreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetStreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetStreamResponse proto.InternalMessageInfo

func (m *GreetStreamResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetStreamResponse)(nil), "greet.GreetStreamResponse")
}

func init() {
	proto.RegisterFile("Grpc_course/server_stream/greetpb/greet.proto", fileDescriptor_eb2c88bdeebc5fd4)
}

var fileDescriptor_eb2c88bdeebc5fd4 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xad, 0xe0, 0x6c, 0xde, 0x04, 0x21, 0x03, 0x91, 0x89, 0x20, 0x39, 0x09, 0xb2, 0x4d,
	0xe6, 0xd1, 0xdb, 0x0e, 0xee, 0xe6, 0x21, 0xbb, 0x79, 0x29, 0x59, 0x79, 0x96, 0xc2, 0xda, 0xc4,
	0xf7, 0xd2, 0xfd, 0xfd, 0xd2, 0xd7, 0x54, 0xea, 0x29, 0x79, 0xdf, 0x97, 0x7c, 0xbf, 0x8f, 0x07,
	0xab, 0x3d, 0x85, 0xb2, 0x28, 0x7d, 0x47, 0x8c, 0x1b, 0x46, 0x3a, 0x23, 0x15, 0x1c, 0x09, 0x5d,
	0xb3, 0xa9, 0x08, 0x31, 0x86, 0xe3, 0x70, 0xae, 0x03, 0xf9, 0xe8, 0xf5, 0x95, 0x0c, 0xe6, 0x03,
	0xf2, 0x7d, 0x7f, 0xa9, 0xdb, 0x4a, 0x3f, 0x02, 0x7c, 0xd7, 0xc4, 0xb1, 0x68, 0x5d, 0x83, 0xf7,
	0xd9, 0x53, 0xf6, 0xac, 0xac, 0x12, 0xe5, 0xd3, 0x35, 0xa8, 0x1f, 0x40, 0x9d, 0xdc, 0xe8, 0x5e,
	0x8a, 0x9b, 0xf7, 0x42, 0x6f, 0x9a, 0x77, 0xb8, 0x91, 0x1c, 0x8b, 0x3f, 0x1d, 0x72, 0xd4, 0x2f,
	0x90, 0x57, 0x29, 0x57, 0x92, 0xe6, 0xdb, 0xdb, 0xf5, 0x80, 0x1f, 0x71, 0xf6, 0xef, 0x81, 0x59,
	0xc1, 0x42, 0xd4, 0x83, 0xd4, 0xb5, 0xc8, 0xc1, 0xb7, 0x8c, 0xfa, 0x0e, 0x66, 0x84, 0xdc, 0x9d,
	0x62, 0xea, 0x92, 0xa6, 0xad, 0x4d, 0xac, 0x03, 0xd2, 0xb9, 0x2e, 0x51, 0xef, 0x60, 0x3e, 0xf9,
	0xae, 0x17, 0x53, 0x50, 0xea, 0xb3, 0x5c, 0x4e, 0xc5, 0xff, 0x1c, 0x73, 0xf1, 0x9a, 0xed, 0xd4,
	0xd7, 0x75, 0xda, 0xd2, 0x71, 0x26, 0x0b, 0x7a, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x53, 0x2d,
	0xae, 0x81, 0x51, 0x01, 0x00, 0x00,
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
	GreetStream(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_GreetStreamClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) GreetStream(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_GreetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetStreamClient interface {
	Recv() (*GreetStreamResponse, error)
	grpc.ClientStream
}

type greetServiceGreetStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetStreamClient) Recv() (*GreetStreamResponse, error) {
	m := new(GreetStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	GreetStream(*GreetRequest, GreetService_GreetStreamServer) error
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_GreetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetStream(m, &greetServiceGreetStreamServer{stream})
}

type GreetService_GreetStreamServer interface {
	Send(*GreetStreamResponse) error
	grpc.ServerStream
}

type greetServiceGreetStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetStreamServer) Send(m *GreetStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetStream",
			Handler:       _GreetService_GreetStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "Grpc_course/server_stream/greetpb/greet.proto",
}
