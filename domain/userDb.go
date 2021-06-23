package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCollection struct {
	Collection *mongo.Collection
}

func NewDatabase(db *mongo.Database) MongoCollection {
	CollectionName := "users"
	return MongoCollection{Collection: db.Collection(CollectionName)}
}

func (m MongoCollection) GetAllDB() ([]UserStruct, error) {
	query := bson.M{}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, er1 := m.Collection.Find(ctx, query)
	if er1 != nil {
		return nil, er1
	}
	var result []UserStruct
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user UserStruct
		cursor.Decode(&user)
		result = append(result, user)
	}
	if er2 := cursor.Err(); er2 != nil {
		return nil, er2
	}
	return result, nil
}

func (m MongoCollection) GetByIdDB(id string) (*UserStruct, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var user UserStruct
	query := bson.M{"id": id}
	er1 := m.Collection.FindOne(ctx, query).Decode(&user)
	if er1 != nil {
		return nil, er1
	}
	return &user, nil
}

func (m MongoCollection) InsertUserDB(newUser *UserStruct) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_, err := m.Collection.InsertOne(ctx, newUser)
	if err != nil {
		panic(err.Error())
	}
	resultString := "Create new user success"
	return resultString, nil
}

func (m MongoCollection) UpdateUserDB(user *UserStruct) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	query := bson.M{"id": user.Id}
	updateQuery := bson.M{
		"$set": user,
	}
	_, err := m.Collection.UpdateOne(ctx, query, updateQuery)
	if err != nil {
		return "", err
	}
	resultString := "Update success"
	return resultString, nil
}

func (m MongoCollection) DeleteUserDB(id string) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	query := bson.M{"id": id}
	_, er1 := m.Collection.DeleteOne(ctx, query)
	if er1 != nil {
		return "", er1
	}
	resultString := "Delete success"
	return resultString, nil
}
