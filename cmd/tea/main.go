package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/twodigitss/apio/configs"
	"github.com/twodigitss/apio/internal/core/finder"
	"github.com/twodigitss/apio/internal/ui"
)

func main() {
	configs.Init()

	thisDir, err := finder.GetFiles(".")
	if err != nil || len(thisDir) <= 0 {
		if len(thisDir) <= 0 {
			err = fmt.Errorf("This dir is empty")
		}
		log.Fatal("Error loading files from given directory:", err)
	}

	file, err := finder.ReadFile(thisDir[0])
	if err != nil {
		log.Fatal("Error decoding file:", err)
	}

	p := tea.NewProgram(ui.New(thisDir, file))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
