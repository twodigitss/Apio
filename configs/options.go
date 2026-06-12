package configs

import (
	"os"

	"github.com/twodigitss/apio/internal/shared"
)

var (
	DefaultPath string = "."
	WorkingDir string = "."
)

func Init() error {
	wd, err := os.Getwd()
	if err != nil { return err }

	WorkingDir = wd

	return nil
}

func SetWorkingDir(path string){
	WorkingDir = shared.ExpandPath(path)
}
