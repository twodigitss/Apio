package runner

import (
	"net/http"
	"strings"
	"github.com/twodigitss/apio/internal/core/parser/models"
)

var client *http.Client = &http.Client{}

func Run(tok models.Tokens) (http.Response, error){
	req, err := http.NewRequest(
		tok.Method, tok.URL, 
		strings.NewReader(tok.Body),
	)

	if err != nil {
		return http.Response{}, err
	}

	for k,v := range tok.Headers{
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return http.Response{}, err
	}

	return *resp, nil

}
