package service

import (

	"github.com/roihan12/h8-mygram/features/comment"
	"github.com/roihan12/h8-mygram/utils"
)

type commentService struct {
	query comment.CommentData
}

func New(data comment.CommentData) comment.CommentService {
	return &commentService{
		query: data,
	}
}

func (cs *commentService) Create(commentEntity comment.CommentEntity) (comment.CommentEntity, error) {
	newComment, err := cs.query.Create(commentEntity)
	if err != nil {
		return comment.CommentEntity{}, utils.ErrInternal
	}
	return cs.query.GetById(newComment.ID)
}

func (cs *commentService) Update(commentEntity comment.CommentEntity, commentId, userId uint) (comment.CommentEntity, error) {
	checkDataExist, err := cs.query.GetById(commentId)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return comment.CommentEntity{}, err
		}
		return comment.CommentEntity{}, utils.ErrInternal
	}

	if checkDataExist.UserID != userId {
		return comment.CommentEntity{}, utils.ErrForbidden
	}

	_, err = cs.query.Update(commentEntity, commentId)
	if err != nil {
		return comment.CommentEntity{}, utils.ErrInternal
	}
	return cs.query.GetById(commentId)
}

func (cs *commentService) Delete(commentId, userId uint) error {
	checkDataExist, err := cs.query.GetById(commentId)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return err
		}
		return utils.ErrInternal
	}

	if checkDataExist.UserID != userId {
		return utils.ErrForbidden
	}

	err = cs.query.Delete(commentId)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return err
		}
		return utils.ErrInternal
	}

	return nil
}

func (cs *commentService) GetByPhotoID(photoId uint) ([]comment.CommentEntity, error) {
	res, err := cs.query.GetByPhotoID(photoId)
	if err != nil {
		return nil, utils.ErrInternal
	}

	return res, nil
}

func (cs *commentService) MyAllComment(userId uint) ([]comment.CommentEntity, error) {
	res, err := cs.query.MyAllComment(userId)
	if err != nil {
		return nil, utils.ErrInternal
	}
	return res, nil
}

func (cs *commentService) GetById(commentId uint) (comment.CommentEntity, error) {
	res, err := cs.query.GetById(commentId)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return comment.CommentEntity{}, err
		}
		return comment.CommentEntity{}, utils.ErrInternal
	}


	return res, nil
}
