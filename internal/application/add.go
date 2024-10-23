package application

import (
	"context"
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"
	"teams_service/internal/core/models"
)

const teamsLimitForUser = 5

func (u *useCase) Add(c context.Context, user *models.User, dto *dto.AddTeam) error {
	teamsCountByOwnerID, err := u.repo.TeamsCountForOwner(c, user.ID)
	if err != nil {
		return err
	}

	if teamsCountByOwnerID >= teamsLimitForUser {
		return cerror.New(cerror.TEAMS_LIMIT, "too many teams created for the user")
	}

	_, err = u.repo.GetTypeByID(c, dto.TypeId)
	if err != nil {
		return err
	}

	return u.repo.Add(c, user, dto)
}
