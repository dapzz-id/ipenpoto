package enums

type ApprovalStatus string

const (
	PendingApprove 		ApprovalStatus = "pending"
	ApprovedApprove 	ApprovalStatus = "approved"
	RejectedApprove 	ApprovalStatus = "rejected"
	CancelledApprove 	ApprovalStatus = "cancelled"
)