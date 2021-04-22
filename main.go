package main

import (
	"fmt"
	api "gitlab.com/insanitywholesale/bookdir/grpc"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	"gitlab.com/insanitywholesale/bookdir/rest"
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

	// set up grpc port
	grpcport := os.Getenv("BOOKDIR_GRPC_PORT")
	if grpcport == "" {
		grpcport = "11000"
	}

	// set up grpc server
	listener, err := net.Listen("tcp", ":"+grpcport)
	if err != nil {
		log.Fatalf("listen failed %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBookDirServer(grpcServer, api.Server{})
	reflection.Register(grpcServer)
	go grpcServer.Serve(listener)

	// set up rest port
	restport := os.Getenv("BOOKDIR_REST_PORT")
	if restport == "" {
		restport = "8080"
	}
	log.Fatal(rest.RunGateway(grpcport, restport))
}
