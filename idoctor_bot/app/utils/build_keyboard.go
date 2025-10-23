package utils

import (
	"encoding/json"
	"log"
)

func MakeReplyMarkup(jsonStr string) json.RawMessage {
	var raw json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &raw)
	if err != nil {
		log.Println("❌ Error in MakeReplyMarkup:", err)
		return nil
	}
	return raw
}
