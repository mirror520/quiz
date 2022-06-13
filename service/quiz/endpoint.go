package quiz

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"

	"github.com/mirror520/quiz/model/comment"
)

var (
	ErrCreateCommentFail   = errors.New("create comment fail")
	ErrCommentUUIDNotFound = errors.New("comment uuid not found")
	ErrModifyFail          = errors.New("modify fail")
	ErrDeleteFail          = errors.New("delete fail")
	ErrInvalidUUID         = errors.New("invalid uuid")
)

type ContextKey int

const UUID ContextKey = iota

// 針對 Service 介面製作端點 Endpoint，其為抽象方法可供內部調用
// 優點：
// 1. 可避免直接相依引用
// 2. 可供 transport 包裝成多種傳輸協定介面，供外部引用，並轉成相同的存取端點

func CreateCommentEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*comment.Comment)
		return svc.CreateComment(req)
	}
}

func GetCommentByUUIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		uuid, ok := ctx.Value(UUID).(string)
		if !ok {
			return nil, ErrInvalidUUID
		}

		return svc.GetCommentByUUID(uuid)
	}
}

func ModifyCommentByUUIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		uuid, ok := ctx.Value(UUID).(string)
		if !ok {
			return nil, ErrInvalidUUID
		}

		req := request.(*comment.Comment)
		return svc.ModifyCommentByUUID(req, uuid)
	}
}

func RemoveCommentByUUIDEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		uuid, ok := ctx.Value(UUID).(string)
		if !ok {
			return nil, ErrInvalidUUID
		}

		err = svc.RemoveCommentByUUID(uuid)
		return nil, err
	}
}
