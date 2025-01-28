// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: command.proto

package twitch

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

type MessageTypeCommand int32

const (
	MessageTypeCommand_TYPE_COMMAND_UNSPECIFIED           MessageTypeCommand = 0
	MessageTypeCommand_TYPE_COMMAND_GET_OAUTH_URL_REQ     MessageTypeCommand = 1
	MessageTypeCommand_TYPE_COMMAND_GET_OAUTH_URL_RES     MessageTypeCommand = 2
	MessageTypeCommand_TYPE_COMMAND_WRITE_PROFILE_REQ     MessageTypeCommand = 3
	MessageTypeCommand_TYPE_COMMAND_WRITE_PROFILE_RESP    MessageTypeCommand = 4
	MessageTypeCommand_TYPE_COMMAND_DELETE_PROFILE_REQ    MessageTypeCommand = 5
	MessageTypeCommand_TYPE_COMMAND_DELETE_PROFILE_RESP   MessageTypeCommand = 6
	MessageTypeCommand_TYPE_COMMAND_CHAT_SET_CONFIG_REQ   MessageTypeCommand = 7
	MessageTypeCommand_TYPE_COMMAND_CHAT_SET_CONFIG_RESP  MessageTypeCommand = 8
	MessageTypeCommand_TYPE_COMMAND_EVENT_SET_CONFIG_REQ  MessageTypeCommand = 9
	MessageTypeCommand_TYPE_COMMAND_EVENT_SET_CONFIG_RESP MessageTypeCommand = 10
)

// Enum value maps for MessageTypeCommand.
var (
	MessageTypeCommand_name = map[int32]string{
		0:  "TYPE_COMMAND_UNSPECIFIED",
		1:  "TYPE_COMMAND_GET_OAUTH_URL_REQ",
		2:  "TYPE_COMMAND_GET_OAUTH_URL_RES",
		3:  "TYPE_COMMAND_WRITE_PROFILE_REQ",
		4:  "TYPE_COMMAND_WRITE_PROFILE_RESP",
		5:  "TYPE_COMMAND_DELETE_PROFILE_REQ",
		6:  "TYPE_COMMAND_DELETE_PROFILE_RESP",
		7:  "TYPE_COMMAND_CHAT_SET_CONFIG_REQ",
		8:  "TYPE_COMMAND_CHAT_SET_CONFIG_RESP",
		9:  "TYPE_COMMAND_EVENT_SET_CONFIG_REQ",
		10: "TYPE_COMMAND_EVENT_SET_CONFIG_RESP",
	}
	MessageTypeCommand_value = map[string]int32{
		"TYPE_COMMAND_UNSPECIFIED":           0,
		"TYPE_COMMAND_GET_OAUTH_URL_REQ":     1,
		"TYPE_COMMAND_GET_OAUTH_URL_RES":     2,
		"TYPE_COMMAND_WRITE_PROFILE_REQ":     3,
		"TYPE_COMMAND_WRITE_PROFILE_RESP":    4,
		"TYPE_COMMAND_DELETE_PROFILE_REQ":    5,
		"TYPE_COMMAND_DELETE_PROFILE_RESP":   6,
		"TYPE_COMMAND_CHAT_SET_CONFIG_REQ":   7,
		"TYPE_COMMAND_CHAT_SET_CONFIG_RESP":  8,
		"TYPE_COMMAND_EVENT_SET_CONFIG_REQ":  9,
		"TYPE_COMMAND_EVENT_SET_CONFIG_RESP": 10,
	}
)

func (x MessageTypeCommand) Enum() *MessageTypeCommand {
	p := new(MessageTypeCommand)
	*p = x
	return p
}

func (x MessageTypeCommand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageTypeCommand) Descriptor() protoreflect.EnumDescriptor {
	return file_command_proto_enumTypes[0].Descriptor()
}

func (MessageTypeCommand) Type() protoreflect.EnumType {
	return &file_command_proto_enumTypes[0]
}

