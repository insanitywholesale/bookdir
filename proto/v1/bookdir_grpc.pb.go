// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookDirClient is the client API for BookDir service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookDirClient interface {
	GetAllPublishers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PublisherList, error)
	GetBooksByPublisher(ctx context.Context, in *Publisher, opts ...grpc.CallOption) (*BookList, error)
	GetAllAuthors(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuthorList, error)
	GetAuthorById(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Author, error)
	GetBooksByAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*BookList, error)
	GetAllBooks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error)
	GetBookByISBN(ctx context.Context, in *ISBN, opts ...grpc.CallOption) (*Book, error)
	AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error)
}

type bookDirClient struct {
	cc grpc.ClientConnInterface
}

func NewBookDirClient(cc grpc.ClientConnInterface) BookDirClient {
	return &bookDirClient{cc}
}

func (c *bookDirClient) GetAllPublishers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PublisherList, error) {
	out := new(PublisherList)
	err := c.cc.Invoke(ctx, "/bookdir.v1.BookDir/GetAllPublishers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookDirClient) GetBooksByPublisher(ctx context.Context, in *Publisher, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/bookdir.v1.BookDir/GetBooksByPublisher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookDirClient) GetAllAuthors(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuthorList, error) {
	out := new(AuthorList)
	err := c.cc.Invoke(ctx, "/bookdir.v1.BookDir/GetAllAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookDirClient) GetAuthorById(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/bookdir.v1.BookDir/GetAuthorById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookDirClient) GetBooksByAuthor(ctx context.Context, in *Author, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/bookdir.v1.BookDir/GetBooksByAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookDirClient) GetAllBooks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/bookdir.v1.BookDir/GetAllBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookDirClient) GetBookByISBN(ctx context.Context, in *ISBN, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/bookdir.v1.BookDir/GetBookByISBN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookDirClient) AddBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/bookdir.v1.BookDir/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookDirServer is the server API for BookDir service.
// All implementations must embed UnimplementedBookDirServer
// for forward compatibility
type BookDirServer interface {
	GetAllPublishers(context.Context, *Empty) (*PublisherList, error)
	GetBooksByPublisher(context.Context, *Publisher) (*BookList, error)
	GetAllAuthors(context.Context, *Empty) (*AuthorList, error)
	GetAuthorById(context.Context, *Author) (*Author, error)
	GetBooksByAuthor(context.Context, *Author) (*BookList, error)
	GetAllBooks(context.Context, *Empty) (*BookList, error)
	GetBookByISBN(context.Context, *ISBN) (*Book, error)
	AddBook(context.Context, *Book) (*Empty, error)
	mustEmbedUnimplementedBookDirServer()
}

// UnimplementedBookDirServer must be embedded to have forward compatible implementations.
type UnimplementedBookDirServer struct {
}

func (UnimplementedBookDirServer) GetAllPublishers(context.Context, *Empty) (*PublisherList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPublishers not implemented")
}
func (UnimplementedBookDirServer) GetBooksByPublisher(context.Context, *Publisher) (*BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooksByPublisher not implemented")
}
func (UnimplementedBookDirServer) GetAllAuthors(context.Context, *Empty) (*AuthorList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAuthors not implemented")
}
func (UnimplementedBookDirServer) GetAuthorById(context.Context, *Author) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorById not implemented")
}
func (UnimplementedBookDirServer) GetBooksByAuthor(context.Context, *Author) (*BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooksByAuthor not implemented")
}
func (UnimplementedBookDirServer) GetAllBooks(context.Context, *Empty) (*BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBooks not implemented")
}
func (UnimplementedBookDirServer) GetBookByISBN(context.Context, *ISBN) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookByISBN not implemented")
}
func (UnimplementedBookDirServer) AddBook(context.Context, *Book) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedBookDirServer) mustEmbedUnimplementedBookDirServer() {}

// UnsafeBookDirServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookDirServer will
// result in compilation errors.
type UnsafeBookDirServer interface {
	mustEmbedUnimplementedBookDirServer()
}

func RegisterBookDirServer(s grpc.ServiceRegistrar, srv BookDirServer) {
	s.RegisterService(&BookDir_ServiceDesc, srv)
}

func _BookDir_GetAllPublishers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookDirServer).GetAllPublishers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookdir.v1.BookDir/GetAllPublishers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookDirServer).GetAllPublishers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookDir_GetBooksByPublisher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Publisher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookDirServer).GetBooksByPublisher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookdir.v1.BookDir/GetBooksByPublisher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookDirServer).GetBooksByPublisher(ctx, req.(*Publisher))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookDir_GetAllAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookDirServer).GetAllAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookdir.v1.BookDir/GetAllAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookDirServer).GetAllAuthors(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookDir_GetAuthorById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookDirServer).GetAuthorById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookdir.v1.BookDir/GetAuthorById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookDirServer).GetAuthorById(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookDir_GetBooksByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookDirServer).GetBooksByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookdir.v1.BookDir/GetBooksByAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookDirServer).GetBooksByAuthor(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookDir_GetAllBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookDirServer).GetAllBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookdir.v1.BookDir/GetAllBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookDirServer).GetAllBooks(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookDir_GetBookByISBN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ISBN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookDirServer).GetBookByISBN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookdir.v1.BookDir/GetBookByISBN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookDirServer).GetBookByISBN(ctx, req.(*ISBN))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookDir_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookDirServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookdir.v1.BookDir/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookDirServer).AddBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

// BookDir_ServiceDesc is the grpc.ServiceDesc for BookDir service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookDir_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookdir.v1.BookDir",
	HandlerType: (*BookDirServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPublishers",
			Handler:    _BookDir_GetAllPublishers_Handler,
		},
		{
			MethodName: "GetBooksByPublisher",
			Handler:    _BookDir_GetBooksByPublisher_Handler,
		},
		{
			MethodName: "GetAllAuthors",
			Handler:    _BookDir_GetAllAuthors_Handler,
		},
		{
			MethodName: "GetAuthorById",
			Handler:    _BookDir_GetAuthorById_Handler,
		},
		{
			MethodName: "GetBooksByAuthor",
			Handler:    _BookDir_GetBooksByAuthor_Handler,
		},
		{
			MethodName: "GetAllBooks",
			Handler:    _BookDir_GetAllBooks_Handler,
		},
		{
			MethodName: "GetBookByISBN",
			Handler:    _BookDir_GetBookByISBN_Handler,
		},
		{
			MethodName: "AddBook",
			Handler:    _BookDir_AddBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/bookdir.proto",
}
