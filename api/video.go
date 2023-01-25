package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/themechanicalcoder/fampay-backend-assignment/models"
)

// GET API which returns the stored video data in a paginated response sorted in descending order of published datetime.
func (api API) getVideos(ctx iris.Context) {
	limit := ctx.URLParamIntDefault("limit", 10)
	offset := ctx.URLParamIntDefault("offset", 0)
	
	videos, err := api.manager.FetchVideos(offset, limit)
	if err != nil {
		log.Println("Error while fetching videos ", err)
		responseJson(ctx, http.StatusInternalServerError, models.SearchResponse{
			Videos: []models.YoutubeVideo{},
			Error: models.ErrorResponse{
				Description: fmt.Sprintf("Error while searching for query %s", err.Error()),
			},
		})
	}
	response := models.SearchResponse{
		Status: "SUCCESS",
		Videos: videos,
	}
	responseJson(ctx, http.StatusOK, response)
}