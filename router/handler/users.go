package handler

import (
	"discord-22-api/database"
	"discord-22-api/entity"
	"discord-22-api/router/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func GetUsers(c echo.Context) error {
	users, err := database.Instance().GetAllUsers()
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	user, err := database.Instance().GetUser(c.Param("id"))
	if err == mongo.ErrNoDocuments {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, user)
}

func PostUser(c echo.Context) error {
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user.JoinedAt = time.Now()
	if err := database.Instance().SaveUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

func PutSumMessageCount(c echo.Context) error {
	messageCount := new(model.RequestMessageCount)
	if err := c.Bind(messageCount); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.Instance().SumUserMessageCount(c.Param("id"), messageCount.Count); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}

func PutSumCommandCount(c echo.Context) error {
	messageCount := new(model.RequestMessageCount)
	if err := c.Bind(messageCount); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.Instance().SumUserCommandCount(c.Param("id"), messageCount.Count); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}
