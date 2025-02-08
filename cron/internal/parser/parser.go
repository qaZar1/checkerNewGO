package parser

import (
	"bytes"
	"fmt"
	"os"

	"github.com/antchfx/htmlquery"
	"github.com/qaZar1/checkerNewGO/cron/internal/models"
	"github.com/qaZar1/checkerNewGO/cron/internal/site"
)

type Parser struct {
	url string
	api *site.API
}

func NewSite(url string) *Parser {
	return &Parser{
		url: url,
		api: site.NewAPI(),
	}
}

func (parser *Parser) ParseReleases() error {
	// data, err := parser.api.MakeRequest(parser.url)
	// if err != nil {
	// 	return err
	// }

	data, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Errorf("%s", err)
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

	releases := make([]models.Version, 0)
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

		releases = append(releases, version)
	}

	return nil
}
