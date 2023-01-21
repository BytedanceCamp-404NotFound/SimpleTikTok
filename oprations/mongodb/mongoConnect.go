package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(database string, collection string, url string) (*mongo.Collection,error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err!=nil {
		return nil, errors.New(fmt.Sprintf("connect mongodb failed, err:%v\n",err))
	}

	err = client.Ping(context.Background(), nil)
	if err!=nil {
		return nil, errors.New(fmt.Sprintf("ping mongodb failed, err:%v\n",err))
	}

	collention := client.Database(database).Collection(collection)
	return collention, nil
}