package storage

type StoryPoint struct {
	ID           string   `json:"id"`
	OwnerID      string   `json:"owner"`
	Image        string   `json:"image"`
	RecordingIds []string `json:"recordings"`
	LocationId   string   `json:"location"`
}
