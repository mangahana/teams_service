package auth_service

import (
	pb "teams_service/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthService struct {
	client pb.UserClient
}

func New(socket string) (*AuthService, error) {
	creds := insecure.NewCredentials()

	conn, err := grpc.NewClient(socket, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return &AuthService{
		client: pb.NewUserClient(conn),
	}, nil
}
