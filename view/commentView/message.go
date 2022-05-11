package commentView

// comment
type CreateCommentMsg struct {
	ID			uint		`json:"id"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	CreaterID	uint		`json:"createrID"`
	IssueID		uint		`json:"issueID"`
}