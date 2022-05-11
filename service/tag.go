package service

import (
	"Task-Manager/storage/models"
	"Task-Manager/utils"
	"errors"

	"gorm.io/gorm"
)

type TagService struct {
	_db *gorm.DB
}

func NewTagService() *TagService {
	return &TagService{_db: utils.GetTaskDBConnection()}
}

func (ts *TagService) CreateTag(tag *models.Tag) error {
	err := ts._db.Create(tag).Error
	if err != nil {
		err = errors.New("tag create failed")
	}
	return err
}