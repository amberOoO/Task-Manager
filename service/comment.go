package service

import (
	"Task-Manager/storage/models"
	"Task-Manager/utils"
	"errors"

	"gorm.io/gorm"
)

type CommentService struct {
	_db *gorm.DB
}

func NewCommentService() *CommentService {
	return &CommentService{_db: utils.GetTaskDBConnection()}
}

func (cs *CommentService) CreateComment(comment *models.Comment) error {
	// 判断creater是否存在
	err := cs._db.First(&models.User{}, comment.CreaterID).Error
	if err != nil {
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("creater not found")
		}
		return err
	}
	// 判断issue是否存在
	err = cs._db.First(&models.Issue{}, comment.IssueID).Error
	if err != nil {
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("issue not found")
		}
		return err
	}

	err = cs._db.Create(comment).Error
	if err != nil {
		// TODO: 初始err加入日志
		err = errors.New("issue create failed")
	}
	return err
}

func (cs *CommentService) GetCommentByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := cs._db.First(&comment, id).Error
	return &comment, err
}