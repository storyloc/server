package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/storyloc/server/pkg/graphql/gen"
	"github.com/storyloc/server/pkg/service"
)

func NewSchema(profileService service.ProfileService, storyService service.StoryService) graphql.ExecutableSchema {
	return gen.NewExecutableSchema(
		gen.Config{
			Resolvers: &Resolver{
				profileService: profileService,
				storyService:   storyService,
			},
		},
	)
}
