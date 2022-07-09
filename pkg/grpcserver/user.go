package grpcserver

import (
	"context"
	"fmt"
	"msqrd/sneakers/pkg/api/userapi"
	"os"

	"google.golang.org/grpc"
)

func (s *GRPCServer) CreateUser(ctx context.Context, r *userapi.CreateUserRequest) (*userapi.CreateUserResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:8081", os.Getenv("USER_SERVICE_URL")), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := userapi.NewUserServiceClient(conn)

	res, err := c.CreateUser(
		context.Background(),
		&userapi.CreateUserRequest{
			Email:    r.GetEmail(),
			Name:     r.GetName(),
			Surname:  r.GetSurname(),
			Password: r.GetPassword(),
		},
	)

	if err != nil {
		return nil, err
	}

	return &userapi.CreateUserResponse{
		Result: res.GetResult(),
	}, nil

}
