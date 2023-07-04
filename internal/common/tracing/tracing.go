package tracing

import (
	"github.com/getsentry/sentry-go"
	sentryotel "github.com/getsentry/sentry-go/otel"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"tracing_example/internal/common/config"
)

func InitTracer(cfg config.ServiceConfig) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.SentryDSN,
		Debug:            true,
		AttachStacktrace: true,
		SampleRate:       1.0,
		TracesSampleRate: 1.0,
		EnableTracing:    true,
		ServerName:       "Cart Service",
		Environment:      cfg.Env,
	}); err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
}
