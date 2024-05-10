package service

import (
	"github.com/ctuzelov/weather-api/internal/config"
	"github.com/ctuzelov/weather-api/internal/repository"
	wapi "github.com/ctuzelov/weather-api/pkg/weatherapi"
)

type Weather interface {
	GetWeather(cfg *config.Config, req *wapi.WeatherRequest) (*wapi.WeatherResponse, error)
}

type Service struct {
	Weather
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
