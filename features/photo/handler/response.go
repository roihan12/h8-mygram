package handler

import (
	"time"

	"github.com/roihan12/h8-mygram/features/photo"
	user "github.com/roihan12/h8-mygram/features/user/handler"
)

type PhotoResponse struct {
	ID        uint             `json:"id,omitempty"`
	Title     string           `json:"title,omitempty"`
	Caption   string           `json:"caption,omitempty"`
	PhotoURL  string           `json:"photo_url,omitempty"`
	UserID    uint             `json:"user_id,omitempty"`
	CreatedAt time.Time        `json:"created_at,omitempty"`
	UpdatedAt time.Time        `json:"updated_at,omitempty"`
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

	return photoResponse
}

func ListPhotoEntityToPhotoResponse(data []photo.PhotoEntity) []PhotoResponse {
	var dataResponses []PhotoResponse
	for _, v := range data {
		dataResponses = append(dataResponses, PhotoEntityToPhotoResponse(v))
	}
	return dataResponses
}
