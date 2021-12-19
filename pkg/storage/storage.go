package storage

import "time"

type StoryRepository interface {
	CreateStory(story Story) (*Story, error)
	GetStory(id string) (*Story, error)
	AllStories() ([]*Story, error)
}

type ProfileRepository interface {
	CreateProfile(story Profile) (*Profile, error)
	GetProfile(id string) (*Profile, error)
}

type Story struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	OwnerID   string `json:"owner"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Profile struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	StoryIDs  []string `json:"stories"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
