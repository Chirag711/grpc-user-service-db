package services

import (
	"context"
	"grpc-user-service-db/grpc-user-service/pb"
	"grpc-user-service-db/models"
	"grpc-user-service-db/repository"
	"grpc-user-service-db/utils"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	result, err := repository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	go utils.SendEmail(req.Email)

	return &pb.User{
		Id:    result.ID.Hex(),
		Name:  result.Name,
		Email: result.Email,
	}, nil
}

func (s *Server) GetAllUsers(ctx context.Context, req *pb.Empty) (*pb.UserList, error) {
	users, err := repository.GetAllUsers()

	if err != nil {
		return nil, err
	}
	var userList []*pb.User

	for _, user := range users {
		userList = append(userList, &pb.User{
			Id:    user.ID.Hex(),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &pb.UserList{
		Users: userList,
	}, nil
}
