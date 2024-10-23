package application

import (
	"context"
	"teams_service/internal/core/dto"
)

func (u *useCase) UpdateMember(c context.Context, dto *dto.UpdateMember) error {
	return u.repo.UpdateMember(c, dto)
}
