// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto-service/voting-protos/election.proto

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

type Election struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Election) Reset() {
	*x = Election{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_election_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Election) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Election) ProtoMessage() {}

func (x *Election) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_election_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Election.ProtoReflect.Descriptor instead.
func (*Election) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_election_proto_rawDescGZIP(), []int{0}
}

func (x *Election) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Election) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Election) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date     string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Party    string `protobuf:"bytes,2,opt,name=party,proto3" json:"party,omitempty"`
	Slogan   string `protobuf:"bytes,3,opt,name=slogan,proto3" json:"slogan,omitempty"`
	Gender   string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Nation   string `protobuf:"bytes,5,opt,name=nation,proto3" json:"nation,omitempty"`
	Age      int32  `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"`
	Election string `protobuf:"bytes,7,opt,name=election,proto3" json:"election,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_election_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_election_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_election_proto_rawDescGZIP(), []int{1}
}

func (x *Filter) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Filter) GetParty() string {
	if x != nil {
		return x.Party
	}
	return ""
}

func (x *Filter) GetSlogan() string {
	if x != nil {
		return x.Slogan
	}
	return ""
}

func (x *Filter) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Filter) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

func (x *Filter) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Filter) GetElection() string {
	if x != nil {
		return x.Election
	}
	return ""
}

type GetAllElection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elections []*Election `protobuf:"bytes,1,rep,name=elections,proto3" json:"elections,omitempty"`
}

func (x *GetAllElection) Reset() {
	*x = GetAllElection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_voting_protos_election_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllElection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllElection) ProtoMessage() {}

func (x *GetAllElection) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_voting_protos_election_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllElection.ProtoReflect.Descriptor instead.
func (*GetAllElection) Descriptor() ([]byte, []int) {
	return file_proto_service_voting_protos_election_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllElection) GetElections() []*Election {
	if x != nil {
		return x.Elections
	}
	return nil
}

var File_proto_service_voting_protos_election_proto protoreflect.FileDescriptor

var file_proto_service_voting_protos_election_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x42, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6c, 0x6f, 0x67, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f,
	0x67, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x09, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xd7, 0x02, 0x0a, 0x0f, 0x45, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x79, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x76,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x16, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_voting_protos_election_proto_rawDescOnce sync.Once
	file_proto_service_voting_protos_election_proto_rawDescData = file_proto_service_voting_protos_election_proto_rawDesc
)

func file_proto_service_voting_protos_election_proto_rawDescGZIP() []byte {
	file_proto_service_voting_protos_election_proto_rawDescOnce.Do(func() {
		file_proto_service_voting_protos_election_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_voting_protos_election_proto_rawDescData)
	})
	return file_proto_service_voting_protos_election_proto_rawDescData
}

var file_proto_service_voting_protos_election_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_service_voting_protos_election_proto_goTypes = []interface{}{
	(*Election)(nil),       // 0: voting_proto.Election
	(*Filter)(nil),         // 1: voting_proto.Filter
	(*GetAllElection)(nil), // 2: voting_proto.GetAllElection
	(*ById)(nil),           // 3: voting_proto.ById
	(*Void)(nil),           // 4: voting_proto.Void
}
var file_proto_service_voting_protos_election_proto_depIdxs = []int32{
	0, // 0: voting_proto.GetAllElection.elections:type_name -> voting_proto.Election
	0, // 1: voting_proto.ElectionService.CreateElection:input_type -> voting_proto.Election
	3, // 2: voting_proto.ElectionService.DeleteElection:input_type -> voting_proto.ById
	0, // 3: voting_proto.ElectionService.UpdateElection:input_type -> voting_proto.Election
	3, // 4: voting_proto.ElectionService.GetByIdElection:input_type -> voting_proto.ById
	1, // 5: voting_proto.ElectionService.GetAllElections:input_type -> voting_proto.Filter
	4, // 6: voting_proto.ElectionService.CreateElection:output_type -> voting_proto.Void
	4, // 7: voting_proto.ElectionService.DeleteElection:output_type -> voting_proto.Void
	4, // 8: voting_proto.ElectionService.UpdateElection:output_type -> voting_proto.Void
	0, // 9: voting_proto.ElectionService.GetByIdElection:output_type -> voting_proto.Election
	2, // 10: voting_proto.ElectionService.GetAllElections:output_type -> voting_proto.GetAllElection
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_service_voting_protos_election_proto_init() }
func file_proto_service_voting_protos_election_proto_init() {
	if File_proto_service_voting_protos_election_proto != nil {
		return
	}
	file_proto_service_voting_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_service_voting_protos_election_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Election); i {
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
		file_proto_service_voting_protos_election_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_proto_service_voting_protos_election_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllElection); i {
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
			RawDescriptor: file_proto_service_voting_protos_election_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_voting_protos_election_proto_goTypes,
		DependencyIndexes: file_proto_service_voting_protos_election_proto_depIdxs,
		MessageInfos:      file_proto_service_voting_protos_election_proto_msgTypes,
	}.Build()
	File_proto_service_voting_protos_election_proto = out.File
	file_proto_service_voting_protos_election_proto_rawDesc = nil
	file_proto_service_voting_protos_election_proto_goTypes = nil
	file_proto_service_voting_protos_election_proto_depIdxs = nil
}
