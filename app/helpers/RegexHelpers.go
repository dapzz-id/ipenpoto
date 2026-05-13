package helpers

import (
	"regexp"
	"strings"
)

// RegexHelpers untuk operasi regex
type RegexHelpers struct{}

// IsValidEmail mengecek apakah email valid
func (r RegexHelpers) IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

// IsValidPhone mengecek apakah nomor telepon valid (simple check)
func (r RegexHelpers) IsValidPhone(phone string) bool {
	pattern := `^\+?[0-9]{10,15}$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}

// IsValidUsername mengecek apakah username valid
func (r RegexHelpers) IsValidUsername(username string) bool {
	pattern := `^[a-zA-Z0-9_-]{3,50}$`
	match, _ := regexp.MatchString(pattern, username)
	return match
}

// ContainsSpecialChar mengecek apakah string mengandung special character
func (r RegexHelpers) ContainsSpecialChar(str string) bool {
	pattern := `[!@#$%^&*()_+\-=\[\]{}|;:,.<>?]`
	match, _ := regexp.MatchString(pattern, str)
	return match
}

// RemoveSpaces menghapus semua spasi
func (r RegexHelpers) RemoveSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

// NewRegexHelpers membuat instance RegexHelpers
func NewRegexHelpers() RegexHelpers {
	return RegexHelpers{}
}
