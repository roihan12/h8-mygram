package data

import (
	"time"

	"github.com/roihan12/h8-mygram/features/comment"
	"github.com/roihan12/h8-mygram/features/photo"
	photoModel "github.com/roihan12/h8-mygram/features/photo/data"
	"github.com/roihan12/h8-mygram/features/user"
	userModel "github.com/roihan12/h8-mygram/features/user/data"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message   string `gorm:"not null;type:varchar(255)"`
	UserID    uint
	PhotoID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Photo     photoModel.Photo `gorm:"foreignkey:PhotoID;association_foreignkey:ID"`
	User      userModel.User   `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}

func CommentEntityToComment(data comment.CommentEntity) Comment {
	return Comment{
		Model:     gorm.Model{ID: data.ID},
		Message:   data.Message,
		UserID:    data.UserID,
		PhotoID:   data.PhotoID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func CommentToCommentEntity(data Comment) comment.CommentEntity {
	result := comment.CommentEntity{
		ID:        data.ID,
		Message:   data.Message,
		UserID:    data.UserID,
		PhotoID:   data.PhotoID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	result.Photo = photo.PhotoEntity{
		ID:        data.Photo.ID,
		Title:     data.Photo.Title,
		Caption:   data.Photo.Caption,
		PhotoURL:  data.Photo.PhotoURL,
		UserID:    data.Photo.UserID,
		CreatedAt: data.Photo.CreatedAt,
		UpdatedAt: data.Photo.UpdatedAt,
	}

	result.User = user.UserEntity{
		ID:       data.User.ID,
		Email:    data.User.Email,
		Username: data.User.Username,
	}

	return result
}

func ListCommentToCommentEntity(data []Comment) []comment.CommentEntity {
	var commentEntity []comment.CommentEntity

	for _, v := range data {
		commentEntity = append(commentEntity, CommentToCommentEntity(v))

	}
	return commentEntity
}
