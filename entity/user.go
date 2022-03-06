package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"strings"
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
	Bot          bool               `bson:"bot" json:"bot"`
	Exp          int64              `json:"exp"`
	RequiredExp  int64              `json:"requiredExp"`
	Level        int64              `json:"level"`
	Icons        string             `json:"icons"`
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
	u.Icons = calculateIcons(level)
}

func calculateIcons(level float64) string {
	var icons []string
	var suns float64
	var moons float64
	var stars float64
	var diff float64

	suns = math.Floor(level / 16)
	if suns < 1 {
		diff = level
	} else {
		diff = level - math.Floor(suns*16)
	}

	moons = math.Floor(level / 4)
	if moons < 1 {
		diff = level
	} else {
		diff = level - math.Floor(moons*4)
	}

	stars = diff

	for i := 0.; i < suns; i++ {
		icons = append(icons, "☀")
	}
	for i := 0.; i < moons; i++ {
		icons = append(icons, "🌙")
	}
	for i := 0.; i < stars; i++ {
		icons = append(icons, "⭐")
	}

	return strings.Join(icons, "")
}
