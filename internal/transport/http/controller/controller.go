package controller

import (
	"teams_service/internal/application"
)

type controller struct {
	useCase application.UseCase
}

func New(useCase application.UseCase) *controller {
	return &controller{
		useCase: useCase,
	}
}
