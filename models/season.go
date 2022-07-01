package models

import "time"

type Season struct {
	Id int
	Title string
	Image string
	Outline string
	Anime Anime
	Director string
	StartDatetime time.Time
	Week string
}
