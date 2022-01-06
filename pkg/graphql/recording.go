package graphql

import (
	"context"
	"github.com/storyloc/server/pkg/graphql/gen"
	"github.com/storyloc/server/pkg/storage"
)

func (r *mutationResolver) CreateRecording(ctx context.Context, input gen.RecordingInput) (*storage.Recording, error) {
	return nil, nil
}

func (r *mutationResolver) DeleteRecording(ctx context.Context, id string) (*storage.Recording, error) {
	return nil, nil
}

type recordingResolver struct{ *Resolver }

func (r *Resolver) Recording() gen.RecordingResolver { return &recordingResolver{r} }

func (r *recordingResolver) StoryPoint(ctx context.Context, recording *storage.Recording) (*storage.StoryPoint, error) {
	return nil, nil
}
