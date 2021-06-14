package database

import (
	"fmt"

	"gorm.io/driver/postgres"

	log "github.com/sirupsen/logrus"

	"github.com/albertoadami/go-gin-blog-api-example/pkg/config"
	"gorm.io/gorm"
)

func InitDb(config config.Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Host, config.User, config.Password, config.Name, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error("Error connecting to database : error=%v", err)
		return nil
	}

	log.Info("Connected to database ", config.Name)
	return db
}
