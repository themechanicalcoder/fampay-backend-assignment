package api

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/themechanicalcoder/fampay-backend-assignment/models"
)

// Search API to search the stored videos using their title and description.
func (api API) search(ctx iris.Context) {
	searchRequest, err := readInput(ctx)
	if err != nil {
		api.log.Println("Error while unmarshalling request ", searchRequest)
		responseJson(ctx, http.StatusBadRequest, models.SearchResponse{
			Videos: []models.YoutubeVideo{},
			Error: models.ErrorResponse{
				Description: "Bad Request",
			},
		})
	}
	videos, err := api.manager.Search(searchRequest.Query)
	if err != nil {
		api.log.Println(fmt.Sprintf("Error while searching query %s ", searchRequest.Query), err)
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
