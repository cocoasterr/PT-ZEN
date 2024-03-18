package models

type BaseModel interface {
	TbName() string
	Model() interface{}
	ModelCreate(payload map[string]interface{}, userName string) map[string]interface{}
	ModelUpdate(payload map[string]interface{}, userName string) map[string]interface{}
	ModelSoftDel(payload map[string]interface{}) map[string]interface{}
}
