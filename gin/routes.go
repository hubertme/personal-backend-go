package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"wang.hihubert.personal-backend/routes"
)

func initHandlerGroups() {
	registerTestRoutes()
	registerContactsRoutes()
	registerMonitoringRoutes()
}

func registerTestRoutes() {
	testRoutes := router.Group("/test")

	testRoutes.GET("/helloWorld", routes.SendHelloWorldResponse)
}

func registerContactsRoutes() {
	contactRoutes := router.Group("/contacts")

	contactRoutes.POST("/me", routes.SendMeMessage)
}

func registerMonitoringRoutes() {
	promeRoutes := router.Group("/monitors")

	promeRoutes.GET("/prometheus", gin.WrapH(promhttp.Handler()))
}
