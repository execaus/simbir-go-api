package server

import (
	"context"
	"github.com/execaus/exloggo"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

const (
	serverMaxHeaderBytes = 1 << 20
	serverReadTimeout    = 10 * time.Second
	serverWriteTimeout   = 10 * time.Second
)

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: serverMaxHeaderBytes,
		ReadTimeout:    serverReadTimeout,
		WriteTimeout:   serverWriteTimeout,
	}
	exloggo.Info("server started successfully")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	exloggo.Info("server shutdown process started")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		exloggo.Error(err.Error())
	} else {
		exloggo.Info("http listener shutdown successfully")
	}

	exloggo.Info("server shutdown process completed successfully")
}
