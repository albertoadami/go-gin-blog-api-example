package service

import (
	"fmt"
	"time"

	"github.com/albertoadami/go-gin-blog-api-example/database"
	"github.com/albertoadami/go-gin-blog-api-example/errors"

	log "github.com/sirupsen/logrus"

	"github.com/albertoadami/go-gin-blog-api-example/models"
	"github.com/albertoadami/go-gin-blog-api-example/rest/json"
)

func CreateUser(user json.CreateUserRequest) (uint, error) {

	var existUser models.User

	result := database.DB.Where("email = ?", user.Email).Find(&existUser)

	if result.RowsAffected > 0 {
		log.Error("The user with email %d already exists!", user.Email)
		return 0, errors.EmailAlreadyInUseError{Status: errors.EmailAlreadyInUse, Message: fmt.Sprintf("The user with email %s exist already!", user.Email)}
	}

	userToAdd := models.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password, Gender: user.Gender, Enabled: false, CreatedAt: time.Now()}
	database.DB.Create(&userToAdd)

	return userToAdd.ID, nil
}
