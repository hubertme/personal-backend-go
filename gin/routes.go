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
	//contactRoute := router.Group("/contacts")
}
