package db

import (
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mirror520/quiz/model/comment"
)

type commentRepository struct {
	db *gorm.DB
}

// 以 DB 實作 Comment 領域模型之資料持續性
func NewCommentRepository() comment.Repository {
	db, err := gorm.Open(sqlite.Open("comment.db"), &gorm.Config{})
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	db.AutoMigrate(&comment.Comment{})

	repo := new(commentRepository)
	repo.db = db
	return repo
}

func (repo *commentRepository) Store(c *comment.Comment) error {
	var tx *gorm.DB
	if c.UUID == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		c.UUID = id.String()

		tx = repo.db.Create(c)
	} else {
		tx = repo.db.Save(c)
	}

	return tx.Error
}

func (repo *commentRepository) FindCommentByUUID(uuid string) (*comment.Comment, error) {
	var c *comment.Comment
	tx := repo.db.Take(&c, "uuid = ?", uuid)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, comment.ErrCommentNotFound
		}

		return nil, err
	}

	return c, nil
}

func (repo *commentRepository) Remove(uuid string) error {
	tx := repo.db.Delete(&comment.Comment{}, "uuid = ?", uuid)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return comment.ErrCommentNotFound
		}

		return err
	}

	return nil
}
