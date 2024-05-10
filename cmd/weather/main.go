package main

import (
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/ctuzelov/weather-api/internal/app"
	"github.com/ctuzelov/weather-api/internal/config"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewProduction()

	cfg := &config.Config{
		L:              l,
		AppMode:        config.LocalAppMode,
		GRPCAddr:       ":9000",
		HTTPAddr:       ":8002",
		PgPort:         "5432",
		PgHost:         "localhost",
		PgName:         "weather",
		PgUser:         "root",
		PgPassword:     "123",
		PgPingEnabled:  true,
		PgPingInterval: time.Second * 60,
	}

	if err := env.Parse(cfg); err != nil {
		cfg.L.Fatal("cannot parse config", zap.Error(err))
	}

	if err := app.New(cfg); err != nil {
		l.Fatal("error due run the server", zap.Error(err))
	}
}
