package storageIpfs

import (
	"bytes"
	"encoding/json"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	config "github.com/storyloc/server/pkg/configuration"
	schema "github.com/storyloc/server/pkg/schema_ipld"
	"github.com/storyloc/server/pkg/storage"
)

func NewStoryRepository(conf config.Configuration) StoryRepository {
	return StoryRepository{
		sh: shell.NewShell(conf.Storage.Ipfs.Url),
	}
}

type StoryRepository struct {
	sh *shell.Shell
}

func (sr StoryRepository) CreateStory(ts storage.Story) (*storage.Story, error) {
	builder := schema.Type.Story.NewBuilder()
	vb, err := json.Marshal(ts)
	if err != nil {
		return nil, err
	}

	if err := dagjson.Decode(builder, bytes.NewReader(vb)); err != nil {
		return nil, err
	}

	_ = builder.Build()

	return &storage.Story{Name: "toDo: transport schema Story to Story"}, nil
}

func (sr StoryRepository) GetStory(id string) (*storage.Story, error) {
	return nil, nil
}

func (sr StoryRepository) AllStories() ([]*storage.Story, error) {
	return nil, nil
}
