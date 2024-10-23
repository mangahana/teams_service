package repository

import (
	"context"
	"teams_service/internal/core/dto"
	"teams_service/internal/core/models"

	"github.com/jackc/pgx/v5"
)

// func (r *repo) Add(c context.Context, userId int, dto *dto.AddTeam) error {
// 	sql := "INSERT INTO teams (name, type_id, owner_id) VALUES ($1, $2, $3);"

// 	err := r.db.QueryRow(c, sql, dto.Name, dto.TypeId, userId).Scan()
// 	if err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			return nil
// 		}
// 		return err
// 	}

// 	return nil
// }

func (r *repo) Add(c context.Context, user *models.User, dto *dto.AddTeam) error {
	tx, err := r.db.BeginTx(c, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(c)
		}
	}()

	var teamId int

	sql := "INSERT INTO teams (name, type_id, owner_id) VALUES ($1, $2, $3) RETURNING id;"
	if err := tx.QueryRow(c, sql, dto.Name, dto.TypeId, user.ID).Scan(&teamId); err != nil {
		return nil
	}

	sql = "INSERT INTO members (team_id, user_id, username, user_photo) VALUES ($1, $2, $3, $4);"
	if _, err := tx.Exec(c, sql, teamId, user.ID, user.Username, user.Photo); err != nil {
		return nil
	}

	if err := tx.Commit(c); err != nil {
		return err
	}

	return nil
}
