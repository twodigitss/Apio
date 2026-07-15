package finder

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/twodigitss/apio/configs"
	"github.com/twodigitss/apio/internal/core/parser/lexer"
	"github.com/twodigitss/apio/internal/core/parser/models"
	"github.com/twodigitss/apio/internal/shared"
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
		return nil, err
	}
	return buffer, nil
}

func ReloadFiles(currentFile []byte) ([]models.Tokens, error) {
	tokens, err := lexer.FileToArrTokens(currentFile)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
