// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: protos/peer/peer.proto

package peer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EndorserClient is the client API for Endorser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndorserClient interface {
	ProcessProposal(ctx context.Context, in *SignedProposal, opts ...grpc.CallOption) (*ProposalResponse, error)
}

type endorserClient struct {
	cc grpc.ClientConnInterface
}

func NewEndorserClient(cc grpc.ClientConnInterface) EndorserClient {
	return &endorserClient{cc}
}

func (c *endorserClient) ProcessProposal(ctx context.Context, in *SignedProposal, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := c.cc.Invoke(ctx, "/peer.Endorser/ProcessProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndorserServer is the server API for Endorser service.
// All implementations must embed UnimplementedEndorserServer
// for forward compatibility
type EndorserServer interface {
	ProcessProposal(context.Context, *SignedProposal) (*ProposalResponse, error)
	mustEmbedUnimplementedEndorserServer()
}

// UnimplementedEndorserServer must be embedded to have forward compatible implementations.
type UnimplementedEndorserServer struct {
}

func (UnimplementedEndorserServer) ProcessProposal(context.Context, *SignedProposal) (*ProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessProposal not implemented")
}
func (UnimplementedEndorserServer) mustEmbedUnimplementedEndorserServer() {}

// UnsafeEndorserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndorserServer will
// result in compilation errors.
type UnsafeEndorserServer interface {
	mustEmbedUnimplementedEndorserServer()
}

func RegisterEndorserServer(s grpc.ServiceRegistrar, srv EndorserServer) {
	s.RegisterService(&Endorser_ServiceDesc, srv)
}

func _Endorser_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndorserServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer.Endorser/ProcessProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndorserServer).ProcessProposal(ctx, req.(*SignedProposal))
	}
	return interceptor(ctx, in, info, handler)
}

// Endorser_ServiceDesc is the grpc.ServiceDesc for Endorser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Endorser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "peer.Endorser",
	HandlerType: (*EndorserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessProposal",
			Handler:    _Endorser_ProcessProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/peer/peer.proto",
}
