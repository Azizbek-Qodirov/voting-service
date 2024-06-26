// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto-service/voting-protos/public_vote.proto

package genprotos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Winner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElectionId  string `protobuf:"bytes,1,opt,name=ElectionId,proto3" json:"ElectionId,omitempty"`
	CandidateId string `protobuf:"bytes,2,opt,name=CandidateId,proto3" json:"CandidateId,omitempty"`
	Votes       int32  `protobuf:"varint,3,opt,name=Votes,proto3" json:"Votes,omitempty"`
}

func (x *Winner) Reset() {
	*x = Winner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Winner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Winner) ProtoMessage() {}

func (x *Winner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Winner.ProtoReflect.Descriptor instead.
func (*Winner) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_vote_proto_rawDescGZIP(), []int{0}
}

func (x *Winner) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *Winner) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

func (x *Winner) GetVotes() int32 {
	if x != nil {
		return x.Votes
	}
	return 0
}

type WhichElection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElectionId string `protobuf:"bytes,1,opt,name=ElectionId,proto3" json:"ElectionId,omitempty"`
}

func (x *WhichElection) Reset() {
	*x = WhichElection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_vote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhichElection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhichElection) ProtoMessage() {}

func (x *WhichElection) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_vote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhichElection.ProtoReflect.Descriptor instead.
func (*WhichElection) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_vote_proto_rawDescGZIP(), []int{1}
}

func (x *WhichElection) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

type PublicVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ElectionId  string `protobuf:"bytes,2,opt,name=ElectionId,proto3" json:"ElectionId,omitempty"`
	PublicId    string `protobuf:"bytes,3,opt,name=PublicId,proto3" json:"PublicId,omitempty"`
	CandidateId string `protobuf:"bytes,4,opt,name=CandidateId,proto3" json:"CandidateId,omitempty"`
}

func (x *PublicVote) Reset() {
	*x = PublicVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_vote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicVote) ProtoMessage() {}

func (x *PublicVote) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_vote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicVote.ProtoReflect.Descriptor instead.
func (*PublicVote) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_vote_proto_rawDescGZIP(), []int{2}
}

func (x *PublicVote) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublicVote) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *PublicVote) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *PublicVote) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

type PublicVoteCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElectionId  string `protobuf:"bytes,2,opt,name=ElectionId,proto3" json:"ElectionId,omitempty"`
	PublicId    string `protobuf:"bytes,3,opt,name=PublicId,proto3" json:"PublicId,omitempty"`
	CandidateId string `protobuf:"bytes,4,opt,name=CandidateId,proto3" json:"CandidateId,omitempty"`
}

func (x *PublicVoteCreate) Reset() {
	*x = PublicVoteCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_vote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicVoteCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicVoteCreate) ProtoMessage() {}

func (x *PublicVoteCreate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_vote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicVoteCreate.ProtoReflect.Descriptor instead.
func (*PublicVoteCreate) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_vote_proto_rawDescGZIP(), []int{3}
}

func (x *PublicVoteCreate) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *PublicVoteCreate) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *PublicVoteCreate) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

type GetAllPublicVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicVotes []*PublicVote `protobuf:"bytes,1,rep,name=publicVotes,proto3" json:"publicVotes,omitempty"`
}

func (x *GetAllPublicVote) Reset() {
	*x = GetAllPublicVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_vote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPublicVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPublicVote) ProtoMessage() {}

func (x *GetAllPublicVote) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_vote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPublicVote.ProtoReflect.Descriptor instead.
func (*GetAllPublicVote) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_vote_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllPublicVote) GetPublicVotes() []*PublicVote {
	if x != nil {
		return x.PublicVotes
	}
	return nil
}

var File_proto_service_voting_protos_public_vote_proto protoreflect.FileDescriptor

var file_proto_service_voting_protos_public_vote_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f, 0x74,
	0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x60, 0x0a, 0x06, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x6f,
	0x74, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x0d, 0x57, 0x68, 0x69, 0x63, 0x68, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x22, 0x70, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x22, 0x4e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74,
	0x65, 0x73, 0x32, 0xb4, 0x03, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x76,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x76,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x76, 0x6f, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x56, 0x6f, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x1a, 0x12,
	0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x18, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x75, 0x63, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x14,
	0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x56, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x57, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x68, 0x69, 0x63, 0x68, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_voting_protos_public_vote_proto_rawDescOnce sync.Once
	file_proto_service_voting_protos_public_vote_proto_rawDescData = file_proto_service_voting_protos_public_vote_proto_rawDesc
)

