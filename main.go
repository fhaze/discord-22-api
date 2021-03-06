package main

import (
	"context"
	"discord-22-api/config"
	"discord-22-api/ctx"
	"discord-22-api/router"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func initTracer() func() {
	c := context.Background()
	exporter, err := otlptracegrpc.New(c)

	if err != nil {
		logrus.Fatal("failed to initialise exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
	)

	otel.SetTracerProvider(tp)

	return func() { _ = tp.Shutdown(c) }
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{DisableTimestamp: true})
	cleanup := initTracer()
	defer cleanup()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-22-KEY"},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: `{"id":"${id}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"latency":"${latency}","bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n"}))
	e.Use(otelecho.Middleware("discord-22-api"))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ex := &ctx.Extended{Context: c}
			return next(ex)
		}
	})
	e.Use(middleware.RequestID())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("X-22-KEY")
			if key == config.Instance().RootApiKey {
				return next(c)
			}
			return c.NoContent(http.StatusUnauthorized)
		}
	})
	router.Assign(e)
	e.Logger.Fatal(e.Start(":8888"))
}
