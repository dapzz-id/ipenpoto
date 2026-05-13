package utils

import (
	"regexp"
	"strings"
	"unicode"
)

type ValidationRules struct {
	Username UsernameRules
	Password PasswordRules
	Email    EmailRules
}

type UsernameRules struct {
	MinLength    int
	MaxLength    int
	AllowedChars string
}

type PasswordRules struct {
	MinLength           int
	RequireUppercase    bool
	RequireLowercase    bool
	RequireNumbers      bool
	RequireSpecialChars bool
}

type EmailRules struct {
	MaxLength int
}

func NewValidationRules() *ValidationRules {
	return &ValidationRules{
		Username: UsernameRules{
			MinLength:    3,
			MaxLength:    50,
			AllowedChars: "^[a-zA-Z0-9_-]+$",
		},
		Password: PasswordRules{
			MinLength:           8,
			RequireUppercase:    true,
			RequireLowercase:    true,
			RequireNumbers:      true,
			RequireSpecialChars: true,
		},
		Email: EmailRules{
			MaxLength: 255,
		},
	}
}

func (r *UsernameRules) Validate(username string) []string {
	var errors []string

	username = strings.TrimSpace(username)

	if len(username) < r.MinLength {
		errors = append(errors, "Username must be at least 3 characters")
	}

	if len(username) > r.MaxLength {
		errors = append(errors, "Username must not exceed 50 characters")
	}

	matched, _ := regexp.MatchString(r.AllowedChars, username)
	if !matched {
		errors = append(errors, "Username can only contain letters, numbers, hyphens, and underscores")
	}

	return errors
}

func (r *PasswordRules) Validate(password string) []string {
	var errors []string

	if len(password) < r.MinLength {
		errors = append(errors, "Password must be at least 8 characters")
	}

	if r.RequireUppercase {
		hasUpper := false
		for _, ch := range password {
			if unicode.IsUpper(ch) {
				hasUpper = true
				break
			}
		}
		if !hasUpper {
			errors = append(errors, "Password must contain at least one uppercase letter")
		}
	}

	if r.RequireLowercase {
		hasLower := false
		for _, ch := range password {
			if unicode.IsLower(ch) {
				hasLower = true
				break
			}
		}
		if !hasLower {
			errors = append(errors, "Password must contain at least one lowercase letter")
		}
	}

	if r.RequireNumbers {
		hasNumber := false
		for _, ch := range password {
			if unicode.IsDigit(ch) {
				hasNumber = true
				break
			}
		}
		if !hasNumber {
			errors = append(errors, "Password must contain at least one number")
		}
	}

	if r.RequireSpecialChars {
		specialChars := "!@#$%^&*()_+-=[]{}|;:,.<>?"
		hasSpecial := false
		for _, ch := range password {
			if strings.ContainsRune(specialChars, ch) {
				hasSpecial = true
				break
			}
		}
		if !hasSpecial {
			errors = append(errors, "Password must contain at least one special character")
		}
	}

	return errors
}

func (r *EmailRules) Validate(email string) []string {
	var errors []string

	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	if len(email) > r.MaxLength {
		errors = append(errors, "Email must not exceed 255 characters")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		errors = append(errors, "Invalid email format")
	}

	return errors
}

func SanitizeInput(input string) string {
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	return input
}
