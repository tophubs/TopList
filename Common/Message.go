package Common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Code    int64
	Message string
	Data    interface{}
}

func (Message Message) Success(message string, data interface{}, w http.ResponseWriter) {
	Message.Code = 0
	Message.Message = message
	Message.Data = data
	jsonStr, err := json.Marshal(Message)
	if err != nil {
		log.Fatal("序列化json错误")
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%s", string(jsonStr))
}

func (Message Message) Error(message string, data interface{}, w http.ResponseWriter) {
	Message.Code = 1001
	Message.Message = message
	Message.Data = data
	jsonStr, err := json.Marshal(Message)
	if err != nil {
		log.Fatal("序列化json错误")
	}
	fmt.Fprintf(w, "%s", string(jsonStr))
}
