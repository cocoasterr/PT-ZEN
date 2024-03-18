package services

import (
	"context"
	"fmt"

	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/models"
)

type UserService struct {
	BaseService
}

func NewUserService(mutation PGRepository.Repository) *UserService {
	return &UserService{
		BaseService{
			Repo:  mutation,
			Model: &models.User{},
		},
	}
}

func (s *UserService) FindUser(ctx context.Context, username string) (map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", username)
	datares, err := s.Repo.FindCustumQuery(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(datares) != 1 {
		return nil, err
	}
	return datares[0], nil
}
