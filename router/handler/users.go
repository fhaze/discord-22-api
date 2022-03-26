package handler

import (
	"discord-22-api/ctx"
	"discord-22-api/database"
	"discord-22-api/entity"
	"discord-22-api/router/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func GetUsers(c echo.Context) error {
	users, err := database.Instance().GetAllUsers()
	if err != nil {
		ctx.From(c).LogError(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	ctx.From(c).LogInfof("got users list")
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	discordId := c.Param("id")
	if discordId == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	user, err := database.Instance().GetUser(discordId)
	if err == mongo.ErrNoDocuments {
		ctx.From(c).LogInfof("user not found by id=%d", discordId)
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		ctx.From(c).LogError(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	ctx.From(c).LogInfof("got user id=%v", user.ID.Hex())
	return c.JSON(http.StatusOK, user)
}

func PostUser(c echo.Context) error {
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	user.JoinedAt = time.Now()
	if err := database.Instance().SaveUser(user); err != nil {
		ctx.From(c).LogError(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	ctx.From(c).LogInfof("created user id=%v name=%s", user.ID.Hex(), user.Name)
	return c.JSON(http.StatusOK, user)
}

func PutSumMessageCount(c echo.Context) error {
	messageCount := new(model.RequestMessageCount)
	discordId := c.Param("id")
	if err := c.Bind(messageCount); err != nil || discordId == "" {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.Instance().SumUserMessageCount(discordId, messageCount.Count); err != nil {
		ctx.From(c).LogError(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	ctx.From(c).LogInfof("user sum message discordId=%s count=%d", discordId, messageCount.Count)
	return c.NoContent(http.StatusOK)
}

func PutSumCommandCount(c echo.Context) error {
	messageCount := new(model.RequestMessageCount)
	discordId := c.Param("id")
	if err := c.Bind(messageCount); err != nil || discordId == "" {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := database.Instance().SumUserCommandCount(discordId, messageCount.Count); err != nil {
		ctx.From(c).LogError(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	ctx.From(c).LogInfof("user sum command discordId=%s count=%d", discordId, messageCount.Count)
	return c.NoContent(http.StatusOK)
}
