package repository

import (
	"context"
	"teams_service/internal/core/dto"
	"teams_service/internal/core/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IRepo interface {
	Add(c context.Context, dto *dto.AddTeam) error

	GetOne(c context.Context, teamId int) (models.OneTeam, error)
	GetTypeByID(c context.Context, typeId int) (models.TeamType, error)
	GetMembers(c context.Context, teamID int) ([]models.Member, error)

	TeamsCountForOwner(c context.Context, ownerID int) (int, error)
}

type repo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) IRepo {
	return &repo{db: db}
}
