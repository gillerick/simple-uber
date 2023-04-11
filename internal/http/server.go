package http

import (
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"simple-uber/internal/configs"
	"time"

	"github.com/gorilla/mux"
)

// Server implements an HTTP server and a router for service endpoints
type Server struct {
	httpServer *http.Server
}

// Run starts the HTTP server
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// ToDo: Inject config properties and handlers
func NewServer(config configs.App) *Server {
	log.Infof("server listening on address: %s, port: %s", config.Host, config.Port)

	r := mux.NewRouter()
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	return &Server{httpServer: srv}
}
