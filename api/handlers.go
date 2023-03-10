package api

import (
	"fmt"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/themechanicalcoder/fampay-backend-assignment/business"
	"github.com/themechanicalcoder/fampay-backend-assignment/config"
)

type API struct {
	app *iris.Application
	log *log.Logger
	host string
	port int
	store business.VideoStore
}

func Initialize(config config.Server, store business.VideoStore, log *log.Logger) API {
	app := iris.New()
	return API{app: app, store: store, host: config.Addr, port: config.Port, log: log}
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
