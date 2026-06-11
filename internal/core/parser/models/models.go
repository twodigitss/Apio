package models

import (
	"encoding/json"
	"fmt"
)

var HttpMethods = []string{
	"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS",
}

// TODO: implemment more body types
type Tokens struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    json.RawMessage
}

func (t Tokens) String() string {
	bodyStr := ""
	if len(t.Body) > 0 {
		bodyStr = string(t.Body)
	}
	return fmt.Sprintf(
		"{Method: %q, URL: %q, Headers: %v, Body: %q}", 
		t.Method, t.URL, t.Headers, bodyStr,
	)
}
