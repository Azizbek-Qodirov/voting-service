// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto-service/voting-protos/public_vote.proto

package genprotos

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

// PublicVoteServiceClient is the client API for PublicVoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicVoteServiceClient interface {
	CreatePublicVote(ctx context.Context, in *PublicVoteCreate, opts ...grpc.CallOption) (*Void, error)
	DeletePublicVote(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	UpdatePublicVote(ctx context.Context, in *PublicVote, opts ...grpc.CallOption) (*Void, error)
	GetByIdPublicVote(ctx context.Context, in *ById, opts ...grpc.CallOption) (*PublicVote, error)
	GetAllPublucVotes(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*GetAllPublicVote, error)
	FindWinner(ctx context.Context, in *WhichElection, opts ...grpc.CallOption) (*Winner, error)
}

type publicVoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicVoteServiceClient(cc grpc.ClientConnInterface) PublicVoteServiceClient {
	return &publicVoteServiceClient{cc}
}

func (c *publicVoteServiceClient) CreatePublicVote(ctx context.Context, in *PublicVoteCreate, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/voting_proto.PublicVoteService/CreatePublicVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicVoteServiceClient) DeletePublicVote(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/voting_proto.PublicVoteService/DeletePublicVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicVoteServiceClient) UpdatePublicVote(ctx context.Context, in *PublicVote, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/voting_proto.PublicVoteService/UpdatePublicVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicVoteServiceClient) GetByIdPublicVote(ctx context.Context, in *ById, opts ...grpc.CallOption) (*PublicVote, error) {
	out := new(PublicVote)
	err := c.cc.Invoke(ctx, "/voting_proto.PublicVoteService/GetByIdPublicVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicVoteServiceClient) GetAllPublucVotes(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*GetAllPublicVote, error) {
	out := new(GetAllPublicVote)
	err := c.cc.Invoke(ctx, "/voting_proto.PublicVoteService/GetAllPublucVotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicVoteServiceClient) FindWinner(ctx context.Context, in *WhichElection, opts ...grpc.CallOption) (*Winner, error) {
	out := new(Winner)
	err := c.cc.Invoke(ctx, "/voting_proto.PublicVoteService/FindWinner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicVoteServiceServer is the server API for PublicVoteService service.
// All implementations must embed UnimplementedPublicVoteServiceServer
// for forward compatibility
type PublicVoteServiceServer interface {
	CreatePublicVote(context.Context, *PublicVoteCreate) (*Void, error)
	DeletePublicVote(context.Context, *ById) (*Void, error)
	UpdatePublicVote(context.Context, *PublicVote) (*Void, error)
	GetByIdPublicVote(context.Context, *ById) (*PublicVote, error)
	GetAllPublucVotes(context.Context, *Filter) (*GetAllPublicVote, error)
	FindWinner(context.Context, *WhichElection) (*Winner, error)
	mustEmbedUnimplementedPublicVoteServiceServer()
}

// UnimplementedPublicVoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublicVoteServiceServer struct {
}

func (UnimplementedPublicVoteServiceServer) CreatePublicVote(context.Context, *PublicVoteCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePublicVote not implemented")
}
func (UnimplementedPublicVoteServiceServer) DeletePublicVote(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePublicVote not implemented")
}
func (UnimplementedPublicVoteServiceServer) UpdatePublicVote(context.Context, *PublicVote) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePublicVote not implemented")
}
func (UnimplementedPublicVoteServiceServer) GetByIdPublicVote(context.Context, *ById) (*PublicVote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdPublicVote not implemented")
}
func (UnimplementedPublicVoteServiceServer) GetAllPublucVotes(context.Context, *Filter) (*GetAllPublicVote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPublucVotes not implemented")
}
func (UnimplementedPublicVoteServiceServer) FindWinner(context.Context, *WhichElection) (*Winner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindWinner not implemented")
}
func (UnimplementedPublicVoteServiceServer) mustEmbedUnimplementedPublicVoteServiceServer() {}

// UnsafePublicVoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicVoteServiceServer will
// result in compilation errors.
type UnsafePublicVoteServiceServer interface {
	mustEmbedUnimplementedPublicVoteServiceServer()
}

func RegisterPublicVoteServiceServer(s grpc.ServiceRegistrar, srv PublicVoteServiceServer) {
	s.RegisterService(&PublicVoteService_ServiceDesc, srv)
}

func _PublicVoteService_CreatePublicVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicVoteCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).CreatePublicVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting_proto.PublicVoteService/CreatePublicVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).CreatePublicVote(ctx, req.(*PublicVoteCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicVoteService_DeletePublicVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).DeletePublicVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting_proto.PublicVoteService/DeletePublicVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).DeletePublicVote(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicVoteService_UpdatePublicVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).UpdatePublicVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting_proto.PublicVoteService/UpdatePublicVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).UpdatePublicVote(ctx, req.(*PublicVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicVoteService_GetByIdPublicVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).GetByIdPublicVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting_proto.PublicVoteService/GetByIdPublicVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).GetByIdPublicVote(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicVoteService_GetAllPublucVotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).GetAllPublucVotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting_proto.PublicVoteService/GetAllPublucVotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).GetAllPublucVotes(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicVoteService_FindWinner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhichElection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).FindWinner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting_proto.PublicVoteService/FindWinner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).FindWinner(ctx, req.(*WhichElection))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicVoteService_ServiceDesc is the grpc.ServiceDesc for PublicVoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicVoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "voting_proto.PublicVoteService",
	HandlerType: (*PublicVoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePublicVote",
			Handler:    _PublicVoteService_CreatePublicVote_Handler,
		},
		{
			MethodName: "DeletePublicVote",
			Handler:    _PublicVoteService_DeletePublicVote_Handler,
		},
		{
			MethodName: "UpdatePublicVote",
			Handler:    _PublicVoteService_UpdatePublicVote_Handler,
		},
		{
			MethodName: "GetByIdPublicVote",
			Handler:    _PublicVoteService_GetByIdPublicVote_Handler,
		},
		{
			MethodName: "GetAllPublucVotes",
			Handler:    _PublicVoteService_GetAllPublucVotes_Handler,
		},
		{
			MethodName: "FindWinner",
			Handler:    _PublicVoteService_FindWinner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto-service/voting-protos/public_vote.proto",
}