// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: proto/vatz/manager/v1/manager.proto

package v1

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type CommandStatus int32

const (
	CommandStatus_SUCCESS CommandStatus = 0
	CommandStatus_FAIL    CommandStatus = 1
)

// Enum value maps for CommandStatus.
var (
	CommandStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
	}
	CommandStatus_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
	}
)

func (x CommandStatus) Enum() *CommandStatus {
	p := new(CommandStatus)
	*p = x
	return p
}

func (x CommandStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_vatz_manager_v1_manager_proto_enumTypes[0].Descriptor()
}

func (CommandStatus) Type() protoreflect.EnumType {
	return &file_proto_vatz_manager_v1_manager_proto_enumTypes[0]
}

func (x CommandStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandStatus.Descriptor instead.
func (CommandStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{0}
}

type ExecuteResponse_State int32

const (
	ExecuteResponse_NONE        ExecuteResponse_State = 0
	ExecuteResponse_PENDING     ExecuteResponse_State = 1
	ExecuteResponse_IN_PROGRESS ExecuteResponse_State = 2
	ExecuteResponse_SUCCESS     ExecuteResponse_State = 3
	ExecuteResponse_FAILURE     ExecuteResponse_State = 4
	ExecuteResponse_TIMEOUT     ExecuteResponse_State = 5
)

// Enum value maps for ExecuteResponse_State.
var (
	ExecuteResponse_State_name = map[int32]string{
		0: "NONE",
		1: "PENDING",
		2: "IN_PROGRESS",
		3: "SUCCESS",
		4: "FAILURE",
		5: "TIMEOUT",
	}
	ExecuteResponse_State_value = map[string]int32{
		"NONE":        0,
		"PENDING":     1,
		"IN_PROGRESS": 2,
		"SUCCESS":     3,
		"FAILURE":     4,
		"TIMEOUT":     5,
	}
)

func (x ExecuteResponse_State) Enum() *ExecuteResponse_State {
	p := new(ExecuteResponse_State)
	*p = x
	return p
}

func (x ExecuteResponse_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExecuteResponse_State) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_vatz_manager_v1_manager_proto_enumTypes[1].Descriptor()
}

func (ExecuteResponse_State) Type() protoreflect.EnumType {
	return &file_proto_vatz_manager_v1_manager_proto_enumTypes[1]
}

func (x ExecuteResponse_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExecuteResponse_State.Descriptor instead.
func (ExecuteResponse_State) EnumDescriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{7, 0}
}

// Metadata 레퍼런스로 항상 넣어서 사용하지 않더라도 넣어주는 값
// option 필수값은 아님 항상 변경이 될 수 있는 값
type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string          `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port    int32           `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Version float32         `protobuf:"fixed32,4,opt,name=version,proto3" json:"version,omitempty"` // version >< revision 때문에 정확한 타입지정이 필요
	Option  *_struct.Struct `protobuf:"bytes,5,opt,name=option,proto3" json:"option,omitempty"`     // mapType 변환
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InitRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InitRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *InitRequest) GetVersion() float32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *InitRequest) GetOption() *_struct.Struct {
	if x != nil {
		return x.Option
	}
	return nil
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   CommandStatus   `protobuf:"varint,1,opt,name=result,proto3,enum=vatz.manager.CommandStatus" json:"result,omitempty"`
	Metadata *_struct.Struct `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{1}
}

func (x *InitResponse) GetResult() CommandStatus {
	if x != nil {
		return x.Result
	}
	return CommandStatus_SUCCESS
}

func (x *InitResponse) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

type VerifyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result CommandStatus `protobuf:"varint,1,opt,name=result,proto3,enum=vatz.manager.CommandStatus" json:"result,omitempty"`
}

func (x *VerifyInfo) Reset() {
	*x = VerifyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyInfo) ProtoMessage() {}

func (x *VerifyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyInfo.ProtoReflect.Descriptor instead.
func (*VerifyInfo) Descriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyInfo) GetResult() CommandStatus {
	if x != nil {
		return x.Result
	}
	return CommandStatus_SUCCESS
}

type EndRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *_struct.Struct `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *EndRequest) Reset() {
	*x = EndRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndRequest) ProtoMessage() {}

