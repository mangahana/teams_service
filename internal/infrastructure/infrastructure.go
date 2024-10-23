package infrastructure

import (
	"context"
	"teams_service/internal/core/dto"
	"teams_service/internal/core/models"
)

type AuthorizationService interface {
	GetUser(c context.Context, token string) (models.User, error)
}

type S3 interface {
	Put(c context.Context, object []byte) (string, error)
	Remove(c context.Context, objectName string) error
}

type Repository interface {
	Add(c context.Context, user *models.User, dto *dto.AddTeam) error
	CreateInvite(c context.Context, dto *dto.CreateInvite) error

	GetOne(c context.Context, teamId int) (models.OneTeam, error)
	GetTypeByID(c context.Context, typeId int) (models.TeamType, error)
	GetMember(c context.Context, teamId, memberId int) (models.Member, error)
	GetMembers(c context.Context, teamID int) ([]models.Member, error)

	TeamsCountForOwner(c context.Context, ownerID int) (int, error)

	Update(c context.Context, dto *dto.Update) error
	UpdatePhoto(c context.Context, teamId int, photo string) error
	UpdatePermissions(c context.Context, dto *dto.UpdateMemberPermissions) error
	UpdateMember(c context.Context, dto *dto.UpdateMember) error
}
