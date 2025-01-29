package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/checkerNewGO/versions/internal/models"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) GetAllVersions() ([]models.Version, error) {
	const query = "SELECT version, description, description_ru FROM main.versions;"

	var versions []models.Version
	if err := pg.db.Select(&versions, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.versions: %s", err)
	}

	return versions, nil
}

func (pg *Database) GetVersionByID(version_str string) (models.Version, error) {
	const query = "SELECT version, description, description_ru FROM main.versions WHERE version = $1;"

	var version models.Version
	if err := pg.db.Get(&version, query, version_str); err != nil {
		return models.Version{}, fmt.Errorf("description does not exist in main.versions: %w", err)
	}

	return version, nil
}

func (pg *Database) AddVersion(version models.Version) error {
	const query = `
INSERT INTO main.versions (
	version,
	description,
	description_ru
) VALUES (
	:version,
	:description,
	:description_ru
) ON CONFLICT (version) DO NOTHING;`

	if _, err := pg.db.NamedExec(query, version); err != nil {
		return fmt.Errorf("Invalid INSERT INTO main.versions: %s", err)
	}

	return nil
}

func (pg *Database) RemoveVersion(version_str string) (bool, error) {
	const query = "DELETE FROM main.versions WHERE version = $1"

	exec, err := pg.db.Exec(query, version_str)
	if err != nil {
		return false, fmt.Errorf("Invalid DELETE main.versions: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected description: %s", err)
	}

	return affected == 0, nil
}

func (pg *Database) UpdateVersion(version models.Version) (bool, error) {
	const query = `
UPDATE main.versions
SET description = :description,
	description_ru = :description_ru
WHERE version = :version;`

	exec, err := pg.db.NamedExec(query, version)
	if err != nil {
		return false, fmt.Errorf("Invalid UPDATE main.versions: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected description: %s", err)
	}

	return affected == 1, nil
}
