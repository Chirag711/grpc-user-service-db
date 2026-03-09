package main

import (
	"context"
	"log"

	"grpc-user-service-db/grpc-user-service/pb"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	user := &pb.User{
		Name:  "Chirag Raul amdocs",
		Email: "chiragraul@test.com",
	}

	res, err := client.CreateUser(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
