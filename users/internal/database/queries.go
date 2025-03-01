package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/checkerNewGO/users/internal/models"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) GetAllUsers() ([]models.User, error) {
	const query = "SELECT chat_id, username, name FROM main.users;"

	var users []models.User

	if err := pg.db.Select(&users, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.users: %s", err)
	}

	return users, nil
}

func (pg *Database) GetUserByChatID(chatId int64) (models.User, error) {
	const query = "SELECT chat_id, username, name FROM main.users WHERE chat_id = $1;"

	var user models.User
	if err := pg.db.Get(&user, query, chatId); err != nil {
		return models.User{}, fmt.Errorf("User does not exist in main.users: %w", err)
	}

	return user, nil
}

func (pg *Database) AddUser(user models.User) error {
	const query = `
INSERT INTO main.users (
	chat_id,
	username,
	name
) VALUES (
	:chat_id,
	:username,
	:name
) ON CONFLICT (chat_id) DO NOTHING;`

	if _, err := pg.db.NamedExec(query, user); err != nil {
		return fmt.Errorf("Invalid INSERT INTO main.users: %s", err)
	}

	return nil
}

func (pg *Database) RemoveUser(chatId int64) (bool, error) {
	const query = "DELETE FROM main.users WHERE chat_id = $1"

	exec, err := pg.db.Exec(query, chatId)
	if err != nil {
		return false, fmt.Errorf("Invalid DELETE main.users: %s", err)
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid affected user: %s", err)
	}

	return affected == 0, nil
}
