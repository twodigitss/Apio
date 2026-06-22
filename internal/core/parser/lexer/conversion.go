package lexer

import (
	"fmt"
	"log"

	"github.com/twodigitss/apio/internal/core/parser/models"
	"github.com/twodigitss/apio/internal/core/parser/splitter"
)

func FileToArrTokens(file []byte) ([]models.Tokens, error) {

	if len(file) == 0 {
		return nil, fmt.Errorf("File size is 0Kb, either corrupted or error getting file contents")
	}

	parts := splitter.RequestSplitter(file)
	// for i, part := range parts {
	// 	fmt.Println("\n", "-----",i,"-----", "\n", part, )
	// }

	var requests []models.Tokens
	for _, part := range parts {
		req, err := Lexer(part)
		if err != nil {
			log.Fatal("Error parsing block:", err)
		}

		requests = append(requests, req)
	}
	return requests, nil
}
