package store

import (
	"encoding/json"
	"io/ioutil"

	"github.com/adrianosela/guardduty/categorization"
)

// FileSystemStore is a file system
// implementation of the Store interface
type FileSystemStore struct {
	filepath string
}

// NewFileSystemStore is the FileSystemStore constructor
func NewFileSystemStore(filepath string) Store {
	return &FileSystemStore{filepath: filepath}
}

// GetCategorization returns a fresh categorization from the file system
func (s *FileSystemStore) GetCategorization() (*categorization.Categorization, error) {
	fileByt, err := ioutil.ReadFile(s.filepath)
	if err != nil {
		return nil, err
	}
	var categorization *categorization.Categorization
	if err = json.Unmarshal([]byte(fileByt), &categorization); err != nil {
		return nil, err
	}
	return categorization, nil
}
