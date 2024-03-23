package data

import (
	"errors"
	"log"

	"github.com/roihan12/h8-mygram/features/photo"
	"github.com/roihan12/h8-mygram/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) photo.PhotoData {
	return &query{
		db: db,
	}
}

func (q *query) GetAll() ([]photo.PhotoEntity, error) {
	var photos []Photo

	if err := q.db.Preload("User").Find(&photos).Error; err != nil {
		return nil, err
	}

	return ListPhotoToPhotoEntity(photos), nil
}

func (q *query) Create(photoEntity photo.PhotoEntity) (photo.PhotoEntity, error) {
	newPhoto := PhotoEntityToPhoto(photoEntity)
	if err := q.db.Create(&newPhoto).Preload("User").Error; err != nil {
		log.Println("create photo query error", err.Error())
		return photo.PhotoEntity{}, utils.ErrInternal
	}
	return PhotoToPhotoEntity(newPhoto), nil
}

func (q *query) Update(photoEntity photo.PhotoEntity, id uint) (photo.PhotoEntity, error) {
	_, err := q.GetById(id)
	if err != nil {
		return photo.PhotoEntity{}, err
	}

	updatePhoto := PhotoEntityToPhoto(photoEntity)

	if err := q.db.Where("id = ?", id).Updates(&updatePhoto).Preload("User").Error; err != nil {
		log.Println("Update photo query error", err.Error())
		return photo.PhotoEntity{}, utils.ErrInternal
	}
	return PhotoToPhotoEntity(updatePhoto), nil
}

func (q *query) GetById(id uint) (photo.PhotoEntity, error) {
	var singlePhoto Photo
	if err := q.db.Preload("User").First(&singlePhoto, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return photo.PhotoEntity{}, utils.ErrDataNotFound
		}
		return photo.PhotoEntity{}, err
	}

	return PhotoToPhotoEntity(singlePhoto), nil
}

func (q *query) Delete(id uint) error {
	var photo Photo
	if err := q.db.Select(clause.Associations).Delete(&photo, id).Error; err != nil {
		log.Println("delete photo query error", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return utils.ErrInternal
	}

	return nil
}
