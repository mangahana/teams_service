package application

import (
	"teams_service/internal/infrastructure/repository"
)

type UseCase interface{}

type useCase struct {
	repo repository.IRepo
}

func New(repo repository.IRepo) UseCase {
	return &useCase{
		repo: repo,
	}
}
