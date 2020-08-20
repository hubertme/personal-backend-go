package gin

import "github.com/gin-gonic/gin"

const LOCAL_ADDRESS = "localhost:3030"

var router = gin.Default()

func SetupGinServer() error {
	//gin.SetMode(gin.ReleaseMode)
	router.Use(allowCORSMiddleware())
	initHandlerGroups()

	return router.Run(LOCAL_ADDRESS)
}
