package graphql

import (
	"context"
	"github.com/storyloc/server/pkg/graphql/gen"
	"github.com/storyloc/server/pkg/storage"
)

func (r *mutationResolver) CreateStoryPoint(ctx context.Context, input gen.StoryPointInput) (*storage.StoryPoint, error) {
	return nil, nil
}

func (r *mutationResolver) DeleteStoryPoint(ctx context.Context, id string) (*storage.StoryPoint, error) {
	return nil, nil
}

type storyPointResolver struct{ *Resolver }

func (r *Resolver) StoryPoint() gen.StoryPointResolver { return &storyPointResolver{r} }

func (r *storyPointResolver) Owner(ctx context.Context, storyPoint *storage.StoryPoint) (*storage.Profile, error) {
	return nil, nil
}

func (r *storyPointResolver) Recordings(ctx context.Context, storyPoint *storage.StoryPoint) ([]*storage.Recording, error) {
	return nil, nil
}

func (r *storyPointResolver) Location(ctx context.Context, storyPoint *storage.StoryPoint) (*storage.Location, error) {
	return nil, nil
}
