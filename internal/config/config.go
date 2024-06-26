package config

import (
	"time"

	"go.uber.org/zap"
)

type AppMode string

const (
	LocalAppMode AppMode = "local"
	ProdAppMode  AppMode = "prod"
)

type Config struct {
	L              *zap.Logger
	AppMode        AppMode       `env:"APP_MODE" envDefault:"local"`
	GRPCAddr       string        `env:"GRPC_ADDR" envDefault:":9003"`
	HTTPAddr       string        `env:"HTTP_ADDR" envDefault:":8003"`
	PgPort         string        `env:"PG_PORT" envDefault:"5432"`
	PgHost         string        `env:"PG_HOST" envDefault:"postgres-db"`
	PgName         string        `env:"PG_NAME" envDefault:"postgres"`
	PgUser         string        `env:"PG_USER" envDefault:"postgres"`
	PgPassword     string        `env:"PG_PASSWORD" envDefault:"123"`
	PgPingEnabled  bool          `env:"PG_PING_ENABLED" envDefault:"true"`
	PgPingInterval time.Duration `env:"PG_PING_INTERVAL" envDefault:"40m"`
	PgMaxOpenConn  int           `env:"PG_MAX_OPEN_CONN" envDefault:"10"`
	PgIdleConn     int           `env:"PG_MAX_IDLE_CONN" envDefault:"10"`
	PgSSLMode      string        `env:"PG_SSL_MODE" envDefault:"disable"`

	// migration path
	// MigrationPath string `env:"MIGRATION_PATH" envDefault:"file://scheme"`
	// Weather API Token
	WeatherToken string `env:"WEATHER_API_TOKEN" envDefault:"8277586fe0654412ac4182943240905"`
}
