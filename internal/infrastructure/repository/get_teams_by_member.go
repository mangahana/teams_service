package repository

import (
	"context"
	"teams_service/internal/core/models"
)

func (r *repo) GetTeamsByMember(c context.Context, memberId int) ([]models.Team, error) {
	output := []models.Team{}

	sql := `SELECT id, name, photo FROM teams WHERE id = any(COALESCE((SELECT array_agg(team_id) FROM members WHERE user_id = $1), '{}'));`

	rows, err := r.db.Query(c, sql, memberId)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var t models.Team
		if err := rows.Scan(&t.ID, &t.Name, &t.Photo); err != nil {
			return output, err
		}
		output = append(output, t)
	}

	return output, nil
}
