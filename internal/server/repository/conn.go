package repository

import (
	"github.com/jmoiron/sqlx"
)

func NewPGRepository(conn *sqlx.DB) *PGRepository {
	return &PGRepository{
		conn: conn,
	}
}

type PGRepository struct {
	conn *sqlx.DB
}
