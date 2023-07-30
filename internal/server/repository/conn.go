package repository

import (
	"github.com/jmoiron/sqlx"
)

func NewPGRepository(conn *sqlx.DB) *pgRepository {

	return &pgRepository{
		conn: conn,
	}
}

type pgRepository struct {
	conn *sqlx.DB
}
