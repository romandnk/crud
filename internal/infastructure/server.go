package infastructure

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Start(port string, connString string, sourceURL string, handler http.Handler) error {
	s.srv = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	migration(connString, sourceURL)
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func migration(connString string, sourceURL string) {
	m, err := migrate.New(
		"file:/"+sourceURL,
		connString,
	)
	if err != nil {
		logrus.Fatalf("error reading migration: %s", err.Error())
	}
	if err := m.Up(); err != nil {
		logrus.Fatalf("error up migration: %s", err.Error())
	}
}
