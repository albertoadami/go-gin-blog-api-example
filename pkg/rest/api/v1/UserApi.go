package v1

import (
	"fmt"
	"net/http"

	json "github.com/albertoadami/go-gin-blog-api-example/pkg/rest/json"
	service "github.com/albertoadami/go-gin-blog-api-example/pkg/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateUser(c *gin.Context) {
	var createUserRequest json.CreateUserRequest

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {

		c.JSON(http.StatusBadRequest, json.GenericErrorResponse{Error: "Invalid Request", Description: err.Error()}) //bad request if the request is not a CreateUserRequest
		return
	}

	id, error := service.AddUser(createUserRequest)

	if error != nil {
		log.Info(fmt.Sprintf("The user with email %s exists already", createUserRequest.Email))
		c.Writer.WriteHeader(409)
	} else {
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}

}
