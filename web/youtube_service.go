package web

import (
	"net/http"

	"github.com/themechanicalcoder/fampay-backend-assignment/config"
	"github.com/themechanicalcoder/fampay-backend-assignment/models"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type YouTubeService struct {
	service           *youtube.Service
	query             string
	maxResults        int
	relevanceLanguage string
}

func Initialise(cfg config.YoutubeConfig) (WebService, error) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: cfg.ApiKey},
	}
	youtubeService, err := youtube.New(client)
	return YouTubeService{query: cfg.Query, maxResults: cfg.MaxResults, relevanceLanguage: cfg.RelevanceLanguage, service: youtubeService}, err
}

func (youtube YouTubeService) FetchYoutubeVideos() ([]models.YoutubeVideo, error) {
	call := youtube.service.Search.List([]string{"snippet"}).
		Q(youtube.query).
		MaxResults(int64(youtube.maxResults)).
		RelevanceLanguage(youtube.relevanceLanguage)

	response, err := call.Do()
	if err != nil {
		return []models.YoutubeVideo{}, err
	}
	return getYoutubeVideos(response), nil
}

func getYoutubeVideos(response *youtube.SearchListResponse) []models.YoutubeVideo {
	videos := []models.YoutubeVideo{}
	for _, item := range response.Items {
		video := models.YoutubeVideo{
			Title:       item.Snippet.Title,
			ChannelId:   item.Snippet.ChannelId,
			Description: item.Snippet.Description,
			VideoId:     item.Id.VideoId,
			Kind:        item.Kind,
			Thumbnails: models.Thumbnail{
				Default: item.Snippet.Thumbnails.Default.Url,
				Medium:  item.Snippet.Thumbnails.Medium.Url,
				High:    item.Snippet.Thumbnails.High.Url,
			},
			PublishedAt: item.Snippet.PublishedAt,
		}
		videos = append(videos, video)
	}
	return videos
}
