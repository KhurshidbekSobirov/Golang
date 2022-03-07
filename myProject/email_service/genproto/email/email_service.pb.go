// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: email/email_service.proto

package email

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

func init() { proto.RegisterFile("email/email_service.proto", fileDescriptor_b57e157ce99a3bae) }

var fileDescriptor_b57e157ce99a3bae = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xcd, 0x4d, 0xcc,
	0xcc, 0xd1, 0x07, 0x93, 0xf1, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0xac, 0x60, 0x41, 0x29, 0x41, 0x24, 0x15, 0x10, 0x19, 0xa3, 0x70, 0x2e, 0x1e,
	0x57, 0x10, 0x37, 0x18, 0xa2, 0x5e, 0x48, 0x89, 0x8b, 0x25, 0x38, 0x35, 0x2f, 0x45, 0x88, 0x47,
	0x0f, 0xa2, 0x0a, 0x2c, 0x29, 0x85, 0xe0, 0x15, 0x94, 0x54, 0x0a, 0x29, 0x73, 0xb1, 0x83, 0xd4,
	0x04, 0xe7, 0x16, 0x0b, 0x71, 0x41, 0x25, 0x82, 0x73, 0x8b, 0x51, 0x15, 0x39, 0x09, 0x9c, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24,
	0xb1, 0x81, 0x6d, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xa4, 0xd1, 0xc2, 0xa8, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailServiceClient interface {
	Send(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Empty, error)
	SendSms(ctx context.Context, in *Sms, opts ...grpc.CallOption) (*Empty, error)
}

type emailServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmailServiceClient(cc *grpc.ClientConn) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) Send(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/email.EmailService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) SendSms(ctx context.Context, in *Sms, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/email.EmailService/SendSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
type EmailServiceServer interface {
	Send(context.Context, *Email) (*Empty, error)
	SendSms(context.Context, *Sms) (*Empty, error)
}

// UnimplementedEmailServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (*UnimplementedEmailServiceServer) Send(ctx context.Context, req *Email) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedEmailServiceServer) SendSms(ctx context.Context, req *Sms) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}

func RegisterEmailServiceServer(s *grpc.Server, srv EmailServiceServer) {
	s.RegisterService(&_EmailService_serviceDesc, srv)
}

func _EmailService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.EmailService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Send(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sms)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.EmailService/SendSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendSms(ctx, req.(*Sms))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "email.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _EmailService_Send_Handler,
		},
		{
			MethodName: "SendSms",
			Handler:    _EmailService_SendSms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email/email_service.proto",
}