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
// 	protoc        v3.12.2
// source: google/cloud/saasaccelerator/management/logs/v1/saas_instance_payload.proto

package logs

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

type InstanceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the event, e.g. Create, Update, etc.
	Verb string `protobuf:"bytes,1,opt,name=verb,proto3" json:"verb,omitempty"`
	// The state of the instance, e.g. "RETRYING_CREATE_INSTANCE".
	Stage string `protobuf:"bytes,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// A human-readable log message, e.g. "error in stage: CREATING, err: location
	// not available".
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	// The ID to uniquely locate all logs associated with a given request.
	TraceId string `protobuf:"bytes,4,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// The instance node which is the subject of the operation, if known.
	// Currently unused as tf actuation does not manage nodes.
	NodeId string `protobuf:"bytes,5,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *InstanceEvent) Reset() {
	*x = InstanceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceEvent) ProtoMessage() {}

func (x *InstanceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceEvent.ProtoReflect.Descriptor instead.
func (*InstanceEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceEvent) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *InstanceEvent) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

func (x *InstanceEvent) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *InstanceEvent) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *InstanceEvent) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

var File_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto protoreflect.FileDescriptor

var file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x61, 0x61, 0x73, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x61, 0x61, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x61, 0x61, 0x73,
	0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x7f,
	0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76,
	0x65, 0x72, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x42,
	0xa6, 0x01, 0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x53, 0x61, 0x61, 0x73, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x73, 0x61, 0x61, 0x73, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDescOnce sync.Once
	file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDescData = file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDesc
)

func file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDescGZIP() []byte {
	file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDescOnce.Do(func() {
		file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDescData)
	})
	return file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDescData
}

var file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_goTypes = []interface{}{
	(*InstanceEvent)(nil), // 0: google.cloud.saasaccelerator.management.logs.v1.InstanceEvent
}
var file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_init() }
func file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_init() {
	if File_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceEvent); i {
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
			RawDescriptor: file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_goTypes,
		DependencyIndexes: file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_depIdxs,
		MessageInfos:      file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_msgTypes,
	}.Build()
	File_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto = out.File
	file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_rawDesc = nil
	file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_goTypes = nil
	file_google_cloud_saasaccelerator_management_logs_v1_saas_instance_payload_proto_depIdxs = nil
}
