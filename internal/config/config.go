package config

import (
	"os"
	"path/filepath"
)

func GetStorageDir() string {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".ctx")
	os.MkdirAll(path, 0755) // Ensure it exists
	return path
}
