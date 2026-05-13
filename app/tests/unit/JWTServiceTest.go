package unit

import (
	"os"
	"testing"
	"time"

	"ipenpoto/app/enums"
	"ipenpoto/app/services"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestJWTService_GenerateAccessToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key-min-32-chars-long")
	defer os.Unsetenv("JWT_SECRET")

	service := services.NewJWTService()
	userID := uuid.New()
	username := "testuser"
	role := enums.Customer

	token, err := service.GenerateAccessToken(userID, username, role)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Fatal("Expected non-empty token")
	}

	parsedToken, _ := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("test-secret-key-min-32-chars-long"), nil
	})

	claims := parsedToken.Claims.(*jwt.MapClaims)
	if (*claims)["username"] != username {
		t.Fatalf("Expected username %s, got %v", username, (*claims)["username"])
	}

	if (*claims)["role"] != string(role) {
		t.Fatalf("Expected role %v, got %v", role, (*claims)["role"])
	}
}

func TestJWTService_GenerateAccessTokenExpiration(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key-min-32-chars-long")
	defer os.Unsetenv("JWT_SECRET")

	service := services.NewJWTService()
	userID := uuid.New()

	token, _ := service.GenerateAccessToken(userID, "test", enums.Customer)

	parsedToken, _ := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("test-secret-key-min-32-chars-long"), nil
	})

	claims := parsedToken.Claims.(*jwt.MapClaims)
	exp := time.Unix(int64((*claims)["exp"].(float64)), 0)

	if exp.Before(time.Now()) {
		t.Fatal("Token should not be expired")
	}

	if exp.After(time.Now().Add(16 * time.Minute)) {
		t.Fatal("Token expiration should be around 15 minutes")
	}
}

func TestJWTService_GenerateRefreshToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key-min-32-chars-long")
	defer os.Unsetenv("JWT_SECRET")

	service := services.NewJWTService()
	userID := uuid.New()

	token, err := service.GenerateRefreshToken(userID)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Fatal("Expected non-empty token")
	}

	parsedToken, _ := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("test-secret-key-min-32-chars-long"), nil
	})

	claims := parsedToken.Claims.(*jwt.MapClaims)
	if (*claims)["user_id"] == nil {
		t.Fatal("Expected user_id in refresh token")
	}
}
