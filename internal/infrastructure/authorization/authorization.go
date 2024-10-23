package authorization

import (
	pb "teams_service/proto/authorization"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type service struct {
	client pb.AuthorizationClient
}

func New(socket string) (*service, error) {
	conn, err := grpc.NewClient(socket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &service{}, err
	}

	return &service{client: pb.NewAuthorizationClient(conn)}, nil
}
