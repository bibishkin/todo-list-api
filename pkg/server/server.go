package server

import (
	"context"
	"net/http"
	"todo-list-api/pkg/logger"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Run() {
	err := s.srv.ListenAndServe()
	if err != nil {
		logger.ErrorLog.Fatalln(err)
	}
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func NewServer(addr string, router http.Handler) *Server {
	s := &Server{
		srv: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}

	return s
}
