package database

import (
	"fmt"

	"context"
	"time"

	"github.com/themechanicalcoder/fampay-backend-assignment/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	databaseName string
	collectionName string
	client *mongo.Client
}

func Connect(config config.DBConfig) (DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// "mongodb+srv://gourav:gourav@cluster0.1tstn9p.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		return nil, err
	}
	
	return MongoDB{client: client, databaseName: config.Database, collectionName: config.Collection}, err
}

func (db MongoDB) InsertMany(ctx context.Context, data []interface{}) (int, error) {
	results, err := db.client.Database(db.databaseName).Collection(db.collectionName).InsertMany(context.Background(), data)
	if err != nil {
		fmt.Println("Error %v ", err)
	}
	return len(results.InsertedIDs), nil
}

func (db MongoDB) Find(ctx context.Context, filter interface{}, options *options.FindOptions, results interface{}) error {
	cur, err := db.client.Database(db.databaseName).Collection(db.collectionName).Find(ctx, filter, options)
	if err != nil {
		return err
	}
	if err = cur.All(ctx, results); err != nil {
		return err
	}
	return nil
}
