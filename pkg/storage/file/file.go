package file

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

var storageHome string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	storageHome = filepath.Join(home, ".storylock")
	if _, err := os.Stat(storageHome); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(storageHome, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}
