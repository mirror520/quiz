package quiz

import (
	"github.com/mirror520/quiz/model/comment"
)

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
	svc := new(service)
	svc.repo = repo
	return svc
}

func (svc *service) CreateComment(c *comment.Comment) (*comment.Comment, error) {
	return nil, nil
}

func (svc *service) GetCommentByUUID(uuid string) (*comment.Comment, error) {
	return nil, nil
}

func (svc *service) ModifyCommentByUUID(c *comment.Comment, uuid string) (*comment.Comment, error) {
	return nil, nil
}

func (svc *service) RemoveCommentByUUID(uuid string) error {
	return nil
}
