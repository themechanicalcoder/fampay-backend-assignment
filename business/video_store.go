package business

import (
	"github.com/themechanicalcoder/fampay-backend-assignment/models"
)

type VideoStore interface {
	Search(query string) (videos []models.YoutubeVideo, err error)
	FetchVideos(offset int, limit int) (videos []models.YoutubeVideo, err error)
	InsertVideos(videos []models.YoutubeVideo) error
}
