package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) UpdatePhoto(c context.Context, teamId int, photo string) error {
	sql := "UPDATE teams SET photo = $2 WHERE id = $1;"
	err := r.db.QueryRow(c, sql, teamId, photo).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}

	return nil
}
