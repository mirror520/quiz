package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/mirror520/quiz/model"
	"github.com/mirror520/quiz/model/comment"
)

type commentRepository struct {
	db *gorm.DB
}

// 以 DB 實作 Comment 領域模型之資料持續性
func NewCommentRepository() comment.Repository {
	cfg := model.Config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
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

func (repo *commentRepository) DB() *gorm.DB {
	return repo.db
}
