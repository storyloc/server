package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/storyloc/server/pkg/storage"
)

func (s schema) mutationCreateStory() *graphql.Field {
	return &graphql.Field{
		Type:        typeStory,
		Description: "create new story",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			name, _ := params.Args["name"].(string)
			transportProfile := storage.Story{Name: name}

			return s.storyService.CreateStory(transportProfile)
		},
	}

}
