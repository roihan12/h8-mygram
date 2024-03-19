package user

import (
	"time"
)

type UserEntity struct {
	ID        uint
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Age       int64  `json:"age" validate:"required,gt=8"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


type UserService interface {
	Login(email, password string) (string, error)
	Register(newUser UserEntity) (UserEntity, error)
	Profile(UserID uint) (UserEntity, error)
	Update(UserID uint, updateData UserEntity) (UserEntity, error)
	Delete(UserID uint) error
}

type UserData interface {
	Login(username string) (UserEntity, error)
	Register(newUser UserEntity) (UserEntity, error)
	Profile(userID uint) (UserEntity, error)
	Update(userID uint, updateData UserEntity) (UserEntity, error)
	Delete(userID uint) error
}
