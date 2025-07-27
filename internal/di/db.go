package di

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	"kumparan/config"
	"log"

	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Configuration) *sql.DB {
	dsn := "host=" + cfg.DbHost + " user=" + cfg.DbUser + " password=" + cfg.DbPassword +
		" dbname=" + cfg.DbName + " port=" + cfg.DbPort + " sslmode=" + cfg.DbSSL

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func NewRedisClient(cfg *config.Configuration) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisServer,
	})
	result, err := client.Ping(context.Background()).Result()
	fmt.Printf("Redis ping result: %v\n", result)
	if err != nil {
		panic(err)
	}
	return client
}
