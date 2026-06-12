package main

import (
	"fmt"
	"io"
	"log"

	"github.com/twodigitss/apio/configs"
	"github.com/twodigitss/apio/internal/core/finder"
	"github.com/twodigitss/apio/internal/core/parser/lexer"
	"github.com/twodigitss/apio/internal/core/parser/splitter"
	"github.com/twodigitss/apio/internal/core/runner"
)

func main() {
	// first line ever of the program. respect it
	// fmt.Println("2nd world")
	configs.Init()

	thisDir, err := finder.GetFiles("~/Projects/apio/cmd/apio/")
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
	// for i, part := range parts {
	// 	fmt.Println("\n", "-----",i,"-----", "\n", part, )
	// }

	tokenizedReq, err := lexer.Lexer(parts[1])
	if err != nil {
		log.Fatal("Error parsing block:", err)
	}
	fmt.Println(tokenizedReq)

	res, err := runner.Run(tokenizedReq)
	if err != nil {
		log.Fatal("Error running block:", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading body:", err)
	}
	fmt.Println(string(body))

}
