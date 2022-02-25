package database

import (
	"eagle-jump-api/database/entity"
	"github.com/labstack/gommon/log"
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
		log.Info("Initialising Database...")
		db = NewMongo()
		log.Info("Database OK!")
	}
	return db
}
