package http

import (
	"fmt"
	"net/http"
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
func NewServer() *Server {

	r := mux.NewRouter()
	addr := fmt.Sprintf("%s:%s")
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	return &Server{httpServer: srv}
}
