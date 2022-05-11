package service

import (
	"Task-Manager/storage/models"
	"Task-Manager/utils"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IssueService struct {
	_db *gorm.DB
}

func NewIssueService() *IssueService {
	return &IssueService{_db: utils.GetTaskDBConnection()}
}

func (iss *IssueService) CreateIssue(issue *models.Issue) error {
	// 判断是否存在用户
	err := iss._db.First(&models.User{}, issue.CreaterID).Error
	if err != nil{
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("user not found")
		}
		return err
	}
	// 判断里程碑是否存在（0为默认值，即为空）
	if issue.MilestoneID != 0 {
		err := iss._db.First(&models.Milestone{}, issue.MilestoneID).Error
		if err != nil{
			// TODO: 初始err加入日志
			// 不存在milestone会特判
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = errors.New("milestone not found")
			}
			return err
		}
	}

	err = iss._db.Create(&issue).Error
	if err != nil {
		// TODO: 初始err加入日志
		err = errors.New("issue create failed")
	}

	return err
}

func (iss *IssueService) UpdateIssue(issue *models.Issue) error {
	// 判断是否存在issue
	err := iss._db.First(&models.Issue{}, issue.ID).Error
	if err != nil{
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("issue not found")
		}
		return err
	}
	// 判断是否存在用户
	err = iss._db.First(&models.User{}, issue.CreaterID).Error
	if err != nil{
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("user not found")
		}
		return err
	}
	// 判断里程碑是否存在（0为默认值，即为空）
	if issue.MilestoneID != 0 {
		err := iss._db.First(&models.Milestone{}, issue.MilestoneID).Error
		if err != nil{
			// TODO: 初始err加入日志
			// 不存在milestone会特判
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = errors.New("milestone not found")
			}
			return err
		}
	}

	err = iss._db.Save(&issue).Error
	if err != nil {
		err = errors.New("issue update failed")
	}

	return err
}

func (iss *IssueService) UpdateIssueTags(issue *models.Issue) error {
	// 判断是否存在issue
	err := iss._db.First(&models.Issue{}, issue.ID).Error
	if err != nil{
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("issue not found")
		}
		return err
	}
	err = iss._db.Model(&issue).Association("Tags").Replace(issue.Tags)
	if err != nil {
		err = errors.New("issue tags update failed")
	}
	return err
}

func (iss *IssueService) DeleteIssue(issue *models.Issue) error {
	// 判断issue是否存在
	err := iss._db.Preload("Tags").First(&issue).Error
	if err != nil {
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("issue not found")
		}
		return err
	}
	err = iss._db.Select(clause.Associations).Delete(&issue).Error
	if err != nil {
		err = errors.New("issue delete failed")
	}
	return err
}

func (iss *IssueService) GetIssueByID(id uint) (*models.Issue, error) {
	var issue models.Issue
	err := iss._db.First(&issue, id).Error
	return &issue, err
}