package main

import (
	"github.com/FabricioCosati/go-service-monitor/internal/routes"
	"github.com/gin-gonic/gin"
	_ "go.opentelemetry.io/otel"
	_ "go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	_ "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	_ "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	_ "go.opentelemetry.io/otel/log/global"
	_ "go.opentelemetry.io/otel/propagation"
	_ "go.opentelemetry.io/otel/sdk/log"
	_ "go.opentelemetry.io/otel/sdk/metric"
	_ "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	r := gin.New()
	routes.InitHealthCheckRoutes(r)
	r.Run(":8081")
}
