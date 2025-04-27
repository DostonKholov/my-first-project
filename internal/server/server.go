package server

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) ServerRun(handlers http.Handler, port string) error {
	s.server = &http.Server{
		Addr:         "localhost:" + port,
		Handler:      handlers,
		ReadTimeout:  10 * time.Second, // 10secont patom zakroetsa
		WriteTimeout: 10 * time.Second,
	}
	log.Println("server is running on localhost:" + port)
	return s.server.ListenAndServe()
}
