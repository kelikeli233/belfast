// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_62012.proto

package protobuf

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

type SC_62012 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   *uint32        `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	Inclog   []*CAPITAL_LOG `protobuf:"bytes,2,rep,name=inclog" json:"inclog,omitempty"`
	Declog   []*CAPITAL_LOG `protobuf:"bytes,3,rep,name=declog" json:"declog,omitempty"`
	Otherlog []*CAPITAL_LOG `protobuf:"bytes,4,rep,name=otherlog" json:"otherlog,omitempty"`
}

func (x *SC_62012) Reset() {
	*x = SC_62012{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_62012_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_62012) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_62012) ProtoMessage() {}

func (x *SC_62012) ProtoReflect() protoreflect.Message {
	mi := &file_SC_62012_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_62012.ProtoReflect.Descriptor instead.
func (*SC_62012) Descriptor() ([]byte, []int) {
	return file_SC_62012_proto_rawDescGZIP(), []int{0}
}

func (x *SC_62012) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *SC_62012) GetInclog() []*CAPITAL_LOG {
	if x != nil {
		return x.Inclog
	}
	return nil
}

func (x *SC_62012) GetDeclog() []*CAPITAL_LOG {
	if x != nil {
		return x.Declog
	}
	return nil
}

func (x *SC_62012) GetOtherlog() []*CAPITAL_LOG {
	if x != nil {
		return x.Otherlog
	}
	return nil
}

var File_SC_62012_proto protoreflect.FileDescriptor

var file_SC_62012_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x36, 0x32, 0x30, 0x31, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x11, 0x43, 0x41, 0x50, 0x49, 0x54,
	0x41, 0x4c, 0x5f, 0x4c, 0x4f, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a,
	0x08, 0x53, 0x43, 0x5f, 0x36, 0x32, 0x30, 0x31, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x2c, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x41, 0x50, 0x49,
	0x54, 0x41, 0x4c, 0x5f, 0x4c, 0x4f, 0x47, 0x52, 0x06, 0x69, 0x6e, 0x63, 0x6c, 0x6f, 0x67, 0x12,
	0x2c, 0x0a, 0x06, 0x64, 0x65, 0x63, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x41, 0x50, 0x49, 0x54, 0x41,
	0x4c, 0x5f, 0x4c, 0x4f, 0x47, 0x52, 0x06, 0x64, 0x65, 0x63, 0x6c, 0x6f, 0x67, 0x12, 0x30, 0x0a,
	0x08, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x43, 0x41, 0x50, 0x49, 0x54, 0x41,
	0x4c, 0x5f, 0x4c, 0x4f, 0x47, 0x52, 0x08, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_62012_proto_rawDescOnce sync.Once
	file_SC_62012_proto_rawDescData = file_SC_62012_proto_rawDesc
)

func file_SC_62012_proto_rawDescGZIP() []byte {
	file_SC_62012_proto_rawDescOnce.Do(func() {
		file_SC_62012_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_62012_proto_rawDescData)
	})
	return file_SC_62012_proto_rawDescData
}

var file_SC_62012_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_62012_proto_goTypes = []interface{}{
	(*SC_62012)(nil),    // 0: belfast.SC_62012
	(*CAPITAL_LOG)(nil), // 1: belfast.CAPITAL_LOG
}
var file_SC_62012_proto_depIdxs = []int32{
	1, // 0: belfast.SC_62012.inclog:type_name -> belfast.CAPITAL_LOG
	1, // 1: belfast.SC_62012.declog:type_name -> belfast.CAPITAL_LOG
	1, // 2: belfast.SC_62012.otherlog:type_name -> belfast.CAPITAL_LOG
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_SC_62012_proto_init() }
func file_SC_62012_proto_init() {
	if File_SC_62012_proto != nil {
		return
	}
	file_CAPITAL_LOG_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_62012_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_62012); i {
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
			RawDescriptor: file_SC_62012_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_62012_proto_goTypes,
		DependencyIndexes: file_SC_62012_proto_depIdxs,
		MessageInfos:      file_SC_62012_proto_msgTypes,
	}.Build()
	File_SC_62012_proto = out.File
	file_SC_62012_proto_rawDesc = nil
	file_SC_62012_proto_goTypes = nil
	file_SC_62012_proto_depIdxs = nil
}