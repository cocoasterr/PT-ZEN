package models

import (
	"time"

	"github.com/google/uuid"
)

type PesananCreate struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Status    bool   `json:"status"`
	Qty       int    `json:"qty"`
}

type Pesanan struct {
	Id string `json:"id"`
	PesananCreate
	DeletedAt time.Time `json:"deleted_at"`
}

func (m *Pesanan) Model() interface{} {
	return &Pesanan{}
}

func (m *Pesanan) ModelCreate(payload map[string]interface{}, userName string) map[string]interface{} {
	payload["id"] = uuid.New().String()
	return payload
}

func (m *Pesanan) ModelUpdate(payload map[string]interface{}, userName string) map[string]interface{} {
	return payload
}

func (m *Pesanan) ModelSoftDel(payload map[string]interface{}) map[string]interface{} {
	payload["deleted_at"] = time.Now()
	return payload
}

func (m *Pesanan) TbName() string {
	return "Pesanan"
}
