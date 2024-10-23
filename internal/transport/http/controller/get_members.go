package controller

import (
	"log"
	"strconv"
	"teams_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *controller) GetMembers(c echo.Context) error {
	queryTeamID := c.QueryParam("id")

	teamId, err := strconv.Atoi(queryTeamID)
	if err != nil {
		log.Println("debug #1")
		return c.JSON(400, cerror.BadRequest())
	}

	members, err := h.useCase.GetMembers(c.Request().Context(), teamId)
	if err != nil {
		log.Println(err)
		return c.JSON(400, cerror.BadRequest())
	}

	return c.JSON(200, members)
}
