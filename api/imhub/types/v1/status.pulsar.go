// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package typesv1

import (
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: imhub/types/v1/status.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_STATUS_UNSPECIFIED      Status = 0
	Status_STATUS_ACTIVE           Status = 1
	Status_STATUS_INACTIVE_PENDING Status = 2
	Status_STATUS_INACTIVE         Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_ACTIVE",
		2: "STATUS_INACTIVE_PENDING",
		3: "STATUS_INACTIVE",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED":      0,
		"STATUS_ACTIVE":           1,
		"STATUS_INACTIVE_PENDING": 2,
		"STATUS_INACTIVE":         3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_imhub_types_v1_status_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_imhub_types_v1_status_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_imhub_types_v1_status_proto_rawDescGZIP(), []int{0}
}

var File_imhub_types_v1_status_proto protoreflect.FileDescriptor

var file_imhub_types_v1_status_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6d, 0x68, 0x75, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x69,
	0x6d, 0x68, 0x75, 0x62, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xbd, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x15, 0x8a, 0x9d, 0x20, 0x11, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01,
	0x1a, 0x10, 0x8a, 0x9d, 0x20, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x1a,
	0x19, 0x8a, 0x9d, 0x20, 0x15, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x1a,
	0x12, 0x8a, 0x9d, 0x20, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x42, 0xbb, 0x01, 0xc8, 0xe1, 0x1e, 0x00, 0xd0, 0xe1, 0x1e, 0x00, 0xe8, 0xe2,
	0x1e, 0x00, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6d, 0x68, 0x75, 0x62, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x75, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x69, 0x6d, 0x68, 0x75, 0x62, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x6d, 0x68, 0x75, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x54, 0x58, 0xaa,
	0x02, 0x0e, 0x49, 0x6d, 0x68, 0x75, 0x62, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0e, 0x49, 0x6d, 0x68, 0x75, 0x62, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1a, 0x49, 0x6d, 0x68, 0x75, 0x62, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x10, 0x49, 0x6d, 0x68, 0x75, 0x62, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imhub_types_v1_status_proto_rawDescOnce sync.Once
	file_imhub_types_v1_status_proto_rawDescData = file_imhub_types_v1_status_proto_rawDesc
)

func file_imhub_types_v1_status_proto_rawDescGZIP() []byte {
	file_imhub_types_v1_status_proto_rawDescOnce.Do(func() {
		file_imhub_types_v1_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_imhub_types_v1_status_proto_rawDescData)
	})
	return file_imhub_types_v1_status_proto_rawDescData
}

var file_imhub_types_v1_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_imhub_types_v1_status_proto_goTypes = []interface{}{
	(Status)(0), // 0: imhub.types.v1.Status
}
var file_imhub_types_v1_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_imhub_types_v1_status_proto_init() }
func file_imhub_types_v1_status_proto_init() {
	if File_imhub_types_v1_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_imhub_types_v1_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_imhub_types_v1_status_proto_goTypes,
		DependencyIndexes: file_imhub_types_v1_status_proto_depIdxs,
		EnumInfos:         file_imhub_types_v1_status_proto_enumTypes,
	}.Build()
	File_imhub_types_v1_status_proto = out.File
	file_imhub_types_v1_status_proto_rawDesc = nil
	file_imhub_types_v1_status_proto_goTypes = nil
	file_imhub_types_v1_status_proto_depIdxs = nil
}
