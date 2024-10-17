package http

import (
	"context"
	"teams_service/internal/application"
	"teams_service/internal/infrastructure/auth_service"
	"teams_service/internal/transport/http/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	server      *echo.Echo
	useCase     application.UseCase
	authService *auth_service.AuthService
}

func New(useCase application.UseCase, authService *auth_service.AuthService) *server {
	return &server{
		server:      echo.New(),
		useCase:     useCase,
		authService: authService,
	}
}

func (h *server) ListenAndServe(socket string) {
	h.server.Start(socket)
}

func (h *server) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}

func (h *server) Register() {
	h.server.Use(middleware.CORS())
	h.server.Use(middleware.BodyLimit("10M"))

	controller := controller.New(h.useCase)

	api := h.server.Group("/api/v1/teams")

	api.GET("", controller.GetOne)
	api.GET("/members", controller.GetMembers)

	private := api.Group("", h.Authenticate)
	private.POST("/add", controller.Add)
}
