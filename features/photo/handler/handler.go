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

func (pc *PhotoController) GetAll(ctx *gin.Context) {

	photos, err := pc.photoService.GetAll()
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	listPhotoResponse := ListPhotoEntityToPhotoResponse(photos)

	utils.HandleSuccess(ctx, "Get all photo successfully", listPhotoResponse)
}

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
