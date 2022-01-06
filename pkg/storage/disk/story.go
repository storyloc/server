package disk

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/storyloc/server/pkg/storage"
)

func NewStoryRepository() StoryRepository {
	return StoryRepository{
		store: newFileStore("story"),
	}
}

type StoryRepository struct {
	store fileStore
}

func (sr StoryRepository) CreateStory(ts storage.Story) (*storage.Story, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	ts.Id = id

	if err := sr.store.Create(id, ts); err != nil {
		return nil, err
	}

	return &ts, nil
}

func (sr StoryRepository) GetStory(id string) (*storage.Story, error) {
	story := &storage.Story{}

	if err := sr.store.Read(id, story); err != nil {
		return nil, err
	}

	return story, nil
}

func (sr StoryRepository) AllStories() ([]*storage.Story, error) {
	ids, err := sr.store.Ids()
	if err != nil {
		return nil, err
	}

	var stories []*storage.Story
	for _, id := range ids {
		story, err := sr.GetStory(id)
		if err != nil {
			return nil, err
		}

		stories = append(stories, story)
	}

	return stories, nil
}
