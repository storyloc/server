package disk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type fileStore struct {
	namespace string
}

func newFileStore(ns string) fileStore {
	return fileStore{namespace: ns}
}

func (fs fileStore) Exist(id string) bool {
	if info, err := os.Stat(fs.IdPath(id)); errors.Is(err, os.ErrNotExist) || info.IsDir() {
		return false
	}

	return true
}

func (fs fileStore) Create(id string, body interface{}) error {
	jsonBody, err := json.MarshalIndent(body, "", " ")
	if err != nil {
		return err
	}

	to := fs.IdPath(id)

	if err := os.MkdirAll(filepath.Dir(to), 0777); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	if err = ioutil.WriteFile(to, jsonBody, 0644); err != nil {
		return err
	}

	return nil
}

func (fs fileStore) Read(id string, to interface{}) error {
	from := fs.IdPath(id)
	if exist := fs.Exist(id); !exist {
		return fmt.Errorf(`storageFile %s does not exist`, from)
	}

	b, err := ioutil.ReadFile(from)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, to)
}

func (fs fileStore) Ids() ([]string, error) {
	files, err := ioutil.ReadDir(fs.Path())
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() || path.Ext(fileName) != ".json" {
			continue
		}

		ids = append(ids, strings.TrimSuffix(file.Name(), path.Ext(fileName)))
	}

	return ids, nil
}

func (fs fileStore) IdPath(id string) string {
	return filepath.Join(fs.Path(), fmt.Sprintf("%s.json", id))
}

func (fs fileStore) Path() string {
	return filepath.Join(storageHome, fs.namespace)
}
