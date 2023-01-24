package api

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/themechanicalcoder/fampay-backend-assignment/models"
)

func (api API) search(ctx iris.Context) {
	searchRequest, err := readInput(ctx)
	if err != nil {
		responseJson(ctx, http.StatusBadRequest, models.SearchResponse{
			Videos: []models.YoutubeVideo{},
			Error: models.ErrorResponse{
				Description: "Bad Request",
			},
		})
	}
	videos, err := api.manager.Search(searchRequest.Query)
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

func readInput(ctx iris.Context) (models.SearchRequest, error) {
	var request models.SearchRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		return models.SearchRequest{}, err
	}
	return request, err
}

func responseJson(ctx iris.Context, statusCode int, resp interface{}) {
	ctx.StatusCode(statusCode)
	ctx.JSON(resp)
}
