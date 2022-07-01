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

func (d *database) GetAllUsers(context context.Context) (users []*entity.User, err error) {
	users = make([]*entity.User, 0)
	cur, err := d.client.
		Database(config.Instance().DbName).
		Collection("user").
		Find(context, bson.D{{}})
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

func (d *database) GetUser(context context.Context, id string) (user *entity.User, err error) {
	if err != nil {
		return
	}
	err = d.client.
		Database(config.Instance().DbName).
		Collection("user").
		FindOne(context, bson.D{{"discordId", id}}).
		Decode(&user)
	if user != nil {
		user.Calculate()
	}
	return
}

func (d *database) SaveUser(context context.Context, user *entity.User) error {
	if user.DiscordId == "" {
		insertedClient, err := d.client.
			Database(config.Instance().DbName).
			Collection("user").
			InsertOne(context, user)
		if err != nil {
			return err
		}
		user.ID = insertedClient.InsertedID.(primitive.ObjectID)

	} else {
		upsert := true
		_, err := d.client.
			Database(config.Instance().DbName).
			Collection("user").
			UpdateOne(context, bson.M{"discordId": user.DiscordId}, bson.D{{"$set", user}}, &options.UpdateOptions{Upsert: &upsert})
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *database) SumUserMessageCount(context context.Context, discordId string, sum int64) (err error) {
	_, err = d.client.
		Database(config.Instance().DbName).
		Collection("user").
		UpdateOne(context, bson.M{"discordId": discordId}, bson.D{{"$inc", bson.M{"messageCount": sum}}})
	return
}

func (d *database) SumUserCommandCount(context context.Context, discordId string, sum int64) (err error) {
	_, err = d.client.
		Database(config.Instance().DbName).
		Collection("user").
		UpdateOne(context, bson.M{"discordId": discordId}, bson.D{{"$inc", bson.M{"commandCount": sum}}})
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
