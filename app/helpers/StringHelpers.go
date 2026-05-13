package helpers

import "strings"

// StringHelpers untuk operasi string
type StringHelpers struct{}

// TrimAndLower membersihkan dan konversi string ke lowercase
func (s StringHelpers) TrimAndLower(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

// IsEmpty mengecek apakah string kosong
func (s StringHelpers) IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

// HasPrefix mengecek prefix dengan case-insensitive
func (s StringHelpers) HasPrefixNoCase(str, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix))
}

// TruncateString memotong string ke panjang maksimal
func (s StringHelpers) TruncateString(str string, maxLen int) string {
	if len(str) <= maxLen {
		return str
	}
	return str[:maxLen] + "..."
}

// NewStringHelpers membuat instance StringHelpers
func NewStringHelpers() StringHelpers {
	return StringHelpers{}
}
