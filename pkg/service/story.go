package service

import "github.com/storyloc/server/pkg/storage"

type StoryService struct {
	storyRepository storage.StoryRepository
}

func NewStoryService(storyRepository storage.StoryRepository) StoryService {
	return StoryService{
		storyRepository: storyRepository,
	}
}

func (ss StoryService) CreateStory(ts storage.Story) (*storage.Story, error) {
	return ss.storyRepository.CreateStory(ts)
}

func (ss StoryService) GetStory(id string) (*storage.Story, error) {
	return ss.storyRepository.GetStory(id)
}

func (ss StoryService) AllStories() ([]*storage.Story, error) {
	return ss.storyRepository.AllStories()
}
