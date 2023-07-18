// Copyright 2023 The Bucketeer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: proto/feature/reason.proto

package feature

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

type Reason_Type int32

const (
	Reason_TARGET        Reason_Type = 0
	Reason_RULE          Reason_Type = 1
	Reason_DEFAULT       Reason_Type = 3
	Reason_CLIENT        Reason_Type = 4
	Reason_OFF_VARIATION Reason_Type = 5
	Reason_PREREQUISITE  Reason_Type = 6
)

// Enum value maps for Reason_Type.
var (
	Reason_Type_name = map[int32]string{
		0: "TARGET",
		1: "RULE",
		3: "DEFAULT",
		4: "CLIENT",
		5: "OFF_VARIATION",
		6: "PREREQUISITE",
	}
	Reason_Type_value = map[string]int32{
		"TARGET":        0,
		"RULE":          1,
		"DEFAULT":       3,
		"CLIENT":        4,
		"OFF_VARIATION": 5,
		"PREREQUISITE":  6,
	}
)

func (x Reason_Type) Enum() *Reason_Type {
	p := new(Reason_Type)
	*p = x
	return p
}

func (x Reason_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reason_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_feature_reason_proto_enumTypes[0].Descriptor()
}

func (Reason_Type) Type() protoreflect.EnumType {
	return &file_proto_feature_reason_proto_enumTypes[0]
}

func (x Reason_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reason_Type.Descriptor instead.
func (Reason_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_feature_reason_proto_rawDescGZIP(), []int{0, 0}
}

type Reason struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   Reason_Type `protobuf:"varint,1,opt,name=type,proto3,enum=bucketeer.feature.Reason_Type" json:"type,omitempty"`
	RuleId string      `protobuf:"bytes,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
}

func (x *Reason) Reset() {
	*x = Reason{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feature_reason_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reason) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reason) ProtoMessage() {}

func (x *Reason) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feature_reason_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reason.ProtoReflect.Descriptor instead.
func (*Reason) Descriptor() ([]byte, []int) {
	return file_proto_feature_reason_proto_rawDescGZIP(), []int{0}
}

func (x *Reason) GetType() Reason_Type {
	if x != nil {
		return x.Type
	}
	return Reason_TARGET
}

func (x *Reason) GetRuleId() string {
	if x != nil {
		return x.RuleId
	}
	return ""
}

var File_proto_feature_reason_proto protoreflect.FileDescriptor

var file_proto_feature_reason_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0xb1, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x65, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52,
	0x55, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x11,
	0x0a, 0x0d, 0x4f, 0x46, 0x46, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x05, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x45, 0x52, 0x45, 0x51, 0x55, 0x49, 0x53, 0x49, 0x54,
	0x45, 0x10, 0x06, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_feature_reason_proto_rawDescOnce sync.Once
	file_proto_feature_reason_proto_rawDescData = file_proto_feature_reason_proto_rawDesc
)

func file_proto_feature_reason_proto_rawDescGZIP() []byte {
	file_proto_feature_reason_proto_rawDescOnce.Do(func() {
		file_proto_feature_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_feature_reason_proto_rawDescData)
	})
	return file_proto_feature_reason_proto_rawDescData
}

var file_proto_feature_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_feature_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_feature_reason_proto_goTypes = []interface{}{
	(Reason_Type)(0), // 0: bucketeer.feature.Reason.Type
	(*Reason)(nil),   // 1: bucketeer.feature.Reason
}
var file_proto_feature_reason_proto_depIdxs = []int32{
	0, // 0: bucketeer.feature.Reason.type:type_name -> bucketeer.feature.Reason.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_feature_reason_proto_init() }
func file_proto_feature_reason_proto_init() {
	if File_proto_feature_reason_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_feature_reason_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reason); i {
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
			RawDescriptor: file_proto_feature_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_feature_reason_proto_goTypes,
		DependencyIndexes: file_proto_feature_reason_proto_depIdxs,
		EnumInfos:         file_proto_feature_reason_proto_enumTypes,
		MessageInfos:      file_proto_feature_reason_proto_msgTypes,
	}.Build()
	File_proto_feature_reason_proto = out.File
	file_proto_feature_reason_proto_rawDesc = nil
	file_proto_feature_reason_proto_goTypes = nil
	file_proto_feature_reason_proto_depIdxs = nil
}
