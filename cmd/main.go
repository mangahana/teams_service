package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"teams_service/internal/application"
	"teams_service/internal/core/configuration"
	authservice "teams_service/internal/infrastructure/auth_service"
	"teams_service/internal/infrastructure/postgresql"
	"teams_service/internal/infrastructure/repository"
	"teams_service/internal/infrastructure/s3"
	"teams_service/internal/transport/grpc"
	"teams_service/internal/transport/http"
	"time"
)

func main() {
	cfg, err := configuration.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgresql.New(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	authService, err := authservice.New(cfg.Services.AuthServiceSocket)
	if err != nil {
		log.Fatal(err)
	}

	s3Client, err := s3.New(&cfg.S3)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)
	useCase := application.New(repo, s3Client)

	httpServer := http.New(useCase, authService)
	httpServer.Register()

	go httpServer.ListenAndServe(cfg.Server.HttpSocket)

	grpsServer := grpc.New(useCase)
	go grpsServer.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdown()

	httpServer.Shutdown(ctx)
	db.Close()
}
