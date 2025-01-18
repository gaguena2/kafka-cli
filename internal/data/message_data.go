package data

import (
	"encoding/json"
	"log"
)

type MessageData struct {
	Data map[string]interface{}
}

func NewMessageData(message *string) *MessageData {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(*message), &data)
	if err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}
	return &MessageData{data}
}
