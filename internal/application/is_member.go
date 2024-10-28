package application

import (
	"context"
	"teams_service/internal/core/models"
)

func (u *useCase) IsMember(c context.Context, user *models.User, teamId int) error {
	_, err := u.repo.GetMember(c, teamId, user.ID)
	return err
}
