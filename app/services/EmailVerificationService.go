package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type EmailVerificationService struct {
	Redis *redis.Client
	TTL   time.Duration
}

func NewEmailVerificationService(redis *redis.Client) *EmailVerificationService {
	return &EmailVerificationService{
		Redis: redis,
		TTL:   24 * time.Hour,
	}
}

func (s *EmailVerificationService) GenerateToken(userID uuid.UUID) (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	tokenString := hex.EncodeToString(token)
	key := fmt.Sprintf("verify:email:%s", tokenString)
	ctx := context.Background()

	err = s.Redis.Set(
		ctx,
		key,
		userID.String(),
		s.TTL,
	).Err()

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *EmailVerificationService) VerifyToken(token string) (uuid.UUID, error) {
	key := fmt.Sprintf("verify:email:%s", token)
	ctx := context.Background()

	userIDString, err := s.Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return uuid.Nil, fmt.Errorf("verification token expired or invalid")
	}
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return uuid.Nil, err
	}

	s.Redis.Del(ctx, key)

	return userID, nil
}
