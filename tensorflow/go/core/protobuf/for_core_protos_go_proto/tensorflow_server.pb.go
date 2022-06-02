// Copyright 2016 The TensorFlow Authors. All Rights Reserved.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//==============================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.8.0
// source: tensorflow/core/protobuf/tensorflow_server.proto

package for_core_protos_go_proto

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

// Defines the configuration of a single TensorFlow server.
type ServerDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The cluster of which this server is a member.
	Cluster *ClusterDef `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// The name of the job of which this server is a member.
	//
	// NOTE(mrry): The `cluster` field must contain a `JobDef` with a `name` field
	// that matches this name.
	JobName string `protobuf:"bytes,2,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	// The task index of this server in its job.
	//
	// NOTE: The `cluster` field must contain a `JobDef` with a matching `name`
	// and a mapping in its `tasks` field for this index.
	TaskIndex int32 `protobuf:"varint,3,opt,name=task_index,json=taskIndex,proto3" json:"task_index,omitempty"`
	// The default configuration for sessions that run on this server.
	DefaultSessionConfig *ConfigProto `protobuf:"bytes,4,opt,name=default_session_config,json=defaultSessionConfig,proto3" json:"default_session_config,omitempty"`
	// The protocol to be used by this server.
	//
	// Acceptable values include: "grpc", "grpc+verbs".
	Protocol string `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// The server port. If not set, then we identify the port from the job_name.
	Port int32 `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	// Device filters for remote tasks in the cluster.
	// NOTE: This is an experimental feature and only effective in TensorFlow 2.x.
	ClusterDeviceFilters *ClusterDeviceFilters `protobuf:"bytes,7,opt,name=cluster_device_filters,json=clusterDeviceFilters,proto3" json:"cluster_device_filters,omitempty"`
}

func (x *ServerDef) Reset() {
	*x = ServerDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_tensorflow_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerDef) ProtoMessage() {}

func (x *ServerDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_tensorflow_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerDef.ProtoReflect.Descriptor instead.
func (*ServerDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_tensorflow_server_proto_rawDescGZIP(), []int{0}
}

func (x *ServerDef) GetCluster() *ClusterDef {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *ServerDef) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *ServerDef) GetTaskIndex() int32 {
	if x != nil {
		return x.TaskIndex
	}
	return 0
}

func (x *ServerDef) GetDefaultSessionConfig() *ConfigProto {
	if x != nil {
		return x.DefaultSessionConfig
	}
	return nil
}

func (x *ServerDef) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ServerDef) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ServerDef) GetClusterDeviceFilters() *ClusterDeviceFilters {
	if x != nil {
		return x.ClusterDeviceFilters
	}
	return nil
}

var File_tensorflow_core_protobuf_tensorflow_server_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_tensorflow_server_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x26,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x02, 0x0a,
	0x09, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x66, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x44, 0x65, 0x66, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x4d, 0x0a, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
	0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x56, 0x0a, 0x16, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x14, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x42, 0x86, 0x01,
	0x0a, 0x1a, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_tensorflow_server_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_tensorflow_server_proto_rawDescData = file_tensorflow_core_protobuf_tensorflow_server_proto_rawDesc
)

func file_tensorflow_core_protobuf_tensorflow_server_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_tensorflow_server_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_tensorflow_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_tensorflow_server_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_tensorflow_server_proto_rawDescData
}

var file_tensorflow_core_protobuf_tensorflow_server_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_tensorflow_server_proto_goTypes = []interface{}{
	(*ServerDef)(nil),            // 0: tensorflow.ServerDef
	(*ClusterDef)(nil),           // 1: tensorflow.ClusterDef
	(*ConfigProto)(nil),          // 2: tensorflow.ConfigProto
	(*ClusterDeviceFilters)(nil), // 3: tensorflow.ClusterDeviceFilters
}
var file_tensorflow_core_protobuf_tensorflow_server_proto_depIdxs = []int32{
	1, // 0: tensorflow.ServerDef.cluster:type_name -> tensorflow.ClusterDef
	2, // 1: tensorflow.ServerDef.default_session_config:type_name -> tensorflow.ConfigProto
	3, // 2: tensorflow.ServerDef.cluster_device_filters:type_name -> tensorflow.ClusterDeviceFilters
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_tensorflow_server_proto_init() }
func file_tensorflow_core_protobuf_tensorflow_server_proto_init() {
	if File_tensorflow_core_protobuf_tensorflow_server_proto != nil {
		return
	}
	file_tensorflow_core_protobuf_cluster_proto_init()
	file_tensorflow_core_protobuf_config_proto_init()
	file_tensorflow_core_protobuf_device_filters_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_tensorflow_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerDef); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_tensorflow_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_tensorflow_server_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_tensorflow_server_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_tensorflow_server_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_tensorflow_server_proto = out.File
	file_tensorflow_core_protobuf_tensorflow_server_proto_rawDesc = nil
	file_tensorflow_core_protobuf_tensorflow_server_proto_goTypes = nil
	file_tensorflow_core_protobuf_tensorflow_server_proto_depIdxs = nil
}
