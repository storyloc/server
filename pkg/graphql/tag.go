package graphql

import (
	"context"
	"github.com/storyloc/server/pkg/graphql/gen"
	"github.com/storyloc/server/pkg/storage"
)

func (r *mutationResolver) CreateTag(ctx context.Context, input gen.TagInput) (*storage.Tag, error) {
	return nil, nil
}
