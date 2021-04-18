// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: v1/bookdir.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type BookList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BookList) Reset() {
	*x = BookList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookdir_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookList) ProtoMessage() {}

func (x *BookList) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookdir_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookList.ProtoReflect.Descriptor instead.
func (*BookList) Descriptor() ([]byte, []int) {
	return file_v1_bookdir_proto_rawDescGZIP(), []int{0}
}

func (x *BookList) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ISBN      string     `protobuf:"bytes,1,opt,name=ISBN,proto3" json:"ISBN,omitempty"`
	Title     string     `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Author    *Author    `protobuf:"bytes,3,opt,name=Author,proto3" json:"Author,omitempty"`
	Year      uint32     `protobuf:"varint,4,opt,name=Year,proto3" json:"Year,omitempty"`
	Edition   uint32     `protobuf:"varint,5,opt,name=Edition,proto3" json:"Edition,omitempty"`
	Publisher *Publisher `protobuf:"bytes,6,opt,name=Publisher,proto3" json:"Publisher,omitempty"`
	Pages     uint32     `protobuf:"varint,7,opt,name=Pages,proto3" json:"Pages,omitempty"`
	Category  string     `protobuf:"bytes,8,opt,name=Category,proto3" json:"Category,omitempty"`
	PDF       bool       `protobuf:"varint,9,opt,name=PDF,proto3" json:"PDF,omitempty"`
	Owned     bool       `protobuf:"varint,10,opt,name=Owned,proto3" json:"Owned,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookdir_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookdir_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_v1_bookdir_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetISBN() string {
	if x != nil {
		return x.ISBN
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Book) GetYear() uint32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Book) GetEdition() uint32 {
	if x != nil {
		return x.Edition
	}
	return 0
}

func (x *Book) GetPublisher() *Publisher {
	if x != nil {
		return x.Publisher
	}
	return nil
}

func (x *Book) GetPages() uint32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *Book) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Book) GetPDF() bool {
	if x != nil {
		return x.PDF
	}
	return false
}

