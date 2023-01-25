package workers

import (
	"log"
	"time"

	//"github.com/davecgh/go-spew/spew"
	"github.com/themechanicalcoder/fampay-backend-assignment/business"
	"github.com/themechanicalcoder/fampay-backend-assignment/web"
)

type Worker struct {
	duration       time.Duration
	youtubeService web.WebService
	videoStore  business.VideoStore
	log            *log.Logger
}

func Initialize(interval int, youtubeService web.WebService, videoStore business.VideoStore, log *log.Logger) Worker {
	return Worker{duration: time.Duration(time.Second * time.Duration(interval)), youtubeService: youtubeService, videoStore: videoStore, log: log}
}

func (w Worker) Start() {
	for {
		youtubeVideos, err := w.youtubeService.FetchYoutubeVideos()
		if err != nil {
			log.Printf("Error while fetching youtube videos  %v", err)
		} else {
			if len(youtubeVideos) != 0 {
				err = w.videoStore.InsertVideos(youtubeVideos)
				if err != nil {
					w.log.Println("Error while inserting the videos :", err.Error())
				}
			}
		}
		//spew.Dump(youtubeVideos)
		time.Sleep(w.duration)
	}
}