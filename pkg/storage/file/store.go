package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type fileStore struct {
	namespace string
}

func newFileStore(ns string) fileStore {
	return fileStore{namespace: ns}
}

func (fs fileStore) Exist(id string) bool {
	if info, err := os.Stat(fs.Path(id)); errors.Is(err, os.ErrNotExist) || info.IsDir() {
		return false
	}

	return true
}

func (fs fileStore) Create(id string, body interface{}) error {
	jsonBody, err := json.MarshalIndent(body, "", " ")
	if err != nil {
		return err
	}

	to := fs.Path(id)

	if err := os.MkdirAll(filepath.Dir(to), 0777); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	if err = ioutil.WriteFile(to, jsonBody, 0644); err != nil {
		return err
	}

	return nil
}

func (fs fileStore) Read(id string, to interface{}) error {
	from := fs.Path(id)
	if exist := fs.Exist(id); !exist {
		return fmt.Errorf(`storageFile %s does not exist`, from)
	}

	b, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, to)
}

func (fs fileStore) Path(id string) string {
	return filepath.Join(storageHome, fmt.Sprintf("%s_%s.json", fs.namespace, id))
}
