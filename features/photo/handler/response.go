package handler

import (
	"time"

	"github.com/roihan12/h8-mygram/features/photo"
	user "github.com/roihan12/h8-mygram/features/user/handler"
)

type PhotoResponse struct {
	ID        uint             `json:"id"`
	Title     string           `json:"title"`
	Caption   string           `json:"caption"`
	PhotoURL  string           `json:"photo_url"`
	UserID    uint             `json:"user_id"`
	CreatedAt time.Time        `json:"created_at,omitempty"`
	UpdatedAt time.Time        `json:"updated_at,omitempty"`
	User      user.UserReponse `json:"user,omitempty"`
	Comments  []CommentRes     `json:"comments,omitempty"`
}
type CommentRes struct {
	ID        uint             `json:"id"`
	Message   string           `json:"message"`
	UserID    uint             `json:"user_id"`
	PhotoID   uint             `json:"photo_id"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	User      user.UserReponse `json:"user,omitempty"`
}

func PhotoEntityToPhotoResponse(data photo.PhotoEntity) PhotoResponse {
	photoResponse := PhotoResponse{
		ID:        data.ID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoURL:  data.PhotoURL,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	photoResponse.User = user.UserReponse{
		Username: data.User.Username,
		Email:    data.User.Email,
	}
	for _, v := range data.Comments {
		var comment = CommentRes{
			ID:        v.ID,
			Message:   v.Message,
			UserID:    v.UserID,
			PhotoID:   v.PhotoID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: user.UserReponse{
				Username: v.User.Username,
				Email:    v.User.Email,
			},
		}

		photoResponse.Comments = append(photoResponse.Comments, comment)
	}

	return photoResponse
}

func ListPhotoEntityToPhotoResponse(data []photo.PhotoEntity) []PhotoResponse {
	var dataResponses []PhotoResponse
	for _, v := range data {
		dataResponses = append(dataResponses, PhotoEntityToPhotoResponse(v))
	}
	return dataResponses
}
