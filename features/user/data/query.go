package data

import (
	"log"

	"github.com/roihan12/h8-mygram/features/user"
	"github.com/roihan12/h8-mygram/utils"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) CheckEmail(newUser user.UserEntity) error {
	u := User{}
	uq.db.Where("email = ?", newUser.Email).First(&u)
	if u.ID != 0 {
		if u.Email == newUser.Email {
			return utils.ErrConflictingData
		}
	}
	return nil
}
func (uq *userQuery) CheckUsername(newUser user.UserEntity) error {
	u := User{}
	uq.db.Where("username = ?", newUser.Username).First(&u)
	if u.ID != 0 {
		if u.Email == newUser.Email {
			return utils.ErrConflictingData
		}
	}
	return nil
}

func (uq *userQuery) Login(email string) (user.UserEntity, error) {

	res := User{}
	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.UserEntity{}, utils.ErrDataNotFound
	}

	return ToUserEntity(res), nil
}

func (uq *userQuery) Register(newUser user.UserEntity) (user.UserEntity, error) {
	// Chek User
	if err := uq.CheckEmail(newUser); err != nil {
		log.Println("error create new user: ", err.Error())
		return user.UserEntity{}, err
	}

	if err := uq.CheckUsername(newUser); err != nil {
		log.Println("error create new user: ", err.Error())
		return user.UserEntity{}, err
	}

	cnv := UserEntityToUser(newUser)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.UserEntity{}, utils.ErrInternal
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) Profile(userID uint) (user.UserEntity, error) {
	res := User{}
	if err := uq.db.Where("id=?", userID).First(&res).Error; err != nil {
		log.Println("get profile query error", err.Error())
		return user.UserEntity{}, utils.ErrInternal
	}

	return ToUserEntity(res), nil
}

func (uq *userQuery) Update(userID uint, updateData user.UserEntity) (user.UserEntity, error) {
	cnv := UserEntityToUser(updateData)
	res := User{}
	qry := uq.db.Model(&res).Where("id = ?", userID).Updates(&cnv)

	if qry.RowsAffected <= 0 {
		log.Println("\tupdate user query error: data not found")
		return user.UserEntity{}, utils.ErrDataNotFound
	}

	if err := qry.Error; err != nil {
		log.Println("\tupdate user query error: ", err.Error())
		return user.UserEntity{}, utils.ErrDataNotFound
	}

	return ToUserEntity(cnv), nil
}

func (uq *userQuery) Delete(userID uint) error {
	res := User{}
	qry := uq.db.Delete(&res, userID)

	if qry.RowsAffected <= 0 {
		log.Println("\tDelete user query error: data not found")
		return utils.ErrDataNotFound
	}
	if err := qry.Error; err != nil {
		log.Println("\tDelete user query error: ", err.Error())
		return utils.ErrInternal
	}

	return nil
}
