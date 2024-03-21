package handler

import (
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/roihan12/h8-mygram/features/photo"
	"github.com/roihan12/h8-mygram/utils"
)

type PhotoController struct {
	photoService photo.PhotoService
}

func New(srv photo.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: srv,
	}
}

// ListPhotos godoc
//
//	@Summary		List photos
//	@Description	List photos
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		PhotoResponse		"Photos retrieved"
//	@Failure		500	{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/photos [get]
//	@Security		BearerAuth
func (pc *PhotoController) GetAll(ctx *gin.Context) {

	photos, err := pc.photoService.GetAll()
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	listPhotoResponse := ListPhotoEntityToPhotoResponse(photos)

	utils.HandleSuccess(ctx, "Get all photo successfully", listPhotoResponse)
}

// GetPhoto godoc
//
//	@Summary		Get a photo
//	@Description	get a photo by id
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
//	@Param			photoId	path		uint				true	"Photo ID"
//	@Success		200		{object}	PhotoResponse		"Photo retrieved"
//	@Failure		400		{object}	utils.ErrorResponse	"Validation error"
//	@Failure		404		{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/photos/{photoId} [get]
//	@Security		BearerAuth
func (pc *PhotoController) GetById(ctx *gin.Context) {
	Id := ctx.Param("photoId")
	photoId, _ := strconv.Atoi(Id)
	photo, err := pc.photoService.GetById(uint(photoId))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	response := PhotoEntityToPhotoResponse(photo)
	utils.HandleSuccess(ctx, "Get by id photo successfully", response)
}

// CreatePhoto
//
//	@Summary		Create a new photo Upload file
//	@Description	create a new photo with title, caption and photo url
//	@Tags			Photos
//	@ID				file.upload
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			photo_url			formData	file				true	"this is image file"
//	@Param			CreatePhotoRequest	formData	CreatePhotoRequest	true	"Create photo request"
//	@Success		200					{object}	PhotoResponse		"Photo retrieved"
//	@Failure		400					{object}	utils.ErrorResponse	"Validation error"
//	@Failure		404					{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500					{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/photos [post]
//	@Security		BearerAuth
func (pc *PhotoController) Create(ctx *gin.Context) {
	var req CreatePhotoRequest
	var photoImage *multipart.FileHeader
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

	photoEntity := CreatePhotoRequestToPhotoEntity(&req)
	photoEntity.UserID = userID.(uint)

	file, err := ctx.FormFile("photo_url")
	if err != nil {
		if err == http.ErrMissingFile {
			// Kembalikan error jika file gambar tidak diunggah
			utils.HandleError(ctx, utils.ErrImageRequired)
			return
		}
		// Penanganan kesalahan lainnya
		utils.HandleError(ctx, err)
		return
	}
	err = utils.CheckFile(file)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	photoImage = file

	response, err := pc.photoService.Create(photoEntity, photoImage)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Create photo successfully", PhotoEntityToPhotoResponse(response))
}
// UpdatePhoto
//
//	@Summary		Updated a  photo Upload file
//	@Description	update a photo with title, caption and photo url
//	@Tags			Photos
//	@ID				file.update
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			photoId				path		uint				true	"Photo ID"
//	@Param			photo_url			formData	file				false	"this is image file"
//	@Param			CreatePhotoRequest	formData	CreatePhotoRequest	true	"Update photo request"
//	@Success		200					{object}	PhotoResponse		"Photo retrieved"
//	@Failure		400					{object}	utils.ErrorResponse	"Validation error"
//	@Failure		404					{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500					{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/photos/{photoId} [put]
//	@Security		BearerAuth
func (pc *PhotoController) Update(ctx *gin.Context) {

	Id := ctx.Param("photoId")
	photoId, _ := strconv.Atoi(Id)
	var req UpdatePhotoRequest
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

	photoEntity := UpdatePhotoRequestToPhotoEntity(&req)

	// Periksa apakah ada file gambar yang diunggah
	file, err := ctx.FormFile("photo_url")
	if file != nil {
		req.PhotoURL = file
	} else {
		req.Title = ctx.Request.FormValue("title")
		req.Caption = ctx.Request.FormValue("caption")
		req.PhotoURL = nil
	}

	if err != nil {
		if err != http.ErrMissingFile {
			// Penanganan kesalahan lainnya
			utils.HandleError(ctx, err)
			return
		}
	}

	response, err := pc.photoService.Update(photoEntity, uint(photoId), userID.(uint), req.PhotoURL)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	utils.HandleSuccess(ctx, "Update photo successfully", PhotoEntityToPhotoResponse(response))
}

// DeletePhoto godoc
//
//	@Summary		Delete a photo
//	@Description	delete a photo by id
//	@Tags			Photos
//	@Accept			json
//	@Produce		json
//	@Param			photoId	path		uint				true	"Photo ID"
//	@Success		200		{object}	utils.Response		"Photo deleted"
//	@Failure		400		{object}	utils.ErrorResponse	"Validation error"
//	@Failure		404		{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/photos/{photoId} [delete]
//	@Security		BearerAuth
func (pc *PhotoController) Delete(ctx *gin.Context) {
	Id := ctx.Param("photoId")
	photoId, _ := strconv.Atoi(Id)

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		utils.HandleAbort(ctx, err)
		return
	}

	err := pc.photoService.Delete(uint(photoId), userID.(uint))
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}
	utils.HandleSuccess(ctx, "Your photo has been successfully deleted", nil)
}
