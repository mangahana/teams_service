package application

import (
	"context"
	"teams_service/internal/core/dto"
	"teams_service/internal/core/models"
)

const UPDATE_TEAM_PERMISSION = "update_team"

func (u *useCase) Update(c context.Context, user *models.User, dto *dto.Update) error {
	err := u.checkPermission(c, dto.TeamId, user.ID, UPDATE_TEAM_PERMISSION)
	if err != nil {
		return err
	}

	return u.repo.Update(c, dto)
}
