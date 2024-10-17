package application

import (
	"context"
	"teams_service/internal/core/models"
)

func (u *useCase) GetMembers(c context.Context, teamID int) ([]models.Member, error) {
	return u.repo.GetMembers(c, teamID)
}
