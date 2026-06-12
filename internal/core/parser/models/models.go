package models

import (
	"fmt"
	"time"
)

var HttpMethods = []string{
	"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS",
}

type Tokens struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

func (t Tokens) String() string {
	bodyStr := ""
	if len(t.Body) > 0 {
		bodyStr = string(t.Body)
	}
	return fmt.Sprintf(
		"{\n  Method: %q, \n  URL: %q, \n  Headers: %v, \n  Body: %q\n}",
		t.Method, t.URL, t.Headers, bodyStr,
	)
}

type Response struct {
    StatusCode int
    Headers    map[string]string
    Body       string
    Duration   time.Duration
}

type ResponseReceivedMsg struct {
    Response Response
    Err      error
}
