package lexer

import (
	"fmt"
	"strings"
	"github.com/twodigitss/apio/internal/core/parser/models"
)

func Lexer(block string) (models.Tokens, error) {
	lines := strings.Split(block, "\n")
	i := 0

	// parse primera línea no vacía como METHOD URL
	parts := strings.Fields(lines[i])
	if len(parts) < 2 {
		return models.Tokens{}, fmt.Errorf("invalid request line: %q", lines[i])
	}
	req := models.Tokens{
		Method: parts[0], 
		URL: parts[1], 
		Headers: make(map[string]string),
		Body: "",
	}
	i++

	// headers hasta línea vacía
	for i < len(lines) && strings.TrimSpace(lines[i]) != "" {
		k, v, ok := strings.Cut(lines[i], ":")
		if ok {
			req.Headers[strings.TrimSpace(k)] = strings.TrimSpace(v)
		}
		i++
	}

	// body = resto
	req.Body = strings.TrimSpace(strings.Join(lines[i:], "\n"))
	return req, nil
}
