// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto-service/voting-protos/public.proto

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

type Public struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Phone     string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Birthday  string `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender    string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	PartyId   string `protobuf:"bytes,8,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	CreatedAt string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt int64  `protobuf:"varint,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Public) Reset() {
	*x = Public{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Public) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Public) ProtoMessage() {}

func (x *Public) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Public.ProtoReflect.Descriptor instead.
func (*Public) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_proto_rawDescGZIP(), []int{0}
}

func (x *Public) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Public) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Public) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Public) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Public) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Public) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *Public) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Public) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *Public) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Public) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Public) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type PublicCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Birthday string `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender   string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	PartyId  string `protobuf:"bytes,8,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *PublicCreate) Reset() {
	*x = PublicCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicCreate) ProtoMessage() {}

func (x *PublicCreate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicCreate.ProtoReflect.Descriptor instead.
func (*PublicCreate) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_proto_rawDescGZIP(), []int{1}
}

func (x *PublicCreate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublicCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PublicCreate) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PublicCreate) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *PublicCreate) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PublicCreate) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *PublicCreate) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *PublicCreate) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type PublicUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Birthday string `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender   string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	PartyId  string `protobuf:"bytes,8,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *PublicUpdate) Reset() {
	*x = PublicUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicUpdate) ProtoMessage() {}

func (x *PublicUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicUpdate.ProtoReflect.Descriptor instead.
func (*PublicUpdate) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_proto_rawDescGZIP(), []int{2}
}

func (x *PublicUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublicUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PublicUpdate) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PublicUpdate) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *PublicUpdate) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PublicUpdate) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *PublicUpdate) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *PublicUpdate) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type PublicDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PublicDelete) Reset() {
	*x = PublicDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicDelete) ProtoMessage() {}

func (x *PublicDelete) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicDelete.ProtoReflect.Descriptor instead.
func (*PublicDelete) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_proto_rawDescGZIP(), []int{3}
}

func (x *PublicDelete) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PublicGetById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PublicGetById) Reset() {
	*x = PublicGetById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicGetById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicGetById) ProtoMessage() {}

func (x *PublicGetById) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicGetById.ProtoReflect.Descriptor instead.
func (*PublicGetById) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_proto_rawDescGZIP(), []int{4}
}

func (x *PublicGetById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PublicGetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartyId  string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Birthday string `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender   string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *PublicGetAllReq) Reset() {
	*x = PublicGetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicGetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicGetAllReq) ProtoMessage() {}

func (x *PublicGetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicGetAllReq.ProtoReflect.Descriptor instead.
func (*PublicGetAllReq) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_proto_rawDescGZIP(), []int{5}
}

func (x *PublicGetAllReq) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PublicGetAllReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PublicGetAllReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PublicGetAllReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *PublicGetAllReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PublicGetAllReq) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *PublicGetAllReq) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type PublicGetAllResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Publics []*Public `protobuf:"bytes,1,rep,name=publics,proto3" json:"publics,omitempty"`
}

func (x *PublicGetAllResp) Reset() {
	*x = PublicGetAllResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_public_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicGetAllResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicGetAllResp) ProtoMessage() {}

func (x *PublicGetAllResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_public_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicGetAllResp.ProtoReflect.Descriptor instead.
func (*PublicGetAllResp) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_public_proto_rawDescGZIP(), []int{6}
}

func (x *PublicGetAllResp) GetPublics() []*Public {
	if x != nil {
		return x.Publics
	}
	return nil
}

var File_proto_service_voting_protos_public_proto protoreflect.FileDescriptor

var file_proto_service_voting_protos_public_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa1, 0x02, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49,
	0x64, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x22, 0x1e,
	0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f,
	0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xbd, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22,
	0x42, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x07, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x73, 0x32, 0xce, 0x02, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x14, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_voting_protos_public_proto_rawDescOnce sync.Once
	file_proto_service_voting_protos_public_proto_rawDescData = file_proto_service_voting_protos_public_proto_rawDesc
)

func file_proto_service_voting_protos_public_proto_rawDescGZIP() []byte {
	file_proto_service_voting_protos_public_proto_rawDescOnce.Do(func() {
		file_proto_service_voting_protos_public_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_voting_protos_public_proto_rawDescData)
	})
	return file_proto_service_voting_protos_public_proto_rawDescData
}

var file_proto_service_voting_protos_public_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_service_voting_protos_public_proto_goTypes = []interface{}{
	(*Public)(nil),           // 0: voting_proto.Public
	(*PublicCreate)(nil),     // 1: voting_proto.PublicCreate
	(*PublicUpdate)(nil),     // 2: voting_proto.PublicUpdate
	(*PublicDelete)(nil),     // 3: voting_proto.PublicDelete
	(*PublicGetById)(nil),    // 4: voting_proto.PublicGetById
	(*PublicGetAllReq)(nil),  // 5: voting_proto.PublicGetAllReq
	(*PublicGetAllResp)(nil), // 6: voting_proto.PublicGetAllResp
	(*Void)(nil),             // 7: voting_proto.Void
}
var file_proto_service_voting_protos_public_proto_depIdxs = []int32{
	0, // 0: voting_proto.PublicGetAllResp.publics:type_name -> voting_proto.Public
	1, // 1: voting_proto.PublicService.Create:input_type -> voting_proto.PublicCreate
	2, // 2: voting_proto.PublicService.Update:input_type -> voting_proto.PublicUpdate
	3, // 3: voting_proto.PublicService.Delete:input_type -> voting_proto.PublicDelete
	4, // 4: voting_proto.PublicService.GetById:input_type -> voting_proto.PublicGetById
	5, // 5: voting_proto.PublicService.GetAll:input_type -> voting_proto.PublicGetAllReq
	7, // 6: voting_proto.PublicService.Create:output_type -> voting_proto.Void
	7, // 7: voting_proto.PublicService.Update:output_type -> voting_proto.Void
	7, // 8: voting_proto.PublicService.Delete:output_type -> voting_proto.Void
	0, // 9: voting_proto.PublicService.GetById:output_type -> voting_proto.Public
	6, // 10: voting_proto.PublicService.GetAll:output_type -> voting_proto.PublicGetAllResp
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_service_voting_protos_public_proto_init() }
func file_proto_service_voting_protos_public_proto_init() {
	if File_proto_service_voting_protos_public_proto != nil {
		return
	}
	file_proto_service_voting_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_service_voting_protos_public_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Public); i {
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
		file_proto_service_voting_protos_public_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicCreate); i {
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
		file_proto_service_voting_protos_public_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicUpdate); i {
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
		file_proto_service_voting_protos_public_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicDelete); i {
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
		file_proto_service_voting_protos_public_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicGetById); i {
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
		file_proto_service_voting_protos_public_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicGetAllReq); i {
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
		file_proto_service_voting_protos_public_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicGetAllResp); i {
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
			RawDescriptor: file_proto_service_voting_protos_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_voting_protos_public_proto_goTypes,
		DependencyIndexes: file_proto_service_voting_protos_public_proto_depIdxs,
		MessageInfos:      file_proto_service_voting_protos_public_proto_msgTypes,
	}.Build()
	File_proto_service_voting_protos_public_proto = out.File
	file_proto_service_voting_protos_public_proto_rawDesc = nil
	file_proto_service_voting_protos_public_proto_goTypes = nil
	file_proto_service_voting_protos_public_proto_depIdxs = nil
}
