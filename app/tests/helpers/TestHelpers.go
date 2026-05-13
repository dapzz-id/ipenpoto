package unit

import (
	"os"
	"testing"
)

// SetupTestEnv initializes environment variables for testing
func SetupTestEnv(t *testing.T) {
	t.Helper()
	os.Setenv("JWT_SECRET", "test-secret-key-min-32-chars-long")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "test")
	os.Setenv("DB_NAME", "api_ipenpoto_test")
	os.Setenv("REDIS_HOST", "localhost")
	os.Setenv("REDIS_PORT", "6379")
}

// CleanupTestEnv removes test environment variables
func CleanupTestEnv(t *testing.T) {
	t.Helper()
	os.Unsetenv("JWT_SECRET")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_USERNAME")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("REDIS_HOST")
	os.Unsetenv("REDIS_PORT")
}
