package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/roihan12/h8-mygram/features/socialMedia"
	"github.com/roihan12/h8-mygram/utils"
)

type SocialMediaController struct {
	socialService socialMedia.SocialMediaService
}

func New(srv socialMedia.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		socialService: srv,
	}
}

// ListSocialMedia godoc
//
//	@Summary		List social media
//	@Description	List social media
//	@Tags			SocialMedia
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		SocialMediaResponse	"Social media retrieved"
//	@Failure		500	{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/socialmedias [get]
//	@Security		BearerAuth
func (sc *SocialMediaController) GetAll(ctx *gin.Context) {

	socialMedia, err := sc.socialService.GetAll()
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	listSocialResponse := ListSocialMediaEntityToSocialMediaResponse(socialMedia)

	utils.HandleSuccess(ctx, "Get all Social media successfully", listSocialResponse)
}

// GetSocialMedia godoc
//
//	@Summary		Get a social media
//	@Description	get a social media by id
//	@Tags			SocialMedia
//	@Accept			json
//	@Produce		json
//	@Param			socialMediaId	path		uint				true	"Social Media ID"
//	@Success		200				{object}	SocialMediaResponse	"Social media retrieved"
//	@Failure		400				{object}	utils.ErrorResponse	"Validation error"
//	@Failure		404				{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500				{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/socialmedias/{socialMediaId} [get]
//	@Security		BearerAuth
func (sc *SocialMediaController) GetById(ctx *gin.Context) {
	Id := ctx.Param("socialMediaId")
	socialId, _ := strconv.Atoi(Id)
	socialMedia, err := sc.socialService.GetById(uint(socialId))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	response := SocialMediaEntityToSocialMediaResponse(socialMedia)
	utils.HandleSuccess(ctx, "Get by id social media successfully", response)
}

// CreateSocialMedia godoc
//
//	@Summary		Create a new social media
//	@Description	create a new social media  with name and social media url
//	@Tags			SocialMedia
//	@Accept			json
//	@Produce		json
//	@Param			CreateSocialMediaRequest	body		CreateSocialMediaRequest	true	"Create social media request"
//	@Success		200							{object}	SocialMediaResponse			"Social media created"
//	@Failure		400							{object}	utils.ErrorResponse			"Validation error"
//	@Failure		401							{object}	utils.ErrorResponse			"Unauthorized error"
//	@Failure		403							{object}	utils.ErrorResponse			"Forbidden error"
//	@Failure		404							{object}	utils.ErrorResponse			"Data not found error"
//	@Failure		409							{object}	utils.ErrorResponse			"Data conflict error"
//	@Failure		500							{object}	utils.ErrorResponse			"Internal server error"
//	@Router			/socialmedias [post]
//	@Security		BearerAuth
func (sc *SocialMediaController) Create(ctx *gin.Context) {
	var req CreateSocialMediaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}
	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}

	socialEntity := CreateSocialRequestToSocialEntity(&req)
	socialEntity.UserID = userID.(uint)

	response, err := sc.socialService.Create(socialEntity)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Create social media successfully", SocialMediaEntityToSocialMediaResponse(response))
}

// UpdateSocialMedia godoc
//
//	@Summary		Update a social media
//	@Description	update a social media  with name and social media url
//	@Tags			SocialMedia
//	@Accept			json
//	@Produce		json
//	@Param			socialMediaId				path		uint						true	"Social Media ID"
//	@Param			CreateSocialMediaRequest	body		CreateSocialMediaRequest	true	"Update social media request"
//	@Success		200							{object}	SocialMediaResponse			"Social media created"
//	@Failure		400							{object}	utils.ErrorResponse			"Validation error"
//	@Failure		401							{object}	utils.ErrorResponse			"Unauthorized error"
//	@Failure		403							{object}	utils.ErrorResponse			"Forbidden error"
//	@Failure		404							{object}	utils.ErrorResponse			"Data not found error"
//	@Failure		409							{object}	utils.ErrorResponse			"Data conflict error"
//	@Failure		500							{object}	utils.ErrorResponse			"Internal server error"
//	@Router			/socialmedias/{socialMediaId} [put]
//	@Security		BearerAuth
func (sc *SocialMediaController) Update(ctx *gin.Context) {

	Id := ctx.Param("socialMediaId")
	socialId, _ := strconv.Atoi(Id)
	var req CreateSocialMediaRequest
	if err := ctx.ShouldBind(&req); err != nil {
		utils.ValidationError(ctx, err)
		return
	}

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}

	socialEntity := CreateSocialRequestToSocialEntity(&req)

	response, err := sc.socialService.Update(socialEntity, uint(socialId), userID.(uint))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	utils.HandleSuccess(ctx, "Update social media successfully", SocialMediaEntityToSocialMediaResponse(response))
}

// DeleteSocialMedia godoc
//
//	@Summary		Delete a social media
//	@Description	Delete a social media by id
//	@Tags			SocialMedia
//	@Accept			json
//	@Produce		json
//	@Param			socialMediaId	path		uint				true	"Social Media ID"
//	@Success		200				{object}	utils.Response		"Social media deleted"
//	@Failure		400				{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401				{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		403				{object}	utils.ErrorResponse	"Forbidden error"
//	@Failure		404				{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500				{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/socialmedias/{socialMediaId} [delete]
//	@Security		BearerAuth
func (sc *SocialMediaController) Delete(ctx *gin.Context) {
	Id := ctx.Param("socialMediaId")
	socialId, _ := strconv.Atoi(Id)

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}

	err := sc.socialService.Delete(uint(socialId), userID.(uint))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Your social media has been successfully deleted", nil)
}
