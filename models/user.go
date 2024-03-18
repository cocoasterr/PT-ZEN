package models

import (
	"github.com/google/uuid"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RespUser struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type User struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phone_number"`
	LoginUser
}

func (m *User) Model() interface{} {
	return &User{}
}

func (m *User) ModelCreate(payload map[string]interface{}, userName string) map[string]interface{} {
	payload["id"] = uuid.New().String()
	return payload
}

func (m *User) ModelUpdate(payload map[string]interface{}, userName string) map[string]interface{} {

	return payload
}
func (m *User) ModelSoftDel(payload map[string]interface{}) map[string]interface{} {

	return nil
}
func (m *User) TbName() string {
	return "users"
}
