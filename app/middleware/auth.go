package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/roihan12/h8-mygram/utils"
)

func AuthMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	if bearerIsExist := strings.Contains(auth, "Bearer"); !bearerIsExist {
		err := utils.ErrEmptyAuthorizationHeader
		utils.HandleAbort(ctx, err)
		return
	}

	token := strings.Split(auth, " ")
	if len(token) < 2 {
		err := utils.ErrInvalidAuthorizationHeader
		utils.HandleAbort(ctx, err)
		return
	}

	claims, err := utils.VerifyAccessToken(token[1])
	if err != nil {
		utils.HandleAbort(ctx, err)
		return
	}

	ctx.Set("username", claims.Username)
	ctx.Set("userID", claims.UserID)

	ctx.Next()
}
