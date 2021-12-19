package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/storyloc/server/pkg/schema_graphql/generated"
	"github.com/storyloc/server/pkg/service"
)

func NewSchema(profileService service.ProfileService, storyService service.StoryService) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &Resolver{
				profileService: profileService,
				storyService:   storyService,
			},
		},
	)
}
