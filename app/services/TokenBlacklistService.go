package services

import (
	"time"

	"ipenpoto/config"

	"github.com/golang-jwt/jwt/v5"
)

type TokenBlacklistService struct{}

func NewTokenBlacklistService() *TokenBlacklistService {
	return &TokenBlacklistService{}
}

func (s *TokenBlacklistService) BlacklistToken(token string, expiration time.Time) error {
	ttl := time.Until(expiration)
	if ttl <= 0 {
		return nil
	}

	key := "blacklist:jwt:" + token
	return config.RDB.Set(config.Ctx, key, "1", ttl).Err()
}

func (s *TokenBlacklistService) IsBlacklisted(token string) bool {
	val, err := config.RDB.Get(config.Ctx, "blacklist:jwt:"+token).Result()
	if err != nil {
		return false
	}
	return val == "1"
}

func (s *TokenBlacklistService) BlacklistTokenFromClaims(claims jwt.MapClaims) error {
	exp, ok := claims["exp"]
	if !ok {
		return nil
	}

	expTime := time.Unix(int64(exp.(float64)), 0)
	return s.BlacklistToken("", expTime)
}
