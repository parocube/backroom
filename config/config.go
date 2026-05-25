package config

import (
	"time"

	"github.com/parocube/backroom/xenv"
)

type Config struct {
	Server   Server
	Postgres Postgres
}

type Server struct {
	ServerPort string
}

type Postgres struct {
	PostgresURL             string
	PostgresMaxOpenConn     int
	PostgresMaxIdleConn     int
	PostgresConnMaxLifetime time.Duration
}

func Load() Config {
	return Config{
		Server: Server{
			ServerPort: xenv.GetEnv("SERVER_PORT", ":8080"),
		},
		Postgres: Postgres{
			PostgresURL:             xenv.GetEnv("POSTGRES_URL", ""),
			PostgresMaxOpenConn:     xenv.GetEnvAsInt("POSTGRES_MAX_OPEN_CONN", 20),
			PostgresMaxIdleConn:     xenv.GetEnvAsInt("POSTGRES_MAX_IDLE_CONN", 10),
			PostgresConnMaxLifetime: xenv.GetEnvAsDuration("POSTGRES_CONN_MAX_LIFETIME", 30*time.Minute),
		},
	}
}
