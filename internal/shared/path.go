package shared

import (
	"os"
	"path/filepath"
	"strings"
)

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") || path == "~" {
		home, err := os.UserHomeDir()
		if err != nil { return path }
		if path == "~" { return home }
		return filepath.Join(home, path[2:])
	}
	return path
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	switch err {
		case nil: return true
		case os.ErrNotExist: return false

		// Schrodinger's path: permission denied, broken symlink, or other error.
		// The path might still exist, but you cannot read its metadata.
		default: return false
	}
}

