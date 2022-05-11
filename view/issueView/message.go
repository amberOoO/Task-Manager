package issueView

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

type UpdateIssueMilestoneMsg struct {
	ID			uint		`json:"id"`
	MilestoneID	uint		`json:"milestoneID"`
}

type UpdateIssueTagsMsg struct {
	ID			uint		`json:"id"`
	Tags		[]string	`json:"tags"`
}

type DeleteIssueMsg struct {
	ID			uint		`json:"id"`
}

type GetIssueWithConditionMsg struct {
	Page			int			`json:"page"`
	PageSize		int			`json:"pageSize"`
	Tags			[]string	`json:"tags"`
	MilestoneIDs	[]uint		`json:"milestoneIDs"`
	// SortKey		string		`json:"sortKey"`
	// Ascending	bool		`json:"ascending"`
}