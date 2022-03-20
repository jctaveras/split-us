package rest

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jctaveras/split-us/pkg/http/rest/routes/auth"
	"github.com/jctaveras/split-us/pkg/http/rest/routes/user"
	"github.com/jctaveras/split-us/pkg/http/router"
)

type server struct {
	handlers *http.ServeMux
}

func NewServer() *server {
	return &server{handlers: router.Routes.Handlers()}
}

func (s *server) Start(ctx context.Context) error {
	auth.InitAuthHandlers(ctx)
	user.InitUserHandlers(ctx)
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
