package database

import (
	"context"
	"discord-22-api/entity"
)

type Database interface {
	GetAllUsers(context context.Context) ([]*entity.User, error)
	GetUser(context context.Context, id string) (*entity.User, error)
	SaveUser(context context.Context, user *entity.User) error
	SumUserMessageCount(context context.Context, discordId string, sum int64) error
	SumUserCommandCount(context context.Context, discordId string, sum int64) error
	Disconnect() error
}

var db Database

func Instance() Database {
	if db == nil {
		db = NewMongo()
	}
	return db
}
