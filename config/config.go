package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Server struct {
	Addr string
	Port int
}

type Config struct {
	Name          string
	DBConfig      DBConfig
	YoutubeConfig YoutubeConfig
	WorkerConfig  WorkerConfig
	Server        Server
}

type WorkerConfig struct {
	QueryInterval int
}

type DBConfig struct {
	Database   string
	Collection string
	MongoURI   string
}

type YoutubeConfig struct {
	ApiKey            string
	QueryInterval     int
	Query             string
	MaxResults        int
	RelevanceLanguage string
}

func LoadConfig(build, path string) (Config, error) {
	cfg := Config{}
	viper.SetConfigName(build)
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return cfg, err
	}

	if err := viper.GetViper().UnmarshalExact(&cfg); err != nil {
		fmt.Println(err)
		return cfg, err
	}

	return cfg, nil
}
