package main

import (
	"flag"
	"log"
	"context"
	"os"

	"github.com/themechanicalcoder/fampay-backend-assignment/api"
	"github.com/themechanicalcoder/fampay-backend-assignment/business"
	"github.com/themechanicalcoder/fampay-backend-assignment/config"
	"github.com/themechanicalcoder/fampay-backend-assignment/database"
	"github.com/themechanicalcoder/fampay-backend-assignment/workers"
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
	log := log.New(os.Stdout, "fampay-backend-assignment : ", log.LstdFlags|log.Lshortfile)
	if err := run(cmdParams, log); err != nil {
		log.Printf("main : error: %+v, exiting ", err)
		os.Exit(1)
	}
}

func run(cmdParams CmdLineParams, log *log.Logger) error {
	cfg, err := config.LoadConfig(cmdParams.Build, cmdParams.ConfigPath)
	
	log.Println("Loading Config ")
	if err != nil {
		log.Println("Error while loading config :")
		return err
	}

	// initialize youtube service
	log.Println("Initializing youtube service :")
	youtubeservice, err := web.Initialise(cfg.YoutubeConfig)
	if err != nil {
		log.Println("Error while initializing youtube service :")
		return err
	}

	//connect to database
	log.Println("Initializing database :")
	db, err := database.Connect(cfg.DBConfig)
	if err != nil {
		log.Println("Error while connecting to database :")
		return err
	}

	// create text index for search
	log.Println("Creating Indexes : ")
	err = db.CreateIndex()
	if err != nil {
		log.Println("Error while creating index ", err)
	}
	
	defer func() {
		log.Println("Database Stopping :")
		db.Disconnect(context.Background())
	}()
	
	// initialize business layer
	log.Println("Initializing Store :")
	store := business.Initialize(db)

	// initialize job
	worker := workers.Initialize(cfg.WorkerConfig.QueryInterval, youtubeservice, store, log)
	go worker.Start()

	api := api.Initialize(cfg.Server, store, log)
	api.Run()
	return nil
}
