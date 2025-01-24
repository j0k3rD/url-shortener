package config

import (
	"os"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	MongoDB  MongoDBConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port        string
	FrontendURL string
}

type PostgresConfig struct {
	User     string
	Password string
	Database string
	Host     string
}

type MongoDBConfig struct {
	URI string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port:        os.Getenv("SERVER_PORT"),
			FrontendURL: os.Getenv("FRONTEND_URL"),
		},
		Postgres: PostgresConfig{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
			Host:     os.Getenv("POSTGRES_HOST"),
		},
		MongoDB: MongoDBConfig{
			URI: os.Getenv("MONGODB_URI"),
		},
		Redis: RedisConfig{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		},
	}

	return cfg, nil
}

// func os.Getenv(key, defaultValue string) string {
// 	if value := os.Getenv(key); value != "" {
// 		return value
// 	}
// 	return defaultValue
// }

// // Métodos auxiliares para obtener cadenas de conexión formateadas
// func (c PostgresConfig) GetConnectionString() string {
// 	return fmt.Sprintf(
// 		"user=%s dbname=%s password=%s host=%s sslmode=disable",
// 		c.User,
// 		c.Database,
// 		c.Password,
// 		c.Host,
// 	)
// }