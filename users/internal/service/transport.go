package service

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Impisigmatus/service_core/utils"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"github.com/qaZar1/checkerNewGO/users/autogen/server"
	"github.com/qaZar1/checkerNewGO/users/internal/models"
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
// @Router /api/users/addUsers [post]
// @Summary Добавление юзера в БД
// @Description При обращении, добавляется отклик в БД по телу запроса
//
// @Tags APIs
// @Accept       application/json
// @Produce      application/json
// @Param 	request	body	user	true	"Тело запроса"
//
// @Success 200 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) PostApiUsersAddUsers(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var user models.User
	if err := jsoniter.Unmarshal(data, &user); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	if err := transport.srv.AddUser(user); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить пользователя")
		return
	}

	utils.WriteNoContent(w)
}

// Set godoc
//
// @Router /api/users/delete-{chat_id} [delete]
// @Summary Удаление юзера из БД
// @Description При обращении, удаляет юзера из БД по его chat_id
//
// @Tags APIs
// @Produce      application/json
// @Param	chat_id	path	int	true	"Chat_id пользователя"
//
// @Success 200 {object} nil "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) DeleteApiUsersDeleteChatId(w http.ResponseWriter, r *http.Request, chatId int) {
	ok, err := transport.srv.RemoveUser(int64(chatId))
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Пользователя не существует")
		return
	}

	if ok {
		utils.WriteString(w, http.StatusOK, nil, "Пользователя не существует")
		return
	} else {
		utils.WriteNoContent(w)
		return
	}
}

// Set godoc
//
// @Router /api/users/get-{chat_id} [get]
// @Summary Получение юзера по его chat_id
// @Description При обращении, возвращает юзера по его chat_id
//
// @Tags APIs
// @Produce      application/json
// @Param	chat_id	path	int	true	"Chat_id пользователя"
//
// @Success 200 {object} user "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiUsersGetChatId(w http.ResponseWriter, r *http.Request, chatId int) {
	user, err := transport.srv.GetUserByChatID(int64(chatId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteNoContent(w)
			return
		}

		utils.WriteString(w, http.StatusNoContent, err, "Не удалось получить пользователя")
		return
	}

	utils.WriteObject(w, user)
}

// Set godoc
//
// @Router /api/users/get [get]
// @Summary Получение всех юзеров
// @Description При обращении, возвращает всех юзеров
//
// @Tags APIs
// @Produce      application/json
//
// @Success 200 {array} user "Запрос выполнен успешно"
// @Failure 400 {object} nil "Ошибка валидации данных"
// @Failure 401 {object} nil "Ошибка авторизации"
// @Failure 500 {object} nil "Произошла внутренняя ошибка сервера"
func (transport *Transport) GetApiUsersGet(w http.ResponseWriter, r *http.Request) {
	users, err := transport.srv.GetAllUsers()
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось получить пользователей")
		return
	}
	if len(users) == 0 {
		utils.WriteString(w, http.StatusInternalServerError, err, "В базе нет пользователей")
		return
	}

	utils.WriteObject(w, users)
}
