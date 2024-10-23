package application

import (
	"context"
	"slices"
	"teams_service/internal/core/cerror"
)

func (u *useCase) checkPermission(c context.Context, teamId, memberId int, permission string) error {
	team, err := u.repo.GetOne(c, teamId)
	if err != nil {
		return err
	}

	if team.OwnerID != memberId {
		member, err := u.repo.GetMember(c, teamId, memberId)
		if err != nil {
			return err
		}

		if !slices.Contains(member.Permissions, permission) {
			return cerror.Forbidden()
		}
	}

	return nil
}
