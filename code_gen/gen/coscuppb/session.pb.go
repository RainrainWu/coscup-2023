// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: code_gen/src/coscup/session.proto

package coscuppb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speaker *Speaker `protobuf:"bytes,1,opt,name=speaker,proto3" json:"speaker,omitempty"`
	Intro   string   `protobuf:"bytes,2,opt,name=intro,proto3" json:"intro,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_gen_src_coscup_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_code_gen_src_coscup_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_code_gen_src_coscup_session_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetSpeaker() *Speaker {
	if x != nil {
		return x.Speaker
	}
	return nil
}

func (x *Session) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

var file_code_gen_src_coscup_session_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50000,
		Name:          "coscup_2023.len_limit",
		Tag:           "varint,50000,opt,name=len_limit",
		Filename:      "code_gen/src/coscup/session.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional int32 len_limit = 50000;
	E_LenLimit = &file_code_gen_src_coscup_session_proto_extTypes[0]
)

var File_code_gen_src_coscup_session_proto protoreflect.FileDescriptor

var file_code_gen_src_coscup_session_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63,
	0x6f, 0x73, 0x63, 0x75, 0x70, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x73, 0x63, 0x75, 0x70, 0x5f, 0x32, 0x30, 0x32, 0x33,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x63, 0x6f, 0x73, 0x63, 0x75, 0x70, 0x2f, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x2e, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x73, 0x63, 0x75, 0x70, 0x5f, 0x32, 0x30, 0x32, 0x33, 0x2e,
	0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x07, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0x80, 0xb5, 0x18, 0x0f, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x3a, 0x3c, 0x0a, 0x09,
	0x6c, 0x65, 0x6e, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6c, 0x65, 0x6e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x65,
	0x6e, 0x2f, 0x63, 0x6f, 0x73, 0x63, 0x75, 0x70, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_code_gen_src_coscup_session_proto_rawDescOnce sync.Once
	file_code_gen_src_coscup_session_proto_rawDescData = file_code_gen_src_coscup_session_proto_rawDesc
)

func file_code_gen_src_coscup_session_proto_rawDescGZIP() []byte {
	file_code_gen_src_coscup_session_proto_rawDescOnce.Do(func() {
		file_code_gen_src_coscup_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_code_gen_src_coscup_session_proto_rawDescData)
	})
	return file_code_gen_src_coscup_session_proto_rawDescData
}

var file_code_gen_src_coscup_session_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_code_gen_src_coscup_session_proto_goTypes = []interface{}{
	(*Session)(nil),                   // 0: coscup_2023.Session
	(*Speaker)(nil),                   // 1: coscup_2023.Speaker
	(*descriptorpb.FieldOptions)(nil), // 2: google.protobuf.FieldOptions
}
var file_code_gen_src_coscup_session_proto_depIdxs = []int32{
	1, // 0: coscup_2023.Session.speaker:type_name -> coscup_2023.Speaker
	2, // 1: coscup_2023.len_limit:extendee -> google.protobuf.FieldOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_code_gen_src_coscup_session_proto_init() }
func file_code_gen_src_coscup_session_proto_init() {
	if File_code_gen_src_coscup_session_proto != nil {
		return
	}
	file_code_gen_src_coscup_speaker_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_code_gen_src_coscup_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
			RawDescriptor: file_code_gen_src_coscup_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_code_gen_src_coscup_session_proto_goTypes,
		DependencyIndexes: file_code_gen_src_coscup_session_proto_depIdxs,
		MessageInfos:      file_code_gen_src_coscup_session_proto_msgTypes,
		ExtensionInfos:    file_code_gen_src_coscup_session_proto_extTypes,
	}.Build()
	File_code_gen_src_coscup_session_proto = out.File
	file_code_gen_src_coscup_session_proto_rawDesc = nil
	file_code_gen_src_coscup_session_proto_goTypes = nil
	file_code_gen_src_coscup_session_proto_depIdxs = nil
}