func (x *EndRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndRequest.ProtoReflect.Descriptor instead.
func (*EndRequest) Descriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{4}
}

func (x *EndRequest) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type EndResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result CommandStatus `protobuf:"varint,1,opt,name=result,proto3,enum=vatz.manager.CommandStatus" json:"result,omitempty"`
}

func (x *EndResponse) Reset() {
	*x = EndResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndResponse) ProtoMessage() {}

func (x *EndResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndResponse.ProtoReflect.Descriptor instead.
func (*EndResponse) Descriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{5}
}

func (x *EndResponse) GetResult() CommandStatus {
	if x != nil {
		return x.Result
	}
	return CommandStatus_SUCCESS
}

type ExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// is_required: true
	TargetInfo *_struct.Struct `protobuf:"bytes,1,opt,name=target_info,json=targetInfo,proto3" json:"target_info,omitempty"`
	// is_required: true
	Command *_struct.Struct `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{6}
}

func (x *ExecuteRequest) GetTargetInfo() *_struct.Struct {
	if x != nil {
		return x.TargetInfo
	}
	return nil
}

func (x *ExecuteRequest) GetCommand() *_struct.Struct {
	if x != nil {
		return x.Command
	}
	return nil
}

type ExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State    ExecuteResponse_State `protobuf:"varint,1,opt,name=state,proto3,enum=vatz.manager.ExecuteResponse_State" json:"state,omitempty"`
	Message  string                `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Protocol string                `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Options  *_struct.Struct       `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{7}
}

func (x *ExecuteResponse) GetState() ExecuteResponse_State {
	if x != nil {
		return x.State
	}
	return ExecuteResponse_NONE
}

func (x *ExecuteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ExecuteResponse) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ExecuteResponse) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{8}
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vatz_manager_v1_manager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_proto_vatz_manager_v1_manager_proto_rawDescGZIP(), []int{9}
}

var File_proto_vatz_manager_v1_manager_proto protoreflect.FileDescriptor

var file_proto_vatz_manager_v1_manager_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x74, 0x7a, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a,
	0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x78,
	0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x42, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x41, 0x0a, 0x0a,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x76, 0x61, 0x74,
	0x7a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x41, 0x0a, 0x0a, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x42, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x7d, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x31, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x8d, 0x02, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x56, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x46,
	0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45,
	0x4f, 0x55, 0x54, 0x10, 0x05, 0x22, 0x0f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x26, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01,
	0x32, 0xe2, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x04,
	0x49, 0x6e, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x76, 0x61, 0x74, 0x7a,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x61, 0x74, 0x7a, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_vatz_manager_v1_manager_proto_rawDescOnce sync.Once
	file_proto_vatz_manager_v1_manager_proto_rawDescData = file_proto_vatz_manager_v1_manager_proto_rawDesc
)

func file_proto_vatz_manager_v1_manager_proto_rawDescGZIP() []byte {
	file_proto_vatz_manager_v1_manager_proto_rawDescOnce.Do(func() {
		file_proto_vatz_manager_v1_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vatz_manager_v1_manager_proto_rawDescData)
	})
	return file_proto_vatz_manager_v1_manager_proto_rawDescData
}

