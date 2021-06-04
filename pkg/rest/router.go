package rest

import "github.com/gin-gonic/gin"

func InitRouter(engine *gin.Engine) {
	engine.GET("/health", func(c *gin.Context) {
		c.Writer.WriteHeader(204)
	})
}
