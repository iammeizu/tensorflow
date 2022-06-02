// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.8.0
// source: tensorflow/core/framework/full_type.proto

package full_type_go_proto

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

// Experimental. Represents the complete type information of a TensorFlow value.
type FullTypeId int32

const (
	// The default represents an uninitialized values.
	FullTypeId_TFT_UNSET FullTypeId = 0
	// Type variables may serve as placeholder for any other type ID in type
	// templates.
	//
	// Examples:
	//   TFT_DATASET[TFT_VAR["T"]] is a Dataset returning a type indicated by "T".
	//   TFT_TENSOR[TFT_VAR["T"]] is a Tensor of n element type indicated by "T".
	//   TFT_TENSOR[TFT_VAR["T"]], TFT_TENSOR[TFT_VAR["T"]] are two tensors of
	//     identical element types.
	//   TFT_TENSOR[TFT_VAR["P"]], TFT_TENSOR[TFT_VAR["Q"]] are two tensors of
	//     independent element types.
	//
	FullTypeId_TFT_VAR FullTypeId = 1
	// Wildcard type. Describes a parameter of unknown type. In TensorFlow, that
	// can mean either a "Top" type (accepts any type), or a dynamically typed
	// object whose type is unknown in context.
	// Important: "unknown" does not necessarily mean undeterminable!
	FullTypeId_TFT_ANY FullTypeId = 2
	// The algebraic product type. This is an algebraic type that may be used just
	// for logical grouping. Not to confused with TFT_TUPLE which describes a
	// concrete object of several elements.
	//
	// Example:
	//   TFT_DATASET[TFT_PRODUCT[TFT_TENSOR[TFT_INT32], TFT_TENSOR[TFT_FLOAT64]]]
	//     is a Dataset producing two tensors, an integer one and a float one.
	//
	FullTypeId_TFT_PRODUCT FullTypeId = 3
	// Represents a named field, with the name stored in the attribute.
	//
	// Parametrization:
	//   TFT_NAMED[<type>]{<name>}
	//   * <type> is the type of the field
	//   * <name> is the field name, as string (thpugh can theoretically be an int
	//     as well)
	//
	// Example:
	//   TFT_RECORD[
	//     TFT_NAMED[TFT_TENSOR[TFT_INT32]]{'foo'},
	//     TFT_NAMED[TFT_TENSOR[TFT_FLOAT32]]{'bar'},
	//   ]
	//     is a structure with two fields, an int tensor "foo" and a float tensor
	//     "bar".
	FullTypeId_TFT_NAMED FullTypeId = 4
	// Template definition. Expands the variables by repeating a template as
	// arguments of container.
	//
	// Parametrization:
	//   TFT_FOR_EACH[<container_type>, <template>, <expansions>]
	//   * <container_type> is the type of the container that the template will be
	//     expanded into
	//   * <template> is any type definition that potentially contains type
	//     variables
	//   * <expansions> is a TFT_VAR and may include more types in the future
	//
	// Example:
	//   TFT_FOR_EACH[
	//         TFT_PRODUCT,
	//         TFT_TENSOR[TFT_VAR["t"]],
	//         TFT_VAR["t"]
	//     ]
	//     will substitute a T = TFT_INT32 to TFT_PRODUCT[TFT_TENSOR[TFT_INT32]]
	//     and a T = (TFT_INT32, TFT_INT64) to
	//     TFT_PRODUCT[TFT_TENSOR[TFT_INT32], TFT_TENSOR[TFT_INT64]].
	FullTypeId_TFT_FOR_EACH FullTypeId = 20
	// Callable types describe functions and ops.
	//
	// Parametrization:
	//   TFT_CALLABLE[<arg type>, <return type>]
	//   * <arg type> is the type of the arguments; TFT_PRODUCT represents
	//   multiple
	//     arguments.
	//   * <return type> is the return type; TFT_PRODUCT represents multiple
	//     return values (that means that callables returning multiple things
	//     don't necessarily return a single tuple).
	//
	// Example:
	//   TFT_CALLABLE[
	//     TFT_ANY,
	//     TFT_PRODUCT[TFT_TENSOR[TFT_INT32], TFT_TENSOR[TFT_FLOAT64]],
	//   ]
	//     is a callable with unspecified (for now) input arguments, and
	//     two return values of type tensor.
	//
	FullTypeId_TFT_CALLABLE FullTypeId = 100
	// The usual Tensor. This is a parametric type.
	//
	// Parametrization:
	//   TFT_TENSOR[<element type>, <shape type>]
	//   * <element type> is currently limited to one of the element types
	//     defined below.
	//   * <shape type> is not yet defined, and may only be TFT_UNKNOWN for now.
	//
	// A TFT_SHAPE type will be defined in the future.
	//
	// Example:
	//   TFT_TENSOR[TFT_INT32, TFT_UNKNOWN]
	//     is a Tensor of int32 element type and unknown shape.
	//
	// TODO(mdan): Define TFT_SHAPE and add more examples.
	FullTypeId_TFT_TENSOR FullTypeId = 1000
	// Array (or tensorflow::TensorList in the variant type registry).
	// Note: this is not to be confused with the deprecated `TensorArray*` ops
	// which are not supported by FullType.
	// This type represents a random-access list whose elements can be
	// described by a single type. Although immutable, Array is expected to
	// support efficient mutation semantics (i.e. element update) in the
	// user-facing API.
	// The element type may be generic or even TFT_ANY for a heterogenous list.
	//
	// Parametrization:
	//   TFT_ARRAY[<element type>]
	//   * <element type> may be any concrete type.
	//
	// Examples:
	//   TFT_ARRAY[TFT_TENSOR[TFT_INT32]] is a TensorArray holding int32 Tensors
	//     of any shape.
	//   TFT_ARRAY[TFT_TENSOR[TFT_UNKNOWN]] is a TensorArray holding Tensors of
	//     mixed element types.
	//   TFT_ARRAY[TFT_UNKNOWN] is a TensorArray holding any element type.
	//   TFT_ARRAY[] is equivalent to TFT_ARRAY[TFT_UNKNOWN].
	//   TFT_ARRAY[TFT_ARRAY[]] is an array or arrays (of unknown types).
	FullTypeId_TFT_ARRAY FullTypeId = 1001
	// Optional (or tensorflow::OptionalVariant in the variant type registry).
	// This type represents a value that may either hold an element of a single
	// specified type, or nothing at all.
	//
	// Parametrization:
	//   TFT_OPTIONAL[<element type>]
	//   * <element type> may be any concrete type.
	//
	// Examples:
	//   TFT_OPTIONAL[TFT_TENSOR[TFT_INT32]] is an Optional holding an int32
	//     Tensor of any shape.
	FullTypeId_TFT_OPTIONAL FullTypeId = 1002
	// Literal types describe compile-time constant values.
	// Literal types may also participate in dependent types.
	//
	// Parametrization:
	//   TFT_LITERAL[<value type>]{<value>}
	//   * <value type> may be any concrete type compatible that can hold <value>
	//   * <value> is the type's attribute, and holds the actual literal value
	//
	// Examples:
	//   TFT_LITERAL[TFT_INT32]{1} is the compile-time constant 1.
	FullTypeId_TFT_LITERAL FullTypeId = 1003
	// The bool element type.
	// TODO(mdan): Quantized types, legacy representations (e.g. ref)
	FullTypeId_TFT_BOOL FullTypeId = 200
	// Integer element types.
	FullTypeId_TFT_UINT8  FullTypeId = 201
	FullTypeId_TFT_UINT16 FullTypeId = 202
	FullTypeId_TFT_UINT32 FullTypeId = 203
	FullTypeId_TFT_UINT64 FullTypeId = 204
	FullTypeId_TFT_INT8   FullTypeId = 205
	FullTypeId_TFT_INT16  FullTypeId = 206
	FullTypeId_TFT_INT32  FullTypeId = 207
	FullTypeId_TFT_INT64  FullTypeId = 208
	// Floating-point element types.
	FullTypeId_TFT_HALF     FullTypeId = 209
	FullTypeId_TFT_FLOAT    FullTypeId = 210
	FullTypeId_TFT_DOUBLE   FullTypeId = 211
	FullTypeId_TFT_BFLOAT16 FullTypeId = 215
	// Complex element types.
	// TODO(mdan): Represent as TFT_COMPLEX[TFT_DOUBLE] instead?
	FullTypeId_TFT_COMPLEX64  FullTypeId = 212
	FullTypeId_TFT_COMPLEX128 FullTypeId = 213
	// The string element type.
	FullTypeId_TFT_STRING FullTypeId = 214
	// Datasets created by tf.data ops and APIs. Datasets have generator/iterable
	// semantics, that is, one can construct an iterator from them. Like
	// Array, they are considered to return elements that can be described
	// by a single type. Unlike Array, they do not support random access or
	// mutation, and can potentially produce an infinite number of elements.
	// A datasets can produce logical structures (e.g. multiple elements). This
	// is expressed using TFT_PRODUCT.
	//
	//
	// Parametrization: TFT_ARRAY[<element type>].
	//   * <element type> may be a concrete type or a type symbol. It represents
	//     the data type of the elements produced by the dataset.
	//
	// Examples:
	//   TFT_DATSET[TFT_TENSOR[TFT_INT32]] is a Dataset producing single int32
	//     Tensors of unknown shape.
	//   TFT_DATSET[TFT_PRODUCT[TFT_TENSOR[TFT_INT32], TFT_TENSOR[TFT_FLOAT32]] is
	//     a Dataset producing pairs of Tensors, one integer and one float.
	// Note: The high ID number is to prepare for the eventuality that Datasets
	// will be supported by user types in the future.
	FullTypeId_TFT_DATASET FullTypeId = 10102
	// A ragged tensor created by tf.ragged ops and APIs.
	//
	// Parametrization: TFT_RAGGED[<element_type>].
	FullTypeId_TFT_RAGGED FullTypeId = 10103
	// A mutex lock tensor, produced by tf.raw_ops.MutexLock.
	// Unlike strict execution models, where ownership of a lock is denoted by
	// "running after the lock has been acquired", in non-strict mode, lock
	// ownership is in the true sense: "the op argument representing the lock is
	// available".
	// Mutex locks are the dynamic counterpart of control dependencies.
	// TODO(mdan): Properly document this thing.
	//
	// Parametrization: TFT_MUTEX_LOCK[].
	FullTypeId_TFT_MUTEX_LOCK FullTypeId = 10202
	// The equivalent of a Tensor with DT_VARIANT dtype, kept here to simplify
	// translation. This type should not normally appear after type inference.
	// Note that LEGACY_VARIANT != ANY: TENSOR[INT32] is a subtype of ANY, but is
	// not a subtype of LEGACY_VARIANT.
	FullTypeId_TFT_LEGACY_VARIANT FullTypeId = 10203
)

