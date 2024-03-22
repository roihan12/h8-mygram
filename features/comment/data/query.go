package data

import (
	"errors"
	"log"

	"github.com/roihan12/h8-mygram/features/comment"
	"github.com/roihan12/h8-mygram/utils"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.CommentData {
	return &query{
		db: db,
	}
}

func (qf *query) Create(commentEntity comment.CommentEntity) (comment.CommentEntity, error) {
	newComment := CommentEntityToComment(commentEntity)
	err := qf.db.Create(&newComment).Error
	if err != nil {
		log.Println("create photo query error", err.Error())
		return comment.CommentEntity{}, utils.ErrInternal
	}
	return CommentToCommentEntity(newComment), nil
}

func (qf *query) GetById(id uint) (comment.CommentEntity, error) {
	var getComment Comment
	if err := qf.db.Preload("User").Preload("Photo").First(&getComment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return comment.CommentEntity{}, utils.ErrDataNotFound
		}
		return comment.CommentEntity{}, err
	}
	
	return CommentToCommentEntity(getComment), nil
}

func (qf *query) Update(commentEntity comment.CommentEntity, id uint) (comment.CommentEntity, error) {
	updateComment := CommentEntityToComment(commentEntity)
	if err := qf.db.Where("id = ?", id).Updates(&updateComment).Error; err != nil {
		log.Println("Update photo query error", err.Error())
		return comment.CommentEntity{}, utils.ErrInternal
	}
	return CommentToCommentEntity(updateComment), nil
}

func (qf *query) Delete(id uint) error {
	var delComment Comment
	if err := qf.db.Delete(&delComment, id).Error; err != nil {
		log.Println("delete photo query error", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return utils.ErrInternal
	}

	return nil
}

func (qf *query) GetByPhotoID(photoId uint) ([]comment.CommentEntity, error) {
	photoComment := []Comment{}
	if err := qf.db.Where("photo_id = ?", photoId).Preload("User").Preload("Photo").Find(&photoComment).Error; err != nil {
		return []comment.CommentEntity{}, utils.ErrInternal
	}
	return ListCommentToCommentEntity(photoComment), nil
}

func (qf *query) MyAllComment(userId uint) ([]comment.CommentEntity, error) {
	var myComment []Comment
	if err := qf.db.Where("user_id = ?", userId).Preload("User").Preload("Photo").Find(&myComment); err.Error != nil {
		return nil, err.Error
	}
	return ListCommentToCommentEntity(myComment), nil
}
