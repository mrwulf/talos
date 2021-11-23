// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: inspect/inspect.proto

package inspect

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	common "github.com/talos-systems/talos/pkg/machinery/api/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DependencyEdgeType int32

const (
	DependencyEdgeType_OUTPUT_EXCLUSIVE    DependencyEdgeType = 0
	DependencyEdgeType_OUTPUT_SHARED       DependencyEdgeType = 3
	DependencyEdgeType_INPUT_STRONG        DependencyEdgeType = 1
	DependencyEdgeType_INPUT_WEAK          DependencyEdgeType = 2
	DependencyEdgeType_INPUT_DESTROY_READY DependencyEdgeType = 4
)

// Enum value maps for DependencyEdgeType.
var (
	DependencyEdgeType_name = map[int32]string{
		0: "OUTPUT_EXCLUSIVE",
		3: "OUTPUT_SHARED",
		1: "INPUT_STRONG",
		2: "INPUT_WEAK",
		4: "INPUT_DESTROY_READY",
	}
	DependencyEdgeType_value = map[string]int32{
		"OUTPUT_EXCLUSIVE":    0,
		"OUTPUT_SHARED":       3,
		"INPUT_STRONG":        1,
		"INPUT_WEAK":          2,
		"INPUT_DESTROY_READY": 4,
	}
)

func (x DependencyEdgeType) Enum() *DependencyEdgeType {
	p := new(DependencyEdgeType)
	*p = x
	return p
}

func (x DependencyEdgeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DependencyEdgeType) Descriptor() protoreflect.EnumDescriptor {
	return file_inspect_inspect_proto_enumTypes[0].Descriptor()
}

func (DependencyEdgeType) Type() protoreflect.EnumType {
	return &file_inspect_inspect_proto_enumTypes[0]
}

func (x DependencyEdgeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DependencyEdgeType.Descriptor instead.
func (DependencyEdgeType) EnumDescriptor() ([]byte, []int) {
	return file_inspect_inspect_proto_rawDescGZIP(), []int{0}
}

// The ControllerRuntimeDependency message contains the graph of controller-resource dependencies.
type ControllerRuntimeDependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *common.Metadata            `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Edges    []*ControllerDependencyEdge `protobuf:"bytes,2,rep,name=edges,proto3" json:"edges,omitempty"`
}

func (x *ControllerRuntimeDependency) Reset() {
	*x = ControllerRuntimeDependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspect_inspect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerRuntimeDependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerRuntimeDependency) ProtoMessage() {}

func (x *ControllerRuntimeDependency) ProtoReflect() protoreflect.Message {
	mi := &file_inspect_inspect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerRuntimeDependency.ProtoReflect.Descriptor instead.
func (*ControllerRuntimeDependency) Descriptor() ([]byte, []int) {
	return file_inspect_inspect_proto_rawDescGZIP(), []int{0}
}

func (x *ControllerRuntimeDependency) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ControllerRuntimeDependency) GetEdges() []*ControllerDependencyEdge {
	if x != nil {
		return x.Edges
	}
	return nil
}

type ControllerRuntimeDependenciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*ControllerRuntimeDependency `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ControllerRuntimeDependenciesResponse) Reset() {
	*x = ControllerRuntimeDependenciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspect_inspect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerRuntimeDependenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerRuntimeDependenciesResponse) ProtoMessage() {}

func (x *ControllerRuntimeDependenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inspect_inspect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerRuntimeDependenciesResponse.ProtoReflect.Descriptor instead.
func (*ControllerRuntimeDependenciesResponse) Descriptor() ([]byte, []int) {
	return file_inspect_inspect_proto_rawDescGZIP(), []int{1}
}

func (x *ControllerRuntimeDependenciesResponse) GetMessages() []*ControllerRuntimeDependency {
	if x != nil {
		return x.Messages
	}
	return nil
}

type ControllerDependencyEdge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControllerName    string             `protobuf:"bytes,1,opt,name=controller_name,json=controllerName,proto3" json:"controller_name,omitempty"`
	EdgeType          DependencyEdgeType `protobuf:"varint,2,opt,name=edge_type,json=edgeType,proto3,enum=inspect.DependencyEdgeType" json:"edge_type,omitempty"`
	ResourceNamespace string             `protobuf:"bytes,3,opt,name=resource_namespace,json=resourceNamespace,proto3" json:"resource_namespace,omitempty"`
	ResourceType      string             `protobuf:"bytes,4,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceId        string             `protobuf:"bytes,5,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ControllerDependencyEdge) Reset() {
	*x = ControllerDependencyEdge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inspect_inspect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerDependencyEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerDependencyEdge) ProtoMessage() {}

