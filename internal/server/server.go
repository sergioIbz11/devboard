package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

func New(addr string) *Server {
	mux := http.NewServeMux()
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: mux,

			ReadTimeout:       10 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       60 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
		},
		mux: mux,
	}
}

func (server *Server) Start() error {
	return server.httpServer.ListenAndServe()
}
