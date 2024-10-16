package repository

import (
	"context"
	"teams_service/internal/core/models"
)

func (r *repo) GetTypeByID(c context.Context, typeId int) (models.TeamType, error) {
	var output models.TeamType
	sql := "SELECT id, name FROM types WHERE id = $1;"
	err := r.db.QueryRow(c, sql, typeId).Scan(&output.ID, &output.Name)
	return output, err
}
