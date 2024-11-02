package transport

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"timeslot-service/internal/app/handler"
)

type Server struct {
	opts      Options
	endpoints *handler.Handler
	router    *chi.Mux
}

func NewServer(opts Options, endpoints *handler.Handler) *Server {
	return &Server{
		opts:      opts,
		endpoints: endpoints,
		router:    chi.NewRouter(),
	}
}

func (s *Server) Run() error {
	s.configureHTTPHandlers()

	err := s.runHTTPServer()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) configureHTTPHandlers() {
	routes := s.endpoints.ConfigureHTTPEndpoints()
	for method, route := range routes {
		for pattern, handlerFn := range route {
			s.router.MethodFunc(method, pattern, handlerFn)
		}
	}
}

func (s *Server) runHTTPServer() error {
	err := http.ListenAndServe(s.opts.HttpPort, s.router)
	if err != nil {
		log.Fatalf("failed to run HTTP server: %s", err.Error())
		return err
	}

	return nil
}
