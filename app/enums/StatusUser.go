package enums

type StatusUser string

const (
	Online 			StatusUser = "active"
	Offline 		StatusUser = "inactive"
	Suspended 		StatusUser = "banned"
	Pending 		StatusUser = "pending"
)