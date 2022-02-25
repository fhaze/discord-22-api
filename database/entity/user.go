package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	DiscordId    string             `bson:"discordId,omitempty" json:"discordId,omitempty"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	MessageCount int64              `bson:"messageCount,omitempty" json:"messageCount"`
	CommandCount int64              `bson:"commandCount,omitempty" json:"commandCount"`
	EagleCoin    int64              `bson:"eagleCoin,omitempty" json:"eagleCoin"`
	JoinedAt     time.Time          `bson:"joinedAt" json:"joinedAt"`
}
