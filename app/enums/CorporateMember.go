package enums

type CorporateRole string

const (
	CorporateAdmin  	CorporateRole = "owner"
	CorporateUser   	CorporateRole = "admin"
	CorporateStaff  	CorporateRole = "staff"
	CorporateViewer 	CorporateRole = "viewer"
)