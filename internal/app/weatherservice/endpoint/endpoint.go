package endpoint

import (
	"github.com/ctuzelov/weather-api/internal/app/weatherservice/service"
	"github.com/ctuzelov/weather-api/internal/config"
	wapi "github.com/ctuzelov/weather-api/pkg/weatherapi"
)

type Endpoint struct {
	wapi.UnimplementedWeatherServiceServer
	cfg *config.Config
	srv *service.Service
}

func NewEndpoint(srv *service.Service, cfg *config.Config) *Endpoint {
	return &Endpoint{
		cfg: cfg,
		srv: srv,
	}
}
