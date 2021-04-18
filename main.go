package main

import (
	"fmt"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"os"
	"net"
	"log"
	"context"
)

const serviceName string = "bookdir"

type server struct {
	// needed for forward compat for some reason
	pb.UnimplementedBookDirServer
}

func getServiceName() string {
	return serviceName
}

func (server) GetAllBooks(_ context.Context, _ *pb.NoArguments) (*pb.BookList, error) {
	fmt.Println("getallbooks")
	return nil, nil
}

func (server) AddBook(ctx context.Context, book *pb.Book) (*pb.BookList, error) {
	fmt.Println("addbook")
	return nil, nil
}

func main() {
	// shitpost
	fmt.Println(getServiceName())
	// set up port
	port := os.Getenv("BOOKDIR_PORT")
	if port == "" {
		port = "8080"
	}
	// grpc server
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("listen failed %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBookDirServer(grpcServer, server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listener)
}
