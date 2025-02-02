package parser

import (
	"github.com/qaZar1/checkerNewGO/cron/internal/api"
	"github.com/qaZar1/checkerNewGO/cron/internal/golang"
)

type Parser struct {
	url         string
	api         *golang.API
	apiVersions *api.APIVersions
}

func NewSite(url string, urlVersions string) *Parser {
	return &Parser{
		url:         url,
		api:         golang.NewAPI(),
		apiVersions: api.NewAPIVersions(urlVersions),
	}
}
