package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/roihan12/h8-mygram/features/user"
	"github.com/roihan12/h8-mygram/utils"
)

type UserController struct {
	srv user.UserService
}

func New(srv user.UserService) *UserController {
	return &UserController{
		srv: srv,
	}
}

func (uc *UserController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	token, err := uc.srv.Login(req.Email, req.Password)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	rsp := newAuthResponse(token)

	utils.HandleSuccess(ctx, "Login user successfully", rsp)
}

func (uc *UserController) Register(ctx *gin.Context) {
	req := RegisterRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	res, err := uc.srv.Register(*ReqToCore(req))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	rsp := ToResponse(res)

	utils.HandleSuccess(ctx, "Register user successfully", rsp)
}

func (uc *UserController) Profile(ctx *gin.Context) {
	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}
	res, err := uc.srv.Profile(userID.(uint))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	rsp := ToResponse(res)

	utils.HandleSuccess(ctx, "Get profile user successfully", rsp)

}

func (uc *UserController) Update(ctx *gin.Context) {
	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}
	input := UpdateRequest{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	res, err := uc.srv.Update(userID.(uint), *ReqToCore(input))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	rsp := ToResponse(res)

	utils.HandleSuccess(ctx, "Update user successfully", rsp)
}

func (uc *UserController) Delete(ctx *gin.Context) {
	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}
	err := uc.srv.Delete(userID.(uint))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Your account has been deleted successfully", "")

}
