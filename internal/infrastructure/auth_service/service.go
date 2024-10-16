package authservice

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthService struct {
	client *grpc.ClientConn
}

func New(socket string) (*AuthService, error) {
	creds := insecure.NewCredentials()

	conn, err := grpc.NewClient(socket, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return &AuthService{
		client: conn,
	}, nil
}

func (a *AuthService) Close() error {
	return a.client.Close()
}
