package main

import (
	"context"
	"fmt"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

const serviceName string = "bookdir"

type Server struct {
	// needed for forward compat for some reason
	pb.UnimplementedBookDirServer
}

func getServiceName() string {
	return serviceName
}

func (Server) GetAllBooks(_ context.Context, _ *pb.Empty) (*pb.BookList, error) {
	fmt.Println("getallbooks")
	return &pb.BookList{}, nil
}

func (Server) AddBook(ctx context.Context, book *pb.Book) (*pb.BookList, error) {
	fmt.Println("addbook")
	return &pb.BookList{}, nil
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
	pb.RegisterBookDirServer(grpcServer, Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(listener)
}
