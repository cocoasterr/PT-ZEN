package PGRepository

import (
	"database/sql"

	"github.com/cocoasterr/kalbe_test/models"
)

type UserRepository struct {
	Repository
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Repository{
			Db:       db,
			Model:    &models.User{},
		},
	}
}
