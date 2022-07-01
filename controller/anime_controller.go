package controller

import (
	"fmt"
	"go-anime/config"
	"log"

	"github.com/gin-gonic/gin"
)

type Anime struct {
	ID     int
	Title  string
	Image  string
	Author string
}

// GET /animes HTTPのハンドリング
func GetAllAnime(c *gin.Context) {
	db := config.DBOpen()
	rows, err := db.Query("SELECT * FROM ANIME;")
	if err != nil {
		log.Fatal("ErrorgetRows db.Query error err:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		u := &Anime{}
		if err := rows.Scan(&u.ID, &u.Title, &u.Image, &u.Author); err != nil {
			log.Fatalf("getRows rows.Scan error err:%v", err)
		}
		fmt.Println(u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("getRows rows.Err error err:%v", err)
	}

}
