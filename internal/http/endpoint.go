package http

import (
	"context"
	"fmt"
	"net/http"

	wapi "github.com/ctuzelov/weather-api/pkg/weatherapi"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Server) setupRoutes() *http.ServeMux {
	grpcMux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := wapi.RegisterWeatherServiceHandlerFromEndpoint(context.Background(), grpcMux, fmt.Sprintf("localhost%s", s.cfg.GRPCAddr), opts); err != nil {
		s.cfg.L.Error("cannot register gateway handler from endpoint", zap.Error(err))
		return nil
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	mux.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		s.cfg.L.Info("alive endpoint")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is working"))
	})

	return mux
}