var file_proto_vatz_manager_v1_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_vatz_manager_v1_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_vatz_manager_v1_manager_proto_goTypes = []interface{}{
	(CommandStatus)(0),         // 0: vatz.manager.CommandStatus
	(ExecuteResponse_State)(0), // 1: vatz.manager.ExecuteResponse.State
	(*InitRequest)(nil),        // 2: vatz.manager.InitRequest
	(*InitResponse)(nil),       // 3: vatz.manager.InitResponse
	(*VerifyRequest)(nil),      // 4: vatz.manager.VerifyRequest
	(*VerifyInfo)(nil),         // 5: vatz.manager.VerifyInfo
	(*EndRequest)(nil),         // 6: vatz.manager.EndRequest
	(*EndResponse)(nil),        // 7: vatz.manager.EndResponse
	(*ExecuteRequest)(nil),     // 8: vatz.manager.ExecuteRequest
	(*ExecuteResponse)(nil),    // 9: vatz.manager.ExecuteResponse
	(*UpdateRequest)(nil),      // 10: vatz.manager.UpdateRequest
	(*UpdateResponse)(nil),     // 11: vatz.manager.UpdateResponse
	(*_struct.Struct)(nil),     // 12: google.protobuf.Struct
}
var file_proto_vatz_manager_v1_manager_proto_depIdxs = []int32{
	12, // 0: vatz.manager.InitRequest.option:type_name -> google.protobuf.Struct
	0,  // 1: vatz.manager.InitResponse.result:type_name -> vatz.manager.CommandStatus
	12, // 2: vatz.manager.InitResponse.metadata:type_name -> google.protobuf.Struct
	12, // 3: vatz.manager.VerifyRequest.options:type_name -> google.protobuf.Struct
	0,  // 4: vatz.manager.VerifyInfo.result:type_name -> vatz.manager.CommandStatus
	12, // 5: vatz.manager.EndRequest.metadata:type_name -> google.protobuf.Struct
	0,  // 6: vatz.manager.EndResponse.result:type_name -> vatz.manager.CommandStatus
	12, // 7: vatz.manager.ExecuteRequest.target_info:type_name -> google.protobuf.Struct
	12, // 8: vatz.manager.ExecuteRequest.command:type_name -> google.protobuf.Struct
	1,  // 9: vatz.manager.ExecuteResponse.state:type_name -> vatz.manager.ExecuteResponse.State
	12, // 10: vatz.manager.ExecuteResponse.options:type_name -> google.protobuf.Struct
	2,  // 11: vatz.manager.Manager.Init:input_type -> vatz.manager.InitRequest
	4,  // 12: vatz.manager.Manager.Verify:input_type -> vatz.manager.VerifyRequest
	6,  // 13: vatz.manager.Manager.End:input_type -> vatz.manager.EndRequest
	8,  // 14: vatz.manager.Manager.Execute:input_type -> vatz.manager.ExecuteRequest
	10, // 15: vatz.manager.Manager.UpdateConfig:input_type -> vatz.manager.UpdateRequest
	3,  // 16: vatz.manager.Manager.Init:output_type -> vatz.manager.InitResponse
	5,  // 17: vatz.manager.Manager.Verify:output_type -> vatz.manager.VerifyInfo
	7,  // 18: vatz.manager.Manager.End:output_type -> vatz.manager.EndResponse
	9,  // 19: vatz.manager.Manager.Execute:output_type -> vatz.manager.ExecuteResponse
	11, // 20: vatz.manager.Manager.UpdateConfig:output_type -> vatz.manager.UpdateResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proto_vatz_manager_v1_manager_proto_init() }
func file_proto_vatz_manager_v1_manager_proto_init() {
	if File_proto_vatz_manager_v1_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_vatz_manager_v1_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_proto_vatz_manager_v1_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse); i {
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
		file_proto_vatz_manager_v1_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_proto_vatz_manager_v1_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyInfo); i {
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
		file_proto_vatz_manager_v1_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndRequest); i {
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
		file_proto_vatz_manager_v1_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndResponse); i {
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
		file_proto_vatz_manager_v1_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteRequest); i {
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
		file_proto_vatz_manager_v1_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteResponse); i {
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
		file_proto_vatz_manager_v1_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_vatz_manager_v1_manager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_proto_vatz_manager_v1_manager_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_vatz_manager_v1_manager_proto_goTypes,
		DependencyIndexes: file_proto_vatz_manager_v1_manager_proto_depIdxs,
		EnumInfos:         file_proto_vatz_manager_v1_manager_proto_enumTypes,
		MessageInfos:      file_proto_vatz_manager_v1_manager_proto_msgTypes,
	}.Build()
	File_proto_vatz_manager_v1_manager_proto = out.File
	file_proto_vatz_manager_v1_manager_proto_rawDesc = nil
	file_proto_vatz_manager_v1_manager_proto_goTypes = nil
	file_proto_vatz_manager_v1_manager_proto_depIdxs = nil
}
