package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	httpsServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpsServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	logrus.Info("starting server")
	return s.httpsServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpsServer.Shutdown(ctx)
}
