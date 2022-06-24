package router

import (
	"go-anime/controller"

	"github.com/gin-gonic/gin"
)

func BuildRouter() gin.IRouter {
	router := gin.Default()
	router.GET("/", controller.GetAllAnime)
	router.Run()
	return router
}
