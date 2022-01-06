package graphql

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/storyloc/server/pkg/graphql/gen"
	"github.com/storyloc/server/pkg/storage"
)

func (r *mutationResolver) CreateStory(ctx context.Context, input gen.StoryInput) (*storage.Story, error) {
	var ts storage.Story
	if err := mapstructure.Decode(input, &ts); err != nil {
		return nil, err
	}

	return r.storyService.CreateStory(ts)
}

func (r *mutationResolver) DeleteStory(ctx context.Context, id string) (*storage.Story, error) {
	return nil, nil
}

func (r *queryResolver) Story(ctx context.Context, id string) (*storage.Story, error) {
	return r.storyService.GetStory(id)
}

func (r *queryResolver) Stories(ctx context.Context) ([]*storage.Story, error) {
	return r.storyService.AllStories()
}

type storyResolver struct{ *Resolver }

func (r *Resolver) Story() gen.StoryResolver { return &storyResolver{r} }

func (r *storyResolver) Owner(ctx context.Context, story *storage.Story) (*storage.Profile, error) {
	return r.profileService.GetProfile(story.OwnerID)
}

func (r *storyResolver) Collection(ctx context.Context, story *storage.Story) ([]*storage.StoryPoint, error) {
	return nil, nil
}
