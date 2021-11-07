package db

import (
	"context"
	"fmt"
	"gamedata/db/model"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TestGolangMongoDB ...
func TestGolangMongoDB(t *testing.T) {

	var client *mongo.Client
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.9:27017")) //몽고DB 접속클라 만듬
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	user := client.Database("gamedata").Collection("user")

	var who model.Users

	currErr := user.FindOne(ctx, bson.D{}).Decode(&who)
	if currErr != nil {
		panic(err)
	}

	fmt.Println(who)

}

// TestGolangMongoDBBson ...
func TestGolangMongoDBBson(t *testing.T) {

	var client *mongo.Client
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.9:27017")) //몽고DB 접속클라 만듬
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	user := client.Database("gamedata").Collection("user")

	var who model.Users

	findFilter := bson.D{primitive.E{"summoner", "ghost"}}

	currErr := user.FindOne(ctx, findFilter).Decode(&who)
	if currErr != nil {
		panic(err)
	}

	fmt.Println(who)

}
