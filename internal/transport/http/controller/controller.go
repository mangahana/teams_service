package controller

import (
	"teams_service/internal/application"
	"teams_service/internal/core/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	validator *validator.Validate
	useCase   application.UseCase
}

func New(useCase application.UseCase) *controller {
	return &controller{
		validator: validator.New(),
		useCase:   useCase,
	}
}

func (h *controller) getSession(c echo.Context) models.Session {
	session, ok := c.Get("user").(models.Session)
	if ok {
		return session
	}
	return models.Session{}
}
