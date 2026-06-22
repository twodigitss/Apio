package data

import (
	"net/http"
)

// Response struct used for the runner's results
type RunResponseMsg struct {
	Response http.Response
	Body     string
	Err      error
}

// TODO: do a better truncation logic.
func Truncate(s string, max int) string {
	if max <= 0 || len(s) <= max {
		return s
	}
	return s[:max-1] + "…"
}
