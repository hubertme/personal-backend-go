package gin

import "wang.hihubert.personal-backend/routes"

func initHandlerGroups() {
	registerTestRoutes()
	registerContactsRoutes()
}

func registerTestRoutes() {
	testRoutes := router.Group("/test")

	testRoutes.GET("/helloWorld", routes.SendHelloWorldResponse)
}

func registerContactsRoutes() {
	contactRoutes := router.Group("/contacts")

	contactRoutes.POST("/me", routes.SendMeMessage)
}
