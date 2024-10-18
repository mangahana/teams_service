package application

import (
	"context"
	"errors"
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"

	"github.com/jackc/pgx/v5"
)

func (u *useCase) CreateInvite(c context.Context, dto *dto.CreateInvite) error {
	team, err := u.repo.GetOne(c, dto.TeamId)
	if err != nil {

		return err
	}

	if team.OwnerID != dto.OwnerId {
		return cerror.Forbidden()
	}

	_, err = u.repo.GetMember(c, dto.TeamId, dto.UserId)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
	}

	return u.repo.CreateInvite(c, dto)
}
