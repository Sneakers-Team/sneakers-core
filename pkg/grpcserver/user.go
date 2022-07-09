package grpcserver

import (
	"context"
	"msqrd/sneakers/pkg/api/userapi"

	"google.golang.org/grpc"
)

func (s *GRPCServer) CreateUser(ctx context.Context, r *userapi.CreateUserRequest) (*userapi.CreateUserResponse, error) {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
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

	return res, nil
}
