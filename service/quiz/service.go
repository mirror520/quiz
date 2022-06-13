package quiz

import (
	"errors"

	"github.com/mirror520/quiz/model/comment"
)

// Service 專注於業務邏輯，並直接操作領域模型
// 可單純對 Service 業務邏輯做單元測試
type Service interface {
	CreateComment(*comment.Comment) (*comment.Comment, error)               // Create a new comment.
	GetCommentByUUID(string) (*comment.Comment, error)                      // Get comment by UUID.
	ModifyCommentByUUID(*comment.Comment, string) (*comment.Comment, error) // Modify Comment by UUID.
	RemoveCommentByUUID(string) error                                       // Remove comment by UUID.
}

type service struct {
	repo comment.Repository
}

func NewService(repo comment.Repository) Service {
	svc := new(service) // 建立服務實例
	svc.repo = repo     // 資料持續性
	return svc          // 可檢驗實例是否已實作介面所有方法
}

func (svc *service) CreateComment(c *comment.Comment) (*comment.Comment, error) {
	err := svc.repo.Store(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (svc *service) GetCommentByUUID(uuid string) (*comment.Comment, error) {
	return svc.repo.FindCommentByUUID(uuid)
}

func (svc *service) ModifyCommentByUUID(modifiedComment *comment.Comment, uuid string) (*comment.Comment, error) {
	if modifiedComment.UUID != uuid {
		return nil, errors.New("uuid inconsistent")
	}

	c, err := svc.repo.FindCommentByUUID(uuid)
	if err != nil {
		return nil, err
	}

	c.ParentID = modifiedComment.ParentID
	c.Comment = modifiedComment.Comment
	c.Author = modifiedComment.Author
	c.Update = modifiedComment.Update
	c.Favorite = modifiedComment.Favorite

	if err := svc.repo.Store(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (svc *service) RemoveCommentByUUID(uuid string) error {
	return svc.repo.Remove(uuid)
}
