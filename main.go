package main

import (
	"fmt"
	api "gitlab.com/insanitywholesale/bookdir/grpc"
	"gitlab.com/insanitywholesale/bookdir/rest"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

const serviceName string = "bookdir"

func getServiceName() string {
	return serviceName
}

func main() {
	// shitpost
	fmt.Println(getServiceName())
	// set up port
	port := os.Getenv("BOOKDIR_PORT")
	if port == "" {
		port = "11000"
	}
	// grpc server
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("listen failed %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBookDirServer(grpcServer, api.Server{})
	reflection.Register(grpcServer)
	go grpcServer.Serve(listener)
	rest.RunGateway()
}
