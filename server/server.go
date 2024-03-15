package main

import (
	"log"
	"net"

	"grpc_pr3.1/pkg/grpcserv/adder"
	"grpc_pr3.1/pkg/proto/proto"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	srv := &adder.GRPCServer{}
	proto.RegisterAdderServer(s, srv)

	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("Прослуховування не відбулося!")
	}

	s.Serve(listener)
}
