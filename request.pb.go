// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: request.proto

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

type MessageTypeRequest int32

const (
	MessageTypeRequest_TYPE_REQUEST_UNSPECIFIED           MessageTypeRequest = 0
	MessageTypeRequest_TYPE_REQUEST_LIST_PROFILES_REQ     MessageTypeRequest = 1
	MessageTypeRequest_TYPE_REQUEST_LIST_PROFILES_RESP    MessageTypeRequest = 2
	MessageTypeRequest_TYPE_REQUEST_CHAT_GET_CONFIG_REQ   MessageTypeRequest = 3
	MessageTypeRequest_TYPE_REQUEST_CHAT_GET_CONFIG_RESP  MessageTypeRequest = 4
	MessageTypeRequest_TYPE_REQUEST_EVENT_GET_CONFIG_REQ  MessageTypeRequest = 5
	MessageTypeRequest_TYPE_REQUEST_EVENT_GET_CONFIG_RESP MessageTypeRequest = 6
	MessageTypeRequest_TYPE_REQUEST_EVENT_GET_STATUS_REQ  MessageTypeRequest = 7
	MessageTypeRequest_TYPE_REQUEST_EVENT_GET_STATUS_RESP MessageTypeRequest = 8
	MessageTypeRequest_TYPE_REQUEST_GET_USER_REQ          MessageTypeRequest = 9
	MessageTypeRequest_TYPE_REQUEST_GET_USER_RESP         MessageTypeRequest = 10
	MessageTypeRequest_TYPE_REQUEST_SEND_SHOUTOUT_REQ     MessageTypeRequest = 11
	MessageTypeRequest_TYPE_REQUEST_SEND_SHOUTOUT_RESP    MessageTypeRequest = 12
)

// Enum value maps for MessageTypeRequest.
var (
	MessageTypeRequest_name = map[int32]string{
		0:  "TYPE_REQUEST_UNSPECIFIED",
		1:  "TYPE_REQUEST_LIST_PROFILES_REQ",
		2:  "TYPE_REQUEST_LIST_PROFILES_RESP",
		3:  "TYPE_REQUEST_CHAT_GET_CONFIG_REQ",
		4:  "TYPE_REQUEST_CHAT_GET_CONFIG_RESP",
		5:  "TYPE_REQUEST_EVENT_GET_CONFIG_REQ",
		6:  "TYPE_REQUEST_EVENT_GET_CONFIG_RESP",
		7:  "TYPE_REQUEST_EVENT_GET_STATUS_REQ",
		8:  "TYPE_REQUEST_EVENT_GET_STATUS_RESP",
		9:  "TYPE_REQUEST_GET_USER_REQ",
		10: "TYPE_REQUEST_GET_USER_RESP",
		11: "TYPE_REQUEST_SEND_SHOUTOUT_REQ",
		12: "TYPE_REQUEST_SEND_SHOUTOUT_RESP",
	}
	MessageTypeRequest_value = map[string]int32{
		"TYPE_REQUEST_UNSPECIFIED":           0,
		"TYPE_REQUEST_LIST_PROFILES_REQ":     1,
		"TYPE_REQUEST_LIST_PROFILES_RESP":    2,
		"TYPE_REQUEST_CHAT_GET_CONFIG_REQ":   3,
		"TYPE_REQUEST_CHAT_GET_CONFIG_RESP":  4,
		"TYPE_REQUEST_EVENT_GET_CONFIG_REQ":  5,
		"TYPE_REQUEST_EVENT_GET_CONFIG_RESP": 6,
		"TYPE_REQUEST_EVENT_GET_STATUS_REQ":  7,
		"TYPE_REQUEST_EVENT_GET_STATUS_RESP": 8,
		"TYPE_REQUEST_GET_USER_REQ":          9,
		"TYPE_REQUEST_GET_USER_RESP":         10,
		"TYPE_REQUEST_SEND_SHOUTOUT_REQ":     11,
		"TYPE_REQUEST_SEND_SHOUTOUT_RESP":    12,
	}
)