// Enum value maps for FullTypeId.
var (
	FullTypeId_name = map[int32]string{
		0:     "TFT_UNSET",
		1:     "TFT_VAR",
		2:     "TFT_ANY",
		3:     "TFT_PRODUCT",
		4:     "TFT_NAMED",
		20:    "TFT_FOR_EACH",
		100:   "TFT_CALLABLE",
		1000:  "TFT_TENSOR",
		1001:  "TFT_ARRAY",
		1002:  "TFT_OPTIONAL",
		1003:  "TFT_LITERAL",
		200:   "TFT_BOOL",
		201:   "TFT_UINT8",
		202:   "TFT_UINT16",
		203:   "TFT_UINT32",
		204:   "TFT_UINT64",
		205:   "TFT_INT8",
		206:   "TFT_INT16",
		207:   "TFT_INT32",
		208:   "TFT_INT64",
		209:   "TFT_HALF",
		210:   "TFT_FLOAT",
		211:   "TFT_DOUBLE",
		215:   "TFT_BFLOAT16",
		212:   "TFT_COMPLEX64",
		213:   "TFT_COMPLEX128",
		214:   "TFT_STRING",
		10102: "TFT_DATASET",
		10103: "TFT_RAGGED",
		10202: "TFT_MUTEX_LOCK",
		10203: "TFT_LEGACY_VARIANT",
	}
	FullTypeId_value = map[string]int32{
		"TFT_UNSET":          0,
		"TFT_VAR":            1,
		"TFT_ANY":            2,
		"TFT_PRODUCT":        3,
		"TFT_NAMED":          4,
		"TFT_FOR_EACH":       20,
		"TFT_CALLABLE":       100,
		"TFT_TENSOR":         1000,
		"TFT_ARRAY":          1001,
		"TFT_OPTIONAL":       1002,
		"TFT_LITERAL":        1003,
		"TFT_BOOL":           200,
		"TFT_UINT8":          201,
		"TFT_UINT16":         202,
		"TFT_UINT32":         203,
		"TFT_UINT64":         204,
		"TFT_INT8":           205,
		"TFT_INT16":          206,
		"TFT_INT32":          207,
		"TFT_INT64":          208,
		"TFT_HALF":           209,
		"TFT_FLOAT":          210,
		"TFT_DOUBLE":         211,
		"TFT_BFLOAT16":       215,
		"TFT_COMPLEX64":      212,
		"TFT_COMPLEX128":     213,
		"TFT_STRING":         214,
		"TFT_DATASET":        10102,
		"TFT_RAGGED":         10103,
		"TFT_MUTEX_LOCK":     10202,
		"TFT_LEGACY_VARIANT": 10203,
	}
)

