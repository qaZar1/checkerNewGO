package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/checkerNewGO/users/internal/database"
	"github.com/qaZar1/checkerNewGO/users/internal/models"
)

type Service struct {
	db database.Database
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: *database.NewDatabase(db),
	}
}

func (srv *Service) GetAllUsers() ([]models.User, error) {
	return srv.db.GetAllUsers()
}

func (srv *Service) GetUserByChatID(chatId int64) (models.User, error) {
	return srv.db.GetUserByChatID(chatId)
}

func (srv *Service) AddUser(user models.User) error {
	return srv.db.AddUser(user)
}

func (srv *Service) RemoveUser(chatId int64) (bool, error) {
	return srv.db.RemoveUser(chatId)
}
