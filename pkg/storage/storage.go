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
	Id        string
	Name      string
	OwnerId   string
	Owner     Profile
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Profile struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
