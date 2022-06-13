package quiz

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"

	"github.com/mirror520/quiz/model/comment"
)

// 針對 Service 介面各個端點，可提供外部傳輸介面 (如: HTTP API、gRPC、PubSub 等)
// 主要為實作傳輸協議解編碼，即:
// input -> decoding -> endpoint -> service -> endpoint -> encoding -> output

func CreateCommentHandler(e endpoint.Endpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *comment.Comment
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.Abort()
			ctx.String(http.StatusBadRequest, ErrCreateCommentFail.Error())
			return
		}

		resp, err := e(ctx, req)
		if err != nil {
			ctx.Abort()
			ctx.String(http.StatusBadRequest, ErrCreateCommentFail.Error())
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

func GetCommentByUUIDHandler(e endpoint.Endpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")
		ctxWithUUID := context.WithValue(ctx, UUID, uuid)

		resp, err := e(ctxWithUUID, nil)
		if err != nil {
			ctx.Abort()
			ctx.String(http.StatusBadRequest, ErrCommentUUIDNotFound.Error())
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

func ModifyCommentByUUIDHandler(e endpoint.Endpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")
		ctxWithUUID := context.WithValue(ctx, UUID, uuid)

		var req *comment.Comment
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.Abort()
			ctx.String(http.StatusBadRequest, ErrModifyFail.Error())
			return
		}

		resp, err := e(ctxWithUUID, req)
		if err != nil {
			ctx.Abort()
			ctx.String(http.StatusBadRequest, ErrModifyFail.Error())
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

func RemoveCommentByUUIDHandler(e endpoint.Endpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")
		ctxWithUUID := context.WithValue(ctx, UUID, uuid)

		_, err := e(ctxWithUUID, nil)
		if err != nil {
			ctx.Abort()
			ctx.String(http.StatusBadRequest, ErrDeleteFail.Error())
			return
		}

		ctx.String(http.StatusOK, "success")
	}
}
