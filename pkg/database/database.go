package database

import (
	"fmt"

	"gorm.io/driver/postgres"

	log "github.com/sirupsen/logrus"

	"github.com/albertoadami/go-gin-blog-api-example/pkg/config"
	"github.com/albertoadami/go-gin-blog-api-example/pkg/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb(config config.Config) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Host, config.User, config.Password, config.Name, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : error=%v", err)
		return
	}

	log.Info("Connected to database ", config.Name)
	DB = db
}

func MigrateDb() {
	DB.Migrator().CreateTable(&models.User{})
}
