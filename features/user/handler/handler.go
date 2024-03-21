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

// Login godoc
//
//	@Summary		Login and get an access token
//	@Description	Logs in a registered user and returns an access token if the credentials are valid.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		LoginRequest		true	"Login request body"
//	@Success		200		{object}	authResponse		"Succesfully logged in"
//	@Failure		400		{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401		{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		500		{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/users/login [post]
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

// Register godoc
//
//	@Summary		Register a new user
//	@Description	create a new user account
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		RegisterRequest		true	"Register request"
//	@Success		200				{object}	UserReponse			"User created"
//	@Failure		400				{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401				{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		404				{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		409				{object}	utils.ErrorResponse	"Data conflict error"
//	@Failure		500				{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/users/register [post]
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

// ProfileUser godoc
//
//	@Summary		Get a user
//	@Description	Get a profile user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	UserReponse			"Profile User displayed"
//	@Failure		400	{object}	utils.ErrorResponse	"Validation error"
//	@Failure		404	{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/users [get]
//	@Security		BearerAuth
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

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Update a user's username, email, password, age
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			updateUserRequest	body		UpdateRequest		true	"Update user request"
//	@Success		200					{object}	UserReponse			"User updated"
//	@Failure		400					{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401					{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		403					{object}	utils.ErrorResponse	"Forbidden error"
//	@Failure		404					{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500					{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/users [put]
//	@Security		BearerAuth
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

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.Response		"User deleted"
//	@Failure		400	{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401	{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		403	{object}	utils.ErrorResponse	"Forbidden error"
//	@Failure		404	{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/users [delete]
//	@Security		BearerAuth
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
