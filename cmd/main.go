package main

import (
	"github.com/albertoadami/go-gin-blog-api-example/pkg/config"
	"github.com/albertoadami/go-gin-blog-api-example/pkg/database"
	"github.com/albertoadami/go-gin-blog-api-example/pkg/rest"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	const port = "8080"
	r := gin.Default()

	//read the config
	var config = readConfig("config", "config")

	log.Info("here the config", config)

	log.Info("Trying to connect to database...")
	database.InitDb(config)

	rest.InitRouter(r)

	log.Info("Service starting on port " + port)
	r.Run(":" + port)
}

func readConfig(fileName string, directory string) config.Config {
	viper.SetConfigName(fileName)
	viper.AddConfigPath(directory)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var config config.Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return config

}
