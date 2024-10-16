package repository

import (
	"context"
	"errors"
	"teams_service/internal/core/dto"

	"github.com/jackc/pgx/v5"
)

func (r *repo) Add(c context.Context, dto *dto.AddTeam) error {
	sql := "INSERT INTO teams (name, type_id, owner_id) VALUES ($1, $2, $3);"

	err := r.db.QueryRow(c, sql, dto.Name, dto.TypeId, dto.OwnerId).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}

	return nil
}
