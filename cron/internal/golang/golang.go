package golang

import (
	"bytes"
	"sync"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
)

type API struct {
	client *resty.Request
	mu     *sync.Mutex
}

func NewAPI() *API {
	return &API{
		client: resty.New().R(),
		mu:     &sync.Mutex{},
	}
}

func (api *API) MakeRequest(url string) (*html.Node, error) {
	api.mu.Lock()

	defer api.mu.Unlock()

	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}

	body := resp.Body()
	buff := bytes.NewBuffer(body)

	html, err := htmlquery.Parse(buff)
	if err != nil {
		return nil, err
	}

	time.Sleep(1 * time.Second)
	return html, nil
}
