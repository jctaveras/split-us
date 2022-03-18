package rest

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server interface {
	Start() error
}

type server struct {
	handlers *http.ServeMux
}

func NewServer() Server {
	return &server{handlers: Handlers()}
}

func (s *server) Start() error {
	rest := &http.Server{
		Addr: ":8080",
		Handler: s.handlers,
	}

	if error := rest.ListenAndServe(); error != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		rest.Shutdown(ctx)
		log.Fatalf("Server Crashed with Error: %v", error)

		defer cancel()

		return error
	}

	return nil
}
