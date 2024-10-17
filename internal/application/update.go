package application

import (
	"context"
	"teams_service/internal/core/dto"
)

func (u *useCase) Update(c context.Context, dto *dto.Update) error {
	err := u.checkPermission(c, dto.TeamId, dto.MemberId)
	if err != nil {
		return err
	}

	return u.repo.Update(c, dto)
}
