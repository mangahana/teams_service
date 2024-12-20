package repository

import (
	"context"
	"teams_service/internal/core/models"
)

func (r *repo) GetMember(c context.Context, teamId, memberId int) (models.Member, error) {
	var member models.Member
	sql := "SELECT permissions, user_id, username, user_photo FROM members WHERE user_id = $1 AND team_id = $2;"
	err := r.db.QueryRow(c, sql, memberId, teamId).Scan(&member.Permissions, &member.UserID, &member.Username, &member.UserPhoto)
	return member, err
}
