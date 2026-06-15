package data

import "time"

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
