package main

import (
	"fmt"
	"log"

	"github.com/twodigitss/apio/configs"
	"github.com/twodigitss/apio/internal/core/finder"
	"github.com/twodigitss/apio/internal/core/parser"
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

	_, err = core.FileToArrTokens(file)
	if err != nil {
		log.Fatal("Error decoding file:", err)
	}

	// res, err := runner.Run(tokenizedReq)
	// if err != nil {
	// 	log.Fatal("Error running block:", err)
	// }

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal("Error reading body:", err)
	// }
	// fmt.Println(string(body))

}
