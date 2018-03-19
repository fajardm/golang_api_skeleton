package bootstrap

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"github.com/fajardm/example_blog_api/commons/helpers"
)

func InitApp() {
	router := SetupRouter()

	helpers.InstanceValidator()

	router.Run(os.Getenv("APP_ADDRESS"))
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	// you should register module here
	// just invoke module router method to here
	// example article.NewArticleHttpHandler(router)

	return router
}
