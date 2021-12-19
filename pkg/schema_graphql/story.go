package graphql

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/storyloc/server/pkg/schema_graphql/model"
	"github.com/storyloc/server/pkg/storage"
)

func (r *mutationResolver) CreateStory(ctx context.Context, input model.StoryInput) (*storage.Story, error) {
	var ts storage.Story
	if err := mapstructure.Decode(input, &ts); err != nil {
		return nil, err
	}

	return r.storyService.CreateStory(ts)
}

func (r *queryResolver) Story(ctx context.Context, id string) (*storage.Story, error) {
	return r.storyService.GetStory(id)
}

func (r *queryResolver) Stories(ctx context.Context) ([]*storage.Story, error) {
	return r.storyService.AllStories()
}

func (r *storyResolver) Owner(ctx context.Context, obj *storage.Story) (*storage.Profile, error) {
	return r.profileService.GetProfile(obj.OwnerID)
}
