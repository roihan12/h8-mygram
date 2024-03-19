package database

import (
	"fmt"

	user "github.com/roihan12/h8-mygram/features/user/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
