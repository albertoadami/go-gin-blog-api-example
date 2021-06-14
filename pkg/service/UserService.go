package service

import (
	"time"

	"github.com/albertoadami/go-gin-blog-api-example/pkg/database"

	"github.com/albertoadami/go-gin-blog-api-example/pkg/models"
	"github.com/albertoadami/go-gin-blog-api-example/pkg/rest/json"
)

func CreateUser(user json.CreateUserRequest) (uint, error) {

	userToAdd := models.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password, Gender: user.Gender, Enabled: false, CreatedAt: time.Now()}
	database.DB.Create(&userToAdd)

	return userToAdd.ID, nil
}
