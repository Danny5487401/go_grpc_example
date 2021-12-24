// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: 07_grpc_error/proto/err_detail.proto

package grpcErrProtobuf

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

type ErrDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"` //
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // 这里可以写一些业务信息:为用户可读的信息，可作为用户提示内容
}

func (x *ErrDetail) Reset() {
	*x = ErrDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file__07_grpc_error_proto_err_detail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrDetail) ProtoMessage() {}

func (x *ErrDetail) ProtoReflect() protoreflect.Message {
	mi := &file__07_grpc_error_proto_err_detail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrDetail.ProtoReflect.Descriptor instead.
func (*ErrDetail) Descriptor() ([]byte, []int) {
	return file__07_grpc_error_proto_err_detail_proto_rawDescGZIP(), []int{0}
}

func (x *ErrDetail) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ErrDetail) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File__07_grpc_error_proto_err_detail_proto protoreflect.FileDescriptor

var file__07_grpc_error_proto_err_detail_proto_rawDesc = []byte{
	0x0a, 0x24, 0x30, 0x37, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x25, 0x5a, 0x23, 0x30, 0x37, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x67,
	0x72, 0x70, 0x63, 0x45, 0x72, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file__07_grpc_error_proto_err_detail_proto_rawDescOnce sync.Once
	file__07_grpc_error_proto_err_detail_proto_rawDescData = file__07_grpc_error_proto_err_detail_proto_rawDesc
)

func file__07_grpc_error_proto_err_detail_proto_rawDescGZIP() []byte {
	file__07_grpc_error_proto_err_detail_proto_rawDescOnce.Do(func() {
		file__07_grpc_error_proto_err_detail_proto_rawDescData = protoimpl.X.CompressGZIP(file__07_grpc_error_proto_err_detail_proto_rawDescData)
	})
	return file__07_grpc_error_proto_err_detail_proto_rawDescData
}

var file__07_grpc_error_proto_err_detail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file__07_grpc_error_proto_err_detail_proto_goTypes = []interface{}{
	(*ErrDetail)(nil), // 0: errDetail
}
var file__07_grpc_error_proto_err_detail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file__07_grpc_error_proto_err_detail_proto_init() }
func file__07_grpc_error_proto_err_detail_proto_init() {
	if File__07_grpc_error_proto_err_detail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__07_grpc_error_proto_err_detail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrDetail); i {
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
			RawDescriptor: file__07_grpc_error_proto_err_detail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file__07_grpc_error_proto_err_detail_proto_goTypes,
		DependencyIndexes: file__07_grpc_error_proto_err_detail_proto_depIdxs,
		MessageInfos:      file__07_grpc_error_proto_err_detail_proto_msgTypes,
	}.Build()
	File__07_grpc_error_proto_err_detail_proto = out.File
	file__07_grpc_error_proto_err_detail_proto_rawDesc = nil
	file__07_grpc_error_proto_err_detail_proto_goTypes = nil
	file__07_grpc_error_proto_err_detail_proto_depIdxs = nil
}
