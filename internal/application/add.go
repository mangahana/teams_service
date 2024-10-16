package application

import (
	"context"
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"
)

func (u *useCase) Add(c context.Context, dto *dto.AddTeam) error {
	teamsCountByOwnerID, err := u.repo.TeamsCountForOwner(c, dto.OwnerId)
	if err != nil {
		return err
	}

	if teamsCountByOwnerID >= 10 {
		return cerror.New(cerror.TEAMS_LIMIT, "too many teams created for the user")
	}

	_, err = u.repo.GetTypeByID(c, dto.TypeId)
	if err != nil {
		return err
	}

	return u.repo.Add(c, dto)
}
