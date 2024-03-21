package comment

import (
	"time"

	"github.com/roihan12/h8-mygram/features/photo"
	"github.com/roihan12/h8-mygram/features/user"
)

type CommentEntity struct {
	ID        uint
	Message   string
	UserID    uint
	PhotoID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Photo     photo.PhotoEntity
	User      user.UserEntity
}

type CommentService interface {
	Create(commentEntity CommentEntity) (CommentEntity, error)
	Update(commentEntity CommentEntity, id, userId uint) (CommentEntity, error)
	Delete(id, userId uint) error
	GetByPhotoID(photoId uint) ([]CommentEntity, error)
	MyAllComment(userId uint) ([]CommentEntity, error)
	GetById(id uint) (CommentEntity, error)
}

type CommentData interface {
	GetById(id uint) (CommentEntity, error)
	Create(commentEntity CommentEntity) (CommentEntity, error)
	Update(commentEntity CommentEntity, id uint) (CommentEntity, error)
	Delete(id uint) error
	GetByPhotoID(photoId uint) ([]CommentEntity, error)
	MyAllComment(userId uint) ([]CommentEntity, error)
}
