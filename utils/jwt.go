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
	tokeStr, err := token.SignedString([]byte(config.JWT_KEY))
	if err != nil {
		return "", err
	}
	return tokeStr, nil
}

func VerifyAccessToken(tokenStr string) (*authCustomClaims, error) {
	claims := &authCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_KEY), nil
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
