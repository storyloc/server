package storage

type StoryRepository interface {
	CreateStory(story Story) (*Story, error)
	GetStory(id string) (*Story, error)
	AllStories() ([]*Story, error)
}

type Story struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	OwnerID       string   `json:"owner"`
	CollectionIDs []string `json:"collection"`
}
