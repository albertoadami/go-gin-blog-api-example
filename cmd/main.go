package main

import (
	"github.com/albertoadami/go-gin-blog-api-example/pkg/rest"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	const port = "8080"
	r := gin.Default()

	rest.InitRouter(r)

	log.Info("Service starting on port " + port)
	r.Run(":" + port)
}
