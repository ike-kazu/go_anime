package main

import (
	_ "go-anime/config"
	"go-anime/router"
	_ "go-anime/models"
)

func main() {
	router.BuildRouter()
}
