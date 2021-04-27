package main

import (
	"context"
	api "gitlab.com/insanitywholesale/bookdir/grpc"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

// Test shitpost
func TestGetServiceName(t *testing.T) {
	expectedResult := "bookdir"
	actualResult := getServiceName()
	if expectedResult != actualResult {
		t.Fatal("service name is wrong")
	}
	t.Log("service name is:", actualResult)

}

// Test gRPC
func TestGetAllBooks(t *testing.T) {
	const bufsize = 1024 * 1024
	l := bufconn.Listen(bufsize)
	s := grpc.NewServer()
	pb.RegisterBookDirServer(s, api.Server{})
	go s.Serve(l)

	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewBookDirClient(conn)
	res, err := client.GetAllBooks(ctx, &pb.Empty{})
	if err != nil {
		t.Fatalf("failed to get GetAllBooks response %v", err)
	}
	t.Log("response", res)
}

// Test adding a book
func TestAddBook(t *testing.T) {
	const bufsize = 1024 * 1024
	l := bufconn.Listen(bufsize)
	s := grpc.NewServer()
	pb.RegisterBookDirServer(s, api.Server{})
	go s.Serve(l)

	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewBookDirClient(conn)

	testbook := &pb.Book{
		ISBN:  "9605122839",
		Title: "hey",
		Author: &pb.Author{
			FirstName:    "me",
			MiddleName:   "notme",
			LastName:     "notyou",
			YearBorn:     1234,
			YearDied:     7890,
			BooksWritten: 2,
		},
		Year:    1999,
		Edition: 1,
		Publisher: &pb.Publisher{
			Name:           "urmom",
			YearStarted:    1237,
			YearEnded:      2077,
			BooksPublished: 9001,
		},
		Pages:    100,
		Category: "tech",
		PDF:      true,
		Owned:    false,
	}
	_, err = client.AddBook(ctx, testbook)
	if err != nil {
		t.Fatalf("failed to do some AddBook action %v", err)
	}
}

// Test getting the previously added book by ISBN
func TestGetBookByISBN(t *testing.T) {
	const bufsize = 1024 * 1024
	l := bufconn.Listen(bufsize)
	s := grpc.NewServer()
	pb.RegisterBookDirServer(s, api.Server{})
	go s.Serve(l)

	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewBookDirClient(conn)
	res, err := client.GetBookByISBN(ctx, &pb.ISBN{ISBN: "9605122839"})
	if err != nil {
		t.Fatalf("failed to get GetBookByISBN response %v", err)
	}
	t.Log("response", res)
}

