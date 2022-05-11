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
	var (
		issueID uint = issue.ID
	)
	// 判断是否存在issue
	err := iss._db.First(&models.Issue{}, issueID).Error
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
	
	err = iss._db.Model(&issue).Updates(&issue).Error
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
	// 级联删除issue_tag表
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

func (iss *IssueService) UpdateMilestone(issue *models.Issue, milestoneID uint) error {
	// 查看issue是否存在
	err := iss._db.First(issue).Error
	if err != nil {
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("issue not found")
		}else{
			err = errors.New("issue update failed")
		}
		return err
	}
	// 查看milestone是否存在
	err = iss._db.First(&models.Milestone{}, milestoneID).Error
	if err != nil {
		// TODO: 初始err加入日志
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("milestone not found")
		}else{
			err = errors.New("issue update failed")
		}
		return err
	}
	// 开始更新milestone
	err = iss._db.Model(issue).Update("milestone_id", milestoneID).Error
	if err != nil {
		err = errors.New("issue update failed")
	}
	return err
}

// 查询条件
type QueryConfig struct {
	Tags		 []string
	MilestoneIDs []uint
}

func (iss *IssueService) GetIssueWithCondition(page int, pageSize int, config QueryConfig) ([]models.Issue, int64, error) {
	var (
		issues   []models.Issue
		queries  *gorm.DB
		totalNum int64
	)
	// 初始化查询
	queries = iss._db.Model(&models.Issue{})
	// 添加tags过滤
	// if config.Tags != nil && len(config.Tags) != 0 {
	// 	queries = queries.Preload("Tags", "content in (?)", config.Tags).Joins("Tag", "issue.id = issue_tags.issue_id").Where("tags IN (?)", config.Tags)
	// }else{
	// 	queries = queries.Preload("Tags").Model(&models.Issue{})
	// }
	// 添加milestone过滤
	if config.MilestoneIDs != nil && len(config.MilestoneIDs) != 0 {
		queries = queries.Where("milestone_id IN (?)", config.MilestoneIDs)
	}
	// 加入分页选项
	queries = queries.Limit(pageSize).Offset((page - 1) * pageSize)

	// 开始查询数据
	// 获取数量
	err := queries.Count(&totalNum).Error
	if err != nil {
		// return issues, 0, errors.New("issue get failed")
		return issues, 0, err
	}
	// 获取数据
	err = queries.Find(&issues).Error
	if err != nil {
		err = errors.New("issue get failed")
	}
	return issues, totalNum, err
}

func (iss *IssueService) GetTotalNum() (int64, error) {
	var (
		err 	 error
		totalNum int64
	)
	err = iss._db.Model(&models.Issue{}).Count(&totalNum).Error
	if err != nil {
		err = errors.New("issue total num get failed")
	}
	return totalNum, err
}

func (iss *IssueService) GetIssueByID(id uint) (*models.Issue, error) {
	var issue models.Issue
	err := iss._db.First(&issue, id).Error
	return &issue, err
}