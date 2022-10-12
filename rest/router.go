package rest

import (
	v1 "github.com/albertoadami/go-gin-blog-api-example/rest/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	engine.GET("/health", func(c *gin.Context) {
		c.Writer.WriteHeader(204)
	})
	apiV1 := engine.Group("/api/v1")
	apiV1.POST("/users", v1.CreateUser)
}
