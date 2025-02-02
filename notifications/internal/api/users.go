package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/checkerNewGO/notifications/internal/models"
)

type APIUsers struct {
	client *resty.Client
}

func NewAPIUsers(url string) *APIUsers {
	return &APIUsers{
		client: resty.New().SetBaseURL(url).SetTimeout(1*time.Minute).SetBasicAuth("dev", "test").SetDisableWarn(true),
	}
}

func (APIUsers *APIUsers) GetAllUsers() ([]models.User, error) {
	const endpoint = "/users/get"

	resp, err := APIUsers.client.R().Get(endpoint)
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

func (APIUsers *APIUsers) AddUser(user models.User) (bool, error) {
	const endpoint = "/users/addUsers"

	resp, err := APIUsers.client.R().SetBody(user).Post(endpoint)
	if err != nil {
		return false, err
	}

	if len(resp.Body()) == 0 {
		return true, nil
	} else {
		return false, err
	}
}

func (APIUsers *APIUsers) GetUserByChatID(chatId int64) (bool, error) {
	const endpoint = "/users/get-%d"

	resp, err := APIUsers.client.R().Get(fmt.Sprintf(endpoint, chatId))
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

func (APIUsers *APIUsers) DeleteUser(chatId int64) (bool, error) {
	const endpoint = "/users/delete-%d"

	resp, err := APIUsers.client.R().Get(fmt.Sprintf(endpoint, chatId))
	if err != nil {
		return false, err
	}

	if len(resp.Body()) == 0 {
		fmt.Println("Данные удалены")
		return true, nil
	}

	return true, nil
}
