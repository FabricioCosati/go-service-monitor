package handlers

import (
	"github.com/FabricioCosati/go-service-monitor/internal/services"
	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(ctx *gin.Context) {
	url := ctx.Query("url")

	if url == "" {
		ctx.JSON(400, gin.H{"error": "url can not be null"})
		return
	}

	status, duration, err := services.HealthCheckService(url)
	if err != nil {
		ctx.JSON(502, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"status": status, "duration": duration})
}
