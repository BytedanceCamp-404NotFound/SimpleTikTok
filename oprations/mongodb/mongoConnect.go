package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type autoIncrement struct {
	Name  string `bson:"name"`
	Value int64  `bson:"value"`
}

func Connect(database string, Table string, url string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("connect mongodb failed, err:%v\n", err))
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ping mongodb failed, err:%v\n", err))
	}

	collection := client.Database(database).Collection(Table)
	filter := bson.D{{
		Key:   "name",
		Value: "auto_increment",
	}}
	num, _ := collection.CountDocuments(context.Background(), filter)
	if num == 0 {
		increment := autoIncrement{
			Name: "auto_increment",
			Value: 0,
		}
		collection.InsertOne(context.Background(),increment)
	}
	return collection, nil
}
