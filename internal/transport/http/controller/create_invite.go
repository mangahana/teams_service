package controller

import (
	"log"
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) CreateInvite(c echo.Context) error {
	var dto *dto.CreateInvite
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	dto.OwnerId = h.getSession(c).UserID

	err := h.useCase.CreateInvite(c.Request().Context(), dto)
	if err != nil {
		log.Println(err)
		return c.JSON(400, err)
	}

	return c.String(200, "OK")
}
