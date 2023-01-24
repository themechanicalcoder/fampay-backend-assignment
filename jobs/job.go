package jobs

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/themechanicalcoder/fampay-backend-assignment/business"
	"github.com/themechanicalcoder/fampay-backend-assignment/youtubeservice"
)

type Worker struct {
	duration       time.Duration
	youtubeService *youtubeservice.YouTubeService
	videoManager   *business.VideoManager
}

func Initialize(interval int, youtubeService *youtubeservice.YouTubeService, videoManager *business.VideoManager) Worker {
	return Worker{duration: time.Duration(time.Second * time.Duration(interval)), youtubeService: youtubeService, videoManager: videoManager}
}

func (w *Worker) Start() {
	for {
		youtubeVideos, err := w.youtubeService.FetchYoutubeVideos()
		if err != nil {
			log.Printf("Error while fetching youtube videos  %v", err)
		}else if len(youtubeVideos) != 0 {
			err = w.videoManager.InsertVideos(youtubeVideos)
			if err != nil {
				fmt.Println(err);
			}
		}
		spew.Dump(youtubeVideos)
		time.Sleep(w.duration)
	}
}
