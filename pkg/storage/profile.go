package storage

type ProfileRepository interface {
	CreateProfile(story Profile) (*Profile, error)
	GetProfile(id string) (*Profile, error)
}

type Profile struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	StoryIDs []string `json:"stories"`
}
