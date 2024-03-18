package services

import (
	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/models"
)

type ProductService struct {
	BaseService
	Repository *PGRepository.ProductRepository
}

func NewProductService(pgRepo PGRepository.Repository, repo *PGRepository.ProductRepository) *ProductService {
	return &ProductService{
		Repository: repo,
		BaseService: BaseService{
			Repo:  pgRepo,
			Model: &models.Product{},
		},
	}
}
