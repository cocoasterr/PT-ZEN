package services

import (
	"context"
	"fmt"

	PGRepository "github.com/cocoasterr/kalbe_test/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/kalbe_test/models"
)

type PesananService struct {
	BaseService
	Repository     *PGRepository.PesananRepository
	ProductService *ProductService
}

func NewPesananService(pgRepo PGRepository.Repository, repo *PGRepository.PesananRepository, productService *ProductService) *PesananService {
	return &PesananService{
		Repository:     repo,
		ProductService: productService,
		BaseService: BaseService{
			Repo:  pgRepo,
			Model: &models.Pesanan{},
		},
	}
}

func (s *PesananService) FindProduct(ctx context.Context, id string) (map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM PRODUCT WHERE id='%s'", id)
	return s.Repo.FindCustomQuery(ctx, query)
}
