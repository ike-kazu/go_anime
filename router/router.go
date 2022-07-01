package router

import (
	"encoding/json"
	"go-anime/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func BuildRouter() gin.IRouter {
	router := gin.Default()

	router.GET("/accounts", controller.GetAllAccounts)
	router.GET("/accounts/:id", controller.GetOneAccount)
	router.GET("/animes", controller.GetAllAnimes)
	router.GET("/animes/:id", controller.GetOneAnime)
	router.GET("/deployments", controller.GetAllDeployments)
	router.GET("/deployments/:id", controller.GetOneDeployment)
	router.GET("/platforms", controller.GetAllPlatforms)
	router.GET("/platforms/:id", controller.GetOnePlatform)
	router.GET("/seasons", controller.GetAllSeasons)
	router.GET("/seasons/:id", controller.GetOneSeason)
	router.GET("/stories", controller.GetAllStories)
	router.GET("/stories/:id", controller.GetOneStory)

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

	router.Run(":8080")
	return router
}
