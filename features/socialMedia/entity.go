package socialMedia

import (
	"time"

	"github.com/roihan12/h8-mygram/features/user"
)

type SocialMediaEntity struct {
	ID             uint
	Name           string
	SocialMediaURL string
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           user.UserEntity
}

type SocialMediaService interface {
	GetAll() ([]SocialMediaEntity, error)
	GetById(id uint) (SocialMediaEntity, error)
	Create(createSocial SocialMediaEntity) (SocialMediaEntity, error)
	Update(updateSocial SocialMediaEntity, id, userId uint) (SocialMediaEntity, error)
	Delete(id, userId uint) error
}

type SocialMediaData interface {
	GetAll() ([]SocialMediaEntity, error)
	GetById(id uint) (SocialMediaEntity, error)
	Create(createSocial SocialMediaEntity) (SocialMediaEntity, error)
	Update(updateSocial SocialMediaEntity, id uint) (SocialMediaEntity, error)
	Delete(id uint) error
}
