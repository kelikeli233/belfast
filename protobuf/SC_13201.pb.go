// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_13201.proto

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

type SC_13201 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountList             []*EXPEDITION_DAILY_COUNT `protobuf:"bytes,1,rep,name=count_list,json=countList" json:"count_list,omitempty"`
	EliteExpeditionCount  *uint32                   `protobuf:"varint,2,req,name=elite_expedition_count,json=eliteExpeditionCount" json:"elite_expedition_count,omitempty"`
	EscortExpeditionCount *uint32                   `protobuf:"varint,3,req,name=escort_expedition_count,json=escortExpeditionCount" json:"escort_expedition_count,omitempty"`
	ChapterCountList      []*EXPEDITION_DAILY_COUNT `protobuf:"bytes,4,rep,name=chapter_count_list,json=chapterCountList" json:"chapter_count_list,omitempty"`
	QuickExpeditionList   []uint32                  `protobuf:"varint,5,rep,name=quick_expedition_list,json=quickExpeditionList" json:"quick_expedition_list,omitempty"`
}

func (x *SC_13201) Reset() {
	*x = SC_13201{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_13201_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_13201) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_13201) ProtoMessage() {}

func (x *SC_13201) ProtoReflect() protoreflect.Message {
	mi := &file_SC_13201_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_13201.ProtoReflect.Descriptor instead.
func (*SC_13201) Descriptor() ([]byte, []int) {
	return file_SC_13201_proto_rawDescGZIP(), []int{0}
}

func (x *SC_13201) GetCountList() []*EXPEDITION_DAILY_COUNT {
	if x != nil {
		return x.CountList
	}
	return nil
}

func (x *SC_13201) GetEliteExpeditionCount() uint32 {
	if x != nil && x.EliteExpeditionCount != nil {
		return *x.EliteExpeditionCount
	}
	return 0
}

func (x *SC_13201) GetEscortExpeditionCount() uint32 {
	if x != nil && x.EscortExpeditionCount != nil {
		return *x.EscortExpeditionCount
	}
	return 0
}

func (x *SC_13201) GetChapterCountList() []*EXPEDITION_DAILY_COUNT {
	if x != nil {
		return x.ChapterCountList
	}
	return nil
}

func (x *SC_13201) GetQuickExpeditionList() []uint32 {
	if x != nil {
		return x.QuickExpeditionList
	}
	return nil
}

var File_SC_13201_proto protoreflect.FileDescriptor

var file_SC_13201_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x31, 0x33, 0x32, 0x30, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x1c, 0x45, 0x58, 0x50, 0x45, 0x44,
	0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x5f, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x08, 0x53, 0x43, 0x5f, 0x31,
	0x33, 0x32, 0x30, 0x31, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61,
	0x73, 0x74, 0x2e, 0x45, 0x58, 0x50, 0x45, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41,
	0x49, 0x4c, 0x59, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x65, 0x78,
	0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x14, 0x65, 0x6c, 0x69, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x73,
	0x63, 0x6f, 0x72, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x15, 0x65, 0x73, 0x63,
	0x6f, 0x72, 0x74, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x12, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x45, 0x58, 0x50, 0x45, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52,
	0x10, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x15, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x13, 0x71, 0x75, 0x69, 0x63, 0x6b, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66,
}

var (
	file_SC_13201_proto_rawDescOnce sync.Once
	file_SC_13201_proto_rawDescData = file_SC_13201_proto_rawDesc
)

func file_SC_13201_proto_rawDescGZIP() []byte {
	file_SC_13201_proto_rawDescOnce.Do(func() {
		file_SC_13201_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_13201_proto_rawDescData)
	})
	return file_SC_13201_proto_rawDescData
}

var file_SC_13201_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_13201_proto_goTypes = []interface{}{
	(*SC_13201)(nil),               // 0: belfast.SC_13201
	(*EXPEDITION_DAILY_COUNT)(nil), // 1: belfast.EXPEDITION_DAILY_COUNT
}
var file_SC_13201_proto_depIdxs = []int32{
	1, // 0: belfast.SC_13201.count_list:type_name -> belfast.EXPEDITION_DAILY_COUNT
	1, // 1: belfast.SC_13201.chapter_count_list:type_name -> belfast.EXPEDITION_DAILY_COUNT
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SC_13201_proto_init() }
func file_SC_13201_proto_init() {
	if File_SC_13201_proto != nil {
		return
	}
	file_EXPEDITION_DAILY_COUNT_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_13201_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_13201); i {
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
			RawDescriptor: file_SC_13201_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_13201_proto_goTypes,
		DependencyIndexes: file_SC_13201_proto_depIdxs,
		MessageInfos:      file_SC_13201_proto_msgTypes,
	}.Build()
	File_SC_13201_proto = out.File
	file_SC_13201_proto_rawDesc = nil
	file_SC_13201_proto_goTypes = nil
	file_SC_13201_proto_depIdxs = nil
}