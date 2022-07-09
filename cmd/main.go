package main

import (
	"log"
	"msqrd/sneakers/pkg/api/userapi"
	"msqrd/sneakers/pkg/grpcserver"
	"net"

	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	userapi.RegisterUserServiceServer(s, new(grpcserver.GRPCServer))

	log.Println("core server started")

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
