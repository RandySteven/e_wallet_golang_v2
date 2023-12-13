package main

import (
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		return
	}

	opt := []grpc.ServerOption{}

	server := grpc.NewServer(opt...)

	if err = server.Serve(listener); err != nil {
		return
	}

}
