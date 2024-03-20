package data

import (
	"time"

	userModel "github.com/roihan12/h8-mygram/features/user/data"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `gorm:"not null;type:varchar(50)"`
	SocialMediaURL string `gorm:"not null;type:varchar(255)"`
	UserID         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           userModel.User `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}
