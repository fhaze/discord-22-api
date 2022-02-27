package ctx

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Extended struct {
	echo.Context
}

func From(c echo.Context) *Extended {
	return c.(*Extended)
}

func (c *Extended) LogInfo(a ...interface{}) {
	logrus.WithFields(
		logrus.Fields{
			"id": c.Response().Header().Get(echo.HeaderXRequestID),
		},
	).Info(a...)
}

func (c *Extended) LogInfof(message string, a ...interface{}) {
	logrus.WithFields(
		logrus.Fields{
			"id": c.Response().Header().Get(echo.HeaderXRequestID),
		},
	).Info(fmt.Sprintf(message, a...))
}

func (c *Extended) LogError(a ...interface{}) {
	logrus.WithFields(
		logrus.Fields{
			"id": c.Response().Header().Get(echo.HeaderXRequestID),
		},
	).Error(a...)
}
