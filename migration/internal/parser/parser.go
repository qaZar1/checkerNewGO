package parser

import (
	"bytes"

	"github.com/antchfx/htmlquery"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/checkerNewGO/migration/internal/database"
	"github.com/qaZar1/checkerNewGO/migration/internal/models"
	"github.com/qaZar1/checkerNewGO/migration/internal/site"
)

type Parser struct {
	url string
	api *site.API
	db  database.Database
}

func NewSite(url string, db *sqlx.DB) *Parser {
	return &Parser{
		url: url,
		api: site.NewAPI(),
		db:  *database.NewDatabase(db),
	}
}

func (parser *Parser) ParseReleases() error {
	data, err := parser.api.MakeRequest(parser.url)
	if err != nil {
		return err
	}

	buff := bytes.NewBuffer(data)

	html, err := htmlquery.Parse(buff)
	if err != nil {
		return err
	}

	versions, err := parser.parseReleases(html)
	if err != nil {
		return err
	}

	for _, version := range versions {
		desc, err := parser.parseDescription(html, version)
		if err != nil {
			return err
		}

		date, err := parser.parseDateRelease(html, version)
		if err != nil {
			return err
		}

		version := models.Version{
			Version:        version,
			Description:    desc,
			Description_RU: "",
			ReleaseDate:    date,
		}

		parser.db.AddVersion(version)
	}

	return nil
}
