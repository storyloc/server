package gql

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"
	config "github.com/storyloc/server/pkg/configuration"
	"github.com/storyloc/server/pkg/server"
	"github.com/storyloc/server/pkg/service"
)

//go:embed index.html
var index embed.FS

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operationName"`
	Variables map[string]interface{} `json:"variables"`
}

func NewServer(conf *config.Configuration, storyService service.StoryService) (server.Server, error) {
	schema, err := newSchema(conf, storyService).build()
	if err != nil {
		return nil, err
	}

	return &gqlServer{
		schema: schema,
		config: conf,
	}, nil
}

type gqlServer struct {
	schema graphql.Schema
	config *config.Configuration
}

func (s gqlServer) Routes(r *chi.Mux) {
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
	var p postData
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(400)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:         s.schema,
		RequestString:  p.Query,
		VariableValues: p.Variables,
		OperationName:  p.Operation,
	})

	if err := json.NewEncoder(w).Encode(result); err != nil {
		fmt.Printf("could not write result to response: %s", err)
	}
}
