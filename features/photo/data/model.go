package data

import (
	"time"

	"github.com/roihan12/h8-mygram/features/photo"
	"github.com/roihan12/h8-mygram/features/user"
	userModel "github.com/roihan12/h8-mygram/features/user/data"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title     string `gorm:"not null;type:varchar(50)"`
	Caption   string `gorm:"type:varchar(255)"`
	PhotoURL  string `gorm:"not null;type:varchar(255)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      userModel.User `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}

func PhotoEntityToPhoto(data photo.PhotoEntity) Photo {
	return Photo{
		Model:     gorm.Model{ID: data.ID},
		UserID:    data.UserID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoURL:  data.PhotoURL,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func PhotoToPhotoEntity(data Photo) photo.PhotoEntity {
	result := photo.PhotoEntity{
		ID:        data.ID,
		UserID:    data.UserID,
		Title:     data.Title,
		Caption:   data.Caption,
		PhotoURL:  data.PhotoURL,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	result.User = user.UserEntity{
		Email:    data.User.Email,
		Username: data.User.Username,
	}

	return result
}

func ListPhotoToPhotoEntity(data []Photo) []photo.PhotoEntity {
	var photoEntity []photo.PhotoEntity
	for _, v := range data {
		photoEntity = append(photoEntity, PhotoToPhotoEntity(v))
	}
	return photoEntity
}
