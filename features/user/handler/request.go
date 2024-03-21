package handler

import "github.com/roihan12/h8-mygram/features/user"

// loginRequest represents the request body for logging in a user
type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email" example:"test@example.com"`
	Password string `json:"password" form:"password" binding:"required,min=6" example:"123456" minLength:"6"`
}

type RegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required" example:"test12"`
	Email    string `json:"email" form:"email" binding:"required,email" example:"test@example.com"`
	Password string `json:"password" form:"password" binding:"required,min=6" example:"123456" minLength:"6"`
	Age      int64  `json:"age" form:"age" binding:"required,gt=8" example:"20"`
}

type UpdateRequest struct {
	Username string `json:"username" form:"username" binding:"omitempty,required" example:"JohnDoe"`
	Email    string `json:"email" form:"email" binding:"omitempty,required,email" example:"test@example.com"`
	Password string `json:"password" form:"password" binding:"omitempty,required,min=6" example:"123456"`
	Age      int64  `json:"age" form:"age" binding:"omitempty,required,gte=8" example:"21"`
}

func ReqToCore(data interface{}) *user.UserEntity {
	res := user.UserEntity{}

	switch docs := data.(type) {
	case LoginRequest:
		cnv := docs
		res.Email = cnv.Email
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Username = cnv.Username
		res.Email = cnv.Email
		res.Password = cnv.Password
		res.Age = cnv.Age
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Username = cnv.Username
		res.Email = cnv.Email
		res.Password = cnv.Password
		res.Age = cnv.Age
	default:
		return nil
	}

	return &res
}
