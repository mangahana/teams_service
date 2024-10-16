package controller

import (
	"errors"
	"strconv"
	"teams_service/internal/core/cerror"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func (h *controller) GetOne(c echo.Context) error {
	queryTeamId := c.QueryParam("id")

	teamId, err := strconv.Atoi(queryTeamId)
	if err != nil {
		return c.JSON(400, cerror.BadRequest())
	}

	team, err := h.useCase.GetOne(c.Request().Context(), teamId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.JSON(404, cerror.NotFound())
		}
		return c.JSON(500, cerror.InternalServerError())
	}

	return c.JSON(200, team)
}
