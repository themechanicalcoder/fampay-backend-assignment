package database

import (
	"context"
	"time"

	"github.com/themechanicalcoder/fampay-backend-assignment/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	databaseName   string
	collectionName string
	client         *mongo.Client
}

/*
	Connect to the mongodb database
*/
func Connect(config config.DBConfig) (DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		return nil, err
	}

	return MongoDB{client: client, databaseName: config.Database, collectionName: config.Collection}, err
}

/*
 	Create text index for title and description 
 */
func (db MongoDB) CreateIndex() error {
	collection := db.client.Database(db.databaseName).Collection(db.collectionName)
	model := mongo.IndexModel{Keys: bson.D{{"title", "text"}, {"description", "text"} }}
	_, err := collection.Indexes().CreateOne(context.Background(), model)
	return err
}

/*
	Insert the data into mongodb
*/
func (db MongoDB) InsertMany(ctx context.Context, data []interface{}) (int, error) {
	results, err := db.client.Database(db.databaseName).Collection(db.collectionName).InsertMany(context.Background(), data)
	if err != nil {
		return 0, err
	}
	return len(results.InsertedIDs), nil
}

/*
	Fetch data from mongodb for a given filter
*/
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

func (db MongoDB) Disconnect(ctx context.Context) {
	db.client.Disconnect(ctx)
}
