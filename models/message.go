package models

type Message struct {
	SenderName  string `json:"sender_name"`
	SenderEmail string `json:"sender_email"`
	Subject     string `json:"subject"`
	Content     string `json:"content"`
}
