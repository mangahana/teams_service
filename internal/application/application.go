package application

import (
	"context"
	"teams_service/internal/core/dto"
	"teams_service/internal/core/models"
	"teams_service/internal/infrastructure/repository"

	"github.com/minio/minio-go/v7"
)

type UseCase interface {
	Add(c context.Context, dto *dto.AddTeam) error

	GetOne(c context.Context, teamId int) (models.OneTeam, error)
	GetMembers(c context.Context, teamID int) ([]models.Member, error)

	Update(c context.Context, dto *dto.Update) error
	UploadPhoto(c context.Context, dto *dto.UploadPhoto) (string, error)
}

type useCase struct {
	repo repository.IRepo
	s3   *minio.Client
}

func New(repo repository.IRepo, s3 *minio.Client) UseCase {
	return &useCase{
		repo: repo,
		s3:   s3,
	}
}