func (x MessageTypeCommand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageTypeCommand.Descriptor instead.
func (MessageTypeCommand) EnumDescriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

// Token stores a twitch OAuth token
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Access   string   `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
	Scopes   []string `protobuf:"bytes,3,rep,name=scopes,proto3" json:"scopes,omitempty"`
	UserId   string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Token) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

func (x *Token) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *Token) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Profile wraps a token and includes the user login
type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  *Token `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Profile) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profiles   map[string]*Profile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ChatConfig *ChatConfig         `protobuf:"bytes,2,opt,name=chat_config,json=chatConfig,proto3" json:"chat_config,omitempty"`
	EsConfig   *EventSubConfig     `protobuf:"bytes,3,opt,name=es_config,json=esConfig,proto3" json:"es_config,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetProfiles() map[string]*Profile {
	if x != nil {
		return x.Profiles
	}
	return nil
}

func (x *Config) GetChatConfig() *ChatConfig {
	if x != nil {
		return x.ChatConfig
	}
	return nil
}

func (x *Config) GetEsConfig() *EventSubConfig {
	if x != nil {
		return x.EsConfig
	}
	return nil
}

// Request a URL to do the OAuth enrollment dance
type GetOAuthURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOAuthURLRequest) Reset() {
	*x = GetOAuthURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthURLRequest) ProtoMessage() {}

func (x *GetOAuthURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthURLRequest.ProtoReflect.Descriptor instead.
func (*GetOAuthURLRequest) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{3}
}

type GetOAuthURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetOAuthURLResponse) Reset() {
	*x = GetOAuthURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthURLResponse) ProtoMessage() {}

func (x *GetOAuthURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthURLResponse.ProtoReflect.Descriptor instead.
func (*GetOAuthURLResponse) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{4}
}

func (x *GetOAuthURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type WriteProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *WriteProfileRequest) Reset() {
	*x = WriteProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteProfileRequest) ProtoMessage() {}

func (x *WriteProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteProfileRequest.ProtoReflect.Descriptor instead.
func (*WriteProfileRequest) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{5}
}

func (x *WriteProfileRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type WriteProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteProfileResponse) Reset() {
	*x = WriteProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteProfileResponse) ProtoMessage() {}

func (x *WriteProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteProfileResponse.ProtoReflect.Descriptor instead.
func (*WriteProfileResponse) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{6}
}

type DeleteProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteProfileRequest) Reset() {
	*x = DeleteProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileRequest) ProtoMessage() {}

func (x *DeleteProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileRequest.ProtoReflect.Descriptor instead.
func (*DeleteProfileRequest) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProfileResponse) Reset() {
	*x = DeleteProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileResponse) ProtoMessage() {}

func (x *DeleteProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileResponse.ProtoReflect.Descriptor instead.
func (*DeleteProfileResponse) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{8}
}

type ChatSetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *ChatConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ChatSetConfigRequest) Reset() {
	*x = ChatSetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatSetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatSetConfigRequest) ProtoMessage() {}

func (x *ChatSetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatSetConfigRequest.ProtoReflect.Descriptor instead.
func (*ChatSetConfigRequest) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{9}
}

func (x *ChatSetConfigRequest) GetConfig() *ChatConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type ChatSetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatSetConfigResponse) Reset() {
	*x = ChatSetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatSetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatSetConfigResponse) ProtoMessage() {}

func (x *ChatSetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatSetConfigResponse.ProtoReflect.Descriptor instead.
func (*ChatSetConfigResponse) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{10}
}

type EventSubSetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *EventSubConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *EventSubSetConfigRequest) Reset() {
	*x = EventSubSetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSubSetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSubSetConfigRequest) ProtoMessage() {}

func (x *EventSubSetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSubSetConfigRequest.ProtoReflect.Descriptor instead.
func (*EventSubSetConfigRequest) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{11}
}

func (x *EventSubSetConfigRequest) GetConfig() *EventSubConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type EventSubSetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventSubSetConfigResponse) Reset() {
	*x = EventSubSetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSubSetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSubSetConfigResponse) ProtoMessage() {}

func (x *EventSubSetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSubSetConfigResponse.ProtoReflect.Descriptor instead.
func (*EventSubSetConfigResponse) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{12}
}

var File_command_proto protoreflect.FileDescriptor

var file_command_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x1a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x5b, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xfa, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x63,
	0x68, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x33, 0x0a, 0x09, 0x65, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x4c,
	0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x14, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x27, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x40, 0x0a, 0x13, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x16, 0x0a,
	0x14, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x68,
	0x61, 0x74, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x17,
	0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x0a, 0x18, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x75, 0x62, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x75, 0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x1b, 0x0a, 0x19, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x53,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2a, 0xaa, 0x03, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f,
	0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x5f,
	0x55, 0x52, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x4f, 0x41,
	0x55, 0x54, 0x48, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x52, 0x45, 0x53, 0x10, 0x02, 0x12, 0x22, 0x0a,
	0x1e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x57, 0x52,
	0x49, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x10,
	0x03, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e,
	0x44, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f,
	0x52, 0x45, 0x53, 0x50, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x50, 0x52,
	0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x05, 0x12, 0x24, 0x0a, 0x20, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10,
	0x06, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e,
	0x44, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x07, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x53, 0x45, 0x54,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x08, 0x12, 0x25,
	0x0a, 0x21, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f,
	0x52, 0x45, 0x51, 0x10, 0x09, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f,
	0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x0a, 0x42, 0x21, 0x5a,
	0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x6e, 0x6f, 0x6d, 0x6f, 0x75, 0x73, 0x6b, 0x6f, 0x69, 0x2f, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_proto_rawDescOnce sync.Once
	file_command_proto_rawDescData = file_command_proto_rawDesc
)

func file_command_proto_rawDescGZIP() []byte {
	file_command_proto_rawDescOnce.Do(func() {
		file_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_proto_rawDescData)
	})
	return file_command_proto_rawDescData
}

var file_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_command_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_command_proto_goTypes = []any{
	(MessageTypeCommand)(0),           // 0: twitch.MessageTypeCommand
	(*Token)(nil),                     // 1: twitch.Token
	(*Profile)(nil),                   // 2: twitch.Profile
	(*Config)(nil),                    // 3: twitch.Config
	(*GetOAuthURLRequest)(nil),        // 4: twitch.GetOAuthURLRequest
	(*GetOAuthURLResponse)(nil),       // 5: twitch.GetOAuthURLResponse
	(*WriteProfileRequest)(nil),       // 6: twitch.WriteProfileRequest
	(*WriteProfileResponse)(nil),      // 7: twitch.WriteProfileResponse
	(*DeleteProfileRequest)(nil),      // 8: twitch.DeleteProfileRequest
	(*DeleteProfileResponse)(nil),     // 9: twitch.DeleteProfileResponse
	(*ChatSetConfigRequest)(nil),      // 10: twitch.ChatSetConfigRequest
	(*ChatSetConfigResponse)(nil),     // 11: twitch.ChatSetConfigResponse
	(*EventSubSetConfigRequest)(nil),  // 12: twitch.EventSubSetConfigRequest
	(*EventSubSetConfigResponse)(nil), // 13: twitch.EventSubSetConfigResponse
	nil,                               // 14: twitch.Config.ProfilesEntry
	(*ChatConfig)(nil),                // 15: twitch.ChatConfig
	(*EventSubConfig)(nil),            // 16: twitch.EventSubConfig
}
var file_command_proto_depIdxs = []int32{
	1,  // 0: twitch.Profile.token:type_name -> twitch.Token
	14, // 1: twitch.Config.profiles:type_name -> twitch.Config.ProfilesEntry
	15, // 2: twitch.Config.chat_config:type_name -> twitch.ChatConfig
	16, // 3: twitch.Config.es_config:type_name -> twitch.EventSubConfig
	2,  // 4: twitch.WriteProfileRequest.profile:type_name -> twitch.Profile
	15, // 5: twitch.ChatSetConfigRequest.config:type_name -> twitch.ChatConfig
	16, // 6: twitch.EventSubSetConfigRequest.config:type_name -> twitch.EventSubConfig
	2,  // 7: twitch.Config.ProfilesEntry.value:type_name -> twitch.Profile
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_command_proto_init() }
func file_command_proto_init() {
	if File_command_proto != nil {
		return
	}
	file_chat_proto_init()
	file_eventsub_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_command_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Token); i {
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
		file_command_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Profile); i {
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
		file_command_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Config); i {
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
		file_command_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetOAuthURLRequest); i {
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
		file_command_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetOAuthURLResponse); i {
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
		file_command_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*WriteProfileRequest); i {
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
		file_command_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*WriteProfileResponse); i {
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
		file_command_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteProfileRequest); i {
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
		file_command_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteProfileResponse); i {
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
		file_command_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ChatSetConfigRequest); i {
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
		file_command_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ChatSetConfigResponse); i {
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
		file_command_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*EventSubSetConfigRequest); i {
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
		file_command_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*EventSubSetConfigResponse); i {
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
			RawDescriptor: file_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_command_proto_goTypes,
		DependencyIndexes: file_command_proto_depIdxs,
		EnumInfos:         file_command_proto_enumTypes,
		MessageInfos:      file_command_proto_msgTypes,
	}.Build()
	File_command_proto = out.File
	file_command_proto_rawDesc = nil
	file_command_proto_goTypes = nil
	file_command_proto_depIdxs = nil
}
