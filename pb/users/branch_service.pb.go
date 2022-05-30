// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.13.0
// source: users/branch_service.proto

package users

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

type ListBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	RegionId   string      `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
}

func (x *ListBranchRequest) Reset() {
	*x = ListBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_branch_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBranchRequest) ProtoMessage() {}

func (x *ListBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_branch_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBranchRequest.ProtoReflect.Descriptor instead.
func (*ListBranchRequest) Descriptor() ([]byte, []int) {
	return file_users_branch_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListBranchRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListBranchRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

type BranchPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	RegionId   string      `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Count      uint32      `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *BranchPaginationResponse) Reset() {
	*x = BranchPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_branch_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchPaginationResponse) ProtoMessage() {}

func (x *BranchPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_branch_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchPaginationResponse.ProtoReflect.Descriptor instead.
func (*BranchPaginationResponse) Descriptor() ([]byte, []int) {
	return file_users_branch_service_proto_rawDescGZIP(), []int{1}
}

func (x *BranchPaginationResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *BranchPaginationResponse) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *BranchPaginationResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListBranchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *BranchPaginationResponse `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Branch     *Branch                   `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *ListBranchResponse) Reset() {
	*x = ListBranchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_branch_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBranchResponse) ProtoMessage() {}

func (x *ListBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_branch_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBranchResponse.ProtoReflect.Descriptor instead.
func (*ListBranchResponse) Descriptor() ([]byte, []int) {
	return file_users_branch_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListBranchResponse) GetPagination() *BranchPaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListBranchResponse) GetBranch() *Branch {
	if x != nil {
		return x.Branch
	}
	return nil
}

var File_users_branch_service_proto protoreflect.FileDescriptor

var file_users_branch_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x69,
	0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x1a, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x18, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x8e, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x77, 0x69, 0x72,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2e, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x32, 0xcb, 0x02, 0x0a, 0x0d, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x77,
	0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x1a, 0x16, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x1a, 0x16, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x12, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x77, 0x69, 0x72,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x19,
	0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x35,
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65,
	0x72, 0x70, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x0e, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3b,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_branch_service_proto_rawDescOnce sync.Once
	file_users_branch_service_proto_rawDescData = file_users_branch_service_proto_rawDesc
)

func file_users_branch_service_proto_rawDescGZIP() []byte {
	file_users_branch_service_proto_rawDescOnce.Do(func() {
		file_users_branch_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_branch_service_proto_rawDescData)
	})
	return file_users_branch_service_proto_rawDescData
}

var file_users_branch_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_users_branch_service_proto_goTypes = []interface{}{
	(*ListBranchRequest)(nil),        // 0: wiradata.users.ListBranchRequest
	(*BranchPaginationResponse)(nil), // 1: wiradata.users.BranchPaginationResponse
	(*ListBranchResponse)(nil),       // 2: wiradata.users.ListBranchResponse
	(*Pagination)(nil),               // 3: wiradata.users.Pagination
	(*Branch)(nil),                   // 4: wiradata.users.Branch
	(*Id)(nil),                       // 5: wiradata.users.Id
	(*MyBoolean)(nil),                // 6: wiradata.users.MyBoolean
}
var file_users_branch_service_proto_depIdxs = []int32{
	3, // 0: wiradata.users.ListBranchRequest.pagination:type_name -> wiradata.users.Pagination
	3, // 1: wiradata.users.BranchPaginationResponse.pagination:type_name -> wiradata.users.Pagination
	1, // 2: wiradata.users.ListBranchResponse.pagination:type_name -> wiradata.users.BranchPaginationResponse
	4, // 3: wiradata.users.ListBranchResponse.branch:type_name -> wiradata.users.Branch
	4, // 4: wiradata.users.BranchService.Create:input_type -> wiradata.users.Branch
	4, // 5: wiradata.users.BranchService.Update:input_type -> wiradata.users.Branch
	5, // 6: wiradata.users.BranchService.View:input_type -> wiradata.users.Id
	5, // 7: wiradata.users.BranchService.Delete:input_type -> wiradata.users.Id
	0, // 8: wiradata.users.BranchService.List:input_type -> wiradata.users.ListBranchRequest
	4, // 9: wiradata.users.BranchService.Create:output_type -> wiradata.users.Branch
	4, // 10: wiradata.users.BranchService.Update:output_type -> wiradata.users.Branch
	4, // 11: wiradata.users.BranchService.View:output_type -> wiradata.users.Branch
	6, // 12: wiradata.users.BranchService.Delete:output_type -> wiradata.users.MyBoolean
	2, // 13: wiradata.users.BranchService.List:output_type -> wiradata.users.ListBranchResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_users_branch_service_proto_init() }
func file_users_branch_service_proto_init() {
	if File_users_branch_service_proto != nil {
		return
	}
	file_users_branch_message_proto_init()
	file_users_generic_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_users_branch_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBranchRequest); i {
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
		file_users_branch_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchPaginationResponse); i {
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
		file_users_branch_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBranchResponse); i {
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
			RawDescriptor: file_users_branch_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_branch_service_proto_goTypes,
		DependencyIndexes: file_users_branch_service_proto_depIdxs,
		MessageInfos:      file_users_branch_service_proto_msgTypes,
	}.Build()
	File_users_branch_service_proto = out.File
	file_users_branch_service_proto_rawDesc = nil
	file_users_branch_service_proto_goTypes = nil
	file_users_branch_service_proto_depIdxs = nil
}
