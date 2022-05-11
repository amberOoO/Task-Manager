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
			}else{
				err = errors.New("milestone create failed")
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

func (ms *MilestoneService) UpdateMilestone(milestone *models.Milestone) error {
	// 判断milestone是否存在
	err := ms._db.First(&models.Milestone{}, milestone.ID).Error
	if err != nil {
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("milestone not found")
		}else{
			return errors.New("milestone update failed")
		}
	}
	// 尝试更新
	err = ms._db.Model(&milestone).Updates(&milestone).Error
	if err != nil {
		// TODO: 初始err加入日志
		err = errors.New("milestone update failed")
	}
	return err
}

func (ms *MilestoneService) DeleteMilestone(milestone *models.Milestone) error {
	// 判断milestone是否存在
	err := ms._db.First(&models.Milestone{}, milestone.ID).Error
	if err != nil {
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("milestone not found")
		}else{
			return errors.New("milestone update failed")
		}
	}
	// 将issue表中，milestone关联的record恢复默认值0
	// FIXME: 感觉这样做，不是很理想。
	err = ms._db.Model(&models.Issue{}).Where("milestone_id = ?", milestone.ID).Update("milestone_id", 0).Error
	if err != nil {
		// TODO: 初始err加入日志
		return errors.New("milestone delete failed")
	}

	// 尝试删除milestone
	err = ms._db.Model(&milestone).Delete(&milestone).Error
	if err != nil {
		// TODO: 初始err加入日志
		err = errors.New("milestone delete failed")
	}
	return err
}