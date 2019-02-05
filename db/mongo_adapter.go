package db

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github/pedroportasvieira/go-api/errors"
	"time"
)

// Insert a record onto a given collection
func InsertMongoRecord(collectionName string, record bson.M) {

	collection, ctx, client := connect(collectionName)

	_, err := collection.InsertOne(ctx, record)

	errors.HandleError(err)

	errors.HandleError(client.Disconnect(ctx))
}

// Return a cursor with the records of a given collection
func GetMongoRecords(collectionName string) (*mongo.Cursor, context.Context) {

	collection, ctx, client := connect(collectionName)

	results, err := collection.Find(ctx, bson.D{})

	errors.HandleError(err)

	defer results.Close(ctx)

	errors.HandleError(client.Disconnect(ctx))

	return results, ctx
}

// Connect to mongo and return a collection instance and context
func connect(collectionName string) (*mongo.Collection, context.Context, *mongo.Client) {
	client, err := mongo.NewClient("mongodb://mongo:27017")

	errors.HandleError(err)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	errors.HandleError(client.Connect(ctx))

	collection := client.Database("testdb").Collection(collectionName)

	return collection, ctx, client
}
