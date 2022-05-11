package milestoneView

import (
	"Task-Manager/storage/models"
	"gorm.io/gorm"
)

// milestone
type CreateMilestoneMsg struct {
	ID			uint		`json:"id"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	Issues		[]uint		`json:"issues"`
}

func (cmm CreateMilestoneMsg) GetIssueArr() []models.Issue{
	IssueArr := make([]models.Issue, len(cmm.Issues))
	for index, issueID := range cmm.Issues {
		IssueArr[index] = models.Issue{
			Model: gorm.Model{
				ID: issueID,
			},
		}
	}
	return IssueArr
}

type UpdateMilestoneMsg struct {
	ID			uint		`json:"id"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
}