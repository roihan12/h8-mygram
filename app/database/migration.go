package database

import (
	"fmt"

	comment "github.com/roihan12/h8-mygram/features/comment/data"
	photo "github.com/roihan12/h8-mygram/features/photo/data"
	socialMedia "github.com/roihan12/h8-mygram/features/socialMedia/data"
	user "github.com/roihan12/h8-mygram/features/user/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{}, &photo.Photo{}, &comment.Comment{}, &socialMedia.SocialMedia{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
