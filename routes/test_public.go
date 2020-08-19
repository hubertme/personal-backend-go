package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wang.hihubert.personal-backend/result"
)

func SendHelloWorldResponse(c *gin.Context) {
	dataRes := map[string]interface{} {
		"codeName": "Hello, World!",
		"foo": "bar",
	}

	c.JSON(http.StatusAccepted, result.Success(dataRes))
}
