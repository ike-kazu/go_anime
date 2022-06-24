package main

import (
	_ "go-anime/config"
	"go-anime/router"
)

func main() {
	router.BuildRouter()
}
