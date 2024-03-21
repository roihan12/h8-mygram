package handler

import "github.com/roihan12/h8-mygram/features/comment"

type CommentRequest struct {
	Message string `json:"message" form:"message" binding:"required" example:"My Commment Message"`
	PhotoID uint   `json:"photo_id" form:"photo_id" binding:"required" example:"1"`
}


type CommentUpdate struct {
	Message string `json:"message" form:"message" binding:"required" example:"My Commment Message"`
}

func CommentUpdateToCommentEntity(request *CommentUpdate) comment.CommentEntity {
	return comment.CommentEntity{
		Message: request.Message,
	}
}

func CommentRequestToCommentEntity(request *CommentRequest) comment.CommentEntity {
	return comment.CommentEntity{
		Message: request.Message,
		PhotoID: request.PhotoID,
	}
}
