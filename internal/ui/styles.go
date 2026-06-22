package ui

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/shared"
)

var styles map[string]string = map[string]string{
	"GET":     "#A8E6CF",
	"POST":    "#FFD3B6",
	"PUT":     "#A9DEF9",
	"DELETE":  "#FFADAD",
	"PATCH":   "#D8B4F8",
	"HEAD":    "#A0E7E5",
	"OPTIONS": "#FFC6FF",
}

func StyleHttpMethod(line string) lipgloss.Style {
	var prefix string
	for _, v := range shared.HttpMethods {
		if strings.HasPrefix(strings.TrimSpace(line), v) {
			prefix = v
		}
	}

	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color(styles[prefix])).
		PaddingLeft(1).
		PaddingRight(1)

}
