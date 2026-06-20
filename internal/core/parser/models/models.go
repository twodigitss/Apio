package models

import (
	"fmt"
)

type Tokens struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

func (t Tokens) Print() string {
	bodyStr := ""
	if len(t.Body) > 0 {
		bodyStr = string(t.Body)
	}
	return fmt.Sprintf(
		"{\n  Method: %q, \n  URL: %q, \n  Headers: %v, \n  Body: %q\n}",
		t.Method, t.URL, t.Headers, bodyStr,
	)
}

func (t Tokens) Label() string {
	return t.Method + " " + t.URL
}
