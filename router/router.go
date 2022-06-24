package router

import (
	"encoding/json"
	"go-anime/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func BuildRouter() gin.IRouter {
	router := gin.Default()
	router.GET("/animes", controller.GetAllAnime)
	router.GET("/", func(c *gin.Context) {
		response := map[string]string{
			"message": "Hello!",
		}
		body, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		c.Writer.Write(body)
	})
	router.Run()
	return router
}
