package application

import (
	"context"
	"slices"
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"
)

const UPDATE_TEAM_PERMISSION = "update_team"

func (u *useCase) Update(c context.Context, dto *dto.Update) error {
	team, err := u.repo.GetOne(c, dto.TeamId)
	if err != nil {
		return err
	}

	if team.OwnerID != dto.MemberId {
		member, err := u.repo.GetMember(c, dto.TeamId, dto.MemberId)
		if err != nil {
			return err
		}

		if !slices.Contains(member.Permissions, UPDATE_TEAM_PERMISSION) {
			return cerror.Forbidden()
		}
	}

	return u.repo.Update(c, dto)
}
