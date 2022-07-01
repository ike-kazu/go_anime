package models

import "time"

type Deployment struct {
	Id int
	PublishedAt time.Time
	Story Story
	Account Account
	Platform Platform
}
