package entity

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskMessage struct {
	Id    primitive.ObjectID `json:"id"`
	Key   string             `json:"key"`
	Cmd   string             `json:"cmd"`
	Param string             `json:"param"`
}

func (m *TaskMessage) ToString() (string, error) {
	data, err := json.Marshal(&m)
	if err != nil {
		return "", err
	}
	return string(data), err
}
