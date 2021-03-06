package comment

import (
	"errors"
	"time"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
)

type Comment struct {
	UUID     string    `json:"uuid" gorm:"primaryKey"`
	ParentID string    `json:"parentid"`
	Comment  string    `json:"comment"`
	Author   string    `json:"author"`
	Update   time.Time `json:"update"`
	Favorite bool      `json:"favorite"`
}

type Repository interface {
	Store(*Comment) error
	FindCommentByUUID(string) (*Comment, error)
	Remove(string) error
}
