package repository

import (
	"context"
	"teams_service/internal/core/models"
)

func (r *repo) GetMembers(c context.Context, teamID int) ([]models.Member, error) {
	var output []models.Member
	sql := "SELECT user_id, user_name, user_photo FROM members WHERE team_id = $1;"

	rows, err := r.db.Query(c, sql, teamID)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var member models.Member
		if err := rows.Scan(&member.UserId, &member.UserName, &member.UserPhoto); err != nil {
			return output, err
		}
		output = append(output, member)
	}

	return output, nil
}
