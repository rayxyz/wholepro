// Code generated by protoc-gen-go.
// source: cstruct.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	cstruct.proto

It has these top-level messages:
	ClientRequest
	ServerReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type ClientRequest struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg"`
}

func (m *ClientRequest) Reset()         { *m = ClientRequest{} }
func (m *ClientRequest) String() string { return proto1.CompactTextString(m) }
func (*ClientRequest) ProtoMessage()    {}

type ServerReply struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg"`
}

func (m *ServerReply) Reset()         { *m = ServerReply{} }
func (m *ServerReply) String() string { return proto1.CompactTextString(m) }
func (*ServerReply) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for CStruct service

type CStructClient interface {
	GetHiMsg(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ServerReply, error)
}

type cStructClient struct {
	cc *grpc.ClientConn
}

func NewCStructClient(cc *grpc.ClientConn) CStructClient {
	return &cStructClient{cc}
}

func (c *cStructClient) GetHiMsg(ctx context.Context, in *ClientRequest, opts ...grpc.CallOption) (*ServerReply, error) {
	out := new(ServerReply)
	err := grpc.Invoke(ctx, "/proto.CStruct/GetHiMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CStruct service

type CStructServer interface {
	GetHiMsg(context.Context, *ClientRequest) (*ServerReply, error)
}

func RegisterCStructServer(s *grpc.Server, srv CStructServer) {
	s.RegisterService(&_CStruct_serviceDesc, srv)
}

func _CStruct_GetHiMsg_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ClientRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CStructServer).GetHiMsg(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _CStruct_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CStruct",
	HandlerType: (*CStructServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHiMsg",
			Handler:    _CStruct_GetHiMsg_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
