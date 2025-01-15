package main

import (
	"github.com/FabricioCosati/go-service-monitor/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	routes.InitHealthCheckRoutes(r)
	r.Run(":8081")
}
