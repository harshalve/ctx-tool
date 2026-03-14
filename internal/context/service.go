package context

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Store handles saving and retrieving context files
type Store struct {
	StorageDir string
}

func (s *Store) Save(ctx ProjectContext) error {
	path := filepath.Join(s.StorageDir, ctx.Name+".json")
	data, err := json.MarshalIndent(ctx, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
