// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.8.0
// source: tensorflow/core/framework/allocation_description.proto

package allocation_description_go_proto

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

type AllocationDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of bytes requested
	RequestedBytes int64 `protobuf:"varint,1,opt,name=requested_bytes,json=requestedBytes,proto3" json:"requested_bytes,omitempty"`
	// Total number of bytes allocated if known
	AllocatedBytes int64 `protobuf:"varint,2,opt,name=allocated_bytes,json=allocatedBytes,proto3" json:"allocated_bytes,omitempty"`
	// Name of the allocator used
	AllocatorName string `protobuf:"bytes,3,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// Identifier of the allocated buffer if known
	AllocationId int64 `protobuf:"varint,4,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Set if this tensor only has one remaining reference
	HasSingleReference bool `protobuf:"varint,5,opt,name=has_single_reference,json=hasSingleReference,proto3" json:"has_single_reference,omitempty"`
	// Address of the allocation.
	Ptr uint64 `protobuf:"varint,6,opt,name=ptr,proto3" json:"ptr,omitempty"`
}

func (x *AllocationDescription) Reset() {
	*x = AllocationDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_allocation_description_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocationDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationDescription) ProtoMessage() {}

func (x *AllocationDescription) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_allocation_description_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationDescription.ProtoReflect.Descriptor instead.
func (*AllocationDescription) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_allocation_description_proto_rawDescGZIP(), []int{0}
}

func (x *AllocationDescription) GetRequestedBytes() int64 {
	if x != nil {
		return x.RequestedBytes
	}
	return 0
}

func (x *AllocationDescription) GetAllocatedBytes() int64 {
	if x != nil {
		return x.AllocatedBytes
	}
	return 0
}

func (x *AllocationDescription) GetAllocatorName() string {
	if x != nil {
		return x.AllocatorName
	}
	return ""
}

func (x *AllocationDescription) GetAllocationId() int64 {
	if x != nil {
		return x.AllocationId
	}
	return 0
}

func (x *AllocationDescription) GetHasSingleReference() bool {
	if x != nil {
		return x.HasSingleReference
	}
	return false
}

func (x *AllocationDescription) GetPtr() uint64 {
	if x != nil {
		return x.Ptr
	}
	return 0
}

var File_tensorflow_core_framework_allocation_description_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_allocation_description_proto_rawDesc = []byte{
	0x0a, 0x36, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x22, 0xf9, 0x01, 0x0a, 0x15, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14,
	0x68, 0x61, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x68, 0x61, 0x73, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x74, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x74, 0x72,
	0x42, 0x9b, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x1b, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_allocation_description_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_allocation_description_proto_rawDescData = file_tensorflow_core_framework_allocation_description_proto_rawDesc
)

func file_tensorflow_core_framework_allocation_description_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_allocation_description_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_allocation_description_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_allocation_description_proto_rawDescData)
	})
	return file_tensorflow_core_framework_allocation_description_proto_rawDescData
}

var file_tensorflow_core_framework_allocation_description_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_framework_allocation_description_proto_goTypes = []interface{}{
	(*AllocationDescription)(nil), // 0: tensorflow.AllocationDescription
}
var file_tensorflow_core_framework_allocation_description_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_allocation_description_proto_init() }
func file_tensorflow_core_framework_allocation_description_proto_init() {
	if File_tensorflow_core_framework_allocation_description_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_allocation_description_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocationDescription); i {
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
			RawDescriptor: file_tensorflow_core_framework_allocation_description_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_allocation_description_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_allocation_description_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_allocation_description_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_allocation_description_proto = out.File
	file_tensorflow_core_framework_allocation_description_proto_rawDesc = nil
	file_tensorflow_core_framework_allocation_description_proto_goTypes = nil
	file_tensorflow_core_framework_allocation_description_proto_depIdxs = nil
}
