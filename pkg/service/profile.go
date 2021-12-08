package service

import "github.com/storyloc/server/pkg/storage"

type ProfileService struct {
	profileRepository storage.ProfileRepository
}

func NewProfileService(profileRepository storage.ProfileRepository) ProfileService {
	return ProfileService{
		profileRepository: profileRepository,
	}
}

func (ps ProfileService) CreateProfile(ts storage.Profile) (*storage.Profile, error) {
	return ps.profileRepository.CreateProfile(ts)
}

func (ps ProfileService) GetProfile(id string) (*storage.Profile, error) {
	return ps.profileRepository.GetProfile(id)
}
