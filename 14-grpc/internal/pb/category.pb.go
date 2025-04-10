// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: proto/category.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Blank struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Blank) Reset() {
	*x = Blank{}
	mi := &file_proto_category_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Blank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blank) ProtoMessage() {}

func (x *Blank) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blank.ProtoReflect.Descriptor instead.
func (*Blank) Descriptor() ([]byte, []int) {
	return file_proto_category_proto_rawDescGZIP(), []int{0}
}

type Category struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Category) Reset() {
	*x = Category{}
	mi := &file_proto_category_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_proto_category_proto_rawDescGZIP(), []int{1}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	mi := &file_proto_category_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCategoryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CategoryList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Categories    []*Category            `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryList) Reset() {
	*x = CategoryList{}
	mi := &file_proto_category_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryList) ProtoMessage() {}

func (x *CategoryList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryList.ProtoReflect.Descriptor instead.
func (*CategoryList) Descriptor() ([]byte, []int) {
	return file_proto_category_proto_rawDescGZIP(), []int{3}
}

func (x *CategoryList) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type GetCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCategoryRequest) Reset() {
	*x = GetCategoryRequest{}
	mi := &file_proto_category_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRequest) ProtoMessage() {}

func (x *GetCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_proto_rawDescGZIP(), []int{4}
}

func (x *GetCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_category_proto protoreflect.FileDescriptor

const file_proto_category_proto_rawDesc = "" +
	"\n" +
	"\x14proto/category.proto\x12\x02pb\"\a\n" +
	"\x05blank\"P\n" +
	"\bCategory\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"M\n" +
	"\x15CreateCategoryRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\"<\n" +
	"\fCategoryList\x12,\n" +
	"\n" +
	"categories\x18\x01 \x03(\v2\f.pb.CategoryR\n" +
	"categories\"$\n" +
	"\x12GetCategoryRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\xd3\x02\n" +
	"\x0fCategoryService\x12;\n" +
	"\x0eCreateCategory\x12\x19.pb.CreateCategoryRequest\x1a\f.pb.Category\"\x00\x12G\n" +
	"\x14CreateCategoryStream\x12\x19.pb.CreateCategoryRequest\x1a\x10.pb.CategoryList\"\x00(\x01\x12R\n" +
	"!CreateCategoryBidirectionalStream\x12\x19.pb.CreateCategoryRequest\x1a\f.pb.Category\"\x00(\x010\x01\x12/\n" +
	"\x0eListCategories\x12\t.pb.blank\x1a\x10.pb.CategoryList\"\x00\x125\n" +
	"\vGetCategory\x12\x16.pb.GetCategoryRequest\x1a\f.pb.Category\"\x00B\rZ\vinternal/pbb\x06proto3"

var (
	file_proto_category_proto_rawDescOnce sync.Once
	file_proto_category_proto_rawDescData []byte
)

func file_proto_category_proto_rawDescGZIP() []byte {
	file_proto_category_proto_rawDescOnce.Do(func() {
		file_proto_category_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_category_proto_rawDesc), len(file_proto_category_proto_rawDesc)))
	})
	return file_proto_category_proto_rawDescData
}

var file_proto_category_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_category_proto_goTypes = []any{
	(*Blank)(nil),                 // 0: pb.blank
	(*Category)(nil),              // 1: pb.Category
	(*CreateCategoryRequest)(nil), // 2: pb.CreateCategoryRequest
	(*CategoryList)(nil),          // 3: pb.CategoryList
	(*GetCategoryRequest)(nil),    // 4: pb.GetCategoryRequest
}
var file_proto_category_proto_depIdxs = []int32{
	1, // 0: pb.CategoryList.categories:type_name -> pb.Category
	2, // 1: pb.CategoryService.CreateCategory:input_type -> pb.CreateCategoryRequest
	2, // 2: pb.CategoryService.CreateCategoryStream:input_type -> pb.CreateCategoryRequest
	2, // 3: pb.CategoryService.CreateCategoryBidirectionalStream:input_type -> pb.CreateCategoryRequest
	0, // 4: pb.CategoryService.ListCategories:input_type -> pb.blank
	4, // 5: pb.CategoryService.GetCategory:input_type -> pb.GetCategoryRequest
	1, // 6: pb.CategoryService.CreateCategory:output_type -> pb.Category
	3, // 7: pb.CategoryService.CreateCategoryStream:output_type -> pb.CategoryList
	1, // 8: pb.CategoryService.CreateCategoryBidirectionalStream:output_type -> pb.Category
	3, // 9: pb.CategoryService.ListCategories:output_type -> pb.CategoryList
	1, // 10: pb.CategoryService.GetCategory:output_type -> pb.Category
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_category_proto_init() }
func file_proto_category_proto_init() {
	if File_proto_category_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_category_proto_rawDesc), len(file_proto_category_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_category_proto_goTypes,
		DependencyIndexes: file_proto_category_proto_depIdxs,
		MessageInfos:      file_proto_category_proto_msgTypes,
	}.Build()
	File_proto_category_proto = out.File
	file_proto_category_proto_goTypes = nil
	file_proto_category_proto_depIdxs = nil
}
