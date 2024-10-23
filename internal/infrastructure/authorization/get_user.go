package authorization

import (
	"context"
	"teams_service/internal/core/models"
	pb "teams_service/proto/authorization"
)

func (s *service) GetUser(c context.Context, token string) (models.User, error) {
	user, err := s.client.GetUser(c, &pb.GetUserRequest{AccessToken: token})
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:          int(user.ID),
		Username:    user.Username,
		Photo:       user.Photo,
		IsBanned:    user.IsBanned,
		Permissions: user.Permissions,
	}, nil
}
