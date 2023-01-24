package main

import (
	"flag"
	"log"

	"github.com/themechanicalcoder/fampay-backend-assignment/api"
	"github.com/themechanicalcoder/fampay-backend-assignment/business"
	"github.com/themechanicalcoder/fampay-backend-assignment/config"
	"github.com/themechanicalcoder/fampay-backend-assignment/database"
	"github.com/themechanicalcoder/fampay-backend-assignment/jobs"
	"github.com/themechanicalcoder/fampay-backend-assignment/web"
)

type CmdLineParams struct {
	Build      string
	ConfigPath string
}

func ReadCmdLine() CmdLineParams {
	cmdParams := CmdLineParams{}
	flag.StringVar(&cmdParams.Build, "build", "dev", "The build type (dev, uat, prod)")
	flag.StringVar(&cmdParams.ConfigPath, "config_path", ".", "The directory where to look for the config")
	flag.Parse()
	return cmdParams
}

func main() {
	cmdParams := ReadCmdLine()

	cfg, err := config.LoadConfig(cmdParams.Build, cmdParams.ConfigPath)
	if err != nil {
		log.Fatal("Error while loading config", err)
	}

	youtubeservice, err := web.Initialise(cfg.YoutubeConfig)
	if err != nil {
		log.Fatal("Error while initializing youtube service %v", err)
	}

	db, err := database.Connect(cfg.DBConfig)
	if err != nil {
		log.Fatal("Error while connecting to database %v", err)
	}
	

	store := business.Initialize(db)
	worker := jobs.Initialize(cfg.WorkerConfig.QueryInterval, youtubeservice, store)
	go worker.Start()

	api := api.Initialize(cfg.Server, store)
	api.Run()
}
