package internal

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

func GenerateEvent() (name string, data string) {
	type event struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"createdAt"`
		Content   string    `json:"content"`
	}

	id := fmt.Sprintf("event-%s", uuid.New().String())

	msg := event{
		Name:      id,
		CreatedAt: time.Now(),
		Content:   "fake content data",
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	return id, string(msgBytes)
}
