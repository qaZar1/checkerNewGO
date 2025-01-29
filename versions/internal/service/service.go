package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/checkerNewGO/versions/internal/database"
	"github.com/qaZar1/checkerNewGO/versions/internal/models"
)

type Service struct {
	db database.Database
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: *database.NewDatabase(db),
	}
}

func (srv *Service) GetAllVersions() ([]models.Version, error) {
	return srv.db.GetAllVersions()
}

func (srv *Service) GetVersionByID(version string) (models.Version, error) {
	return srv.db.GetVersionByID(version)
}

func (srv *Service) AddVersion(info models.Version) error {
	return srv.db.AddVersion(info)
}

func (srv *Service) RemoveVersion(version string) (bool, error) {
	return srv.db.RemoveVersion(version)
}

func (srv *Service) UpdateVersion(info models.Version) (bool, error) {
	return srv.db.UpdateVersion(info)
}
