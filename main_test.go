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
	var l *bufconn.Listener
	l = bufconn.Listen(bufsize)
	s := grpc.NewServer()
	pb.RegisterBookDirServer(s, api.Server{})
	go func() {
		err := s.Serve(l)
		if err != nil {
			t.Fatalf("server exited with error: %v", err)
		}
	}()

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
