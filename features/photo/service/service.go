package service

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/roihan12/h8-mygram/features/comment"
	"github.com/roihan12/h8-mygram/features/photo"
	"github.com/roihan12/h8-mygram/utils"
)

type photoService struct {
	query  photo.PhotoData
	cquery comment.CommentService
	upload utils.Uploader
}

func New(q photo.PhotoData, u utils.Uploader, c comment.CommentService) photo.PhotoService {
	return &photoService{
		query:  q,
		cquery: c,
		upload: u,
	}
}

func (ps *photoService) Create(newPhoto photo.PhotoEntity, image *multipart.FileHeader) (photo.PhotoEntity, error) {

	if image != nil {
		// chech file upload
		err := utils.CheckFile(image)
		if err != nil {
			return photo.PhotoEntity{}, err
		}
	}

	imageURL, err := ps.upload.Upload(image)
	if err != nil {
		log.Println(err)
		return photo.PhotoEntity{}, utils.ErrInternal
	}

	newPhoto.PhotoURL = imageURL
	resPhoto, err := ps.query.Create(newPhoto)
	if err != nil {
		return photo.PhotoEntity{}, err
	}

	return resPhoto, nil
}

func (ps *photoService) GetAll() ([]photo.PhotoEntity, error) {

	res, err := ps.query.GetAll()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrInternal
	}

	return res, nil
}

func (ps *photoService) GetById(photoID uint) (photo.PhotoEntity, error) {

	res, err := ps.query.GetById(photoID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return photo.PhotoEntity{}, err
		}
		return photo.PhotoEntity{}, utils.ErrInternal
	}

	comments, err := ps.cquery.GetByPhotoID(photoID)
	if err != nil && err != utils.ErrDataNotFound {
		return photo.PhotoEntity{}, utils.ErrInternal
	}

	// Konversi tipe data komentar
	var photoComments []photo.Comment
	for _, c := range comments {
		photoComments = append(photoComments, photo.Comment{
			ID:        c.ID,
			PhotoID:   c.PhotoID,
			UserID:    c.UserID,
			Message:   c.Message,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			User:      c.User,
		})
	}

	// Memperbarui field Comments pada objek foto res dengan komentar yang diperoleh.
	res.Comments = photoComments
	fmt.Println(photoComments)
	return res, nil
}

func (ps *photoService) Update(updatePhoto photo.PhotoEntity, photoID, userId uint, image *multipart.FileHeader) (photo.PhotoEntity, error) {

	res, err := ps.query.GetById(photoID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return photo.PhotoEntity{}, err
		}
		return photo.PhotoEntity{}, utils.ErrInternal
	}

	if userId != res.UserID {
		return photo.PhotoEntity{}, utils.ErrForbidden
	}

	if image != nil {
		// chech file upload
		err := utils.CheckFile(image)
		if err != nil {
			return photo.PhotoEntity{}, err
		}

		imageURL, err := ps.upload.Upload(image)
		if err != nil {
			log.Println(err)
			return photo.PhotoEntity{}, utils.ErrInternal
		}
		updatePhoto.PhotoURL = imageURL
	}

	updateRes, err := ps.query.Update(updatePhoto, photoID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return photo.PhotoEntity{}, err
		}
		return photo.PhotoEntity{}, utils.ErrInternal
	}

	if image != nil && res.PhotoURL != "" {
		publicID := utils.GetPublicID(res.PhotoURL)
		if err := ps.upload.Destroy(publicID); err != nil {
			log.Println("destroy file", err)
			return photo.PhotoEntity{}, utils.ErrInternal
		}
	}

	return updateRes, nil
}

func (ps *photoService) Delete(photoID, userId uint) error {
	res, err := ps.query.GetById(photoID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return err
		}
		return utils.ErrInternal
	}

	if userId != res.UserID {
		return utils.ErrForbidden
	}

	if res.PhotoURL != "" {
		publicID := utils.GetPublicID(res.PhotoURL)
		if err := ps.upload.Destroy(publicID); err != nil {
			log.Println("destroy file", err)
			return utils.ErrInternal
		}
	}

	err = ps.query.Delete(photoID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return err
		}
		return utils.ErrInternal
	}

	return nil
}
