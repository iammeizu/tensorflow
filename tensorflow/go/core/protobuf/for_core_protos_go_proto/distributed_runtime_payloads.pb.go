// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.8.0
// source: tensorflow/core/protobuf/distributed_runtime_payloads.proto

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

// Used to serialize and transmit tensorflow::Status payloads through
// grpc::Status `error_details` since grpc::Status lacks payload API.
// TODO(b/204231601): Use GRPC API once supported.
type GrpcPayloadContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payloads map[string][]byte `protobuf:"bytes,1,rep,name=payloads,proto3" json:"payloads,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GrpcPayloadContainer) Reset() {
	*x = GrpcPayloadContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcPayloadContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcPayloadContainer) ProtoMessage() {}

func (x *GrpcPayloadContainer) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcPayloadContainer.ProtoReflect.Descriptor instead.
func (*GrpcPayloadContainer) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcPayloadContainer) GetPayloads() map[string][]byte {
	if x != nil {
		return x.Payloads
	}
	return nil
}

// If included as a payload, this message flags the Status to have lost payloads
// during the GRPC transmission.
// URI: "type.googleapis.com/tensorflow.distributed_runtime.GrpcPayloadsLost"
type GrpcPayloadsLost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GrpcPayloadsLost) Reset() {
	*x = GrpcPayloadsLost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcPayloadsLost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcPayloadsLost) ProtoMessage() {}

func (x *GrpcPayloadsLost) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcPayloadsLost.ProtoReflect.Descriptor instead.
func (*GrpcPayloadsLost) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDescGZIP(), []int{1}
}

// If included as a payload, this message flags the Status to be a possible
// outcome of a worker restart.
// URI:
// "type.googleapis.com/tensorflow.distributed_runtime.WorkerPossiblyRestarted"
type WorkerPossiblyRestarted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WorkerPossiblyRestarted) Reset() {
	*x = WorkerPossiblyRestarted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerPossiblyRestarted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerPossiblyRestarted) ProtoMessage() {}

func (x *WorkerPossiblyRestarted) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerPossiblyRestarted.ProtoReflect.Descriptor instead.
func (*WorkerPossiblyRestarted) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDescGZIP(), []int{2}
}

var File_tensorflow_core_protobuf_distributed_runtime_payloads_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xb3, 0x01,
	0x0a, 0x14, 0x47, 0x72, 0x70, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x64, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x72, 0x70, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x4c, 0x6f, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x42, 0x5a, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDescData = file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDesc
)

func file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDescData
}

var file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_goTypes = []interface{}{
	(*GrpcPayloadContainer)(nil),    // 0: tensorflow.distributed_runtime.GrpcPayloadContainer
	(*GrpcPayloadsLost)(nil),        // 1: tensorflow.distributed_runtime.GrpcPayloadsLost
	(*WorkerPossiblyRestarted)(nil), // 2: tensorflow.distributed_runtime.WorkerPossiblyRestarted
	nil,                             // 3: tensorflow.distributed_runtime.GrpcPayloadContainer.PayloadsEntry
}
var file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_depIdxs = []int32{
	3, // 0: tensorflow.distributed_runtime.GrpcPayloadContainer.payloads:type_name -> tensorflow.distributed_runtime.GrpcPayloadContainer.PayloadsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_init() }
func file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_init() {
	if File_tensorflow_core_protobuf_distributed_runtime_payloads_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcPayloadContainer); i {
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
		file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcPayloadsLost); i {
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
		file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerPossiblyRestarted); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_distributed_runtime_payloads_proto = out.File
	file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_rawDesc = nil
	file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_goTypes = nil
	file_tensorflow_core_protobuf_distributed_runtime_payloads_proto_depIdxs = nil
}
