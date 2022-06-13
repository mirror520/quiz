package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mirror520/quiz/model/comment"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository() comment.Repository {
	repo := new(commentRepository)

	db, err := gorm.Open(sqlite.Open("comment.db"), &gorm.Config{})
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	repo.db = db

	return repo
}

func (repo *commentRepository) Store(c *comment.Comment) error {
	return nil
}

func (repo *commentRepository) FindCommentByUUID(uuid string) (*comment.Comment, error) {
	return nil, nil
}

func (repo *commentRepository) Remove(uuid string) error {
	return nil
}
