package disk

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/storyloc/server/pkg/storage"
)

func NewProfileRepository() ProfileRepository {
	return ProfileRepository{
		store: newFileStore("profile"),
	}
}

type ProfileRepository struct {
	store fileStore
}

func (sr ProfileRepository) CreateProfile(ts storage.Profile) (*storage.Profile, error) {
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

func (sr ProfileRepository) GetProfile(id string) (*storage.Profile, error) {
	profile := &storage.Profile{}

	if err := sr.store.Read(id, profile); err != nil {
		return nil, err
	}

	return profile, nil
}
