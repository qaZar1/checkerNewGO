package api

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/checkerNewGO/cron/internal/models"
)

type APIVersions struct {
	client *resty.Client
}

func NewAPIVersions(url string) *APIVersions {
	return &APIVersions{
		client: resty.New().SetBaseURL(url).SetTimeout(1*time.Minute).SetBasicAuth("dev", "test").SetDisableWarn(true),
	}
}

func (APIVersions *APIVersions) GetAllVersions() ([]models.Version, error) {
	const endpoint = "/versions/get"

	resp, err := APIVersions.client.R().Get(endpoint)
	if err != nil {
		return nil, err
	}

	resp.Body()
	allVersions := []models.Version{}
	if err := jsoniter.Unmarshal(resp.Body(), &allVersions); err != nil {
		return nil, err
	}

	return allVersions, nil
}

func (APIVersions *APIVersions) AddVersion(version models.Version) (bool, error) {
	const endpoint = "/versions/addVersion"

	resp, err := APIVersions.client.R().SetBody(version).Post(endpoint)
	if err != nil {
		return false, err
	}

	if len(resp.Body()) == 0 {
		return true, nil
	} else {
		return false, err
	}
}

func (APIVersions *APIVersions) GetVersion(version string) (bool, error) {
	const endpoint = "/versions/get-%s"

	resp, err := APIVersions.client.R().Get(fmt.Sprintf(endpoint, version))
	if err != nil {
		return false, err
	}

	if string(resp.Body()) == "Не удалось получить пользователей" {
		return false, nil
	}

	versionModel := models.Version{}
	if err := jsoniter.Unmarshal(resp.Body(), &versionModel); err != nil {
		return false, err
	}

	return true, nil
}

func (APIVersions *APIVersions) DeleteVersion(version string) (bool, error) {
	const endpoint = "/versions/delete-%s"

	resp, err := APIVersions.client.R().Get(fmt.Sprintf(endpoint, version))
	if err != nil {
		return false, err
	}

	if len(resp.Body()) == 0 {
		fmt.Println("Данные удалены")
		return true, nil
	}

	return true, nil
}
