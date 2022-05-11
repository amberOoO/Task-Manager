package service

import (
	"Task-Manager/storage/models"
	"Task-Manager/utils"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	_db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{_db: utils.GetTaskDBConnection()}
}

func (iss *UserService) CreateUser(user *models.User) error {
	err := iss._db.Create(user).Error
	if err != nil {
		err = errors.New("user create failed")
	}
	return err
}

func (iss *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := iss._db.First(&user, id).Error
	return &user, err
}