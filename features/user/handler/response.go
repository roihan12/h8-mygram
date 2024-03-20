package handler

import "github.com/roihan12/h8-mygram/features/user"

type UserReponse struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int64  `json:"age,omitempty"`
}

// authResponse represents an authentication response body
type authResponse struct {
	AccessToken string `json:"token" example:"eGdh5kiOTyyaQ3_bNykYDeYHO21Jg2..."`
}

// newAuthResponse is a helper function to create a response body for handling authentication data
func newAuthResponse(token string) authResponse {
	return authResponse{
		AccessToken: token,
	}
}

func ToResponse(data user.UserEntity) UserReponse {
	return UserReponse{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		Age:      data.Age,
	}
}

func GetToResponse(data user.UserEntity) UserReponse {
	return UserReponse{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		Age:      data.Age,
	}
}
