package handler

import (
	"mime/multipart"

	"github.com/roihan12/h8-mygram/features/photo"
)

type CreatePhotoRequest struct {
	Title    string                `json:"title" form:"title" binding:"required" example:"My photo"`
	Caption  string                `json:"caption" form:"caption"`
	// PhotoURL *multipart.File `json:"photo_url" form:"photo_url" binding:"required"`
}

type UpdatePhotoRequest struct {
	Title   string `json:"title" form:"title" binding:"required" example:"My photo"`
	Caption string `json:"caption" form:"caption"`
	PhotoURL *multipart.FileHeader `json:"photo_url,omitempty" form:"photo_url"`
}

func CreatePhotoRequestToPhotoEntity(request *CreatePhotoRequest) photo.PhotoEntity {
	return photo.PhotoEntity{
		Title:   request.Title,
		Caption: request.Caption,
	}
}
func UpdatePhotoRequestToPhotoEntity(request *UpdatePhotoRequest) photo.PhotoEntity {
	return photo.PhotoEntity{
		Title:   request.Title,
		Caption: request.Caption,
	}
}
