package service

import (
	"log"

	"github.com/roihan12/h8-mygram/features/socialMedia"
	"github.com/roihan12/h8-mygram/utils"
)

type socialMediaService struct {
	query socialMedia.SocialMediaData
}

func New(q socialMedia.SocialMediaData) socialMedia.SocialMediaService {
	return &socialMediaService{
		query: q,
	}
}

func (sm *socialMediaService) Create(newSocial socialMedia.SocialMediaEntity) (socialMedia.SocialMediaEntity, error) {

	res, err := sm.query.Create(newSocial)
	if err != nil {
		return socialMedia.SocialMediaEntity{}, err
	}

	return res, nil
}

func (sm *socialMediaService) GetAll() ([]socialMedia.SocialMediaEntity, error) {

	res, err := sm.query.GetAll()
	if err != nil {
		log.Println(err)
		return nil, utils.ErrInternal
	}

	return res, nil
}

func (sm *socialMediaService) GetById(socialID uint) (socialMedia.SocialMediaEntity, error) {

	res, err := sm.query.GetById(socialID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return socialMedia.SocialMediaEntity{}, err
		}
		return socialMedia.SocialMediaEntity{}, utils.ErrInternal
	}

	return res, nil
}

func (sm *socialMediaService) Update(updateSocial socialMedia.SocialMediaEntity, socialID, userId uint) (socialMedia.SocialMediaEntity, error) {

	res, err := sm.query.GetById(socialID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return socialMedia.SocialMediaEntity{}, err
		}
		return socialMedia.SocialMediaEntity{}, utils.ErrInternal
	}

	if userId != res.UserID {
		return socialMedia.SocialMediaEntity{}, utils.ErrForbidden
	}

	updateRes, err := sm.query.Update(updateSocial, socialID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return socialMedia.SocialMediaEntity{}, err
		}
		return socialMedia.SocialMediaEntity{}, utils.ErrInternal
	}

	return updateRes, nil
}

func (sm *socialMediaService) Delete(socialID, userId uint) error {
	res, err := sm.query.GetById(socialID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return err
		}
		return utils.ErrInternal
	}

	if userId != res.UserID {
		return utils.ErrForbidden
	}

	err = sm.query.Delete(socialID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return err
		}
		return utils.ErrInternal
	}

	return nil
}
