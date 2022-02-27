package database

import (
	"discord-22-api/entity"
)

type Database interface {
	GetAllUsers() ([]*entity.User, error)
	GetUser(id string) (*entity.User, error)
	SaveUser(user *entity.User) error
	SumUserMessageCount(discordId string, sum int64) error
	SumUserCommandCount(discordId string, sum int64) error
	Disconnect() error
}

var db Database

func Instance() Database {
	if db == nil {
		db = NewMongo()
	}
	return db
}
