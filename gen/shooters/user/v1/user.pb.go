// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: shooters/user/v1/user.proto

package userv1

import (
	phone_number "google.golang.org/genproto/googleapis/type/phone_number"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state               protoimpl.MessageState    `protogen:"open.v1"`
	Id                  string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName           string                    `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName            string                    `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email               string                    `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Username            string                    `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Role                string                    `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	PhoneNumber         *phone_number.PhoneNumber `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	IsApproved          bool                      `protobuf:"varint,9,opt,name=is_approved,json=isApproved,proto3" json:"is_approved,omitempty"`
	EmailVerified       bool                      `protobuf:"varint,10,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	CountryCode         string                    `protobuf:"bytes,11,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	PhoneNumberVerified bool                      `protobuf:"varint,12,opt,name=phone_number_verified,json=phoneNumberVerified,proto3" json:"phone_number_verified,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_shooters_user_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_user_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_shooters_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetPhoneNumber() *phone_number.PhoneNumber {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

func (x *User) GetIsApproved() bool {
	if x != nil {
		return x.IsApproved
	}
	return false
}

func (x *User) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *User) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *User) GetPhoneNumberVerified() bool {
	if x != nil {
		return x.PhoneNumberVerified
	}
	return false
}

type UserAttribute struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserAttribute) Reset() {
	*x = UserAttribute{}
	mi := &file_shooters_user_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAttribute) ProtoMessage() {}

