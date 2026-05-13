package requests

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Name     				string 			`json:"name" validate:"required"`
	Username 				string 			`json:"username" validate:"required"`
	Email    				string 			`json:"email" validate:"required,email"`
	Password 				string 			`json:"password" validate:"required,min=6"`
	PasswordConfirmation 	string 			`json:"password_confirmation" validate:"required,eqfield=Password"`
}