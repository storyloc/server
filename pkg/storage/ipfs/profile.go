package ipfs

import (
	shell "github.com/ipfs/go-ipfs-api"
	config "github.com/storyloc/server/pkg/configuration"
	"github.com/storyloc/server/pkg/storage"
)

func NewProfileRepository(conf config.Configuration) ProfileRepository {
	return ProfileRepository{
		sh: shell.NewShell(conf.Storage.Ipfs.Url),
	}
}

type ProfileRepository struct {
	sh *shell.Shell
}

func (pr ProfileRepository) CreateProfile(ts storage.Profile) (*storage.Profile, error) {
	return nil, nil
}

func (pr ProfileRepository) GetProfile(id string) (*storage.Profile, error) {
	return nil, nil
}
