package application

import (
	"context"
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"
)

func (u *useCase) UpdateMemberPermissions(c context.Context, dto *dto.UpdateMemberPermissions) error {
	team, err := u.repo.GetOne(c, dto.TeamId)
	if err != nil {
		return err
	}

	if team.OwnerID != dto.OwnerId {
		return cerror.Forbidden()
	}

	_, err = u.repo.GetMember(c, dto.TeamId, dto.MemberId)
	if err != nil {
		return err
	}

	return u.repo.UpdatePermissions(c, dto)
}
