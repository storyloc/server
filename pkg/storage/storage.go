package storage

import "time"

type StoryRepository interface {
	CreateStory(Story) (*Story, error)
	GetStory(string) (*Story, error)
}

type Story struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