func (x FullTypeId) Enum() *FullTypeId {
	p := new(FullTypeId)
	*p = x
	return p
}

func (x FullTypeId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FullTypeId) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_framework_full_type_proto_enumTypes[0].Descriptor()
}

func (FullTypeId) Type() protoreflect.EnumType {
	return &file_tensorflow_core_framework_full_type_proto_enumTypes[0]
}

func (x FullTypeId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FullTypeId.Descriptor instead.
func (FullTypeId) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_full_type_proto_rawDescGZIP(), []int{0}
}

// Highly experimental and very likely to change.
// This encoding uses tags instead of dedicated messages for regularity. In
// particular the encoding imposes no restrictions on what the parameters of any
// type should be, which in particular needs to be true for type symbols.
type FullTypeDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The principal type represented by this object. This may be a concrete type
	// (Tensor, Dataset) a type variable (used for dependent types) a type
	// symbol (Any, Union). See FullTypeId for details.
	TypeId FullTypeId     `protobuf:"varint,1,opt,name=type_id,json=typeId,proto3,enum=tensorflow.FullTypeId" json:"type_id,omitempty"`
	Args   []*FullTypeDef `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	// Literal values of this type object, if the the type admits one.
	// For example, a type variable admits a string attribute - its name.
	// Shape-related types may admit int attributes - their static shape values.
	// Fields for more data types to be added as needed.
	//
	// Types that are assignable to Attr:
	//	*FullTypeDef_S
	//	*FullTypeDef_I
	Attr isFullTypeDef_Attr `protobuf_oneof:"attr"`
}

func (x *FullTypeDef) Reset() {
	*x = FullTypeDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_full_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullTypeDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullTypeDef) ProtoMessage() {}

func (x *FullTypeDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_full_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullTypeDef.ProtoReflect.Descriptor instead.
func (*FullTypeDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_full_type_proto_rawDescGZIP(), []int{0}
}

func (x *FullTypeDef) GetTypeId() FullTypeId {
	if x != nil {
		return x.TypeId
	}
	return FullTypeId_TFT_UNSET
}

func (x *FullTypeDef) GetArgs() []*FullTypeDef {
	if x != nil {
		return x.Args
	}
	return nil
}

func (m *FullTypeDef) GetAttr() isFullTypeDef_Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (x *FullTypeDef) GetS() string {
	if x, ok := x.GetAttr().(*FullTypeDef_S); ok {
		return x.S
	}
	return ""
}

func (x *FullTypeDef) GetI() int64 {
	if x, ok := x.GetAttr().(*FullTypeDef_I); ok {
		return x.I
	}
	return 0
}

type isFullTypeDef_Attr interface {
	isFullTypeDef_Attr()
}

type FullTypeDef_S struct {
	S string `protobuf:"bytes,3,opt,name=s,proto3,oneof"`
}

type FullTypeDef_I struct {
	I int64 `protobuf:"varint,4,opt,name=i,proto3,oneof"` // TODO(mdan): list/tensor, map? Need to reconcile with TFT_RECORD, etc.
}

func (*FullTypeDef_S) isFullTypeDef_Attr() {}

func (*FullTypeDef_I) isFullTypeDef_Attr() {}

var File_tensorflow_core_framework_full_type_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_full_type_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x93, 0x01, 0x0a, 0x0b, 0x46, 0x75, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x66, 0x12, 0x2f, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x66, 0x52,
	0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x01, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x01, 0x73, 0x12, 0x0e, 0x0a, 0x01, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x01, 0x69, 0x42, 0x06, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x2a, 0x9e, 0x04,
	0x0a, 0x0a, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x0d, 0x0a, 0x09,
	0x54, 0x46, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54,
	0x46, 0x54, 0x5f, 0x56, 0x41, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x46, 0x54, 0x5f,
	0x41, 0x4e, 0x59, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x46, 0x54, 0x5f, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x46, 0x54, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x44, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x46, 0x54, 0x5f, 0x46, 0x4f, 0x52,
	0x5f, 0x45, 0x41, 0x43, 0x48, 0x10, 0x14, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x46, 0x54, 0x5f, 0x43,
	0x41, 0x4c, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x64, 0x12, 0x0f, 0x0a, 0x0a, 0x54, 0x46, 0x54,
	0x5f, 0x54, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x10, 0xe8, 0x07, 0x12, 0x0e, 0x0a, 0x09, 0x54, 0x46,
	0x54, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0xe9, 0x07, 0x12, 0x11, 0x0a, 0x0c, 0x54, 0x46,
	0x54, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0xea, 0x07, 0x12, 0x10, 0x0a,
	0x0b, 0x54, 0x46, 0x54, 0x5f, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x4c, 0x10, 0xeb, 0x07, 0x12,
	0x0d, 0x0a, 0x08, 0x54, 0x46, 0x54, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0xc8, 0x01, 0x12, 0x0e,
	0x0a, 0x09, 0x54, 0x46, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0xc9, 0x01, 0x12, 0x0f,
	0x0a, 0x0a, 0x54, 0x46, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0xca, 0x01, 0x12,
	0x0f, 0x0a, 0x0a, 0x54, 0x46, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0xcb, 0x01,
	0x12, 0x0f, 0x0a, 0x0a, 0x54, 0x46, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0xcc,
	0x01, 0x12, 0x0d, 0x0a, 0x08, 0x54, 0x46, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x38, 0x10, 0xcd, 0x01,
	0x12, 0x0e, 0x0a, 0x09, 0x54, 0x46, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0xce, 0x01,
	0x12, 0x0e, 0x0a, 0x09, 0x54, 0x46, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0xcf, 0x01,
	0x12, 0x0e, 0x0a, 0x09, 0x54, 0x46, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0xd0, 0x01,
	0x12, 0x0d, 0x0a, 0x08, 0x54, 0x46, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0xd1, 0x01, 0x12,
	0x0e, 0x0a, 0x09, 0x54, 0x46, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0xd2, 0x01, 0x12,
	0x0f, 0x0a, 0x0a, 0x54, 0x46, 0x54, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0xd3, 0x01,
	0x12, 0x11, 0x0a, 0x0c, 0x54, 0x46, 0x54, 0x5f, 0x42, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x31, 0x36,
	0x10, 0xd7, 0x01, 0x12, 0x12, 0x0a, 0x0d, 0x54, 0x46, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x58, 0x36, 0x34, 0x10, 0xd4, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x54, 0x46, 0x54, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x31, 0x32, 0x38, 0x10, 0xd5, 0x01, 0x12, 0x0f, 0x0a, 0x0a,
	0x54, 0x46, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0xd6, 0x01, 0x12, 0x10, 0x0a,
	0x0b, 0x54, 0x46, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x53, 0x45, 0x54, 0x10, 0xf6, 0x4e, 0x12,
	0x0f, 0x0a, 0x0a, 0x54, 0x46, 0x54, 0x5f, 0x52, 0x41, 0x47, 0x47, 0x45, 0x44, 0x10, 0xf7, 0x4e,
	0x12, 0x13, 0x0a, 0x0e, 0x54, 0x46, 0x54, 0x5f, 0x4d, 0x55, 0x54, 0x45, 0x58, 0x5f, 0x4c, 0x4f,
	0x43, 0x4b, 0x10, 0xda, 0x4f, 0x12, 0x17, 0x0a, 0x12, 0x54, 0x46, 0x54, 0x5f, 0x4c, 0x45, 0x47,
	0x41, 0x43, 0x59, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x4e, 0x54, 0x10, 0xdb, 0x4f, 0x42, 0x81,
	0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0e, 0x46, 0x75, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x50, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_full_type_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_full_type_proto_rawDescData = file_tensorflow_core_framework_full_type_proto_rawDesc
)

func file_tensorflow_core_framework_full_type_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_full_type_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_full_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_full_type_proto_rawDescData)
	})
	return file_tensorflow_core_framework_full_type_proto_rawDescData
}

var file_tensorflow_core_framework_full_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_framework_full_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_framework_full_type_proto_goTypes = []interface{}{
	(FullTypeId)(0),     // 0: tensorflow.FullTypeId
	(*FullTypeDef)(nil), // 1: tensorflow.FullTypeDef
}
var file_tensorflow_core_framework_full_type_proto_depIdxs = []int32{
	0, // 0: tensorflow.FullTypeDef.type_id:type_name -> tensorflow.FullTypeId
	1, // 1: tensorflow.FullTypeDef.args:type_name -> tensorflow.FullTypeDef
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_full_type_proto_init() }
func file_tensorflow_core_framework_full_type_proto_init() {
	if File_tensorflow_core_framework_full_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_full_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullTypeDef); i {
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
	file_tensorflow_core_framework_full_type_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FullTypeDef_S)(nil),
		(*FullTypeDef_I)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_framework_full_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_full_type_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_full_type_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_framework_full_type_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_framework_full_type_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_full_type_proto = out.File
	file_tensorflow_core_framework_full_type_proto_rawDesc = nil
	file_tensorflow_core_framework_full_type_proto_goTypes = nil
	file_tensorflow_core_framework_full_type_proto_depIdxs = nil
}
