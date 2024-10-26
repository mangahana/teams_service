package controller

import "github.com/labstack/echo/v4"

func (h *controller) MyTeams(c echo.Context) error {
	user := h.getUser(c)
	teams, err := h.useCase.GetMyTeams(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, teams)
}
