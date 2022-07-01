package database

import (
	"context"
	"discord-22-api/config"
	"discord-22-api/entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"log"
)

type database struct {
	client *mongo.Client
	ctx    context.Context
}

func (d *database) GetAllUsers() (users []*entity.User, err error) {
	users = make([]*entity.User, 0)
	cur, err := d.client.
		Database(config.Instance().DbName).
		Collection("user").
		Find(d.ctx, bson.D{{}})
	if err != nil {
		return
	}
	defer cur.Close(d.ctx)
	for cur.Next(d.ctx) {
		var user *entity.User
		err = cur.Decode(&user)
		user.Calculate()
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

func (d *database) GetUser(id string) (user *entity.User, err error) {
	if err != nil {
		return
	}
	err = d.client.
		Database(config.Instance().DbName).
		Collection("user").
		FindOne(d.ctx, bson.D{{"discordId", id}}).
		Decode(&user)
	if user != nil {
		user.Calculate()
	}
	return
}

func (d *database) SaveUser(user *entity.User) error {
	if user.DiscordId == "" {
		insertedClient, err := d.client.
			Database(config.Instance().DbName).
			Collection("user").
			InsertOne(d.ctx, user)
		if err != nil {
			return err
		}
		user.ID = insertedClient.InsertedID.(primitive.ObjectID)

	} else {
		upsert := true
		_, err := d.client.
			Database(config.Instance().DbName).
			Collection("user").
			UpdateOne(d.ctx, bson.M{"discordId": user.DiscordId}, bson.D{{"$set", user}}, &options.UpdateOptions{Upsert: &upsert})
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *database) SumUserMessageCount(discordId string, sum int64) (err error) {
	_, err = d.client.
		Database(config.Instance().DbName).
		Collection("user").
		UpdateOne(d.ctx, bson.M{"discordId": discordId}, bson.D{{"$inc", bson.M{"messageCount": sum}}})
	return
}

func (d *database) SumUserCommandCount(discordId string, sum int64) (err error) {
	_, err = d.client.
		Database(config.Instance().DbName).
		Collection("user").
		UpdateOne(d.ctx, bson.M{"discordId": discordId}, bson.D{{"$inc", bson.M{"commandCount": sum}}})
	return
}

func (d *database) Disconnect() error {
	return d.client.Disconnect(d.ctx)
}

func NewMongo() Database {
	cfg := config.Instance()
	uri := fmt.Sprintf("mongodb://%s:%s@%s", cfg.DbUser, cfg.DbPass, cfg.DbHost)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetMonitor(otelmongo.NewMonitor()))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &database{client: client, ctx: ctx}
}
