package app

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConnect(DatabaseName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ListUserGolang:abcd1234@cluster0.voqyf.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	DB := client.Database(DatabaseName)
	return DB
}
