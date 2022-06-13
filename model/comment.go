package model

import "time"

type Comment struct {
	UUID     string    `json:"uuid"`
	ParentID string    `json:"parentid"`
	Comment  string    `json:"comment"`
	Author   string    `json:"author"`
	Update   time.Time `json:"update"`
	Favorite bool      `json:"favorite"`
}
