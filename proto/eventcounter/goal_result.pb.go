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
// source: proto/eventcounter/goal_result.proto

package eventcounter

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

type GoalResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoalId           string             `protobuf:"bytes,1,opt,name=goal_id,json=goalId,proto3" json:"goal_id,omitempty"`
	VariationResults []*VariationResult `protobuf:"bytes,2,rep,name=variation_results,json=variationResults,proto3" json:"variation_results,omitempty"`
}

func (x *GoalResult) Reset() {
	*x = GoalResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eventcounter_goal_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalResult) ProtoMessage() {}

func (x *GoalResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eventcounter_goal_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalResult.ProtoReflect.Descriptor instead.
func (*GoalResult) Descriptor() ([]byte, []int) {
	return file_proto_eventcounter_goal_result_proto_rawDescGZIP(), []int{0}
}

func (x *GoalResult) GetGoalId() string {
	if x != nil {
		return x.GoalId
	}
	return ""
}

func (x *GoalResult) GetVariationResults() []*VariationResult {
	if x != nil {
		return x.VariationResults
	}
	return nil
}

var File_proto_eventcounter_goal_result_proto protoreflect.FileDescriptor

var file_proto_eventcounter_goal_result_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65,
	0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x29,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x0a, 0x47, 0x6f, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x6f, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x6f, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x54, 0x0a, 0x11, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x10, 0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2d, 0x69,
	0x6f, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_eventcounter_goal_result_proto_rawDescOnce sync.Once
	file_proto_eventcounter_goal_result_proto_rawDescData = file_proto_eventcounter_goal_result_proto_rawDesc
)

func file_proto_eventcounter_goal_result_proto_rawDescGZIP() []byte {
	file_proto_eventcounter_goal_result_proto_rawDescOnce.Do(func() {
		file_proto_eventcounter_goal_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_eventcounter_goal_result_proto_rawDescData)
	})
	return file_proto_eventcounter_goal_result_proto_rawDescData
}

var file_proto_eventcounter_goal_result_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_eventcounter_goal_result_proto_goTypes = []interface{}{
	(*GoalResult)(nil),      // 0: bucketeer.eventcounter.GoalResult
	(*VariationResult)(nil), // 1: bucketeer.eventcounter.VariationResult
}
var file_proto_eventcounter_goal_result_proto_depIdxs = []int32{
	1, // 0: bucketeer.eventcounter.GoalResult.variation_results:type_name -> bucketeer.eventcounter.VariationResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_eventcounter_goal_result_proto_init() }
func file_proto_eventcounter_goal_result_proto_init() {
	if File_proto_eventcounter_goal_result_proto != nil {
		return
	}
	file_proto_eventcounter_variation_result_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_eventcounter_goal_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalResult); i {
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
			RawDescriptor: file_proto_eventcounter_goal_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_eventcounter_goal_result_proto_goTypes,
		DependencyIndexes: file_proto_eventcounter_goal_result_proto_depIdxs,
		MessageInfos:      file_proto_eventcounter_goal_result_proto_msgTypes,
	}.Build()
	File_proto_eventcounter_goal_result_proto = out.File
	file_proto_eventcounter_goal_result_proto_rawDesc = nil
	file_proto_eventcounter_goal_result_proto_goTypes = nil
	file_proto_eventcounter_goal_result_proto_depIdxs = nil
}
