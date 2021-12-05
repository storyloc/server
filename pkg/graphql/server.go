package graphql

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"
	"net/http"
)

//go:embed index.html
var index embed.FS

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func Handle(router *chi.Mux) {
	router.Get("/", handleGraphiQL)
	router.Post("/graphql", handleGraphQL)
}

func handleGraphiQL(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.FS(index))
	fs.ServeHTTP(w, r)
}

func handleGraphQL(w http.ResponseWriter, r *http.Request) {
	var p postData
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(400)
		return
	}

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"World": &graphql.Field{
					Type: World__type,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						fmt.Println("implement World")
						return nil, nil
					},
				},
				"Communities": &graphql.Field{
					Type: Communities__type,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						fmt.Println("implement Communities")
						return nil, nil
					},
				},
				"Community": &graphql.Field{
					Type: Community__type,
					Args: graphql.FieldConfigArgument{
						"UUID": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), Description: "community uuid"},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						fmt.Println("implement Community")
						return nil, nil
					},
				},
			},
		}),
	})

	result := graphql.Do(graphql.Params{
		Context:        r.Context(),
		Schema:         schema,
		RequestString:  p.Query,
		VariableValues: p.Variables,
		OperationName:  p.Operation,
	})

	if err := json.NewEncoder(w).Encode(result); err != nil {
		fmt.Printf("could not write result to response: %s", err)
	}
}