func (x *UserAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_user_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAttribute.ProtoReflect.Descriptor instead.
func (*UserAttribute) Descriptor() ([]byte, []int) {
	return file_shooters_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserAttribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UserAttribute) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_shooters_user_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_user_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_shooters_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3,oneof" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	mi := &file_shooters_user_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_user_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_shooters_user_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ListUsersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          string                 `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	mi := &file_shooters_user_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_user_v1_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_shooters_user_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *ListUsersRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type ListUsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*User                `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	mi := &file_shooters_user_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_user_v1_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_shooters_user_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *ListUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *UserRequest           `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_shooters_user_v1_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_user_v1_user_proto_msgTypes[6]
	if x != nil {
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
	return file_shooters_user_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserRequest) GetUser() *UserRequest {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3,oneof" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_shooters_user_v1_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_user_v1_user_proto_msgTypes[7]
	if x != nil {
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
	return file_shooters_user_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *CreateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserRequest struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	FirstName     string                    `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                    `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email         string                    `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Username      string                    `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	PhoneNumber   *phone_number.PhoneNumber `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Role          string                    `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_shooters_user_v1_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shooters_user_v1_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_shooters_user_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *UserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRequest) GetPhoneNumber() *phone_number.PhoneNumber {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

func (x *UserRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_shooters_user_v1_user_proto protoreflect.FileDescriptor

const file_shooters_user_v1_user_proto_rawDesc = "" +
	"\n" +
	"\x1bshooters/user/v1/user.proto\x12\x10shooters.user.v1\x1a\x1egoogle/type/phone_number.proto\"\xf4\x02\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n" +
	"\n" +
	"first_name\x18\x02 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x03 \x01(\tR\blastName\x12\x14\n" +
	"\x05email\x18\x04 \x01(\tR\x05email\x12\x1a\n" +
	"\busername\x18\x05 \x01(\tR\busername\x12\x12\n" +
	"\x04role\x18\x06 \x01(\tR\x04role\x12;\n" +
	"\fphone_number\x18\a \x01(\v2\x18.google.type.PhoneNumberR\vphoneNumber\x12\x1f\n" +
	"\vis_approved\x18\t \x01(\bR\n" +
	"isApproved\x12%\n" +
	"\x0eemail_verified\x18\n" +
	" \x01(\bR\remailVerified\x12!\n" +
	"\fcountry_code\x18\v \x01(\tR\vcountryCode\x122\n" +
	"\x15phone_number_verified\x18\f \x01(\bR\x13phoneNumberVerified\"7\n" +
	"\rUserAttribute\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\" \n" +
	"\x0eGetUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"K\n" +
	"\x0fGetUserResponse\x12/\n" +
	"\x04user\x18\x01 \x01(\v2\x16.shooters.user.v1.UserH\x00R\x04user\x88\x01\x01B\a\n" +
	"\x05_user\"&\n" +
	"\x10ListUsersRequest\x12\x12\n" +
	"\x04role\x18\x01 \x01(\tR\x04role\"A\n" +
	"\x11ListUsersResponse\x12,\n" +
	"\x05users\x18\x01 \x03(\v2\x16.shooters.user.v1.UserR\x05users\"F\n" +
	"\x11CreateUserRequest\x121\n" +
	"\x04user\x18\x01 \x01(\v2\x1d.shooters.user.v1.UserRequestR\x04user\"N\n" +
	"\x12CreateUserResponse\x12/\n" +
	"\x04user\x18\x01 \x01(\v2\x16.shooters.user.v1.UserH\x00R\x04user\x88\x01\x01B\a\n" +
	"\x05_user\"\xcc\x01\n" +
	"\vUserRequest\x12\x1d\n" +
	"\n" +
	"first_name\x18\x01 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x02 \x01(\tR\blastName\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1a\n" +
	"\busername\x18\x04 \x01(\tR\busername\x12;\n" +
	"\fphone_number\x18\x05 \x01(\v2\x18.google.type.PhoneNumberR\vphoneNumber\x12\x12\n" +
	"\x04role\x18\x06 \x01(\tR\x04role2\x92\x02\n" +
	"\vUserService\x12V\n" +
	"\tListUsers\x12\".shooters.user.v1.ListUsersRequest\x1a#.shooters.user.v1.ListUsersResponse\"\x00\x12Y\n" +
	"\n" +
	"CreateUser\x12#.shooters.user.v1.CreateUserRequest\x1a$.shooters.user.v1.CreateUserResponse\"\x00\x12P\n" +
	"\aGetUser\x12 .shooters.user.v1.GetUserRequest\x1a!.shooters.user.v1.GetUserResponse\"\x00B\xc4\x01\n" +
	"\x14com.shooters.user.v1B\tUserProtoH\x02P\x01Z=github.com/shoot3rs/user/internal/gen/shooters/user/v1;userv1\xa2\x02\x03SUX\xaa\x02\x10Shooters.User.V1\xca\x02\x10Shooters\\User\\V1\xe2\x02\x1cShooters\\User\\V1\\GPBMetadata\xea\x02\x12Shooters::User::V1b\x06proto3"

var (
	file_shooters_user_v1_user_proto_rawDescOnce sync.Once
	file_shooters_user_v1_user_proto_rawDescData []byte
)

func file_shooters_user_v1_user_proto_rawDescGZIP() []byte {
	file_shooters_user_v1_user_proto_rawDescOnce.Do(func() {
		file_shooters_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shooters_user_v1_user_proto_rawDesc), len(file_shooters_user_v1_user_proto_rawDesc)))
	})
	return file_shooters_user_v1_user_proto_rawDescData
}

var file_shooters_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_shooters_user_v1_user_proto_goTypes = []any{
	(*User)(nil),                     // 0: shooters.user.v1.User
	(*UserAttribute)(nil),            // 1: shooters.user.v1.UserAttribute
	(*GetUserRequest)(nil),           // 2: shooters.user.v1.GetUserRequest
	(*GetUserResponse)(nil),          // 3: shooters.user.v1.GetUserResponse
	(*ListUsersRequest)(nil),         // 4: shooters.user.v1.ListUsersRequest
	(*ListUsersResponse)(nil),        // 5: shooters.user.v1.ListUsersResponse
	(*CreateUserRequest)(nil),        // 6: shooters.user.v1.CreateUserRequest
	(*CreateUserResponse)(nil),       // 7: shooters.user.v1.CreateUserResponse
	(*UserRequest)(nil),              // 8: shooters.user.v1.UserRequest
	(*phone_number.PhoneNumber)(nil), // 9: google.type.PhoneNumber
}
var file_shooters_user_v1_user_proto_depIdxs = []int32{
	9, // 0: shooters.user.v1.User.phone_number:type_name -> google.type.PhoneNumber
	0, // 1: shooters.user.v1.GetUserResponse.user:type_name -> shooters.user.v1.User
	0, // 2: shooters.user.v1.ListUsersResponse.users:type_name -> shooters.user.v1.User
	8, // 3: shooters.user.v1.CreateUserRequest.user:type_name -> shooters.user.v1.UserRequest
	0, // 4: shooters.user.v1.CreateUserResponse.user:type_name -> shooters.user.v1.User
	9, // 5: shooters.user.v1.UserRequest.phone_number:type_name -> google.type.PhoneNumber
	4, // 6: shooters.user.v1.UserService.ListUsers:input_type -> shooters.user.v1.ListUsersRequest
	6, // 7: shooters.user.v1.UserService.CreateUser:input_type -> shooters.user.v1.CreateUserRequest
	2, // 8: shooters.user.v1.UserService.GetUser:input_type -> shooters.user.v1.GetUserRequest
	5, // 9: shooters.user.v1.UserService.ListUsers:output_type -> shooters.user.v1.ListUsersResponse
	7, // 10: shooters.user.v1.UserService.CreateUser:output_type -> shooters.user.v1.CreateUserResponse
	3, // 11: shooters.user.v1.UserService.GetUser:output_type -> shooters.user.v1.GetUserResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_shooters_user_v1_user_proto_init() }
func file_shooters_user_v1_user_proto_init() {
	if File_shooters_user_v1_user_proto != nil {
		return
	}
	file_shooters_user_v1_user_proto_msgTypes[3].OneofWrappers = []any{}
	file_shooters_user_v1_user_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shooters_user_v1_user_proto_rawDesc), len(file_shooters_user_v1_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shooters_user_v1_user_proto_goTypes,
		DependencyIndexes: file_shooters_user_v1_user_proto_depIdxs,
		MessageInfos:      file_shooters_user_v1_user_proto_msgTypes,
	}.Build()
	File_shooters_user_v1_user_proto = out.File
	file_shooters_user_v1_user_proto_goTypes = nil
	file_shooters_user_v1_user_proto_depIdxs = nil
}
