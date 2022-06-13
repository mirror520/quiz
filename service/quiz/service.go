package quiz

import (
	"context"

	"github.com/mirror520/quiz/model/comment"
)

type Service interface {
	CreateComment(context.Context, *comment.Comment) (*comment.Comment, error)               // Create a new comment.
	GetCommentByUUID(context.Context, string) (*comment.Comment, error)                      // Get comment by UUID.
	ModifyCommentByUUID(context.Context, *comment.Comment, string) (*comment.Comment, error) // Modify Comment by UUID.
	RemoveCommentByUUID(context.Context, string) error                                       // Remove comment by UUID.
}

type service struct {
	repo comment.Repository
}

func NewService(repo comment.Repository) Service {
	svc := new(service)
	svc.repo = repo
	return svc
}

func (svc *service) CreateComment(ctx context.Context, comment *comment.Comment) (*comment.Comment, error) {
	return nil, nil
}

func (svc *service) GetCommentByUUID(ctx context.Context, uuid string) (*comment.Comment, error) {
	return nil, nil
}

func (svc *service) ModifyCommentByUUID(ctx context.Context, comment *comment.Comment, uuid string) (*comment.Comment, error) {
	return nil, nil
}

func (svc *service) RemoveCommentByUUID(ctx context.Context, uuid string) error {
	return nil
}