func (x MessageTypeRequest) Enum() *MessageTypeRequest {
	p := new(MessageTypeRequest)
	*p = x
	return p
}

func (x MessageTypeRequest) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageTypeRequest) Descriptor() protoreflect.EnumDescriptor {
	return file_request_proto_enumTypes[0].Descriptor()
}

func (MessageTypeRequest) Type() protoreflect.EnumType {
	return &file_request_proto_enumTypes[0]
}

func (x MessageTypeRequest) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageTypeRequest.Descriptor instead.
func (MessageTypeRequest) EnumDescriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0}
}

type ListProfilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProfilesRequest) Reset() {
	*x = ListProfilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProfilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfilesRequest) ProtoMessage() {}

func (x *ListProfilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProfilesRequest.ProtoReflect.Descriptor instead.
func (*ListProfilesRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{0}
}

type ListProfilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *ListProfilesResponse) Reset() {
	*x = ListProfilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProfilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfilesResponse) ProtoMessage() {}

func (x *ListProfilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProfilesResponse.ProtoReflect.Descriptor instead.
func (*ListProfilesResponse) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{1}
}

func (x *ListProfilesResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type ChatGetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatGetConfigRequest) Reset() {
	*x = ChatGetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatGetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatGetConfigRequest) ProtoMessage() {}

func (x *ChatGetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatGetConfigRequest.ProtoReflect.Descriptor instead.
func (*ChatGetConfigRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{2}
}

type ChatGetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *ChatConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ChatGetConfigResponse) Reset() {
	*x = ChatGetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatGetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatGetConfigResponse) ProtoMessage() {}

func (x *ChatGetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatGetConfigResponse.ProtoReflect.Descriptor instead.
func (*ChatGetConfigResponse) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{3}
}

func (x *ChatGetConfigResponse) GetConfig() *ChatConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type EventSubGetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventSubGetConfigRequest) Reset() {
	*x = EventSubGetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSubGetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSubGetConfigRequest) ProtoMessage() {}

func (x *EventSubGetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSubGetConfigRequest.ProtoReflect.Descriptor instead.
func (*EventSubGetConfigRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{4}
}

type EventSubGetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *EventSubConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *EventSubGetConfigResponse) Reset() {
	*x = EventSubGetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSubGetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSubGetConfigResponse) ProtoMessage() {}

func (x *EventSubGetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSubGetConfigResponse.ProtoReflect.Descriptor instead.
func (*EventSubGetConfigResponse) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{5}
}

func (x *EventSubGetConfigResponse) GetConfig() *EventSubConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type EventSubGetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventSubGetStatusRequest) Reset() {
	*x = EventSubGetStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSubGetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSubGetStatusRequest) ProtoMessage() {}

