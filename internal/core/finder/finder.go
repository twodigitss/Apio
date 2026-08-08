package finder

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/twodigitss/apio/internal/core/shared"
)

var (
	WorkingDir string = "."
)

func GetFiles(path string) ([]os.DirEntry, error) {
	if path == "" {
		path = WorkingDir
	}
	_path := shared.ExpandPath(path)

	WorkingDir = shared.ExpandPath(_path)

	// ponytail: inlined shared.PathExists
	if _, err := os.Stat(_path); err != nil {
		return nil, fmt.Errorf("This path seems to not exist")
	}

	thisDir, err := os.ReadDir(_path)
	if err != nil {
		return nil, err
	}

	var restfiles []os.DirEntry
	for _, v := range thisDir {
		if v.IsDir() {
			continue
		}

		ext := filepath.Ext(v.Name())
		switch ext {
		case ".http", ".rest":
			restfiles = append(restfiles, v)
		default:
			continue
		}
	}

	return restfiles, nil
}

func ReadFile(file os.DirEntry) ([]byte, error) {
	buffer, err := os.ReadFile(filepath.Join(WorkingDir, file.Name()))
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
