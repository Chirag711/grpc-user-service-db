package main

import (
	"context"
	"log"

	"grpc-user-service-db/grpc-user-service/pb"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	user := &pb.User{
		Name:  "Chirag Raul amdocs",
		Email: "chiragdraul@gmail.com",
	}

	res, err := client.CreateUser(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	// res, err := client.GetAllUsers(context.Background(), &pb.Empty{})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, user := range res.Users {
	// 	fmt.Println(user.Id, user.Name, user.Email)
	// }

}
