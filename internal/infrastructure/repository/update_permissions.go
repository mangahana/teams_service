package repository

import (
	"context"
	"errors"
	"teams_service/internal/core/dto"

	"github.com/jackc/pgx/v5"
)

func (r *repo) UpdatePermissions(c context.Context, dto *dto.UpdateMemberPermissions) error {
	sql := "UPDATE members SET permissions = $3 WHERE team_id = $1 AND user_id = $2;"
	err := r.db.QueryRow(c, sql, dto.TeamId, dto.MemberId, dto.Permissions).Scan()
	if err == nil || errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	return err
}
