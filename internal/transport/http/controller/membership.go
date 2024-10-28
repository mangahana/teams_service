package controller

import (
	"strconv"
	"teams_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *controller) Membership(c echo.Context) error {
	queryTeamId := c.QueryParam("id")

	teamId, err := strconv.Atoi(queryTeamId)
	if err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	user := h.getUser(c)

	if err := h.useCase.IsMember(c.Request().Context(), &user, teamId); err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	return c.JSON(200, true)
}
