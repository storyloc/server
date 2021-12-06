package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/storyloc/server/pkg/configuration"
	"net/http"
)

type Server interface {
	Routes(router *chi.Mux)
}

func Start(conf *config.Configuration, srvs []Server) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	for _, srv := range srvs {
		srv.Routes(r)
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Server.Port), r)
}
