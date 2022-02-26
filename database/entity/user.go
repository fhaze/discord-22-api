package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	DiscordId    string             `bson:"discordId,omitempty" json:"discordId,omitempty"`
	Avatar       string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	MessageCount int64              `bson:"messageCount,omitempty" json:"messageCount"`
	CommandCount int64              `bson:"commandCount,omitempty" json:"commandCount"`
	EagleCoin    int64              `bson:"eagleCoin,omitempty" json:"eagleCoin"`
	JoinedAt     time.Time          `bson:"joinedAt" json:"joinedAt"`
	Exp          int64              `json:"exp"`
	RequiredExp  int64              `json:"requiredExp"`
	Level        int64              `json:"level"`
}

func (u *User) Calculate() {
	msgScore := float64(u.MessageCount)*.18 + float64(u.CommandCount)*.23
	timeScore := time.Now().Sub(u.JoinedAt).Hours()
	u.Exp = int64(msgScore + timeScore)

	level := float64(0)
	for true {
		level++
		required := level * 100
		u.RequiredExp = int64(required)
		if float64(u.Exp) < required {
			u.Level = int64(level)
			break
		}
	}
}
