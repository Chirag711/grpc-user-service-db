package main

import (
	"log"
	"net"

	"grpc-user-service-db/config"
	"grpc-user-service-db/grpc-user-service/pb"
	"grpc-user-service-db/services"

	"google.golang.org/grpc"
)

func main() {

	config.ConnectDB()

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &services.Server{})

	log.Println("gRPC Server running on port 50051")

	err = s.Serve(lis)

	if err != nil {
		log.Fatal(err)
	}
}
