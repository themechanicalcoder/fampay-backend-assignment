package business

import (
	"context"
	"fmt"

	"github.com/themechanicalcoder/fampay-backend-assignment/database"
	"github.com/themechanicalcoder/fampay-backend-assignment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoManager struct {
	DB database.DB
}

// {$or: [{ title: { $regex: "est", $options: 'i' } },
//{ description: { $regex: "est", $options: 'i' } }]
//}

func (manager VideoManager) Search(query string) (videos []models.YoutubeVideo, err error) {
	filter := generateRegexFilter(query)
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"publishedat", -1}})
	err = manager.DB.Find(context.Background(), filter, findOptions, &videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (manager VideoManager) FetchVideos(offset int, limit int) (videos []models.YoutubeVideo, err error) {
	filter := bson.M{}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"publishedat", -1}})
	findOptions.SetSkip(int64(offset))
	findOptions.SetLimit(int64(limit))

	err = manager.DB.Find(context.Background(), filter, findOptions, &videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (manager VideoManager) InsertVideos(videos []models.YoutubeVideo) error {
	var interfaces []interface{}
	for _, data := range videos {
		interfaces = append(interfaces, data)
	}
	_, err := manager.DB.InsertMany(context.Background(), interfaces)
	if err == nil {
		return fmt.Errorf("Error while inserting videos in database %w", err)
	}
	return nil
}

func generateRegexFilter(keyword string) bson.M {
	return bson.M{"$or": bson.A{
		bson.M{"title": bson.M{"$regex": primitive.Regex{Pattern: keyword, Options: "i"}}},
		bson.M{"description": bson.M{"$regex": primitive.Regex{Pattern: keyword, Options: "i"}}},
	}}
}
