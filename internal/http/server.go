package http

import (
	"net/http"

	"github.com/ctuzelov/weather-api/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	cfg *config.Config
	srv *http.Server
}

func New(cfg *config.Config) *Server {
	srv := &http.Server{
		Addr: cfg.HTTPAddr,
	}

	return &Server{
		cfg: cfg,
		srv: srv,
	}
}

func (s *Server) Start() {
	s.srv.Handler = s.setupRoutes()

	go func() {
		s.cfg.L.Info("start listening http", zap.String("PORT", s.cfg.HTTPAddr))
		if err := http.ListenAndServe(s.cfg.HTTPAddr, s.srv.Handler); err != nil {
			s.cfg.L.Fatal("error run http & gateway server", zap.Error(err))
		}
	}()
}
