package service

import (
	"Task-Manager/storage/models"
	"Task-Manager/utils"
	"errors"

	"gorm.io/gorm"
)

type MilestoneService struct {
	_db	*gorm.DB
}

func NewMilestoneService() *MilestoneService {
	return &MilestoneService{_db: utils.GetTaskDBConnection()}
}

func (ms *MilestoneService) CreateMilestone(milestone *models.Milestone) error {
	// 查看关联的Issue是否存在
	for _, issue := range milestone.Issues {
		err := ms._db.First(&models.Issue{}, issue.ID).Error
		if err != nil {
			// TODO: 初始err加入日志
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = errors.New("issue not found")
			}
			return err
		}
	}
	// 尝试创建
	err := ms._db.Create(milestone).Error
	if err != nil {
		// TODO: 初始err加入日志
		err = errors.New("milestone create failed")
	}
	return err
}