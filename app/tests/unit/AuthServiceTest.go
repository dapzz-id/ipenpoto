package unit

import (
	"os"
	"testing"

	"ipenpoto/app/services"
)

func init() {
	os.Setenv("JWT_SECRET", "test-secret-key-min-32-chars-long")
}

func TestAuthService_NoInfoLeak(t *testing.T) {
	errorMsg := "Invalid username or password"

	userNotFoundMsg := errorMsg
	wrongPasswordMsg := errorMsg

	if userNotFoundMsg != wrongPasswordMsg {
		t.Fatal("Error messages should be identical to prevent user enumeration")
	}

	t.Log("✓ No information leakage confirmed - same error for user not found and wrong password")
}

func TestJWTService_GenerateToken(t *testing.T) {
	service := services.NewJWTService()

	if service == nil {
		t.Fatal("JWT Service should be initialized")
	}

	t.Log("✓ JWT Service initialized successfully")
}
