package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wang.hihubert.personal-backend/models"
	"wang.hihubert.personal-backend/result"
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

	c.JSON(http.StatusAccepted, result.Success(message))
}
