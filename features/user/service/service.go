package services

import (
	"log"

	"github.com/roihan12/h8-mygram/features/user"
	"github.com/roihan12/h8-mygram/utils"
)

type userUseCase struct {
	query user.UserData
}

func New(query user.UserData) user.UserService {
	return &userUseCase{
		query: query,
	}
}

func (uuc *userUseCase) Register(newUser user.UserEntity) (user.UserEntity, error) {
	hashed, err := utils.GeneratePassword(newUser.Password)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return user.UserEntity{}, utils.ErrInternal
	}

	newUser.Password = string(hashed)
	res, err := uuc.query.Register(newUser)
	if err != nil {
		if err == utils.ErrConflictingData {
			return user.UserEntity{}, utils.ErrDuplicateData
		}
		return user.UserEntity{}, utils.ErrInternal
	}

	return res, nil
}

func (uuc *userUseCase) Login(email, password string) (string, error) {
	res, err := uuc.query.Login(email)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return "", utils.ErrInvalidCredentials
		}
		log.Println("error login query: ", err.Error())
		return "", utils.ErrInternal
	}

	if err := utils.CheckPassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", utils.ErrInvalidCredentials
	}

	useToken, err := utils.GenerateToken(res.Username, res.ID)
	if err != nil {
		return "", utils.ErrTokenCreation
	}
	return useToken, nil
}

func (uuc *userUseCase) Profile(UserID uint) (user.UserEntity, error) {

	res, err := uuc.query.Profile(UserID)
	if err != nil {
		log.Println("data not found")
		if err == utils.ErrDataNotFound {
			return user.UserEntity{}, err
		}
		return user.UserEntity{}, utils.ErrInternal
	}

	return res, nil
}

func (uuc *userUseCase) Update(UserID uint, updateData user.UserEntity) (user.UserEntity, error) {
	if updateData.Password != "" {
		hashed, err := utils.GeneratePassword(updateData.Password)
		if err != nil {
			log.Println("bcrypt error ", err.Error())
			return user.UserEntity{}, utils.ErrInternal
		}

		updateData.Password = string(hashed)
	}

	res, err := uuc.query.Update(UserID, updateData)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return user.UserEntity{}, err
		}
		return user.UserEntity{}, utils.ErrInternal
	}

	return res, nil
}

func (uuc *userUseCase) Delete(UserID uint) error {
	err := uuc.query.Delete(UserID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return err
		}
		return utils.ErrInternal
	}

	return nil
}
