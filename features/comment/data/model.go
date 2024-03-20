package data

import (
	"time"

	photoModel "github.com/roihan12/h8-mygram/features/photo/data"
	userModel "github.com/roihan12/h8-mygram/features/user/data"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message   string `gorm:"not null;type:varchar(255)"`
	UserID    string
	PhotoID   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Photo     photoModel.Photo `gorm:"foreignkey:PhotoID;association_foreignkey:ID"`
	User      userModel.User   `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}
