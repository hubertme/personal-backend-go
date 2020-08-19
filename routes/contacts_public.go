package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	emailTitle := fmt.Sprintf("A Message from %v <%v>", message.SenderName, message.SenderEmail)
	sender := fmt.Sprintf("%v <%v>", message.SenderName, message.SenderEmail)
	res, id, mgErr := services.MailgunSendMessage(sender, services.DefaultReceiver, emailTitle, message.Content)
	if mgErr != nil {
		c.JSON(http.StatusInternalServerError, result.DevError(mgErr.Error()))
		return
	}

	dataRes := map[string]interface{}{
		"mg_id":  id,
		"mg_msg": res,
	}
	c.JSON(http.StatusAccepted, result.Success(dataRes))
}
