package chat

import (
	"time"
)

type Message struct {
	SenderID   SenderID  `json:"sender_id"`
	SenderName string    `json:"sender_name"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
}
