package handler

import (
	"github.com/roihan12/h8-mygram/features/socialMedia"
)

type CreateSocialMediaRequest struct {
	Name           string `json:"title" form:"name" binding:"required" example:"My social"`
	SocialMediaURL string `json:"caption" form:"social_media_url" binding:"required"`
}

func CreatePhotoRequestToPhotoEntity(request *CreateSocialMediaRequest) socialMedia.SocialMediaEntity {
	return socialMedia.SocialMediaEntity{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
	}
}
