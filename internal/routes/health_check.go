package routes

import (
	"github.com/FabricioCosati/go-service-monitor/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitHealthCheckRoutes(r *gin.Engine) {

	r.GET("/health-check", handlers.HealthCheckHandler)
}
