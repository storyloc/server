package graphql

import (
	"github.com/storyloc/server/pkg/graphql/gen"
	"github.com/storyloc/server/pkg/service"
)

type Resolver struct {
	profileService service.ProfileService
	storyService   service.StoryService
}

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }
