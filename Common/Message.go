package Common

import (
	"encoding/json"
	"log"
)

type Message struct {
	Code    int64
	Message string
	Data    interface{}
}

func (Message Message) Success(message string, data interface{}) []byte {
	Message.Code = 0
	Message.Message = message
	Message.Data = data
	jsonBytes, err := json.Marshal(Message)
	if err != nil {
		log.Fatal("序列号json错误")
	}
	return jsonBytes
}

func (Message Message) Error(message string, data interface{}) []byte {
	Message.Code = 1001
	Message.Message = message
	Message.Data = data
	jsonBytes, err := json.Marshal(Message)
	if err != nil {
		log.Fatal("序列号json错误")
	}
	return jsonBytes
}
