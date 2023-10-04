// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: proto/book.proto

package __

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

type BookCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=Author,proto3" json:"Author,omitempty"`
}

func (x *BookCreateRequest) Reset() {
	*x = BookCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookCreateRequest) ProtoMessage() {}

func (x *BookCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookCreateRequest.ProtoReflect.Descriptor instead.
func (*BookCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{0}
}

func (x *BookCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookCreateRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type BookUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDBook int64  `protobuf:"varint,1,opt,name=IDBook,proto3" json:"IDBook,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
}

func (x *BookUpdateRequest) Reset() {
	*x = BookUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookUpdateRequest) ProtoMessage() {}

func (x *BookUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookUpdateRequest.ProtoReflect.Descriptor instead.
func (*BookUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{1}
}

func (x *BookUpdateRequest) GetIDBook() int64 {
	if x != nil {
		return x.IDBook
	}
	return 0
}

func (x *BookUpdateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookUpdateRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type BookDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDBook int64 `protobuf:"varint,1,opt,name=IDBook,proto3" json:"IDBook,omitempty"`
}

func (x *BookDeleteRequest) Reset() {
	*x = BookDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookDeleteRequest) ProtoMessage() {}

func (x *BookDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookDeleteRequest.ProtoReflect.Descriptor instead.
func (*BookDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{2}
}

func (x *BookDeleteRequest) GetIDBook() int64 {
	if x != nil {
		return x.IDBook
	}
	return 0
}

type ListBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDBook int64  `protobuf:"varint,1,opt,name=IDBook,proto3" json:"IDBook,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
}

func (x *ListBooksRequest) Reset() {
	*x = ListBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksRequest) ProtoMessage() {}

func (x *ListBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksRequest.ProtoReflect.Descriptor instead.
func (*ListBooksRequest) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{3}
}

func (x *ListBooksRequest) GetIDBook() int64 {
	if x != nil {
		return x.IDBook
	}
	return 0
}

func (x *ListBooksRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListBooksRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type ListBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages string          `protobuf:"bytes,1,opt,name=Messages,proto3" json:"Messages,omitempty"`
	Code     string          `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Data     []*BookResponse `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ListBooksResponse) Reset() {
	*x = ListBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksResponse) ProtoMessage() {}

func (x *ListBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksResponse.ProtoReflect.Descriptor instead.
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{4}
}

func (x *ListBooksResponse) GetMessages() string {
	if x != nil {
		return x.Messages
	}
	return ""
}

func (x *ListBooksResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ListBooksResponse) GetData() []*BookResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type BookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDBook int64  `protobuf:"varint,1,opt,name=IDBook,proto3" json:"IDBook,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
}

func (x *BookResponse) Reset() {
	*x = BookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_book_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookResponse) ProtoMessage() {}

func (x *BookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_book_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookResponse.ProtoReflect.Descriptor instead.
func (*BookResponse) Descriptor() ([]byte, []int) {
	return file_proto_book_proto_rawDescGZIP(), []int{5}
}

func (x *BookResponse) GetIDBook() int64 {
	if x != nil {
		return x.IDBook
	}
	return 0
}

func (x *BookResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

var File_proto_book_proto protoreflect.FileDescriptor

var file_proto_book_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x41, 0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x6b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x59, 0x0a, 0x11, 0x42,
	0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x49, 0x44, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x49, 0x44, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x49,
	0x44, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x49, 0x44, 0x42,
	0x6f, 0x6f, 0x6b, 0x22, 0x58, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x44, 0x42, 0x6f, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x49, 0x44, 0x42, 0x6f, 0x6f, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x6b, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x54, 0x0a, 0x0c, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x44,
	0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x49, 0x44, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x32, 0x8b, 0x02, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x17,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x16, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03,
	0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_book_proto_rawDescOnce sync.Once
	file_proto_book_proto_rawDescData = file_proto_book_proto_rawDesc
)

func file_proto_book_proto_rawDescGZIP() []byte {
	file_proto_book_proto_rawDescOnce.Do(func() {
		file_proto_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_book_proto_rawDescData)
	})
	return file_proto_book_proto_rawDescData
}

var file_proto_book_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_book_proto_goTypes = []interface{}{
	(*BookCreateRequest)(nil), // 0: book.BookCreateRequest
	(*BookUpdateRequest)(nil), // 1: book.BookUpdateRequest
	(*BookDeleteRequest)(nil), // 2: book.BookDeleteRequest
	(*ListBooksRequest)(nil),  // 3: book.ListBooksRequest
	(*ListBooksResponse)(nil), // 4: book.ListBooksResponse
	(*BookResponse)(nil),      // 5: book.BookResponse
}
var file_proto_book_proto_depIdxs = []int32{
	5, // 0: book.ListBooksResponse.Data:type_name -> book.BookResponse
	0, // 1: book.BookService.CreateBook:input_type -> book.BookCreateRequest
	1, // 2: book.BookService.UpdateBook:input_type -> book.BookUpdateRequest
	2, // 3: book.BookService.DeleteBook:input_type -> book.BookDeleteRequest
	3, // 4: book.BookService.ListBooks:input_type -> book.ListBooksRequest
	4, // 5: book.BookService.CreateBook:output_type -> book.ListBooksResponse
	4, // 6: book.BookService.UpdateBook:output_type -> book.ListBooksResponse
	4, // 7: book.BookService.DeleteBook:output_type -> book.ListBooksResponse
	4, // 8: book.BookService.ListBooks:output_type -> book.ListBooksResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_book_proto_init() }
func file_proto_book_proto_init() {
	if File_proto_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookCreateRequest); i {
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
		file_proto_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookUpdateRequest); i {
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
		file_proto_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookDeleteRequest); i {
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
		file_proto_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksRequest); i {
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
		file_proto_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksResponse); i {
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
		file_proto_book_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookResponse); i {
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
			RawDescriptor: file_proto_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_book_proto_goTypes,
		DependencyIndexes: file_proto_book_proto_depIdxs,
		MessageInfos:      file_proto_book_proto_msgTypes,
	}.Build()
	File_proto_book_proto = out.File
	file_proto_book_proto_rawDesc = nil
	file_proto_book_proto_goTypes = nil
	file_proto_book_proto_depIdxs = nil
}
