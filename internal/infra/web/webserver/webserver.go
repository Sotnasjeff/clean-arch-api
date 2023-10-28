package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Route struct {
	Method string
	Path   string
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[Route]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[Route]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	route := Route{
		Method: method,
		Path:   path,
	}
	s.Handlers[route] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for route, handler := range s.Handlers {
		if route.Method == "GET" {
			s.Router.Get(route.Path, handler)
		}
		s.Router.Post(route.Path, handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
