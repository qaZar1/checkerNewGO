package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Impisigmatus/service_core/log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	Hostname string
	Port     uint64
	Database string
	User     string
	Password string
}

func NewPostgres(cfg Config) *sql.DB {
	const driver = "pgx"

	pattern := fmt.Sprintf(
		"host=%s port=%d database=%s user=%s password=%s sslmode=disable",
		cfg.Hostname, cfg.Port, cfg.Database, cfg.User, cfg.Password,
	)

	config, err := pgx.ParseConfig(pattern)
	if err != nil {
		log.Panicf("Invalid postgres config: %s", err)
	}
	config.Logger = &pgxLogger{}

	connection := stdlib.RegisterConnConfig(config)
	db, err := sql.Open(driver, connection)
	if err != nil {
		log.Panicf("Invalid postgres connect: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Panicf("Invalid postgres ping: %s", err)
	}

	return db
}
