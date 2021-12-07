package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	config "github.com/storyloc/server/pkg/configuration"
)

type Server interface {
	Routes(router *chi.Mux)
}

func Start(conf config.Configuration, srvs ...Server) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	for _, srv := range srvs {
		srv.Routes(r)
	}
	fmt.Printf("Server is runningon port: %s \r\n", conf.Server.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Server.Port), r)
}
