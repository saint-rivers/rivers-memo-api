// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: memo.proto

package memo

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Last *string  `protobuf:"bytes,1,opt,name=last,proto3,oneof" json:"last,omitempty"`
	Size int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Tags []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *PageRequest) Reset() {
	*x = PageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRequest) ProtoMessage() {}

func (x *PageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRequest.ProtoReflect.Descriptor instead.
func (*PageRequest) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{0}
}

func (x *PageRequest) GetLast() string {
	if x != nil && x.Last != nil {
		return *x.Last
	}
	return ""
}

func (x *PageRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PageRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type MemoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string   `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Tags []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *MemoRequest) Reset() {
	*x = MemoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoRequest) ProtoMessage() {}

func (x *MemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoRequest.ProtoReflect.Descriptor instead.
func (*MemoRequest) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{1}
}

func (x *MemoRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *MemoRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Memo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Link        string   `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Tags        []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Title       *string  `protobuf:"bytes,4,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Description *string  `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	ImageUrl    *string  `protobuf:"bytes,6,opt,name=image_url,json=imageUrl,proto3,oneof" json:"image_url,omitempty"`
}

func (x *Memo) Reset() {
	*x = Memo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memo) ProtoMessage() {}

func (x *Memo) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memo.ProtoReflect.Descriptor instead.
func (*Memo) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{2}
}

func (x *Memo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Memo) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Memo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Memo) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Memo) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Memo) GetImageUrl() string {
	if x != nil && x.ImageUrl != nil {
		return *x.ImageUrl
	}
	return ""
}

type MemoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memos []*Memo `protobuf:"bytes,1,rep,name=memos,proto3" json:"memos,omitempty"`
}

func (x *MemoReply) Reset() {
	*x = MemoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoReply) ProtoMessage() {}

func (x *MemoReply) ProtoReflect() protoreflect.Message {
	mi := &file_memo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoReply.ProtoReflect.Descriptor instead.
func (*MemoReply) Descriptor() ([]byte, []int) {
	return file_memo_proto_rawDescGZIP(), []int{3}
}

func (x *MemoReply) GetMemos() []*Memo {
	if x != nil {
		return x.Memos
	}
	return nil
}

var File_memo_proto protoreflect.FileDescriptor

var file_memo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0b, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6c, 0x61,
	0x73, 0x74, 0x22, 0x35, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x04, 0x4d, 0x65,
	0x6d, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x28, 0x0a, 0x09, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x05, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x05, 0x6d, 0x65, 0x6d, 0x6f, 0x73,
	0x32, 0x6c, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x26, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x12, 0x0c, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x4d, 0x65, 0x6d, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x12, 0x0c, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x15,
	0x5a, 0x13, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x6d, 0x65, 0x6d, 0x6f,
	0x3b, 0x6d, 0x65, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memo_proto_rawDescOnce sync.Once
	file_memo_proto_rawDescData = file_memo_proto_rawDesc
)

func file_memo_proto_rawDescGZIP() []byte {
	file_memo_proto_rawDescOnce.Do(func() {
		file_memo_proto_rawDescData = protoimpl.X.CompressGZIP(file_memo_proto_rawDescData)
	})
	return file_memo_proto_rawDescData
}

var file_memo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_memo_proto_goTypes = []any{
	(*PageRequest)(nil), // 0: PageRequest
	(*MemoRequest)(nil), // 1: MemoRequest
	(*Memo)(nil),        // 2: Memo
	(*MemoReply)(nil),   // 3: MemoReply
	(*empty.Empty)(nil), // 4: google.protobuf.Empty
}
var file_memo_proto_depIdxs = []int32{
	2, // 0: MemoReply.memos:type_name -> Memo
	0, // 1: MemoService.GetMemos:input_type -> PageRequest
	0, // 2: MemoService.InsertMemos:input_type -> PageRequest
	3, // 3: MemoService.GetMemos:output_type -> MemoReply
	4, // 4: MemoService.InsertMemos:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_memo_proto_init() }
func file_memo_proto_init() {
	if File_memo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_memo_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PageRequest); i {
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
		file_memo_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MemoRequest); i {
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
		file_memo_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Memo); i {
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
		file_memo_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MemoReply); i {
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
	file_memo_proto_msgTypes[0].OneofWrappers = []any{}
	file_memo_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_memo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memo_proto_goTypes,
		DependencyIndexes: file_memo_proto_depIdxs,
		MessageInfos:      file_memo_proto_msgTypes,
	}.Build()
	File_memo_proto = out.File
	file_memo_proto_rawDesc = nil
	file_memo_proto_goTypes = nil
	file_memo_proto_depIdxs = nil
}
