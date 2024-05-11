package app

import (
	"fmt"
	"net"

	"github.com/ctuzelov/weather-api/internal/app/weatherservice/endpoint"
	"github.com/ctuzelov/weather-api/internal/app/weatherservice/service"
	"github.com/ctuzelov/weather-api/internal/config"
	"github.com/ctuzelov/weather-api/internal/db"
	"github.com/ctuzelov/weather-api/internal/http"
	"github.com/ctuzelov/weather-api/internal/repository"
	wapi "github.com/ctuzelov/weather-api/pkg/weatherapi"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(cfg *config.Config) error {
	s := grpc.NewServer()
	reflection.Register(s)

	dbClient, err := db.NewClient(cfg)
	if err != nil {
		cfg.L.Fatal("error with db", zap.Error(err))
	}

	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		return fmt.Errorf("cannot create listener, %w", err)
	}

	httpServer := http.New(cfg)
	httpServer.Start()

	repo := repository.NewRepository(dbClient.GetDB())
	srv := service.NewService(repo)
	endpoints := endpoint.NewEndpoint(srv, cfg)

	wapi.RegisterWeatherServiceServer(s, endpoints)

	cfg.L.Info("starting listening grpc server", zap.Any("PORT", cfg.GRPCAddr))
	if err := s.Serve(lis); err != nil {
		cfg.L.Error("error serve grpc server", zap.Error(err))
		return fmt.Errorf("error serve grpc server, %w", err)
	}

	return nil
}
