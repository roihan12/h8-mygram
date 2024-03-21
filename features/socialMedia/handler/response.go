package handler

import (
	"time"

	"github.com/roihan12/h8-mygram/features/socialMedia"
	user "github.com/roihan12/h8-mygram/features/user/handler"
)

type SocialMediaResponse struct {
	ID             uint             `json:"id"`
	Name           string           `json:"name"`
	SocialMediaURL string           `json:"social_media_url"`
	UserID         uint             `json:"user_id"`
	CreatedAt      time.Time        `json:"created_at,omitempty"`
	UpdatedAt      time.Time        `json:"updated_at,omitempty"`
	User           user.UserReponse `json:"user,omitempty"`
}

func SocialMediaEntityToSocialMediaResponse(data socialMedia.SocialMediaEntity) SocialMediaResponse {
	socialMediaResponse := SocialMediaResponse{
		ID:             data.ID,
		Name:           data.Name,
		SocialMediaURL: data.SocialMediaURL,
		UserID:         data.UserID,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	}

	socialMediaResponse.User = user.UserReponse{
		ID:       data.User.ID,
		Username: data.User.Username,
	}

	return socialMediaResponse
}

func ListSocialMediaEntityToSocialMediaResponse(data []socialMedia.SocialMediaEntity) []SocialMediaResponse {
	var dataResponses []SocialMediaResponse
	for _, v := range data {
		dataResponses = append(dataResponses, SocialMediaEntityToSocialMediaResponse(v))
	}
	return dataResponses
}
