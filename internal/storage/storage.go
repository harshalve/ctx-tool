package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/harshalve/ctx-tool/internal/context"
)

type Storage struct {
	StorageDir string
}

func NewStorage() (*Storage, error) {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".ctx")

	// Ensure the directory exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	return &Storage{StorageDir: dir}, nil
}

func (s *Storage) ListAll() ([]string, error) {
	files, err := os.ReadDir(s.StorageDir)
	if err != nil {
		return nil, err
	}

	var projects []string
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".json" {
			projects = append(projects, strings.TrimSuffix(f.Name(), ".json"))
		}
	}
	return projects, nil
}

func (s *Storage) Load(name string) (context.ProjectContext, error) {
	var ctx context.ProjectContext
	path := filepath.Join(s.StorageDir, name+".json")

	data, err := os.ReadFile(path)
	if err != nil {
		return ctx, err
	}

	err = json.Unmarshal(data, &ctx)
	return ctx, err
}

// Delete removes a context file
func (s *Storage) Delete(name string) error {
	path := filepath.Join(s.StorageDir, name+".json")
	return os.Remove(path)
}
