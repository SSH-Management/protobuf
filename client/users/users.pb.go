// Code generated by protoc-gen-go-tag. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: users/users.proto

package users

import (
	_ "github.com/SSH-Management/protobuf/common"
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

type LinuxUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password     string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty" conform:"trim" validate:"required,max=50"`
	Shell        string   `protobuf:"bytes,5,opt,name=shell,proto3" json:"shell,omitempty" conform:"trim" validate:"required,min=6,max=50"`
	Username     string   `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty" validate:"required,max=50" conform:"trim"`
	SystemGroups []string `protobuf:"bytes,8,rep,name=system_groups,json=systemGroups,proto3" json:"system_groups,omitempty" validate:"dive,required,max=50"`
}

func (x *LinuxUser) Reset() {
	*x = LinuxUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinuxUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinuxUser) ProtoMessage() {}

func (x *LinuxUser) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinuxUser.ProtoReflect.Descriptor instead.
func (*LinuxUser) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{0}
}

func (x *LinuxUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LinuxUser) GetShell() string {
	if x != nil {
		return x.Shell
	}
	return ""
}

func (x *LinuxUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LinuxUser) GetSystemGroups() []string {
	if x != nil {
		return x.SystemGroups
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *LinuxUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" validate:"required"`
	PublicKey string     `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty" conform:"trim" validate:"required"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetUser() *LinuxUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateUserRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{2}
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteUserRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_users_users_proto protoreflect.FileDescriptor

var file_users_users_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x09, 0x4c, 0x69, 0x6e,
	0x75, 0x78, 0x55, 0x73, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xc2, 0xde, 0x1f, 0x2b, 0x0a, 0x29,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x74, 0x72, 0x69, 0x6d, 0x22, 0x20, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x2c, 0x6d, 0x61, 0x78, 0x3d, 0x35, 0x30, 0x22, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x4b, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x35, 0xc2, 0xde, 0x1f, 0x31, 0x0a, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x22, 0x74, 0x72, 0x69, 0x6d, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6d, 0x69, 0x6e, 0x3d,
	0x36, 0x2c, 0x6d, 0x61, 0x78, 0x3d, 0x35, 0x30, 0x22, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c,
	0x12, 0x4b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x2f, 0xc2, 0xde, 0x1f, 0x2b, 0x0a, 0x29, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x22, 0x74, 0x72, 0x69, 0x6d, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6d, 0x61, 0x78, 0x3d,
	0x35, 0x30, 0x22, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a,
	0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x25, 0xc2, 0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x64, 0x69, 0x76, 0x65, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x2c, 0x6d, 0x61, 0x78, 0x3d, 0x35, 0x30, 0x22, 0x52, 0x0c, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x75, 0x78, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x19, 0xc2, 0xde, 0x1f, 0x15, 0x0a, 0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a,
	0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x47, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xc2, 0xde, 0x1f, 0x24, 0x0a, 0x22, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x74, 0x72, 0x69, 0x6d, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x4c, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x53, 0x48, 0x2d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_users_proto_rawDescOnce sync.Once
	file_users_users_proto_rawDescData = file_users_users_proto_rawDesc
)

func file_users_users_proto_rawDescGZIP() []byte {
	file_users_users_proto_rawDescOnce.Do(func() {
		file_users_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_users_proto_rawDescData)
	})
	return file_users_users_proto_rawDescData
}

var file_users_users_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_users_users_proto_goTypes = []interface{}{
	(*LinuxUser)(nil),          // 0: users.LinuxUser
	(*CreateUserRequest)(nil),  // 1: users.CreateUserRequest
	(*CreateUserResponse)(nil), // 2: users.CreateUserResponse
	(*DeleteUserRequest)(nil),  // 3: users.DeleteUserRequest
}
var file_users_users_proto_depIdxs = []int32{
	0, // 0: users.CreateUserRequest.user:type_name -> users.LinuxUser
	1, // 1: users.UserService.Create:input_type -> users.CreateUserRequest
	2, // 2: users.UserService.Create:output_type -> users.CreateUserResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_users_users_proto_init() }
func file_users_users_proto_init() {
	if File_users_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinuxUser); i {
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
		file_users_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_users_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_users_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
			RawDescriptor: file_users_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_users_proto_goTypes,
		DependencyIndexes: file_users_users_proto_depIdxs,
		MessageInfos:      file_users_users_proto_msgTypes,
	}.Build()
	File_users_users_proto = out.File
	file_users_users_proto_rawDesc = nil
	file_users_users_proto_goTypes = nil
	file_users_users_proto_depIdxs = nil
}
