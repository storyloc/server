package graphql

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/storyloc/server/pkg/graphql/gen"
	"github.com/storyloc/server/pkg/storage"
)

func (r *mutationResolver) CreateProfile(ctx context.Context, input gen.ProfileInput) (*storage.Profile, error) {
	var tp storage.Profile
	if err := mapstructure.Decode(input, &tp); err != nil {
		return nil, err
	}

	return r.profileService.CreateProfile(tp)
}

func (r *queryResolver) Profile(ctx context.Context, id string) (*storage.Profile, error) {
	return r.profileService.GetProfile(id)
}

type profileResolver struct{ *Resolver }

func (r *Resolver) Profile() gen.ProfileResolver { return &profileResolver{r} }

func (r *profileResolver) Stories(ctx context.Context, profile *storage.Profile) ([]*storage.Story, error) {
	var profileStories []*storage.Story

	stories, err := r.storyService.AllStories()
	if err != nil {
		return nil, err
	}

	for _, story := range stories {
		if story.OwnerID != profile.Id {
			continue
		}

		profileStories = append(profileStories, story)
	}

	return profileStories, nil
}
