package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func NewExporter(ctx *gin.Context) (sdktrace.SpanExporter, error) {
	return stdouttrace.New()
}

func NewTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("HealthCheckService"),
		),
	)
	if err != nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}

func ServiceTracingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		exp, err := NewExporter(ctx)
		if err != nil {
			log.Fatalf("failed to initialize exporter: %v", err)
		}
		tp := NewTraceProvider(exp)
		defer func() {
			_ = tp.Shutdown(ctx)
		}()

		otel.SetTracerProvider(tp)
		tracer := tp.Tracer("go.opentelemetry.io")
		ctx.Next()

		_, span := tracer.Start(ctx, "service-monitor")
		defer span.End()
	}
}
