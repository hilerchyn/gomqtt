// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	BPushMsg
	SPushMsg
	PChatMsg
	GChatMsg
	LoginMsg
	LogoutMsg
	SubMsg
	UnSubMsg
	Reply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// 广播
type BPushMsg struct {
}

func (m *BPushMsg) Reset()                    { *m = BPushMsg{} }
func (m *BPushMsg) String() string            { return proto1.CompactTextString(m) }
func (*BPushMsg) ProtoMessage()               {}
func (*BPushMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 单播
type SPushMsg struct {
	Tt [][]byte `protobuf:"bytes,1,rep,name=tt,proto3" json:"tt,omitempty"`
	Tu [][]byte `protobuf:"bytes,2,rep,name=tu,proto3" json:"tu,omitempty"`
}

func (m *SPushMsg) Reset()                    { *m = SPushMsg{} }
func (m *SPushMsg) String() string            { return proto1.CompactTextString(m) }
func (*SPushMsg) ProtoMessage()               {}
func (*SPushMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 私聊
type PChatMsg struct {
}

func (m *PChatMsg) Reset()                    { *m = PChatMsg{} }
func (m *PChatMsg) String() string            { return proto1.CompactTextString(m) }
func (*PChatMsg) ProtoMessage()               {}
func (*PChatMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 群播
type GChatMsg struct {
}

func (m *GChatMsg) Reset()                    { *m = GChatMsg{} }
func (m *GChatMsg) String() string            { return proto1.CompactTextString(m) }
func (*GChatMsg) ProtoMessage()               {}
func (*GChatMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 登陆消息
type LoginMsg struct {
	An  []byte `protobuf:"bytes,1,opt,name=an,proto3" json:"an,omitempty"`
	Un  []byte `protobuf:"bytes,2,opt,name=un,proto3" json:"un,omitempty"`
	Cid int32  `protobuf:"varint,3,opt,name=cid" json:"cid,omitempty"`
	Gip []byte `protobuf:"bytes,4,opt,name=gip,proto3" json:"gip,omitempty"`
}

func (m *LoginMsg) Reset()                    { *m = LoginMsg{} }
func (m *LoginMsg) String() string            { return proto1.CompactTextString(m) }
func (*LoginMsg) ProtoMessage()               {}
func (*LoginMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 登出消息
type LogoutMsg struct {
	An []byte `protobuf:"bytes,1,opt,name=an,proto3" json:"an,omitempty"`
	Un []byte `protobuf:"bytes,2,opt,name=un,proto3" json:"un,omitempty"`
}

func (m *LogoutMsg) Reset()                    { *m = LogoutMsg{} }
func (m *LogoutMsg) String() string            { return proto1.CompactTextString(m) }
func (*LogoutMsg) ProtoMessage()               {}
func (*LogoutMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// 订阅主题消息
type SubMsg struct {
}

func (m *SubMsg) Reset()                    { *m = SubMsg{} }
func (m *SubMsg) String() string            { return proto1.CompactTextString(m) }
func (*SubMsg) ProtoMessage()               {}
func (*SubMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// 取消订阅主题消息
type UnSubMsg struct {
}

func (m *UnSubMsg) Reset()                    { *m = UnSubMsg{} }
func (m *UnSubMsg) String() string            { return proto1.CompactTextString(m) }
func (*UnSubMsg) ProtoMessage()               {}
func (*UnSubMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Reply struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto1.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto1.RegisterType((*BPushMsg)(nil), "proto.BPushMsg")
	proto1.RegisterType((*SPushMsg)(nil), "proto.SPushMsg")
	proto1.RegisterType((*PChatMsg)(nil), "proto.PChatMsg")
	proto1.RegisterType((*GChatMsg)(nil), "proto.GChatMsg")
	proto1.RegisterType((*LoginMsg)(nil), "proto.LoginMsg")
	proto1.RegisterType((*LogoutMsg)(nil), "proto.LogoutMsg")
	proto1.RegisterType((*SubMsg)(nil), "proto.SubMsg")
	proto1.RegisterType((*UnSubMsg)(nil), "proto.UnSubMsg")
	proto1.RegisterType((*Reply)(nil), "proto.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rpc service

type RpcClient interface {
	// 推送接口
	BPush(ctx context.Context, in *BPushMsg, opts ...grpc.CallOption) (*Reply, error)
	SPush(ctx context.Context, in *SPushMsg, opts ...grpc.CallOption) (*Reply, error)
	PChat(ctx context.Context, in *PChatMsg, opts ...grpc.CallOption) (*Reply, error)
	GChat(ctx context.Context, in *GChatMsg, opts ...grpc.CallOption) (*Reply, error)
	// 用户相关接口
	Login(ctx context.Context, in *LoginMsg, opts ...grpc.CallOption) (*Reply, error)
	Logout(ctx context.Context, in *LogoutMsg, opts ...grpc.CallOption) (*Reply, error)
	// 用户订阅相关
	Subscribe(ctx context.Context, in *SubMsg, opts ...grpc.CallOption) (*Reply, error)
	UnSubscribe(ctx context.Context, in *UnSubMsg, opts ...grpc.CallOption) (*Reply, error)
}

type rpcClient struct {
	cc *grpc.ClientConn
}

func NewRpcClient(cc *grpc.ClientConn) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) BPush(ctx context.Context, in *BPushMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/BPush", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) SPush(ctx context.Context, in *SPushMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/SPush", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) PChat(ctx context.Context, in *PChatMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/PChat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GChat(ctx context.Context, in *GChatMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/GChat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Login(ctx context.Context, in *LoginMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Logout(ctx context.Context, in *LogoutMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Subscribe(ctx context.Context, in *SubMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) UnSubscribe(ctx context.Context, in *UnSubMsg, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/proto.Rpc/UnSubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rpc service

type RpcServer interface {
	// 推送接口
	BPush(context.Context, *BPushMsg) (*Reply, error)
	SPush(context.Context, *SPushMsg) (*Reply, error)
	PChat(context.Context, *PChatMsg) (*Reply, error)
	GChat(context.Context, *GChatMsg) (*Reply, error)
	// 用户相关接口
	Login(context.Context, *LoginMsg) (*Reply, error)
	Logout(context.Context, *LogoutMsg) (*Reply, error)
	// 用户订阅相关
	Subscribe(context.Context, *SubMsg) (*Reply, error)
	UnSubscribe(context.Context, *UnSubMsg) (*Reply, error)
}

func RegisterRpcServer(s *grpc.Server, srv RpcServer) {
	s.RegisterService(&_Rpc_serviceDesc, srv)
}

func _Rpc_BPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BPushMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).BPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/BPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).BPush(ctx, req.(*BPushMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_SPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SPushMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).SPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/SPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).SPush(ctx, req.(*SPushMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_PChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PChatMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).PChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/PChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).PChat(ctx, req.(*PChatMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GChatMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/GChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GChat(ctx, req.(*GChatMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Login(ctx, req.(*LoginMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Logout(ctx, req.(*LogoutMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Subscribe(ctx, req.(*SubMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_UnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSubMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).UnSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Rpc/UnSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).UnSubscribe(ctx, req.(*UnSubMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BPush",
			Handler:    _Rpc_BPush_Handler,
		},
		{
			MethodName: "SPush",
			Handler:    _Rpc_SPush_Handler,
		},
		{
			MethodName: "PChat",
			Handler:    _Rpc_PChat_Handler,
		},
		{
			MethodName: "GChat",
			Handler:    _Rpc_GChat_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Rpc_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Rpc_Logout_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _Rpc_Subscribe_Handler,
		},
		{
			MethodName: "UnSubscribe",
			Handler:    _Rpc_UnSubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x2d, 0x08, 0x81, 0xb1, 0x6a, 0xb3, 0x27, 0xf4, 0x64, 0xf6, 0xd4, 0xd4, 0x84, 0x83,
	0xbe, 0x81, 0x1e, 0xb8, 0xa8, 0x69, 0x20, 0x3e, 0x00, 0x60, 0x43, 0x49, 0x74, 0xd9, 0xc0, 0xee,
	0xc1, 0xa7, 0xf5, 0x55, 0xdc, 0x19, 0x76, 0x6b, 0x6a, 0x1b, 0xd2, 0x13, 0xff, 0x4c, 0xbe, 0xf9,
	0xb3, 0x7c, 0x10, 0xf7, 0xb2, 0x4e, 0x65, 0xdf, 0xa9, 0x8e, 0x05, 0xf4, 0xe1, 0x00, 0xd1, 0xd3,
	0x5a, 0x0f, 0xdb, 0xd7, 0xa1, 0xe1, 0x2b, 0x88, 0x0a, 0x9b, 0xd9, 0x15, 0x78, 0x4a, 0x25, 0xb3,
	0x3b, 0x7f, 0x39, 0xcf, 0x4d, 0xa2, 0x59, 0x27, 0x9e, 0x9d, 0x35, 0xde, 0xad, 0x9f, 0xb7, 0xa5,
	0xc2, 0x3b, 0x93, 0x33, 0x97, 0xdf, 0x20, 0x7a, 0xe9, 0x9a, 0x56, 0xd8, 0x8e, 0x52, 0x98, 0x8e,
	0x19, 0xde, 0x94, 0x02, 0x67, 0x2d, 0x4c, 0x07, 0xcd, 0x5a, 0xb0, 0x05, 0xf8, 0x75, 0xfb, 0x91,
	0xf8, 0x66, 0x11, 0xe4, 0x18, 0x71, 0xd3, 0xb4, 0x32, 0x39, 0x27, 0x04, 0x23, 0xbf, 0x87, 0xd8,
	0xf4, 0x75, 0x5a, 0x9d, 0x50, 0xc8, 0x23, 0x08, 0x0b, 0x5d, 0xd9, 0x27, 0xbd, 0x0b, 0x9b, 0x6f,
	0x20, 0xc8, 0x37, 0xf2, 0xf3, 0x1b, 0xdb, 0xbf, 0x86, 0xc6, 0xde, 0x63, 0x7c, 0xf8, 0xf1, 0xc0,
	0xcf, 0x65, 0xcd, 0x96, 0x10, 0x90, 0x05, 0x76, 0x3d, 0xda, 0x49, 0x9d, 0x93, 0xdb, 0xb9, 0x5d,
	0x50, 0x03, 0x3f, 0x43, 0xb2, 0xd8, 0x23, 0x8b, 0x09, 0x92, 0x0c, 0xed, 0x48, 0xe7, 0xeb, 0x18,
	0x99, 0xed, 0x91, 0xd9, 0x04, 0x49, 0x76, 0x77, 0xa4, 0x73, 0x7d, 0x40, 0xae, 0x20, 0x1c, 0xbd,
	0xb1, 0xc5, 0x1f, 0x3a, 0x6a, 0x3c, 0xc2, 0xc6, 0x46, 0xd5, 0x50, 0xf7, 0x6d, 0xb5, 0x61, 0x97,
	0xee, 0xbf, 0x48, 0xde, 0x01, 0x9b, 0xc2, 0x05, 0x89, 0xb5, 0xb4, 0x7b, 0x87, 0x93, 0xfd, 0x9f,
	0xaf, 0x42, 0x1a, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xaf, 0x56, 0xd3, 0x7a, 0x02,
	0x00, 0x00,
}
