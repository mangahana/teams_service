package repository

import (
	"context"
	"teams_service/internal/core/models"
)

func (r *repo) GetOne(c context.Context, teamId int) (models.OneTeam, error) {
	var team models.OneTeam
	sql := "SELECT id, name, description, photo, is_verified FROM teams WHERE id = $1;"
	err := r.db.QueryRow(c, sql, teamId).Scan(&team.ID, &team.Name, &team.Description, &team.Photo, &team.IsVerified)
	return team, err
}
