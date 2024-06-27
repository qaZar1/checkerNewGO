package service

import (
	"fmt"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/checkerNewGO/users/autogen"
)

type Application struct {
	db *sqlx.DB
}

func NewApplication(db *sqlx.DB) *Application {
	return &Application{db: db}
}

func (app *Application) GetAllUsers() ([]autogen.Info, error) {
	const query = "SELECT chat_id, username, name FROM main.users;"

	var users []autogen.Info
	if err := app.db.Select(&users, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.users: %s", err)
	}

	return users, nil
}

func (app *Application) GetUserByChatID(chatId int64) (*autogen.Info, error) {
	const query = "SELECT chat_id, username, name FROM main.users WHERE chat_id = $1;"

	var users autogen.Info
	if err := app.db.Get(&users, query, chatId); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.users: %s", err)
	}

	return &users, nil
}

func (app *Application) AddUser(user autogen.User) error {
	const query = `
INSERT INTO main.users (
	chat_id,
	username,
	name
) VALUES (
	:chatid,
	:username,
	:name
) ON CONFLICT (chat_id) DO NOTHING;`

	if _, err := app.db.NamedExec(query, user); err != nil {
		return fmt.Errorf("Invalid INSERT INTO main.users: %s", err)
	}

	return nil
}

func (app *Application) RemoveUser(chatId int64) (bool, error) {
	const query = "DELETE FROM main.users WHERE chat_id = $1"

	exec, err := app.db.Exec(query, chatId)
	if err != nil {
		return false, fmt.Errorf("Invalid DELETE main.users: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected user: %s", err)
	}

	return affected == 0, nil
}
