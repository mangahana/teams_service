package repository

import (
	"context"
	"errors"
	"teams_service/internal/core/dto"

	"github.com/jackc/pgx/v5"
)

func (r *repo) Update(c context.Context, dto *dto.Update) error {
	sql := "UPDATE teams SET name = $1, description = $2 WHERE id = $3;"

	err := r.db.QueryRow(c, sql, dto.Name, dto.Description, dto.TeamId).Scan()
	if errors.Is(err, pgx.ErrNoRows) || err == nil {
		return nil
	}

	return err
}
