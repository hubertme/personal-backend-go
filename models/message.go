package models

import "time"

type Message struct {
	ID          string    `json:"id"`
	SenderName  string    `json:"sender_name"`
	SenderEmail string    `json:"sender_email"`
	Subject     string    `json:"subject"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
}
