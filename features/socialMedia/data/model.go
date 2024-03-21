package data

import (
	"time"

	"github.com/roihan12/h8-mygram/features/socialMedia"
	"github.com/roihan12/h8-mygram/features/user"
	userModel "github.com/roihan12/h8-mygram/features/user/data"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `gorm:"not null;type:varchar(50)"`
	SocialMediaURL string `gorm:"not null;type:varchar(255)"`
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           userModel.User `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}

func SocialMediaEntityToSocialMedia(data socialMedia.SocialMediaEntity) SocialMedia {
	return SocialMedia{
		Model:          gorm.Model{ID: data.ID},
		UserID:         data.UserID,
		Name:           data.Name,
		SocialMediaURL: data.SocialMediaURL,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	}
}

func SocialMediaToSocialMediaEntity(data SocialMedia) socialMedia.SocialMediaEntity {
	result := socialMedia.SocialMediaEntity{
		ID:             data.ID,
		UserID:         data.UserID,
		Name:           data.Name,
		SocialMediaURL: data.SocialMediaURL,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	}
	result.User = user.UserEntity{
		ID: data.User.ID,
		Email:    data.User.Email,
		Username: data.User.Username,
	}

	return result
}

func ListSocialMediaToSocialMediaEntity(data []SocialMedia) []socialMedia.SocialMediaEntity {
	var socialMediaEntity []socialMedia.SocialMediaEntity
	for _, v := range data {
		socialMediaEntity = append(socialMediaEntity, SocialMediaToSocialMediaEntity(v))
	}
	return socialMediaEntity
}
