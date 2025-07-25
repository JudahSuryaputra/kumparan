package shared

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"go.uber.org/dig"
	"kumparan/config"
)

type (
	Deps struct {
		dig.In

		DB          *sql.DB
		Config      *config.Configuration
		RedisClient *redis.Client
	}
)
