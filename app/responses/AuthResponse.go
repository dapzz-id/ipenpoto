package responses

import "github.com/google/uuid"

type LoginResponse struct {
	ID       uuid.UUID  	`json:"id"`
	Username string 		`json:"username"`
	Role     string 		`json:"role"`
}