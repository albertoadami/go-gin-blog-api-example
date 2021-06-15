package v1

import (
	"net/http"

	"github.com/albertoadami/go-gin-blog-api-example/errors"
	"github.com/albertoadami/go-gin-blog-api-example/rest/json"
	"github.com/albertoadami/go-gin-blog-api-example/service"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var createUserRequest json.CreateUserRequest

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {

		c.JSON(http.StatusBadRequest, json.GenericErrorResponse{Error: "Invalid Request", Description: err.Error()}) //bad request if the request is not a CreateUserRequest
		return
	}

	id, error := service.CreateUser(createUserRequest)

	if error != nil {
		switch error.(type) {
		case errors.EmailAlreadyInUseError:
			c.Writer.WriteHeader(409)
		default:
			c.Writer.WriteHeader(500)
		}
	} else {
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}

}
