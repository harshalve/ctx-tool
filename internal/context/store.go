package context

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func (s *Store) Load(name string) (ProjectContext, error) {
	path := filepath.Join(s.StorageDir, name+".json")
	var ctx ProjectContext

	data, err := os.ReadFile(path)
	if err != nil {
		return ctx, err
	}

	err = json.Unmarshal(data, &ctx)
	return ctx, err
}
