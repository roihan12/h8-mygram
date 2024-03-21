package data

import (
	"errors"
	"log"

	"github.com/roihan12/h8-mygram/features/socialMedia"
	"github.com/roihan12/h8-mygram/utils"
	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) socialMedia.SocialMediaData {
	return &query{
		db: db,
	}
}

func (q *query) GetAll() ([]socialMedia.SocialMediaEntity, error) {
	var social []SocialMedia

	if err := q.db.Preload("User").Find(&social).Error; err != nil {
		return nil, utils.ErrInternal
	}

	return ListSocialMediaToSocialMediaEntity(social), nil
}

func (q *query) Create(socialMediaEntity socialMedia.SocialMediaEntity) (socialMedia.SocialMediaEntity, error) {
	newSocial := SocialMediaEntityToSocialMedia(socialMediaEntity)
	if err := q.db.Create(&newSocial).Error; err != nil {
		log.Println("create social media query error", err.Error())
		return socialMedia.SocialMediaEntity{}, utils.ErrInternal
	}
	return SocialMediaToSocialMediaEntity(newSocial), nil
}

func (q *query) Update(socialMediaEntity socialMedia.SocialMediaEntity, id uint) (socialMedia.SocialMediaEntity, error) {
	_, err := q.GetById(id)
	if err != nil {
		return socialMedia.SocialMediaEntity{}, err
	}

	updateSocial := SocialMediaEntityToSocialMedia(socialMediaEntity)

	if err := q.db.Where("id = ?", id).Updates(&updateSocial).Error; err != nil {
		log.Println("Update photo query error", err.Error())
		return socialMedia.SocialMediaEntity{}, utils.ErrInternal
	}
	return SocialMediaToSocialMediaEntity(updateSocial), nil
}

func (q *query) GetById(id uint) (socialMedia.SocialMediaEntity, error) {
	var social SocialMedia
	if err := q.db.Preload("User").First(&social, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return socialMedia.SocialMediaEntity{}, utils.ErrDataNotFound
		}
		return socialMedia.SocialMediaEntity{}, err
	}

	return SocialMediaToSocialMediaEntity(social), nil
}

func (q *query) Delete(id uint) error {
	var social SocialMedia
	if err := q.db.Delete(&social, id).Error; err != nil {
		log.Println("delete social media query error", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return utils.ErrInternal
	}

	return nil
}
