// Code generated by protoc-gen-go. DO NOT EDIT.
// source: integrator.proto

package integrator

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

type IntegratorRequest struct {
	FirstNumber          int64    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int64    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	ProductId            int64    `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegratorRequest) Reset()         { *m = IntegratorRequest{} }
func (m *IntegratorRequest) String() string { return proto.CompactTextString(m) }
func (*IntegratorRequest) ProtoMessage()    {}
func (*IntegratorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc8675db8316c943, []int{0}
}

func (m *IntegratorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegratorRequest.Unmarshal(m, b)
}
func (m *IntegratorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegratorRequest.Marshal(b, m, deterministic)
}
func (m *IntegratorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegratorRequest.Merge(m, src)
}
func (m *IntegratorRequest) XXX_Size() int {
	return xxx_messageInfo_IntegratorRequest.Size(m)
}
func (m *IntegratorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegratorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IntegratorRequest proto.InternalMessageInfo

func (m *IntegratorRequest) GetFirstNumber() int64 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *IntegratorRequest) GetSecondNumber() int64 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

func (m *IntegratorRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type IntegratorResponse struct {
	SumResponseCount     int64    `protobuf:"varint,1,opt,name=sum_response_count,json=sumResponseCount,proto3" json:"sum_response_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntegratorResponse) Reset()         { *m = IntegratorResponse{} }
func (m *IntegratorResponse) String() string { return proto.CompactTextString(m) }
func (*IntegratorResponse) ProtoMessage()    {}
func (*IntegratorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc8675db8316c943, []int{1}
}

func (m *IntegratorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegratorResponse.Unmarshal(m, b)
}
func (m *IntegratorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegratorResponse.Marshal(b, m, deterministic)
}
func (m *IntegratorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegratorResponse.Merge(m, src)
}
func (m *IntegratorResponse) XXX_Size() int {
	return xxx_messageInfo_IntegratorResponse.Size(m)
}
func (m *IntegratorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegratorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntegratorResponse proto.InternalMessageInfo

func (m *IntegratorResponse) GetSumResponseCount() int64 {
	if m != nil {
		return m.SumResponseCount
	}
	return 0
}

func init() {
	proto.RegisterType((*IntegratorRequest)(nil), "Integrator.IntegratorRequest")
	proto.RegisterType((*IntegratorResponse)(nil), "Integrator.IntegratorResponse")
}

func init() {
	proto.RegisterFile("integrator.proto", fileDescriptor_fc8675db8316c943)
}

var fileDescriptor_fc8675db8316c943 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xbf, 0x7e, 0x01, 0xc1, 0xb1, 0x42, 0xdd, 0x53, 0x11, 0x2a, 0x1a, 0x2f, 0x1e, 0x24,
	0x8b, 0xfa, 0x0f, 0xea, 0xa9, 0x07, 0x45, 0xaa, 0x27, 0x2f, 0x21, 0xcd, 0x8e, 0x12, 0x61, 0x77,
	0xd6, 0x9d, 0x19, 0x04, 0x7f, 0xbd, 0x74, 0x93, 0x1a, 0x41, 0x3c, 0xbe, 0xcf, 0x3e, 0xcc, 0xec,
	0xbc, 0x30, 0xeb, 0x82, 0xe0, 0x6b, 0x6a, 0x84, 0x52, 0x15, 0x13, 0x09, 0x19, 0x58, 0x7d, 0x93,
	0xf2, 0x13, 0x8e, 0xc6, 0xb4, 0xc6, 0x77, 0x45, 0x16, 0x73, 0x06, 0xd3, 0x97, 0x2e, 0xb1, 0xd4,
	0x41, 0xfd, 0x06, 0xd3, 0x7c, 0x72, 0x3a, 0xb9, 0x28, 0xd6, 0x07, 0x99, 0xdd, 0x67, 0x64, 0xce,
	0xe1, 0x90, 0xb1, 0xa5, 0xe0, 0x76, 0xce, 0xff, 0xec, 0x4c, 0x7b, 0x38, 0x48, 0x0b, 0x80, 0x98,
	0xc8, 0x69, 0x2b, 0x75, 0xe7, 0xe6, 0x45, 0x36, 0xf6, 0x07, 0xb2, 0x72, 0xe5, 0x12, 0xcc, 0xcf,
	0xdd, 0x1c, 0x29, 0x30, 0x9a, 0x4b, 0x30, 0xac, 0xbe, 0x4e, 0x43, 0xae, 0x5b, 0xd2, 0x20, 0xc3,
	0x17, 0x66, 0xac, 0x7e, 0x27, 0xde, 0x6e, 0xf9, 0xf5, 0x13, 0x14, 0x8f, 0xea, 0xcd, 0x1d, 0x40,
	0xbf, 0x93, 0xb7, 0x69, 0x51, 0x8d, 0x73, 0xab, 0x5f, 0xe7, 0x1d, 0x9f, 0xfc, 0xf5, 0xdc, 0x0f,
	0x2e, 0xff, 0x2d, 0xaf, 0x9e, 0x6d, 0xf3, 0x81, 0x4c, 0x1e, 0x1f, 0x12, 0xbd, 0x61, 0x2b, 0xb6,
	0x71, 0x4d, 0x14, 0x4c, 0x96, 0x54, 0xa2, 0x8a, 0xcd, 0x45, 0xb2, 0x1d, 0xab, 0xdd, 0xec, 0x65,
	0x74, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff, 0x85, 0x85, 0x79, 0x4f, 0x6f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SumClient is the client API for Sum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumClient interface {
	NumbersSum(ctx context.Context, in *IntegratorRequest, opts ...grpc.CallOption) (*IntegratorResponse, error)
}

type sumClient struct {
	cc grpc.ClientConnInterface
}

func NewSumClient(cc grpc.ClientConnInterface) SumClient {
	return &sumClient{cc}
}

func (c *sumClient) NumbersSum(ctx context.Context, in *IntegratorRequest, opts ...grpc.CallOption) (*IntegratorResponse, error) {
	out := new(IntegratorResponse)
	err := c.cc.Invoke(ctx, "/Integrator.Sum/NumbersSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServer is the server API for Sum service.
type SumServer interface {
	NumbersSum(context.Context, *IntegratorRequest) (*IntegratorResponse, error)
}

// UnimplementedSumServer can be embedded to have forward compatible implementations.
type UnimplementedSumServer struct {
}

func (*UnimplementedSumServer) NumbersSum(ctx context.Context, req *IntegratorRequest) (*IntegratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NumbersSum not implemented")
}

func RegisterSumServer(s *grpc.Server, srv SumServer) {
	s.RegisterService(&_Sum_serviceDesc, srv)
}

func _Sum_NumbersSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServer).NumbersSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Integrator.Sum/NumbersSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServer).NumbersSum(ctx, req.(*IntegratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sum_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Integrator.Sum",
	HandlerType: (*SumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NumbersSum",
			Handler:    _Sum_NumbersSum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "integrator.proto",
}
