package controller

import (
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) Update(c echo.Context) error {
	var dto *dto.Update
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	err := h.useCase.Update(c.Request().Context(), dto)
	if err != nil {
		return c.JSON(400, err)
	}

	return nil
}
