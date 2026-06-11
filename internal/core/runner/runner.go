package runner

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/twodigitss/apio/internal/core/parser/models"
)

// NOTE: obviously this is going to be finished.
func Run(tok models.Tokens){
	switch tok.Method {
		case "GET": {
			res, err := http.Get(tok.URL)
			if err != nil { fmt.Print("Error:", err); return }
			defer res.Body.Close()

			bodyBytes, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatalf("Failed to read body: %v", err)
			}

			fmt.Println(string(bodyBytes))
		}

		case "POST": {
			res, err := http.Post(tok.URL, tok.Headers["Content-Type"], nil)
			if err != nil { fmt.Print("Error:", err); return }
			defer res.Body.Close()

			bodyBytes, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatalf("Failed to read body: %v", err)
			}

			fmt.Println(string(bodyBytes))
		}
		default: return
	}
}
