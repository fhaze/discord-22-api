package handler

import (
	"discord-22-api/config"
	"discord-22-api/router/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetVersion(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Version{Tag: config.Instance().CommitHash})
}
