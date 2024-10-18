package repository

import (
	"context"
	"errors"
	"teams_service/internal/core/dto"

	"github.com/jackc/pgx/v5"
)

func (r *repo) CreateInvite(c context.Context, dto *dto.CreateInvite) error {
	sql := "INSERT INTO invites (team_id, user_id) VALUES ($1, $2);"
	err := r.db.QueryRow(c, sql, dto.TeamId, dto.UserId).Scan()
	if err == nil || errors.Is(err, pgx.ErrNoRows) {
		return nil
	}
	return err
}
