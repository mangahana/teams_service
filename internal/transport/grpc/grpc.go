package grpc

import (
	"context"
	"log"
	"net"
	"teams_service/internal/application"
	pb "teams_service/proto/teams"

	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedTeamsServer
	server  *grpc.Server
	useCase application.UseCase
}

func New(useCase application.UseCase) *grpcServer {
	return &grpcServer{
		useCase: useCase,
		server:  grpc.NewServer(),
	}
}

func (s *grpcServer) Run() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterTeamsServer(s.server, s)
	s.server.Serve(listener)
}

func (s *grpcServer) GetOne(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	log.Println("inc")
	team, err := s.useCase.GetOne(ctx, int(r.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Id:      int32(team.ID),
		Name:    team.Name,
		Photo:   *team.Photo,
		OwnerId: int32(team.OwnerID),
	}, nil
}
