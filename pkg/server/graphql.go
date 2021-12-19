package server

import (
	"embed"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	config "github.com/storyloc/server/pkg/configuration"
	"net/http"
)

//go:embed index.html
var index embed.FS

type gqlServer struct {
	schema graphql.ExecutableSchema
	config config.Configuration
}

func NewGraphqlServer(conf config.Configuration, schema graphql.ExecutableSchema) (Server, error) {
	srv := &gqlServer{
		config: conf,
		schema: schema,
	}

	return srv, nil
}

func (s *gqlServer) Routes(r *chi.Mux) {
	if s.config.Server.GraphiQl {
		r.Get("/", s.handleGraphiQL)
	}

	r.Post("/graphql", s.handleGraphQL)
}

func (s gqlServer) handleGraphiQL(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.FS(index))
	fs.ServeHTTP(w, r)
}

func (s gqlServer) handleGraphQL(w http.ResponseWriter, r *http.Request) {
	h := handler.NewDefaultServer(s.schema)
	h.ServeHTTP(w, r)
}
