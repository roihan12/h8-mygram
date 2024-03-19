package data

import (
	"github.com/roihan12/h8-mygram/features/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;uniqueIndex;type:varchar(50)"`
	Email    string `gorm:"not null;uniqueIndex;type:varchar(50)"`
	Password string `gorm:"not null;type:varchar(255)"`
	Age      int64  `gorm:"not null;type:integer"`
}

func ToUserEntity(data User) user.UserEntity {
	return user.UserEntity{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		Age:      data.Age,
		Password: data.Password,
	}
}

func UserEntityToUser(userEntity user.UserEntity) User {
	return User{
		Model:    gorm.Model{ID: userEntity.ID},
		Username: userEntity.Username,
		Email:    userEntity.Email,
		Password: userEntity.Password,
		Age:      userEntity.Age,
	}
}
