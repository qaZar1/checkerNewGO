package site

import (
	"sync"

	"github.com/go-resty/resty/v2"
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

func (api *API) MakeRequest(url string) ([]byte, error) {
	api.mu.Lock()

	defer api.mu.Unlock()

	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
