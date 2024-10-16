package auth_service

import (
	"context"
	"teams_service/internal/core/models"
	pb "teams_service/proto/auth"
)

func (service *AuthService) GetSession(c context.Context, token string) (*models.Session, error) {
	req := &pb.Request{
		Token: token,
	}

	resp, err := service.client.GetSession(c, req)
	if err != nil {
		return nil, err
	}

	return &models.Session{
		UserID:      int(resp.UserId),
		IsBanned:    resp.IsBanned,
		Permissions: resp.Permissions,
	}, nil
}
