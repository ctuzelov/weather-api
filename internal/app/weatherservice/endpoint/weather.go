package endpoint

import (
	"context"

	wapi "github.com/ctuzelov/weather-api/pkg/weatherapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (e *Endpoint) GetWeather(ctx context.Context, req *wapi.WeatherRequest) (*wapi.WeatherResponse, error) {
	rsp, err := e.srv.GetWeather(e.cfg, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get weather %v", err)
	}

	return rsp, nil
}
