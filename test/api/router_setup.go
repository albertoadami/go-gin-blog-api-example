package api

import (
	"github.com/albertoadami/go-gin-blog-api-example/pkg/rest"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	rest.InitRouter(r)

	return r
}
