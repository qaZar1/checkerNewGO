package models

type Version struct {
	Version        string `validate:"required"`
	Description    string `validate:"required"`
	Description_RU string `validate:"required"`
}
