package models

import (
	"encoding/json"
)

type Tokens struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

type printing struct {
	Method  string            `json:"Method"`
	Url     string            `json:"Url"`
	Headers map[string]string `json:"Headers"`
	Body    any               `json:"Body"`
}

func (t Tokens) Print() string {
	var bodyVal any = ""
	if len(t.Body) > 0 {
		var parsed any
		if err := json.Unmarshal([]byte(t.Body), &parsed); err == nil {
			bodyVal = parsed
		} else {
			bodyVal = t.Body
		}
	}

	b := printing{
		Method:  t.Method,
		Url:     t.URL,
		Headers: t.Headers,
		Body:    bodyVal,
	}

	b_, err := json.MarshalIndent(b, "", "    ")
	if err != nil {
		return ""
	}

	return string(b_)
}

func (t Tokens) Label() string {
	return t.Method + " " + t.URL
}
