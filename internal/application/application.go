package application

import (
	"context"
	"teams_service/internal/core/dto"
	"teams_service/internal/core/models"
	"teams_service/internal/infrastructure"
)

type UseCase interface {
	Add(c context.Context, user *models.User, dto *dto.AddTeam) (int, error)
	CreateInvite(c context.Context, dto *dto.CreateInvite) error

	IsMember(c context.Context, user *models.User, teamId int) error

	GetOne(c context.Context, teamId int) (models.OneTeam, error)
	GetMembers(c context.Context, teamID int) ([]models.Member, error)
	GetMyTeams(c context.Context, user *models.User) ([]models.Team, error)

	Update(c context.Context, user *models.User, dto *dto.Update) error
	UploadPhoto(c context.Context, user *models.User, dto *dto.UploadPhoto) (string, error)
	UpdateMemberPermissions(c context.Context, dto *dto.UpdateMemberPermissions) error
	UpdateMember(c context.Context, dto *dto.UpdateMember) error
}

type useCase struct {
	repo infrastructure.Repository
	s3   infrastructure.S3
}

func New(repo infrastructure.Repository, s3 infrastructure.S3) UseCase {
	return &useCase{
		repo: repo,
		s3:   s3,
	}
}