func file_proto_service_voting_protos_public_vote_proto_rawDescGZIP() []byte {
	file_proto_service_voting_protos_public_vote_proto_rawDescOnce.Do(func() {
		file_proto_service_voting_protos_public_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_voting_protos_public_vote_proto_rawDescData)
	})
	return file_proto_service_voting_protos_public_vote_proto_rawDescData
}

var file_proto_service_voting_protos_public_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_service_voting_protos_public_vote_proto_goTypes = []interface{}{
	(*Winner)(nil),           // 0: voting_proto.Winner
	(*WhichElection)(nil),    // 1: voting_proto.WhichElection
	(*PublicVote)(nil),       // 2: voting_proto.PublicVote
	(*PublicVoteCreate)(nil), // 3: voting_proto.PublicVoteCreate
	(*GetAllPublicVote)(nil), // 4: voting_proto.GetAllPublicVote
	(*ById)(nil),             // 5: voting_proto.ById
	(*Filter)(nil),           // 6: voting_proto.Filter
	(*Void)(nil),             // 7: voting_proto.Void
}
var file_proto_service_voting_protos_public_vote_proto_depIdxs = []int32{
	2, // 0: voting_proto.GetAllPublicVote.publicVotes:type_name -> voting_proto.PublicVote
	3, // 1: voting_proto.PublicVoteService.CreatePublicVote:input_type -> voting_proto.PublicVoteCreate
	5, // 2: voting_proto.PublicVoteService.DeletePublicVote:input_type -> voting_proto.ById
	2, // 3: voting_proto.PublicVoteService.UpdatePublicVote:input_type -> voting_proto.PublicVote
	5, // 4: voting_proto.PublicVoteService.GetByIdPublicVote:input_type -> voting_proto.ById
	6, // 5: voting_proto.PublicVoteService.GetAllPublucVotes:input_type -> voting_proto.Filter
	1, // 6: voting_proto.PublicVoteService.FindWinner:input_type -> voting_proto.WhichElection
	7, // 7: voting_proto.PublicVoteService.CreatePublicVote:output_type -> voting_proto.Void
	7, // 8: voting_proto.PublicVoteService.DeletePublicVote:output_type -> voting_proto.Void
	7, // 9: voting_proto.PublicVoteService.UpdatePublicVote:output_type -> voting_proto.Void
	2, // 10: voting_proto.PublicVoteService.GetByIdPublicVote:output_type -> voting_proto.PublicVote
	4, // 11: voting_proto.PublicVoteService.GetAllPublucVotes:output_type -> voting_proto.GetAllPublicVote
	0, // 12: voting_proto.PublicVoteService.FindWinner:output_type -> voting_proto.Winner
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_service_voting_protos_public_vote_proto_init() }
func file_proto_service_voting_protos_public_vote_proto_init() {
	if File_proto_service_voting_protos_public_vote_proto != nil {
		return
	}
	file_proto_service_voting_protos_void_proto_init()
	file_proto_service_voting_protos_election_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_service_voting_protos_public_vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Winner); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_voting_protos_public_vote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhichElection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_voting_protos_public_vote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicVote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_voting_protos_public_vote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicVoteCreate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_service_voting_protos_public_vote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPublicVote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_voting_protos_public_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_voting_protos_public_vote_proto_goTypes,
		DependencyIndexes: file_proto_service_voting_protos_public_vote_proto_depIdxs,
		MessageInfos:      file_proto_service_voting_protos_public_vote_proto_msgTypes,
	}.Build()
	File_proto_service_voting_protos_public_vote_proto = out.File
	file_proto_service_voting_protos_public_vote_proto_rawDesc = nil
	file_proto_service_voting_protos_public_vote_proto_goTypes = nil
	file_proto_service_voting_protos_public_vote_proto_depIdxs = nil
}
