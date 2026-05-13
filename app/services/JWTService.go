package services

import (
	"ipenpoto/app/enums"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (s *JWTService) GenerateAccessToken(
	userID uuid.UUID,
	username string,
	role enums.UserRole,
) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"username": username,
		"role": role,
		"exp": time.Now().
			Add(time.Minute * 15).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}

func (s *JWTService) GenerateRefreshToken(
	userID uuid.UUID,
) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().
			Add(time.Hour * 24 * 7).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}