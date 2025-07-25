package config

type (
	DbConfig struct {
		DbHost     string `envconfig:"DB_HOST" default:"localhost"`
		DbUser     string `envconfig:"DB_USER" default:"postgres"`
		DbName     string `envconfig:"DB_NAME" default:"postgres"`
		DbPassword string `envconfig:"DB_PASSWORD" default:"postgres"`
		DbSSL      string `envconfig:"DB_SSL" default:"disable"`
		DbPort     string `envconfig:"DB_PORT" default:"5432"`
	}
)
