package runner

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/twodigitss/apio/internal/core/config"
	"github.com/twodigitss/apio/internal/core/parser/models"
)

func resolveTimeout() time.Duration {
	var timeout, err = time.ParseDuration(config.Default().Core.Timeout)
	if err != nil {
		timeout = 5 * time.Second
	}
	return timeout
}

var client = &http.Client{}

func Run(tok models.Tokens) (*http.Response, error) {
	req, err := http.NewRequest(
		tok.Method, tok.URL,
		strings.NewReader(tok.Body),
	)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(req.Context(), resolveTimeout())
	defer cancel()
	req = req.WithContext(ctx)

	for k, v := range tok.Headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
