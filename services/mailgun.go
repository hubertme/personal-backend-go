package services

import (
	"context"
	"github.com/mailgun/mailgun-go/v4"
	"os"
	"time"
)

var mg *mailgun.MailgunImpl

const DefaultSender = "Hubert Wang <no-reply@postman.hihubert.wang>"
const DefaultReceiver = "Hubert Wang <pingme@hihubert.wang>"

func MailgunSetup() {
	domainName := "postmaster.hihubert.wang"
	apiKey := os.Getenv("MAILGUN_API_KEY")
	mg = mailgun.NewMailgun(domainName, apiKey)
}

func MailgunSendMessage(sender string, receiver string, subject string, content string) (string, string, error) {
	message := mg.NewMessage(sender, subject, content, receiver)

	// Define background context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send message
	mes, id, err := mg.Send(ctx, message)

	if err != nil {
		return "", "", err
	}

	return mes, id, nil
}
