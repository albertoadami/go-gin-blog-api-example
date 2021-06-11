package service

import (
	"fmt"

	json "github.com/albertoadami/go-gin-blog-api-example/pkg/rest/json"
)

func AddUser(user json.CreateUserRequest) (int, error) {

	fmt.Println("here the request ", user)
	return 123, nil

}
