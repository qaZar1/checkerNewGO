package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/checkerNewGO/notifications/internal/models"
)

type API struct {
	client *resty.Client
}

func NewApi(url string) *API {
	return &API{
		client: resty.New().SetBaseURL(url).SetTimeout(1*time.Minute).SetBasicAuth("dev", "test").SetDisableWarn(true),
	}
}

func (api *API) GetAllUsers() ([]models.User, error) {
	const endpoint = "/users"

	resp, err := api.client.R().Get(endpoint)
	if err != nil {
		return nil, err
	}

	resp.Body()
	allUsers := []models.User{}
	if err := jsoniter.Unmarshal(resp.Body(), &allUsers); err != nil {
		return nil, err
	}

	return allUsers, nil
}

func (api *API) AddUser(user models.User) (bool, error) {
	const endpoint = "/users/add"

	_, err := api.client.R().SetBody(user).Post(endpoint)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (api *API) GetUserByChatID(chatId int64) (bool, error) {
	const endpoint = "/users/%d/get"

	resp, err := api.client.R().Get(fmt.Sprintf(endpoint, chatId))
	if err != nil {
		return false, err
	}

	if len(resp.Body()) == 0 {
		return false, nil
	}

	user := models.User{}
	if err := jsoniter.Unmarshal(resp.Body(), &user); err != nil {
		return false, err
	}

	return true, nil
}
