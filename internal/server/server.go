package server

import "net/http"

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr: port,
		Handler: handler,
		MaxHeaderBytes: 1 << 20,
	}

	return s.httpServer.ListenAndServe()
}