package http

import (
	"github.com/labstack/echo/v4"
)

func (h *server) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		user, err := h.authService.GetUser(c.Request().Context(), token)
		if err != nil {
			return c.String(403, "")
		}

		c.Set("user", user)

		return next(c)
	}
}
