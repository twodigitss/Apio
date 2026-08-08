package models

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/ui/data"
)

type Tokens struct {
	Method   string
	URL      string
	Protocol string
	Headers  map[string]string
	Body     string
}

type printing struct {
	Method   string            `json:"Method"`
	Url      string            `json:"Url"`
	Protocol string            `json:"Protocol,omitempty"`
	Headers  map[string]string `json:"Headers"`
	Body     any               `json:"Body"`
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
		Method:   t.Method,
		Url:      t.URL,
		Protocol: t.Protocol,
		Headers:  t.Headers,
		Body:     bodyVal,
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

func (t Tokens) PrintV2() string {
	var s strings.Builder

	var protocol string = "HTTP/2.0 (default)"
	if t.Protocol != "" {
		protocol = t.Protocol
	}

	//fucking hell, the go lsp telling me to change it for fprintf is making me insane
	s.WriteString(fmt.Sprintf(" Method: %s\n",
		lipgloss.NewStyle().Foreground(lipgloss.Color(data.GetColorByHttpMethod(t.Method))).Render(t.Method),
	))
	s.WriteString(fmt.Sprintf("󰿘 Protocol: %s\n",
		lipgloss.NewStyle().Foreground(lipgloss.Color("#acacac")).Render(protocol),
	))
	s.WriteString(fmt.Sprintf("󰌷 Url: %s\n\n",
		lipgloss.NewStyle().Foreground(lipgloss.Color("#dedede")).Render(t.URL),
	))

	s.WriteString("Body: ")
	if len(t.Body) > 0 {
		var parsed any
		if err := json.Unmarshal([]byte(t.Body), &parsed); err == nil {
			b, _ := json.MarshalIndent(parsed, "", "  ")
			s.WriteString(
				lipgloss.NewStyle().Foreground(lipgloss.Color("#dedede")).Render(string(b)),
			)

		} else {
			s.WriteString("\n" +
				lipgloss.NewStyle().Foreground(lipgloss.Color("#acacac")).Render(t.Body),
			)
		}
		s.WriteString("\n")
	} else {
		s.WriteString(
			lipgloss.NewStyle().Foreground(lipgloss.Color("#dedede")).Render("{}\n"),
		)
	}

	s.WriteString("\n")
	s.WriteString(" Headers:\n")
	s.WriteString(prettyHeaders(t.Headers))

	return s.String()
}

func prettyHeaders(h map[string]string) string {
	keys := make([]string, 0, len(h))
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sb strings.Builder
	for _, k := range keys {
		fmt.Fprintf(&sb, " • %s: %s\n",
			lipgloss.NewStyle().Render(k),
			lipgloss.NewStyle().Foreground(lipgloss.Color("#b0b0b0")).Italic(true).Render(h[k]),
		)
	}
	return sb.String()
}