func (x *ControllerDependencyEdge) ProtoReflect() protoreflect.Message {
	mi := &file_inspect_inspect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerDependencyEdge.ProtoReflect.Descriptor instead.
func (*ControllerDependencyEdge) Descriptor() ([]byte, []int) {
	return file_inspect_inspect_proto_rawDescGZIP(), []int{2}
}

func (x *ControllerDependencyEdge) GetControllerName() string {
	if x != nil {
		return x.ControllerName
	}
	return ""
}

func (x *ControllerDependencyEdge) GetEdgeType() DependencyEdgeType {
	if x != nil {
		return x.EdgeType
	}
	return DependencyEdgeType_OUTPUT_EXCLUSIVE
}

func (x *ControllerDependencyEdge) GetResourceNamespace() string {
	if x != nil {
		return x.ResourceNamespace
	}
	return ""
}

func (x *ControllerDependencyEdge) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *ControllerDependencyEdge) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

var File_inspect_inspect_proto protoreflect.FileDescriptor

var file_inspect_inspect_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x37, 0x0a, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x45, 0x64,
	0x67, 0x65, 0x52, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x22, 0x69, 0x0a, 0x25, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x22, 0xf2, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x45, 0x64, 0x67,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x65, 0x64,
	0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x45, 0x64, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x65, 0x64, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x2a, 0x78, 0x0a, 0x12, 0x44, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x45, 0x64, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x5f, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x53,
	0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x5f,
	0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x50, 0x55,
	0x54, 0x5f, 0x53, 0x54, 0x52, 0x4f, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e,
	0x50, 0x55, 0x54, 0x5f, 0x57, 0x45, 0x41, 0x4b, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e,
	0x50, 0x55, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x52, 0x4f, 0x59, 0x5f, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x10, 0x04, 0x32, 0x79, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2e,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6c,
	0x6f, 0x73, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_inspect_inspect_proto_rawDescOnce sync.Once
	file_inspect_inspect_proto_rawDescData = file_inspect_inspect_proto_rawDesc
)

func file_inspect_inspect_proto_rawDescGZIP() []byte {
	file_inspect_inspect_proto_rawDescOnce.Do(func() {
		file_inspect_inspect_proto_rawDescData = protoimpl.X.CompressGZIP(file_inspect_inspect_proto_rawDescData)
	})
	return file_inspect_inspect_proto_rawDescData
}

var (
	file_inspect_inspect_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_inspect_inspect_proto_msgTypes  = make([]protoimpl.MessageInfo, 3)
	file_inspect_inspect_proto_goTypes   = []interface{}{
		(DependencyEdgeType)(0),                       // 0: inspect.DependencyEdgeType
		(*ControllerRuntimeDependency)(nil),           // 1: inspect.ControllerRuntimeDependency
		(*ControllerRuntimeDependenciesResponse)(nil), // 2: inspect.ControllerRuntimeDependenciesResponse
		(*ControllerDependencyEdge)(nil),              // 3: inspect.ControllerDependencyEdge
		(*common.Metadata)(nil),                       // 4: common.Metadata
		(*emptypb.Empty)(nil),                         // 5: google.protobuf.Empty
	}
)

var file_inspect_inspect_proto_depIdxs = []int32{
	4, // 0: inspect.ControllerRuntimeDependency.metadata:type_name -> common.Metadata
	3, // 1: inspect.ControllerRuntimeDependency.edges:type_name -> inspect.ControllerDependencyEdge
	1, // 2: inspect.ControllerRuntimeDependenciesResponse.messages:type_name -> inspect.ControllerRuntimeDependency
	0, // 3: inspect.ControllerDependencyEdge.edge_type:type_name -> inspect.DependencyEdgeType
	5, // 4: inspect.InspectService.ControllerRuntimeDependencies:input_type -> google.protobuf.Empty
	2, // 5: inspect.InspectService.ControllerRuntimeDependencies:output_type -> inspect.ControllerRuntimeDependenciesResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_inspect_inspect_proto_init() }
func file_inspect_inspect_proto_init() {
	if File_inspect_inspect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inspect_inspect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerRuntimeDependency); i {
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
		file_inspect_inspect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerRuntimeDependenciesResponse); i {
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
		file_inspect_inspect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerDependencyEdge); i {
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
			RawDescriptor: file_inspect_inspect_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inspect_inspect_proto_goTypes,
		DependencyIndexes: file_inspect_inspect_proto_depIdxs,
		EnumInfos:         file_inspect_inspect_proto_enumTypes,
		MessageInfos:      file_inspect_inspect_proto_msgTypes,
	}.Build()
	File_inspect_inspect_proto = out.File
	file_inspect_inspect_proto_rawDesc = nil
	file_inspect_inspect_proto_goTypes = nil
	file_inspect_inspect_proto_depIdxs = nil
}