func (x *EventSubGetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSubGetStatusRequest.ProtoReflect.Descriptor instead.
func (*EventSubGetStatusRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{6}
}

type EventSubGetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status EventSubStatus `protobuf:"varint,1,opt,name=status,proto3,enum=twitch.EventSubStatus" json:"status,omitempty"`
	Detail string         `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *EventSubGetStatusResponse) Reset() {
	*x = EventSubGetStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSubGetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSubGetStatusResponse) ProtoMessage() {}

func (x *EventSubGetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSubGetStatusResponse.ProtoReflect.Descriptor instead.
func (*EventSubGetStatusResponse) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{7}
}

func (x *EventSubGetStatusResponse) GetStatus() EventSubStatus {
	if x != nil {
		return x.Status
	}
	return EventSubStatus_EVENT_SUB_STATUS_UNKNOWN
}

func (x *EventSubGetStatusResponse) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type GetAvatarPathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *GetAvatarPathRequest) Reset() {
	*x = GetAvatarPathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvatarPathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvatarPathRequest) ProtoMessage() {}

func (x *GetAvatarPathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvatarPathRequest.ProtoReflect.Descriptor instead.
func (*GetAvatarPathRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{8}
}

func (x *GetAvatarPathRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type GetAvatarPathResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Path  string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetAvatarPathResponse) Reset() {
	*x = GetAvatarPathResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvatarPathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvatarPathResponse) ProtoMessage() {}

func (x *GetAvatarPathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvatarPathResponse.ProtoReflect.Descriptor instead.
func (*GetAvatarPathResponse) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{9}
}

func (x *GetAvatarPathResponse) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetAvatarPathResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile string `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Login   string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_request_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *GetUserRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	User  *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_request_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserResponse) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type SendShoutoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromProfile     string `protobuf:"bytes,1,opt,name=from_profile,json=fromProfile,proto3" json:"from_profile,omitempty"`
	ToBroadcasterId string `protobuf:"bytes,2,opt,name=to_broadcaster_id,json=toBroadcasterId,proto3" json:"to_broadcaster_id,omitempty"`
	ModeratorId     string `protobuf:"bytes,3,opt,name=moderator_id,json=moderatorId,proto3" json:"moderator_id,omitempty"`
}

func (x *SendShoutoutRequest) Reset() {
	*x = SendShoutoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendShoutoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendShoutoutRequest) ProtoMessage() {}

func (x *SendShoutoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendShoutoutRequest.ProtoReflect.Descriptor instead.
func (*SendShoutoutRequest) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{12}
}

func (x *SendShoutoutRequest) GetFromProfile() string {
	if x != nil {
		return x.FromProfile
	}
	return ""
}

func (x *SendShoutoutRequest) GetToBroadcasterId() string {
	if x != nil {
		return x.ToBroadcasterId
	}
	return ""
}

func (x *SendShoutoutRequest) GetModeratorId() string {
	if x != nil {
		return x.ModeratorId
	}
	return ""
}

type SendShoutoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendShoutoutResponse) Reset() {
	*x = SendShoutoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendShoutoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendShoutoutResponse) ProtoMessage() {}

func (x *SendShoutoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendShoutoutResponse.ProtoReflect.Descriptor instead.
func (*SendShoutoutResponse) Descriptor() ([]byte, []int) {
	return file_request_proto_rawDescGZIP(), []int{13}
}

var File_request_proto protoreflect.FileDescriptor

var file_request_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x1a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x74, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43,
	0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x1a, 0x0a, 0x18, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4b, 0x0a, 0x19, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x1a, 0x0a, 0x18,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x63, 0x0a, 0x19, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x75, 0x62, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x2c, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x41, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x40,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x22, 0x49, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x87, 0x01, 0x0a, 0x13,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x75, 0x74, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x6f, 0x5f, 0x62, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x74, 0x6f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x6f,
	0x75, 0x74, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0xee, 0x03,
	0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x53,
	0x5f, 0x52, 0x45, 0x51, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x46,
	0x49, 0x4c, 0x45, 0x53, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x54,
	0x5f, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x52, 0x45, 0x51, 0x10,
	0x03, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x47, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x04, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x47,
	0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x05, 0x12,
	0x26, 0x0a, 0x22, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x06, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x47, 0x45,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x07, 0x12, 0x26,
	0x0a, 0x22, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x52, 0x45, 0x53, 0x50, 0x10, 0x08, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x52, 0x45, 0x51, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x10, 0x0a, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x53, 0x48, 0x4f, 0x55, 0x54,
	0x4f, 0x55, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x0b, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x53,
	0x48, 0x4f, 0x55, 0x54, 0x4f, 0x55, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x0c, 0x42, 0x21,
	0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6e, 0x6f, 0x6d, 0x6f, 0x75, 0x73, 0x6b, 0x6f, 0x69, 0x2f, 0x74, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_request_proto_rawDescOnce sync.Once
	file_request_proto_rawDescData = file_request_proto_rawDesc
)

