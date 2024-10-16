package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) TeamsCountForOwner(c context.Context, ownerID int) (int, error) {
	var count int
	sql := "SELECT COUNT(id) as count FROM teams WHERE owner_id = $1;"
	err := r.db.QueryRow(c, sql, ownerID).Scan(&count)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}
