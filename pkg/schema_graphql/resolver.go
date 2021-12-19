package graphql

import (
	"github.com/storyloc/server/pkg/schema_graphql/generated"
	"github.com/storyloc/server/pkg/service"
)

type Resolver struct {
	profileService service.ProfileService
	storyService   service.StoryService
}

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type profileResolver struct{ *Resolver }

func (r *Resolver) Profile() generated.ProfileResolver { return &profileResolver{r} }

type storyResolver struct{ *Resolver }

func (r *Resolver) Story() generated.StoryResolver { return &storyResolver{r} }
