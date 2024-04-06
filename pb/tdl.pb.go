// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: tdl.proto

package pb

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

type RunScriptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Script string `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
}

func (x *RunScriptRequest) Reset() {
	*x = RunScriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunScriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunScriptRequest) ProtoMessage() {}

func (x *RunScriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tdl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunScriptRequest.ProtoReflect.Descriptor instead.
func (*RunScriptRequest) Descriptor() ([]byte, []int) {
	return file_tdl_proto_rawDescGZIP(), []int{0}
}

func (x *RunScriptRequest) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

type RunScriptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccessStart bool `protobuf:"varint,1,opt,name=is_success_start,json=isSuccessStart,proto3" json:"is_success_start,omitempty"`
}

func (x *RunScriptResponse) Reset() {
	*x = RunScriptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunScriptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunScriptResponse) ProtoMessage() {}

func (x *RunScriptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tdl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunScriptResponse.ProtoReflect.Descriptor instead.
func (*RunScriptResponse) Descriptor() ([]byte, []int) {
	return file_tdl_proto_rawDescGZIP(), []int{1}
}

func (x *RunScriptResponse) GetIsSuccessStart() bool {
	if x != nil {
		return x.IsSuccessStart
	}
	return false
}

type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tdl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_tdl_proto_rawDescGZIP(), []int{2}
}

type ActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionType string `protobuf:"bytes,1,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	Payload    string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	RawMessage string `protobuf:"bytes,3,opt,name=raw_message,json=rawMessage,proto3" json:"raw_message,omitempty"`
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tdl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_tdl_proto_rawDescGZIP(), []int{3}
}

func (x *ActionResponse) GetActionType() string {
	if x != nil {
		return x.ActionType
	}
	return ""
}

func (x *ActionResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *ActionResponse) GetRawMessage() string {
	if x != nil {
		return x.RawMessage
	}
	return ""
}

type AuthInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *AuthInputRequest) Reset() {
	*x = AuthInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInputRequest) ProtoMessage() {}

func (x *AuthInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tdl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInputRequest.ProtoReflect.Descriptor instead.
func (*AuthInputRequest) Descriptor() ([]byte, []int) {
	return file_tdl_proto_rawDescGZIP(), []int{4}
}

func (x *AuthInputRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type AuthInputResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthInputResponse) Reset() {
	*x = AuthInputResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tdl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInputResponse) ProtoMessage() {}

func (x *AuthInputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tdl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInputResponse.ProtoReflect.Descriptor instead.
func (*AuthInputResponse) Descriptor() ([]byte, []int) {
	return file_tdl_proto_rawDescGZIP(), []int{5}
}

var File_tdl_proto protoreflect.FileDescriptor

var file_tdl_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x64, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x74, 0x64, 0x6c,
	0x22, 0x2a, 0x0a, 0x10, 0x52, 0x75, 0x6e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x22, 0x3d, 0x0a, 0x11,
	0x52, 0x75, 0x6e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x0e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x61, 0x77,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x61, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x41, 0x75,
	0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb9, 0x01, 0x0a, 0x03, 0x54, 0x64,
	0x6c, 0x12, 0x36, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x2e, 0x74, 0x64, 0x6c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x64, 0x6c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x52, 0x75, 0x6e,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x15, 0x2e, 0x74, 0x64, 0x6c, 0x2e, 0x52, 0x75, 0x6e,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x74, 0x64, 0x6c, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x15, 0x2e, 0x74, 0x64, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x64,
	0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x4d, 0x4c, 0x48, 0x65, 0x78, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x2f,
	0x74, 0x64, 0x6c, 0x5f, 0x66, 0x66, 0x69, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tdl_proto_rawDescOnce sync.Once
	file_tdl_proto_rawDescData = file_tdl_proto_rawDesc
)

func file_tdl_proto_rawDescGZIP() []byte {
	file_tdl_proto_rawDescOnce.Do(func() {
		file_tdl_proto_rawDescData = protoimpl.X.CompressGZIP(file_tdl_proto_rawDescData)
	})
	return file_tdl_proto_rawDescData
}

var file_tdl_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tdl_proto_goTypes = []interface{}{
	(*RunScriptRequest)(nil),  // 0: tdl.RunScriptRequest
	(*RunScriptResponse)(nil), // 1: tdl.RunScriptResponse
	(*ActionRequest)(nil),     // 2: tdl.ActionRequest
	(*ActionResponse)(nil),    // 3: tdl.ActionResponse
	(*AuthInputRequest)(nil),  // 4: tdl.AuthInputRequest
	(*AuthInputResponse)(nil), // 5: tdl.AuthInputResponse
}
var file_tdl_proto_depIdxs = []int32{
	2, // 0: tdl.Tdl.GetAction:input_type -> tdl.ActionRequest
	0, // 1: tdl.Tdl.RunScript:input_type -> tdl.RunScriptRequest
	4, // 2: tdl.Tdl.AuthInput:input_type -> tdl.AuthInputRequest
	3, // 3: tdl.Tdl.GetAction:output_type -> tdl.ActionResponse
	1, // 4: tdl.Tdl.RunScript:output_type -> tdl.RunScriptResponse
	5, // 5: tdl.Tdl.AuthInput:output_type -> tdl.AuthInputResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tdl_proto_init() }
func file_tdl_proto_init() {
	if File_tdl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tdl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunScriptRequest); i {
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
		file_tdl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunScriptResponse); i {
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
		file_tdl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRequest); i {
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
		file_tdl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResponse); i {
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
		file_tdl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInputRequest); i {
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
		file_tdl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInputResponse); i {
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
			RawDescriptor: file_tdl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tdl_proto_goTypes,
		DependencyIndexes: file_tdl_proto_depIdxs,
		MessageInfos:      file_tdl_proto_msgTypes,
	}.Build()
	File_tdl_proto = out.File
	file_tdl_proto_rawDesc = nil
	file_tdl_proto_goTypes = nil
	file_tdl_proto_depIdxs = nil
}