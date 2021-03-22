package main

import (
	"flag"
	"log"
	"net"

	v1 "github.com/davidterranova/homework/api/protobuf/v1"
	"github.com/davidterranova/homework/store/adapters"
	"github.com/davidterranova/homework/store/repository/inmemory"
	"google.golang.org/grpc"
)

var (
	bindAddress = flag.String("bind-address", ":5000", "gRPC server bind address")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *bindAddress)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	portRepository := inmemory.NewPortRepository()
	server := adapters.NewGRPCServer(portRepository)

	grpcServer := grpc.NewServer()
	v1.RegisterPortServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start grpc server: %s", err)
	}
}
