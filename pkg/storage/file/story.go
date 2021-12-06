package file

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/storyloc/server/pkg/storage"
	"time"
)

func NewStoryRepository() StoryRepository {
	return StoryRepository{
		fs: newFileStore("story"),
	}
}

type StoryRepository struct {
	fs fileStore
}

func (sr StoryRepository) CreateStory(ts storage.Story) (*storage.Story, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	ts.Id = id
	ts.CreatedAt = time.Now()
	ts.UpdatedAt = time.Now()

	if err := sr.fs.Create(id, ts); err != nil {
		return nil, err
	}

	return &ts, nil
}

func (sr StoryRepository) GetStory(id string) (*storage.Story, error) {
	story := &storage.Story{}

	if err := sr.fs.Read(id, story); err != nil {
		return nil, err
	}

	return story, nil
}
