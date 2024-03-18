package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductCreate struct {
	Name  string `json:"name"`
	Harga int    `json:"harga"`
	Stock int    `json:"stock"`
}
type Product struct {
	Id string `json:"id"`
	ProductCreate
	Created_at time.Time `json:"created_at"`
	Created_by string    `json:"created_by"`
	Updated_at time.Time `json:"updated_at"`
	Updated_by string    `json:"updated_by"`
	Deleted_at time.Time `json:"deleted_at"`
}

func (m *Product) Model() interface{} {
	return &Product{}
}

func (m *Product) ModelCreate(payload map[string]interface{}, userName string) map[string]interface{} {
	payload["id"] = uuid.New().String()
	payload["created_at"] = time.Now()
	payload["created_by"] = userName
	payload["updated_at"] = time.Now()
	payload["updated_by"] = userName
	return payload
}

func (m *Product) ModelUpdate(payload map[string]interface{}, userName string) map[string]interface{} {
	payload["updated_at"] = time.Now()
	payload["updated_by"] = userName
	return payload
}

func (m *Product) ModelSoftDel(payload map[string]interface{}) map[string]interface{} {
	payload["deleted_at"] = time.Now()
	return payload
}

func (m *Product) TbName() string {
	return "Product"
}
