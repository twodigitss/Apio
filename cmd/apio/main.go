package main

import (
	"fmt"
	"log"

	"github.com/twodigitss/apio/configs"
	"github.com/twodigitss/apio/internal/core/finder"
	"github.com/twodigitss/apio/internal/core/parser/splitter"
)

func main() {
	// first line ever of the program. respect it
	// fmt.Println("2nd world")
	configs.Init()

	thisDir, err := finder.GetFiles("")
	if err != nil || len(thisDir)<=0 {
		if len(thisDir) <= 0 { 
			err = fmt.Errorf("This dir is empty")
		}

		log.Fatal("Error loading files from given directory:", err)
	}

	file, err := finder.ReadFile(thisDir[0])
	if err != nil {
		log.Fatal("Error decoding file:", err)
	}

	parts := splitter.RequestSplitter(file)
	for _, part := range parts {
		fmt.Println(part)
	}

}
