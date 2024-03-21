package handler

import (
	"github.com/roihan12/h8-mygram/features/socialMedia"
)

type CreateSocialMediaRequest struct {
	Name           string `json:"name" form:"name" binding:"required" example:"My social"`
	SocialMediaURL string `json:"social_media_url" form:"social_media_url" binding:"required"`
}

func CreateSocialRequestToSocialEntity(request *CreateSocialMediaRequest) socialMedia.SocialMediaEntity {
	return socialMedia.SocialMediaEntity{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
	}
}
