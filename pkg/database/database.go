package database

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"

	log "github.com/sirupsen/logrus"

	"github.com/albertoadami/go-gin-blog-api-example/pkg/config"
	"gorm.io/gorm"
)

func InitDb(config config.Config) *gorm.DB {

	dsn := url.URL{
		User:     url.UserPassword(config.User, config.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Path:     config.Name,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		log.Error("Error connecting to database : error=%v", err)
		return nil
	}

	log.Info("Connected to database ", config.Name)
	return db
}
