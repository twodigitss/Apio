package finder

import (
	"fmt"
	"github.com/twodigitss/apio/configs"
	"github.com/twodigitss/apio/internal/shared"
	"os"
	"path/filepath"
)

// Would it be better if i do include the full path?
func GetFiles(path string) ([]os.DirEntry, error) {
	if path == "" {
		path = configs.WorkingDir
	}
	_path := shared.ExpandPath(path)

	configs.SetWorkingDir(_path)

	{
		pathexists := shared.PathExists(_path)
		if !pathexists {
			return nil, fmt.Errorf("This path seems to not exist")
		}
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

	// fmt.Println("REST/HTTP FILES FOUND IN THIS DIR: ", restfiles)
	return restfiles, nil
}

func ReadFile(file os.DirEntry) ([]byte, error) {
	buffer, err := os.ReadFile(configs.WorkingDir + "/" + file.Name())
	if err != nil {
		return nil, nil
	}
	return buffer, nil
}

func ReloadFiles(dir []os.DirEntry) {
	// file, err := ReadFile(dir[0])
	// if err != nil {
	// 	log.Fatal("Error decoding file:", err)
	// }

}
