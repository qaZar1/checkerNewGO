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
	"github.com/qaZar1/checkerNewGO/users/autogen"
)

type Transport struct {
	app *Application
}

func NewTransport(db *sqlx.DB) autogen.ServerInterface {
	return &Transport{
		app: NewApplication(db),
	}
}

func (transport *Transport) PostApiUsersAdd(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid read body: %s", err), "Не удалось прочитать тело запроса")
		return
	}

	var user autogen.User
	if err := jsoniter.Unmarshal(data, &user); err != nil {
		utils.WriteString(w, http.StatusBadRequest, fmt.Errorf("Invalid parse body: %s", err), "Не удалось распарсить тело запроса формата JSON")
		return
	}

	if err := transport.app.AddUser(user); err != nil {
		utils.WriteString(w, http.StatusInternalServerError, err, "Не удалось добавить пользователя")
		return
	}

	utils.WriteNoContent(w)
}

func (transport *Transport) DeleteApiUsersChatIdRemove(w http.ResponseWriter, r *http.Request, chatId int64) {
	ok, err := transport.app.RemoveUser(chatId)
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

func (transport *Transport) GetApiUsersChatIdGet(w http.ResponseWriter, r *http.Request, chatId int64) {
	user, err := transport.app.GetUserByChatID(chatId)
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

func (transport *Transport) GetApiUsersGet(w http.ResponseWriter, r *http.Request) {
	users, err := transport.app.GetAllUsers()
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
