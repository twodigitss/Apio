package ui

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/shared"
)

var styles map[string]string = map[string]string{
	"GET": "#85f2a8",
	"POST": "#85baf2",
	"PUT": "#e68bf7",
	"DELETE": "#f78bab",
	"PATCH": "#f2af85",
	"HEAD": "#ca85f2",
	"OPTIONS": "#e9f285",
}

func StyleHttpMethod(line string) lipgloss.Style {
	var prefix string;
	for _, v := range shared.HttpMethods {
		if strings.HasPrefix( strings.TrimSpace(line), v) {
			prefix = v
		}
	}

	return lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#121212")).
	Background(lipgloss.Color(styles[prefix])).
	PaddingLeft(1).
	PaddingRight(1)

}
