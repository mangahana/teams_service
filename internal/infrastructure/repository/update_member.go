package repository

import (
	"context"
	"teams_service/internal/core/dto"
)

func (r *repo) UpdateMember(c context.Context, dto *dto.UpdateMember) error {
	sql := "UPDATE members SET username = $1, user_photo = $2 WHERE user_id = $3;"
	_, err := r.db.Exec(c, sql, dto.Username, dto.Photo, dto.ID)
	return err
}
