package application

import (
	"context"
	"teams_service/internal/core/models"
)

func (u *useCase) GetOne(c context.Context, teamId int) (models.OneTeam, error) {
	return u.repo.GetOne(c, teamId)
}
