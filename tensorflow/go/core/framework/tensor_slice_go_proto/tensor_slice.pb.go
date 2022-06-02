// Protocol buffer representing slices of a tensor

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.8.0
// source: tensorflow/core/framework/tensor_slice.proto

package tensor_slice_go_proto

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

// Can only be interpreted if you know the corresponding TensorShape.
type TensorSliceProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Extent of the slice in all tensor dimensions.
	//
	// Must have one entry for each of the dimension of the tensor that this
	// slice belongs to.  The order of sizes is the same as the order of
	// dimensions in the TensorShape.
	Extent []*TensorSliceProto_Extent `protobuf:"bytes,1,rep,name=extent,proto3" json:"extent,omitempty"`
}

func (x *TensorSliceProto) Reset() {
	*x = TensorSliceProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_tensor_slice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorSliceProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorSliceProto) ProtoMessage() {}

func (x *TensorSliceProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_tensor_slice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorSliceProto.ProtoReflect.Descriptor instead.
func (*TensorSliceProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_tensor_slice_proto_rawDescGZIP(), []int{0}
}

func (x *TensorSliceProto) GetExtent() []*TensorSliceProto_Extent {
	if x != nil {
		return x.Extent
	}
	return nil
}

// Extent of the slice in one dimension.
type TensorSliceProto_Extent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start index of the slice, starting at 0.
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// Length of the slice: if the length is missing or -1 we will
	// interpret this as "everything in this dimension".  We use
	// "oneof" to preserve information about whether the length is
	// present without changing the serialization format from the
	// prior proto2 version of this proto.
	//
	// Types that are assignable to HasLength:
	//	*TensorSliceProto_Extent_Length
	HasLength isTensorSliceProto_Extent_HasLength `protobuf_oneof:"has_length"`
}

func (x *TensorSliceProto_Extent) Reset() {
	*x = TensorSliceProto_Extent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_tensor_slice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorSliceProto_Extent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorSliceProto_Extent) ProtoMessage() {}

func (x *TensorSliceProto_Extent) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_tensor_slice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorSliceProto_Extent.ProtoReflect.Descriptor instead.
func (*TensorSliceProto_Extent) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_tensor_slice_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TensorSliceProto_Extent) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (m *TensorSliceProto_Extent) GetHasLength() isTensorSliceProto_Extent_HasLength {
	if m != nil {
		return m.HasLength
	}
	return nil
}

func (x *TensorSliceProto_Extent) GetLength() int64 {
	if x, ok := x.GetHasLength().(*TensorSliceProto_Extent_Length); ok {
		return x.Length
	}
	return 0
}

type isTensorSliceProto_Extent_HasLength interface {
	isTensorSliceProto_Extent_HasLength()
}

type TensorSliceProto_Extent_Length struct {
	Length int64 `protobuf:"varint,2,opt,name=length,proto3,oneof"`
}

func (*TensorSliceProto_Extent_Length) isTensorSliceProto_Extent_HasLength() {}

var File_tensorflow_core_framework_tensor_slice_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_tensor_slice_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x3b, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x46, 0x0a, 0x06,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x0c, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x42, 0x87, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x42, 0x11, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x6c, 0x69,
	0x63, 0x65, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_tensor_slice_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_tensor_slice_proto_rawDescData = file_tensorflow_core_framework_tensor_slice_proto_rawDesc
)

func file_tensorflow_core_framework_tensor_slice_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_tensor_slice_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_tensor_slice_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_tensor_slice_proto_rawDescData)
	})
	return file_tensorflow_core_framework_tensor_slice_proto_rawDescData
}

var file_tensorflow_core_framework_tensor_slice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_framework_tensor_slice_proto_goTypes = []interface{}{
	(*TensorSliceProto)(nil),        // 0: tensorflow.TensorSliceProto
	(*TensorSliceProto_Extent)(nil), // 1: tensorflow.TensorSliceProto.Extent
}
var file_tensorflow_core_framework_tensor_slice_proto_depIdxs = []int32{
	1, // 0: tensorflow.TensorSliceProto.extent:type_name -> tensorflow.TensorSliceProto.Extent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_tensor_slice_proto_init() }
func file_tensorflow_core_framework_tensor_slice_proto_init() {
	if File_tensorflow_core_framework_tensor_slice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_tensor_slice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorSliceProto); i {
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
		file_tensorflow_core_framework_tensor_slice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorSliceProto_Extent); i {
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
	file_tensorflow_core_framework_tensor_slice_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TensorSliceProto_Extent_Length)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_framework_tensor_slice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_tensor_slice_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_tensor_slice_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_tensor_slice_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_tensor_slice_proto = out.File
	file_tensorflow_core_framework_tensor_slice_proto_rawDesc = nil
	file_tensorflow_core_framework_tensor_slice_proto_goTypes = nil
	file_tensorflow_core_framework_tensor_slice_proto_depIdxs = nil
}
