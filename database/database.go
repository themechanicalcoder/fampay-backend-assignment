package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB interface {
	InsertMany(ctx context.Context, data []interface{}) (int, error)
	Find(ctx context.Context, filter interface{},  options *options.FindOptions, results interface{}) error
}