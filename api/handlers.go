package api

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/themechanicalcoder/fampay-backend-assignment/business"
	"github.com/themechanicalcoder/fampay-backend-assignment/config"
)

type API struct {
	app *iris.Application
	host string
	port int
	manager business.VideoStore
}

func Initialize(config config.Server, manager business.VideoStore) API {
	app := iris.New()
	return API{app: app, manager: manager, host: config.Addr, port: config.Port}
}

func (api API) register() {
	api.app.Get("/v1/search/", api.search)
	api.app.Get("/v1/videos/", api.getVideos)
}

func (api API) Run() {
	addr := fmt.Sprintf("%s:%d", api.host, api.port)
	api.register()
	api.app.Listen(addr)
}
