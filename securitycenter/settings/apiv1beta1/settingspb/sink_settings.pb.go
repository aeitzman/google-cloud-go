// Copyright 2020 Google LLC
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
// 	protoc        v3.21.5
// source: google/cloud/securitycenter/settings/v1beta1/sink_settings.proto

package settingspb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Sink Settings for Security Command Center
type SinkSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the project to send logs to. This project must be
	// part of the same organization where the Security Center API is
	// enabled. The format is `projects/{project}`. If it is empty, we do
	// not output logs. If a project ID is provided it will be normalized to a
	// project number.
	LoggingSinkProject string `protobuf:"bytes,1,opt,name=logging_sink_project,json=loggingSinkProject,proto3" json:"logging_sink_project,omitempty"`
}

func (x *SinkSettings) Reset() {
	*x = SinkSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SinkSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SinkSettings) ProtoMessage() {}

func (x *SinkSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SinkSettings.ProtoReflect.Descriptor instead.
func (*SinkSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDescGZIP(), []int{0}
}

func (x *SinkSettings) GetLoggingSinkProject() string {
	if x != nil {
		return x.LoggingSinkProject
	}
	return ""
}

var File_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73,
	0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x22, 0x40, 0x0a, 0x0c, 0x53, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x30, 0x0a, 0x14, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x6e, 0x6b,
	0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x42, 0xa9, 0x02, 0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x09, 0x53, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x3b, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x2c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x2c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x30, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDescData = file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDesc
)

func file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDescData
}

var file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_goTypes = []interface{}{
	(*SinkSettings)(nil), // 0: google.cloud.securitycenter.settings.v1beta1.SinkSettings
}
var file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_init() }
func file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_init() {
	if File_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SinkSettings); i {
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
			RawDescriptor: file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto = out.File
	file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_rawDesc = nil
	file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_goTypes = nil
	file_google_cloud_securitycenter_settings_v1beta1_sink_settings_proto_depIdxs = nil
}