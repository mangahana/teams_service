package application

import (
	"context"
	"slices"
	"teams_service/internal/core/cerror"
)

const UPDATE_TEAM_PERMISSION = "update_team"

func (u *useCase) checkPermission(c context.Context, teamId, memberId int) error {
	team, err := u.repo.GetOne(c, teamId)
	if err != nil {
		return err
	}

	if team.OwnerID != memberId {
		member, err := u.repo.GetMember(c, teamId, memberId)
		if err != nil {
			return err
		}

		if !slices.Contains(member.Permissions, UPDATE_TEAM_PERMISSION) {
			return cerror.Forbidden()
		}
	}

	return nil
}
