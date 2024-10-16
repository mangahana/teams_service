package http

import (
	"github.com/labstack/echo/v4"
)

func (h *server) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		session, err := h.authService.GetSession(c.Request().Context(), token)
		if err != nil {
			return c.String(403, "")
		}

		c.Set("user", session)

		return next(c)
	}
}
