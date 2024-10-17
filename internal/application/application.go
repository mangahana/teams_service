package application

import (
	"context"
	"teams_service/internal/core/dto"
	"teams_service/internal/core/models"
	"teams_service/internal/infrastructure/repository"
)

type UseCase interface {
	Add(c context.Context, dto *dto.AddTeam) error

	GetOne(c context.Context, teamId int) (models.OneTeam, error)
	GetMembers(c context.Context, teamID int) ([]models.Member, error)

	Update(c context.Context, dto *dto.Update) error
}

type useCase struct {
	repo repository.IRepo
}

func New(repo repository.IRepo) UseCase {
	return &useCase{
		repo: repo,
	}
}
