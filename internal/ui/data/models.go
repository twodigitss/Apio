package data

import (
	"net/http"
	"time"
)

var res http.Response

type Response struct {
	Body          string
	ContentLength int64
	Headers       map[string]string
	Status        string
	StatusCode    int
	Duration      time.Duration
}
