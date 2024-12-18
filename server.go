package musiclib

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Run(port string, h http.Handler) error {
	s.server = &http.Server{
		Addr:           ":" + port,
		Handler:        h,
		MaxHeaderBytes: 1 << 20, //1Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
