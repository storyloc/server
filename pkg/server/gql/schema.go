package gql

import (
	"github.com/graphql-go/graphql"
	config "github.com/storyloc/server/pkg/configuration"
	"github.com/storyloc/server/pkg/service"
)

type schema struct {
	config       *config.Configuration
	storyService service.StoryService
}

func newSchema(conf *config.Configuration, storyService service.StoryService) schema {
	return schema{
		config:       conf,
		storyService: storyService,
	}
}

func (s schema) build() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"profile": s.queryStory(),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createProfile": s.mutationCreateStory(),
			},
		}),
	})
}
