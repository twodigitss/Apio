package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/core/config"
	"github.com/twodigitss/apio/internal/core/shared"
	data "github.com/twodigitss/apio/internal/ui/data"
)

// ponytail: Print() deleted — superseded by PrintV2()

type Tokens struct {
	Method   string
	URL      string
	Protocol string
	Headers  map[string]string
	Body     string
}

var cfg = config.Default()

func (t Tokens) Label() string {
	return t.Method + " " + t.URL
}

func (t Tokens) PrintV2(glyphs bool) string {
	var s strings.Builder

	var protocol string = "HTTP/2.0 (default)"
	if t.Protocol != "" {
		protocol = t.Protocol
	}

	//fucking hell, the go lsp telling me to change it for fprintf is making me insane
	fmt.Fprintf(&s, "%s: %s\n",
		shared.Label("", "Method", glyphs),
		lipgloss.NewStyle().Foreground(lipgloss.Color(data.GetColorByHttpMethod(t.Method))).Render(t.Method))
	fmt.Fprintf(&s, "%s: %s\n",
		shared.Label("󰿘", "Protocol", glyphs),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#acacac")).Render(protocol))
	fmt.Fprintf(&s, "%s: %s\n\n",
		shared.Label("󰌷", "Url", glyphs),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#dedede")).Render(t.URL))

	fmt.Fprintf(&s, "%s: ", "Body")
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
	fmt.Fprintf(&s, "%s:\n", shared.Label("", "Headers", glyphs))
	s.WriteString(shared.PrettyHeaders(t.Headers, cfg.Colors.SUBTEXT))

	return s.String()
}
