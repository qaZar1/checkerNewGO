package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/checkerNewGO/migration/internal/models"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) AddVersion(version models.Version) error {
	const query = `
INSERT INTO main.versions (
	version,
	description,
	description_ru,
	release_date
) VALUES (
	:version,
	:description,
	:description_ru,
	:releasedate
) ON CONFLICT (version) DO NOTHING;`

	if _, err := pg.db.NamedExec(query, version); err != nil {
		return fmt.Errorf("Invalid INSERT INTO main.versions: %s", err)
	}

	return nil
}
