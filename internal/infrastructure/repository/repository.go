package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type IRepo interface{}

type repo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) IRepo {
	return &repo{db: db}
}
