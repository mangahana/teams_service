package controller

import (
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) Add(c echo.Context) error {
	var dto dto.AddTeam
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	session := h.getSession(c)

	dto.OwnerId = session.UserID

	err := h.useCase.Add(c.Request().Context(), &dto)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "OK")
}
