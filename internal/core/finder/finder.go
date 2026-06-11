package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/twodigitss/apio/configs"
)

func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") || path == "~" {
		home, err := os.UserHomeDir()
		if err != nil { return path }
		if path == "~" { return home }
		return filepath.Join(home, path[2:])
	}
	return path
}

func routeExists(path string) bool {
	_, err := os.Stat(path)
	switch err {
		case nil: return true
		case os.ErrNotExist: return false

		// Schrodinger's path: permission denied, broken symlink, or other error.
		// The path might still exist, but you cannot read its metadata.
		default: return false
	}
}

func GetFiles(path string) ([]os.DirEntry, error) {
	var _path string = configs.DefaultPath
	if path != "" { _path = path }
	_path = expandPath(_path)

	{
		pathexists := routeExists(_path)
		if !pathexists {
			return nil, fmt.Errorf("This path seems to not exist")
		}
	}

	thisDir, err := os.ReadDir(_path)
	if err != nil {
		return nil, nil
	}

	var restfiles []os.DirEntry
	for _, v := range thisDir {
		if v.IsDir() { continue }

		ext := filepath.Ext(v.Name())
		switch ext {
		case ".http", ".rest":
			restfiles = append(restfiles, v)
		default:
			continue
		}
	}

	// fmt.Println("REST/HTTP FILES FOUND IN THIS DIR: ", restfiles)
	return restfiles, nil
}

func ReadFile(file os.DirEntry) ([]byte, error) {
	buffer, err := os.ReadFile(configs.DefaultPath + "/" + file.Name())
	if err != nil { return nil, nil }
	return buffer, nil
}