func (x *Book) GetOwned() bool {
	if x != nil {
		return x.Owned
	}
	return false
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName    string `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	MiddleName   string `protobuf:"bytes,2,opt,name=MiddleName,proto3" json:"MiddleName,omitempty"`
	LastName     string `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	YearBorn     uint32 `protobuf:"varint,4,opt,name=YearBorn,proto3" json:"YearBorn,omitempty"`
	YearDied     uint32 `protobuf:"varint,5,opt,name=YearDied,proto3" json:"YearDied,omitempty"`
	BooksWritten uint32 `protobuf:"varint,6,opt,name=BooksWritten,proto3" json:"BooksWritten,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookdir_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookdir_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_v1_bookdir_proto_rawDescGZIP(), []int{2}
}

func (x *Author) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Author) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *Author) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Author) GetYearBorn() uint32 {
	if x != nil {
		return x.YearBorn
	}
	return 0
}

func (x *Author) GetYearDied() uint32 {
	if x != nil {
		return x.YearDied
	}
	return 0
}

func (x *Author) GetBooksWritten() uint32 {
	if x != nil {
		return x.BooksWritten
	}
	return 0
}

type Publisher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	YearStarted    uint32 `protobuf:"varint,2,opt,name=YearStarted,proto3" json:"YearStarted,omitempty"`
	YearEnded      uint32 `protobuf:"varint,3,opt,name=YearEnded,proto3" json:"YearEnded,omitempty"`
	BooksPublished uint32 `protobuf:"varint,4,opt,name=BooksPublished,proto3" json:"BooksPublished,omitempty"`
}

func (x *Publisher) Reset() {
	*x = Publisher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookdir_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Publisher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Publisher) ProtoMessage() {}

func (x *Publisher) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookdir_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Publisher.ProtoReflect.Descriptor instead.
func (*Publisher) Descriptor() ([]byte, []int) {
	return file_v1_bookdir_proto_rawDescGZIP(), []int{3}
}

func (x *Publisher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Publisher) GetYearStarted() uint32 {
	if x != nil {
		return x.YearStarted
	}
	return 0
}

func (x *Publisher) GetYearEnded() uint32 {
	if x != nil {
		return x.YearEnded
	}
	return 0
}

func (x *Publisher) GetBooksPublished() uint32 {
	if x != nil {
		return x.BooksPublished
	}
	return 0
}

type NoArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoArguments) Reset() {
	*x = NoArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bookdir_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoArguments) ProtoMessage() {}

func (x *NoArguments) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bookdir_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoArguments.ProtoReflect.Descriptor instead.
func (*NoArguments) Descriptor() ([]byte, []int) {
	return file_v1_bookdir_proto_rawDescGZIP(), []int{4}
}

var File_v1_bookdir_proto protoreflect.FileDescriptor

var file_v1_bookdir_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x64, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x64, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x08,
	0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x64, 0x69,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x22, 0x99, 0x02, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x53, 0x42,
	0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x53, 0x42, 0x4e, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x64, 0x69, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x45, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x64, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x44, 0x46, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x50, 0x44, 0x46, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x22, 0xbe, 0x01, 0x0a,
	0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x59, 0x65, 0x61, 0x72, 0x42, 0x6f, 0x72, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x59, 0x65, 0x61, 0x72, 0x42, 0x6f, 0x72, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x59, 0x65, 0x61, 0x72, 0x44, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x59, 0x65, 0x61, 0x72, 0x44, 0x69, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x57, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x57, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x22, 0x87, 0x01,
	0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x59, 0x65, 0x61, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x59, 0x65, 0x61, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x59, 0x65, 0x61, 0x72, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x59, 0x65, 0x61, 0x72, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x4e, 0x6f, 0x41, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xb1, 0x01, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x44,
	0x69, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x64,
	0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x1a, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x64, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12,
	0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x4a,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x64, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x14, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x64, 0x69, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x3a, 0x01, 0x2a, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x61, 0x6e, 0x69, 0x74,
	0x79, 0x77, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x64,
	0x69, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_bookdir_proto_rawDescOnce sync.Once
	file_v1_bookdir_proto_rawDescData = file_v1_bookdir_proto_rawDesc
)

func file_v1_bookdir_proto_rawDescGZIP() []byte {
	file_v1_bookdir_proto_rawDescOnce.Do(func() {
		file_v1_bookdir_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_bookdir_proto_rawDescData)
	})
	return file_v1_bookdir_proto_rawDescData
}

var file_v1_bookdir_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_bookdir_proto_goTypes = []interface{}{
	(*BookList)(nil),    // 0: bookdir.v1.BookList
	(*Book)(nil),        // 1: bookdir.v1.Book
	(*Author)(nil),      // 2: bookdir.v1.Author
	(*Publisher)(nil),   // 3: bookdir.v1.Publisher
	(*NoArguments)(nil), // 4: bookdir.v1.NoArguments
}
var file_v1_bookdir_proto_depIdxs = []int32{
	1, // 0: bookdir.v1.BookList.books:type_name -> bookdir.v1.Book
	2, // 1: bookdir.v1.Book.Author:type_name -> bookdir.v1.Author
	3, // 2: bookdir.v1.Book.Publisher:type_name -> bookdir.v1.Publisher
	4, // 3: bookdir.v1.BookDirService.GetAllBooks:input_type -> bookdir.v1.NoArguments
	1, // 4: bookdir.v1.BookDirService.AddBook:input_type -> bookdir.v1.Book
	0, // 5: bookdir.v1.BookDirService.GetAllBooks:output_type -> bookdir.v1.BookList
	0, // 6: bookdir.v1.BookDirService.AddBook:output_type -> bookdir.v1.BookList
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_bookdir_proto_init() }
func file_v1_bookdir_proto_init() {
	if File_v1_bookdir_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_bookdir_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookList); i {
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
		file_v1_bookdir_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_v1_bookdir_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_v1_bookdir_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Publisher); i {
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
		file_v1_bookdir_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoArguments); i {
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
			RawDescriptor: file_v1_bookdir_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_bookdir_proto_goTypes,
		DependencyIndexes: file_v1_bookdir_proto_depIdxs,
		MessageInfos:      file_v1_bookdir_proto_msgTypes,
	}.Build()
	File_v1_bookdir_proto = out.File
	file_v1_bookdir_proto_rawDesc = nil
	file_v1_bookdir_proto_goTypes = nil
	file_v1_bookdir_proto_depIdxs = nil
}
