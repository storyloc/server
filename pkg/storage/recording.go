package storage

type Recording struct {
	ID           string `json:"id"`
	StoryPointId string `json:"storyPoint"`
	Title        string `json:"title"`
	Data         string `json:"data"`
}
