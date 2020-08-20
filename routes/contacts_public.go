package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"wang.hihubert.personal-backend/contents"
	"wang.hihubert.personal-backend/database"
	"wang.hihubert.personal-backend/models"
	"wang.hihubert.personal-backend/result"
	"wang.hihubert.personal-backend/services"
)

func SendMeMessage(c *gin.Context) {
	message := models.Message{}
	err := c.BindJSON(&message)

	if err != nil {
		c.JSON(http.StatusBadRequest, result.FormatError(err.Error()))
		return
	} else if len(message.SenderEmail) <= 0 || len(message.SenderName) <= 0 || len(message.Content) <= 0 || len(message.Subject) <= 0 {
		c.JSON(http.StatusBadRequest, result.EmptyError(nil))
		return
	}

	sender := fmt.Sprintf("%v <%v>", message.SenderName, message.SenderEmail)

	// Send email to sender
	// Async, no need to wait
	go func(senderName string, subject string, content string) {
		confirmEmail := contents.EmailReceivedConfirmation(senderName, subject, content)
		res, id, mgSendErr := services.MailgunSendMessage(services.DefaultSender, sender, "Thank You for Reaching Out!", confirmEmail)
		if mgSendErr != nil {
			log.Println("Error sending to sender:", mgSendErr.Error())
		} else {
			log.Println("Success sending to sender:", id, res)
		}
	}(message.SenderName, message.Subject, message.Content)

	// Send email to me
	// Async, no need to wait
	go func(subject string, content string) {
		res, id, mgErr := services.MailgunSendMessage(sender, services.DefaultReceiver, subject, content)
		if mgErr != nil {
			//c.JSON(http.StatusInternalServerError, result.DevError(mgErr.Error()))
			log.Println("Error sending to my mailbox:", mgErr.Error())
		} else {
			log.Println("Success sending to my mailbox:", id, res)
		}
	}(message.Subject, message.Content)

	// Save to MySQL
	state, sqlErr := database.DB.Prepare("INSERT INTO messages (sender_name, sender_email, subject, content) VALUES (?, ?, ?, ?)")
	if sqlErr != nil {
		c.JSON(http.StatusInternalServerError, result.DevError(sqlErr.Error()))
		return
	}

	results, qErr := state.Query(message.SenderName, message.SenderEmail, message.Subject, message.Content)
	if qErr != nil {
		c.JSON(http.StatusInternalServerError, result.DevError(sqlErr.Error()))
		return
	}

	//dataRes := map[string]interface{}{
	//	"sender_name":  message.SenderName,
	//	"sender_email": message.SenderEmail,
	//}
	c.JSON(http.StatusAccepted, result.Success(results))
}
