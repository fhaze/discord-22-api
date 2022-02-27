package main

import (
	"discord-22-api/config"
	"discord-22-api/ctx"
	"discord-22-api/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	fieldMap := logrus.FieldMap{}
	fieldMap[logrus.FieldKeyTime] = "timestamp"
	logrus.SetFormatter(&logrus.JSONFormatter{FieldMap: fieldMap})

	e := echo.New()
	e.Use(middleware.Logger())
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
