package application

import (
	"context"
	"teams_service/internal/core/models"
)

func (u *useCase) GetOne(c context.Context, teamId int) (models.OneTeam, error) {
	team, err := u.repo.GetOne(c, teamId)
	if err != nil {
		return team, err
	}

	members, err := u.repo.GetMembers(c, team.ID)
	if err != nil {
		return team, err
	}

	team.Members = members

	return team, err
}
