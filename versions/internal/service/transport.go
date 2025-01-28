package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/checkerNewGO/versions/autogen/server"
	"github.com/qaZar1/checkerNewGO/versions/internal/models"
)

type Transport struct {
	srv *Service
}

func NewTransport(db *sqlx.DB) server.ServerInterface {
	return &Transport{
		srv: NewService(db),
	}
}

// Set godoc
//
// @Router /api/versions/addVersion [post]
// @Summary Добавление версии в БД
// @Description При обращении, добавляется отклик в БД по телу запроса
//
// @Tags APIs
// @Accept       application/json
// @Param 	request	body	version	true	"Тело запроса"
//
// @Success 200 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) PostApiVersionsAddVersion(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var info models.Version
	if err := jsoniter.Unmarshal(data, &info); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	if err := transport.srv.AddVersion(info); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить версию Go")
		return
	}

	utils.WriteNoContent(w)
}

// Set godoc
//
// @Router /api/versions/get [get]
// @Summary Получение всех версий
// @Description При обращении, возвращает все версии
//
// @Tags APIs
// @Produce      application/json
//
// @Success 200 {array}	version "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiVersionsGet(w http.ResponseWriter, r *http.Request) {
	versions, err := transport.srv.GetAllVersions()
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить пользователей")
		return
	}
	if len(versions) == 0 {
		utils.WriteString(w, http.StatusInternalServerError, err, "В базе нет версий Go")
		return
	}

	utils.WriteObject(w, versions)
}

// Set godoc
//
// @Router /api/versions/get-{version} [get]
// @Summary Получение версии по ее номеру
// @Description При обращении, возвращает версию
//
// @Tags APIs
// @Produce      application/json
// @Param	version	path	string	true	"Версия"
//
// @Success 200 {object} version "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiVersionsGetVersion(w http.ResponseWriter, r *http.Request, version string) {
	info, err := transport.srv.GetVersionByID(version)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить пользователей")
		return
	}
	if info == (models.Version{}) {
		utils.WriteString(w, http.StatusInternalServerError, err, "В базе нет пользователей")
		return
	}

	utils.WriteObject(w, info)
}

// Set godoc
//
// @Router /api/versions/delete-{version} [delete]
// @Summary Удаление версии из БД
// @Description При обращении, удаляет версию из БД
//
// @Tags APIs
// @Produce      application/json
// @Param	version	path	string	true	"Версия обновления"
//
// @Success 200 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) DeleteApiVersionsDeleteVersion(w http.ResponseWriter, r *http.Request, version string) {
	ok, err := transport.srv.RemoveVersion(version)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Данных о версии не существует")
		return
	}

	if ok {
		utils.WriteString(w, http.StatusOK, nil, "Данных о версии не существует")
		return
	} else {
		utils.WriteNoContent(w)
		return
	}
}

// Set godoc
//
// @Router /api/versions/update-{version} [put]
// @Summary Обновление данных о версии
// @Description При обращении, обновляет данные о версии
//
// @Tags APIs
// @Produce      application/json
//
// @Param	version	path	string	true	"Версия"
// @Param	request	body	version	true	"Тело запроса"
//
// @Success 200 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) PutApiVersionsUpdateVersion(w http.ResponseWriter, r *http.Request, version string) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var info models.Version
	if err := jsoniter.Unmarshal(data, &info); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	info.Version = version

	ok, err := transport.srv.UpdateVersion(info)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось обновить данные версии")
	}

	if ok {
		utils.WriteString(w, http.StatusOK, nil, "Данные обновлены!")
		return
	} else {
		utils.WriteNoContent(w)
		return
	}
}
