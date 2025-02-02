package models

type Version struct {
	Version        string `db:"version"`
	Description    string `db:"description"`
	Description_RU string `db:"description_ru"`
	ReleaseDate    string `db:"release_date"`
} // @name version
