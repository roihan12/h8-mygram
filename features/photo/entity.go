package photo

import (
	"mime/multipart"
	"time"

	"github.com/roihan12/h8-mygram/features/user"
)

type PhotoEntity struct {
	ID        uint
	Title     string
	Caption   string
	PhotoURL  string
	UserID    uint
	User      user.UserEntity
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments   []Comment
}

type Comment struct {
	ID        uint
	Message   string
	UserID    uint
	PhotoID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.UserEntity
}
type PhotoService interface {
	GetAll() ([]PhotoEntity, error)
	GetById(id uint) (PhotoEntity, error)
	Create(createPhoto PhotoEntity, image *multipart.FileHeader) (PhotoEntity, error)
	Update(updatePhoto PhotoEntity, id, userId uint, image *multipart.FileHeader) (PhotoEntity, error)
	Delete(id, userId uint) error
}

type PhotoData interface {
	GetAll() ([]PhotoEntity, error)
	GetById(id uint) (PhotoEntity, error)
	Create(createPhoto PhotoEntity) (PhotoEntity, error)
	Update(updatePhoto PhotoEntity, id uint) (PhotoEntity, error)
	Delete(id uint) error
}
