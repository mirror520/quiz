package quiz

import "github.com/mirror520/quiz/model"

type Service interface {
	CreateComment(comment *model.Comment) (*model.Comment, error)                    // Create a new comment.
	GetCommentByUUID(uuid string) (*model.Comment, error)                            // Get comment by UUID.
	ModifyCommentByUUID(comment *model.Comment, uuid string) (*model.Comment, error) // Modify Comment by UUID.
	RemoveCommentByUUID(uuid string) error                                           // Remove comment by UUID.
}

type service struct{}

func NewService() Service {
	return new(service)
}

func (svc *service) CreateComment(comment *model.Comment) (*model.Comment, error) {
	return nil, nil
}

func (svc *service) GetCommentByUUID(uuid string) (*model.Comment, error) {
	return nil, nil
}

func (svc *service) ModifyCommentByUUID(comment *model.Comment, uuid string) (*model.Comment, error) {
	return nil, nil
}

func (svc *service) RemoveCommentByUUID(uuid string) error {
	return nil
}
