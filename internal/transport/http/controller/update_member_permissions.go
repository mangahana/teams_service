package controller

import (
	"teams_service/internal/core/cerror"
	"teams_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) UpdateMemberPermissions(c echo.Context) error {
	var dto *dto.UpdateMemberPermissions
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	dto.MemberId = h.getSession(c).UserID

	err := h.useCase.UpdateMemberPermissions(c.Request().Context(), dto)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.String(200, "OK")
}
