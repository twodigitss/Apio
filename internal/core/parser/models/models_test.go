package models

import (
	"strings"
	"testing"
)

func TestTokens_Print(t *testing.T) {
	t.Run("Valid JSON body", func(t *testing.T) {
		tok := Tokens{
			Method: "POST",
			URL:    "http://example.com/api",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: `{"foo":"bar","baz":[1,2,3]}`,
		}

		res := tok.Print()
		t.Logf("Result:\n%s", res)

		if strings.Contains(res, `\"foo\"`) {
			t.Error("Body seems to be escaped string instead of parsed JSON object")
		}
	})

	t.Run("Invalid JSON body", func(t *testing.T) {
		tok := Tokens{
			Method: "POST",
			URL:    "http://example.com/api",
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
			Body: `plain text body`,
		}

		// This should not panic
		res := tok.Print()
		t.Logf("Result:\n%s", res)

		if !strings.Contains(res, `"Body": "plain text body"`) {
			t.Error("Expected Body field to contain the raw text value")
		}
	})

	t.Run("Empty body", func(t *testing.T) {
		tok := Tokens{
			Method: "GET",
			URL:    "http://example.com/api",
			Headers: map[string]string{
				"Accept": "application/json",
			},
			Body: ``,
		}

		res := tok.Print()
		t.Logf("Result:\n%s", res)

		if !strings.Contains(res, `"Body": ""`) {
			t.Error("Expected Body field to be empty string")
		}
	})
}
