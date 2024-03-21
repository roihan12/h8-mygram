package handler

import (
	"time"

	"github.com/roihan12/h8-mygram/features/comment"
	photo "github.com/roihan12/h8-mygram/features/photo/handler"
	user "github.com/roihan12/h8-mygram/features/user/handler"
)

type CommentResponse struct {
	ID        uint                `json:"id"`
	Message   string              `json:"message"`
	UserID    uint                `json:"user_id"`
	PhotoID   uint                `json:"photo_id"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	User      user.UserReponse    `json:"user,omitempty"`
	Photo     photo.PhotoResponse `json:"photo,omitempty"`
}

func CommentEntityToCommentResponse(commentEntity comment.CommentEntity) CommentResponse {
	commentResponse := CommentResponse{
		ID:        commentEntity.ID,
		Message:   commentEntity.Message,
		UserID:    commentEntity.UserID,
		PhotoID:   commentEntity.PhotoID,
		CreatedAt: commentEntity.CreatedAt,
		UpdatedAt: commentEntity.UpdatedAt,
	}

	commentResponse.User = user.UserReponse{
		ID:       commentEntity.User.ID,
		Email:    commentEntity.User.Email,
		Username: commentEntity.User.Username,
	}

	commentResponse.Photo = photo.PhotoResponse{
		ID:        commentEntity.Photo.ID,
		Title:     commentEntity.Photo.Title,
		Caption:   commentEntity.Photo.Caption,
		PhotoURL:  commentEntity.Photo.PhotoURL,
		UserID:    commentEntity.Photo.UserID,
		CreatedAt: commentEntity.Photo.CreatedAt,
		UpdatedAt: commentEntity.Photo.UpdatedAt,
	}

	return commentResponse
}

func ListCommentToCommentResponse(commentEntity []comment.CommentEntity) []CommentResponse {
	var dataRes []CommentResponse
	for _, v := range commentEntity {
		dataRes = append(dataRes, CommentEntityToCommentResponse(v))
	}
	return dataRes
}
