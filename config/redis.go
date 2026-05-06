package config

import (
    "context"
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
    "github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}

func getEnvInt(key string, fallback int) int {
    v := os.Getenv(key)
    if v == "" {
        return fallback
    }
    n, err := strconv.Atoi(v)
    if err != nil {
        return fallback
    }
    return n
}

func ConnectRedis() {
    _ = godotenv.Load()

    host := getEnv("REDIS_HOST", "localhost")
    port := getEnv("REDIS_PORT", "6379")
    password := getEnv("REDIS_PASSWORD", "")
    db := getEnvInt("REDIS_DB", 0)

    addr := host + ":" + port

    RDB = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })

    if err := RDB.Ping(Ctx).Err(); err != nil {
        log.Fatalf("Redis connection failed: %v", err)
    }

    log.Printf("Redis connected on %s (db=%d)", addr, db)
}