package view

import (
	"Task-Manager/storage/models"

	"gorm.io/gorm"
)

// user
type CreateUserMsg struct {
	Nickname  string	`json:"nickname"`
}

// issue
type CreateIssueMsg struct {
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	CreaterID	uint		`json:"createrID"`
	MilestoneID	uint		`json:"milestoneID"`
	Tags		[]string	`json:"tags"`
}

type UpdateIssueMsg struct {
	ID			uint		`json:"id"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	CreaterID	uint		`json:"createrID"`
	MilestoneID	uint		`json:"milestoneID"`
}

type UpdateIssueTagsMsg struct {
	ID			uint		`json:"id"`
	Tags		[]string	`json:"tags"`
}

type DeleteIssueMsg struct {
	ID			uint		`json:"id"`
}

// comment
type CreateCommentMsg struct {
	ID			uint		`json:"id"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	CreaterID	uint		`json:"createrID"`
	IssueID		uint		`json:"issueID"`
}

// tag
type CreateTagMsg struct {
	Content string `json:"content"`
}

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