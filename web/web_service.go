package web

import (
	"github.com/themechanicalcoder/fampay-backend-assignment/models"

)

type WebService interface {
	FetchYoutubeVideos() ([]models.YoutubeVideo, error)
}

