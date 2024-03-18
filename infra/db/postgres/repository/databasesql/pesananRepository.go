package PGRepository

import (
	"database/sql"
	"fmt"

	"github.com/cocoasterr/kalbe_test/models"
)

type PesananRepository struct {
	Repository
}

func NewPesananRepository(db *sql.DB) *PesananRepository {
	return &PesananRepository{
		Repository{
			Db:    db,
			Model: &models.Pesanan{},
		},
	}
}

func (mp *PesananRepository) CobaDulu() {
	fmt.Println()
}
