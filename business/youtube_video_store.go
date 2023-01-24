package business

import (
	"context"
	"fmt"

	"github.com/themechanicalcoder/fampay-backend-assignment/database"
	"github.com/themechanicalcoder/fampay-backend-assignment/models"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type YoutubeVideoStore struct {
	db database.DB
}

func Initialize(database database.DB) VideoStore { 
	return YoutubeVideoStore{db : database}
}

// Search videos containing partial match for the search query in either video title or description.
func (manager YoutubeVideoStore) Search(query string) (videos []models.YoutubeVideo, err error) {
	filter := generateFuzzyFilter(query)
	sort := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
	findOptions := options.Find().SetSort(sort)
	err = manager.db.Find(context.Background(), filter, findOptions, &videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

/*
	Returns paginated sorted in descending order of published datetime.
*/
func (manager YoutubeVideoStore) FetchVideos(offset int, limit int) (videos []models.YoutubeVideo, err error) {
	filter := bson.M{}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"publishedat", -1}})
	findOptions.SetSkip(int64(offset))
	findOptions.SetLimit(int64(limit))

	err = manager.db.Find(context.Background(), filter, findOptions, &videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

/*
	Insert the youtube videos into database
*/
func (manager YoutubeVideoStore) InsertVideos(videos []models.YoutubeVideo) error {
	var interfaces []interface{}
	for _, data := range videos {
		interfaces = append(interfaces, data)
	}
	_, err := manager.db.InsertMany(context.Background(), interfaces)
	if err == nil {
		return fmt.Errorf("Error while inserting videos in database %w", err)
	}
	return nil
}

/*
	Generate fuzzy filter
*/
func generateFuzzyFilter(query string) bson.D {
	return bson.D{{"$text", bson.D{{"$search", query}}}}
}