func file_request_proto_rawDescGZIP() []byte {
	file_request_proto_rawDescOnce.Do(func() {
		file_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_proto_rawDescData)
	})
	return file_request_proto_rawDescData
}

var file_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_request_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_request_proto_goTypes = []any{
	(MessageTypeRequest)(0),           // 0: twitch.MessageTypeRequest
	(*ListProfilesRequest)(nil),       // 1: twitch.ListProfilesRequest
	(*ListProfilesResponse)(nil),      // 2: twitch.ListProfilesResponse
	(*ChatGetConfigRequest)(nil),      // 3: twitch.ChatGetConfigRequest
	(*ChatGetConfigResponse)(nil),     // 4: twitch.ChatGetConfigResponse
	(*EventSubGetConfigRequest)(nil),  // 5: twitch.EventSubGetConfigRequest
	(*EventSubGetConfigResponse)(nil), // 6: twitch.EventSubGetConfigResponse
	(*EventSubGetStatusRequest)(nil),  // 7: twitch.EventSubGetStatusRequest
	(*EventSubGetStatusResponse)(nil), // 8: twitch.EventSubGetStatusResponse
	(*GetAvatarPathRequest)(nil),      // 9: twitch.GetAvatarPathRequest
	(*GetAvatarPathResponse)(nil),     // 10: twitch.GetAvatarPathResponse
	(*GetUserRequest)(nil),            // 11: twitch.GetUserRequest
	(*GetUserResponse)(nil),           // 12: twitch.GetUserResponse
	(*SendShoutoutRequest)(nil),       // 13: twitch.SendShoutoutRequest
	(*SendShoutoutResponse)(nil),      // 14: twitch.SendShoutoutResponse
	(*ChatConfig)(nil),                // 15: twitch.ChatConfig
	(*EventSubConfig)(nil),            // 16: twitch.EventSubConfig
	(EventSubStatus)(0),               // 17: twitch.EventSubStatus
	(*User)(nil),                      // 18: twitch.User
}
var file_request_proto_depIdxs = []int32{
	15, // 0: twitch.ChatGetConfigResponse.config:type_name -> twitch.ChatConfig
	16, // 1: twitch.EventSubGetConfigResponse.config:type_name -> twitch.EventSubConfig
	17, // 2: twitch.EventSubGetStatusResponse.status:type_name -> twitch.EventSubStatus
	18, // 3: twitch.GetUserResponse.user:type_name -> twitch.User
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_request_proto_init() }
func file_request_proto_init() {
	if File_request_proto != nil {
		return
	}
	file_chat_proto_init()
	file_eventsub_proto_init()
	file_twitch_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_request_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListProfilesRequest); i {
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
		file_request_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListProfilesResponse); i {
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
		file_request_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ChatGetConfigRequest); i {
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
		file_request_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ChatGetConfigResponse); i {
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
		file_request_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*EventSubGetConfigRequest); i {
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
		file_request_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*EventSubGetConfigResponse); i {
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
		file_request_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*EventSubGetStatusRequest); i {
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
		file_request_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*EventSubGetStatusResponse); i {
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
		file_request_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetAvatarPathRequest); i {
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
		file_request_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetAvatarPathResponse); i {
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
		file_request_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserRequest); i {
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
		file_request_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserResponse); i {
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
		file_request_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*SendShoutoutRequest); i {
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
		file_request_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*SendShoutoutResponse); i {
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
			RawDescriptor: file_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_proto_goTypes,
		DependencyIndexes: file_request_proto_depIdxs,
		EnumInfos:         file_request_proto_enumTypes,
		MessageInfos:      file_request_proto_msgTypes,
	}.Build()
	File_request_proto = out.File
	file_request_proto_rawDesc = nil
	file_request_proto_goTypes = nil
	file_request_proto_depIdxs = nil
}
