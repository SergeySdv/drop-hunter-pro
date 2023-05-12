package pg

import (
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"os"
	"path"
	"time"
)

type TimeTamps struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

func NewPGConnection(connString string) (*sql.DB, error) {

	var err error
	const postgresDriverName = "pgx"

	driverName, err := ocsql.Register(postgresDriverName, ocsql.WithOptions(ocsql.TraceOptions{
		AllowRoot:         false,
		Ping:              false,
		RowsNext:          true,
		RowsClose:         true,
		RowsAffected:      false,
		LastInsertID:      false,
		Query:             true,
		QueryParams:       true,
		DefaultAttributes: nil,
	}))
	if err != nil {
		return nil, err
	}

	sqlDB, err := sql.Open(driverName, connString)
	if err != nil {
		return nil, err
	}

	return sqlDB, nil
}

func WrapPgConnWithSqlx(conn *sql.DB) (*sqlx.DB, error) {
	const postgresDriverName = "pgx"
	sqlxDB := sqlx.NewDb(conn, postgresDriverName)

	if err := sqlxDB.Ping(); err != nil {
		return nil, err
	}
	return sqlxDB, nil
}

func RunMigrations(pg *sql.DB, subPath string, force bool) error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	migrationDir := path.Join(dir, subPath)

	if force {
		_ = goose.Down(pg, migrationDir)
	}

	currentMigrationVersion, err := goose.EnsureDBVersion(pg)
	if err != nil {
		return err
	}
	migrations, err := goose.CollectMigrations(migrationDir, currentMigrationVersion, 100)
	if err != nil {
		return err
	}

	for _, m := range migrations {
		err = m.Up(pg)
		if err != nil {
			return err
		}
	}

	return nil
}
