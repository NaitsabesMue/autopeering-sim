// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/peering.proto

package proto

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

type Packet struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b8b9eff9a131834, []int{0}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Packet) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Packet) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Packet)(nil), "proto.Packet")
}

func init() { proto.RegisterFile("proto/peering.proto", fileDescriptor_6b8b9eff9a131834) }

var fileDescriptor_6b8b9eff9a131834 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x48, 0x4d, 0x2d, 0xca, 0xcc, 0x4b, 0xd7, 0x03, 0xf3, 0x84, 0x58, 0xc1, 0x94,
	0x52, 0x24, 0x17, 0x5b, 0x40, 0x62, 0x72, 0x76, 0x6a, 0x89, 0x90, 0x2c, 0x17, 0x57, 0x41, 0x69,
	0x52, 0x4e, 0x66, 0x72, 0x7c, 0x76, 0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x27,
	0x44, 0xc4, 0x3b, 0xb5, 0x52, 0x48, 0x86, 0x8b, 0xb3, 0x38, 0x33, 0x3d, 0x2f, 0xb1, 0xa4, 0xb4,
	0x28, 0x55, 0x82, 0x09, 0x22, 0x0b, 0x17, 0x10, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94,
	0x60, 0x06, 0x4b, 0x80, 0xd9, 0x46, 0xfa, 0x5c, 0xec, 0x01, 0x10, 0x2b, 0x85, 0x54, 0xb8, 0x58,
	0x02, 0x40, 0x34, 0x2f, 0xc4, 0x72, 0x3d, 0x88, 0x95, 0x52, 0xa8, 0x5c, 0x27, 0xd5, 0x28, 0xe5,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xf2, 0xfc, 0x9c, 0x9c, 0xc4,
	0x64, 0xfd, 0xc4, 0xd2, 0x92, 0x7c, 0xa8, 0xcb, 0xf5, 0xc1, 0xaa, 0x93, 0xd8, 0xc0, 0x94, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x12, 0xe1, 0x10, 0x54, 0xd7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PeeringClient is the client API for Peering service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeeringClient interface {
	Ping(ctx context.Context, in *Packet, opts ...grpc.CallOption) (*Packet, error)
}

type peeringClient struct {
	cc *grpc.ClientConn
}

func NewPeeringClient(cc *grpc.ClientConn) PeeringClient {
	return &peeringClient{cc}
}

func (c *peeringClient) Ping(ctx context.Context, in *Packet, opts ...grpc.CallOption) (*Packet, error) {
	out := new(Packet)
	err := c.cc.Invoke(ctx, "/proto.Peering/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeeringServer is the server API for Peering service.
type PeeringServer interface {
	Ping(context.Context, *Packet) (*Packet, error)
}

// UnimplementedPeeringServer can be embedded to have forward compatible implementations.
type UnimplementedPeeringServer struct {
}

func (*UnimplementedPeeringServer) Ping(ctx context.Context, req *Packet) (*Packet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterPeeringServer(s *grpc.Server, srv PeeringServer) {
	s.RegisterService(&_Peering_serviceDesc, srv)
}

func _Peering_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Packet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeeringServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Peering/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeeringServer).Ping(ctx, req.(*Packet))
	}
	return interceptor(ctx, in, info, handler)
}

var _Peering_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Peering",
	HandlerType: (*PeeringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Peering_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/peering.proto",
}
