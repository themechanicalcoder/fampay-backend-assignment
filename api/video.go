package api

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/themechanicalcoder/fampay-backend-assignment/models"
)

func (api API) getVideos(ctx iris.Context) {
	limit := ctx.URLParamIntDefault("limit", 10)
	offset := ctx.URLParamIntDefault("offset", 0)
	
	videos, err := api.manager.FetchVideos(offset, limit)
	if err != nil {
		responseJson(ctx, http.StatusInternalServerError, models.SearchResponse{
			Videos: []models.YoutubeVideo{},
			Error: models.ErrorResponse{
				Description: "Something went wrong",
			},
		})
	}
	response := models.SearchResponse{
		Status: "SUCCESS",
		Videos: videos,
	}
	responseJson(ctx, http.StatusOK, response)
}