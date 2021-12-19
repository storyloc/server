package server

import (
	"fmt"
	"log"
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

	log.Printf("open http://localhost:%s/", conf.Server.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Server.Port), r)
}
