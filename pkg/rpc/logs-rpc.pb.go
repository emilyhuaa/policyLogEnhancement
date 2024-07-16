// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: logs-rpc.proto

package rpc

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

// Metadata represents information about a pod/srv.
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the pod/srv.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The namespace of the pod/srv.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_logs_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_logs_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// MetadataInfoCacheEntry represents an entry in the metadataInfoCache.
type MetadataCacheEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IP address of the pod/srv (key).
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// The metadata information (value).
	Metadata *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *MetadataCacheEntry) Reset() {
	*x = MetadataCacheEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataCacheEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataCacheEntry) ProtoMessage() {}

func (x *MetadataCacheEntry) ProtoReflect() protoreflect.Message {
	mi := &file_logs_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataCacheEntry.ProtoReflect.Descriptor instead.
func (*MetadataCacheEntry) Descriptor() ([]byte, []int) {
	return file_logs_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataCacheEntry) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *MetadataCacheEntry) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// GetCacheRequest represents a request to retrieve the metadataInfoCache.
type GetCacheRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCacheRequest) Reset() {
	*x = GetCacheRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCacheRequest) ProtoMessage() {}

func (x *GetCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logs_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCacheRequest.ProtoReflect.Descriptor instead.
func (*GetCacheRequest) Descriptor() ([]byte, []int) {
	return file_logs_rpc_proto_rawDescGZIP(), []int{2}
}

// GetCacheReply represents the response containing the metadataInfoCache.
type GetCacheReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of metadataCache entries.
	Entries []*MetadataCacheEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *GetCacheReply) Reset() {
	*x = GetCacheReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logs_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCacheReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCacheReply) ProtoMessage() {}

func (x *GetCacheReply) ProtoReflect() protoreflect.Message {
	mi := &file_logs_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCacheReply.ProtoReflect.Descriptor instead.
func (*GetCacheReply) Descriptor() ([]byte, []int) {
	return file_logs_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *GetCacheReply) GetEntries() []*MetadataCacheEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_logs_rpc_proto protoreflect.FileDescriptor

var file_logs_rpc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x73, 0x2d, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x3c, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x22, 0x4f, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x32, 0x4e, 0x0a, 0x0c, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12,
	0x14, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x69, 0x6c, 0x79, 0x68,
	0x75, 0x61, 0x61, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x45, 0x6e,
	0x68, 0x61, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logs_rpc_proto_rawDescOnce sync.Once
	file_logs_rpc_proto_rawDescData = file_logs_rpc_proto_rawDesc
)

func file_logs_rpc_proto_rawDescGZIP() []byte {
	file_logs_rpc_proto_rawDescOnce.Do(func() {
		file_logs_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_logs_rpc_proto_rawDescData)
	})
	return file_logs_rpc_proto_rawDescData
}

var file_logs_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_logs_rpc_proto_goTypes = []any{
	(*Metadata)(nil),           // 0: rpc.Metadata
	(*MetadataCacheEntry)(nil), // 1: rpc.MetadataCacheEntry
	(*GetCacheRequest)(nil),    // 2: rpc.GetCacheRequest
	(*GetCacheReply)(nil),      // 3: rpc.GetCacheReply
}
var file_logs_rpc_proto_depIdxs = []int32{
	0, // 0: rpc.MetadataCacheEntry.metadata:type_name -> rpc.Metadata
	1, // 1: rpc.GetCacheReply.entries:type_name -> rpc.MetadataCacheEntry
	2, // 2: rpc.CacheService.GetMetadataCache:input_type -> rpc.GetCacheRequest
	3, // 3: rpc.CacheService.GetMetadataCache:output_type -> rpc.GetCacheReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_logs_rpc_proto_init() }
func file_logs_rpc_proto_init() {
	if File_logs_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logs_rpc_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Metadata); i {
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
		file_logs_rpc_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MetadataCacheEntry); i {
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
		file_logs_rpc_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetCacheRequest); i {
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
		file_logs_rpc_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetCacheReply); i {
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
			RawDescriptor: file_logs_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logs_rpc_proto_goTypes,
		DependencyIndexes: file_logs_rpc_proto_depIdxs,
		MessageInfos:      file_logs_rpc_proto_msgTypes,
	}.Build()
	File_logs_rpc_proto = out.File
	file_logs_rpc_proto_rawDesc = nil
	file_logs_rpc_proto_goTypes = nil
	file_logs_rpc_proto_depIdxs = nil
}