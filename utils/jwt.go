package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/roihan12/h8-mygram/app/config"
)

type authCustomClaims struct {
	Username string
	UserID   uint
	jwt.StandardClaims
}

func GenerateToken(username string, userID uint) (string, error) {
	key, err := config.NewEnv()
	if err != nil {
		return "", err
	}
	claims := &authCustomClaims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "myGram",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	tokeStr, err := token.SignedString([]byte(key.JWT_KEY))
	if err != nil {
		return "", err
	}
	return tokeStr, nil
}

func VerifyAccessToken(tokenStr string) (*authCustomClaims, error) {
	key, err := config.NewEnv()
	if err != nil {
		return nil, err
	}
	claims := &authCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key.JWT_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrTokenCreation
	}

	claims, ok := token.Claims.(*authCustomClaims)

	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
