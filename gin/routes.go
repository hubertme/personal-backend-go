package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
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
	setupRecordMetrics()
	promeRoutes := router.Group("/monitors")

	promeRoutes.GET("/prometheus", gin.WrapH(promhttp.Handler()))
}

// Prometheus
var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func setupRecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(time.Second * 2)
		}
	}()
}
