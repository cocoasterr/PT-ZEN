package models

import (
	"time"

	"github.com/google/uuid"
)

type MasterLocation struct {
	Location_id   string    `json:"location_id"`
	Location_name string    `json:"location_name"`
	Created_at    time.Time `json:"created_at"`
	Created_by    string    `json:"created_by"`
	Updated_at    time.Time `json:"updated_at"`
	Updated_by    string    `json:"updated_by"`
	Deleted_at    time.Time `json:"deleted_at"`
}

func (m *MasterLocation) Model() interface{} {
	return &MasterLocation{}
}

func (m *MasterLocation) ModelCreate(payload map[string]interface{}, userName string) map[string]interface{} {
	payload["id"] = uuid.New().String()
	payload["created_at"] = time.Now()
	payload["updated_at"] = time.Now()
	return payload
}

func (m *MasterLocation) ModelUpdate(payload map[string]interface{}, userName string) map[string]interface{} {
	if payload["location_name"] != "" {
		payload["updated_at"] = time.Now()
		return payload
	}
	return nil
}

func (m *MasterLocation) ModelSoftDel(payload map[string]interface{}) map[string]interface{} {
	payload["deleted_at"] = time.Now()
	return payload
}

func (m *MasterLocation) TbName() string {
	return "MasterLocation"
}
