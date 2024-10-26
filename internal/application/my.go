package application

import (
	"context"
	"teams_service/internal/core/models"
)

func (u *useCase) GetMyTeams(c context.Context, user *models.User) ([]models.Team, error) {
	return u.repo.GetTeamsByMember(c, user.ID)
}
