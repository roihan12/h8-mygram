package user

import (
	"time"
)

type UserEntity struct {
	ID        uint
	Username  string 
	Email     string 
	Password  string 
	Age       int64 
	CreatedAt time.Time
	UpdatedAt time.Time
}


type UserService interface {
	Login(email, password string) (string, error)
	Register(newUser UserEntity) (UserEntity, error)
	Profile(userID uint) (UserEntity, error)
	Update(userID uint, updateData UserEntity) (UserEntity, error)
	Delete(userID uint) error
}

type UserData interface {
	Login(username string) (UserEntity, error)
	Register(newUser UserEntity) (UserEntity, error)
	Profile(userID uint) (UserEntity, error)
	Update(userID uint, updateData UserEntity) (UserEntity, error)
	Delete(userID uint) error
}
