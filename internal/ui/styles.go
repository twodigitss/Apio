package ui

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/shared"
)

var styles map[string]string = map[string]string{
	"GET":     "#85f2a8",
	"POST":    "#E8B208", //#85baf2
	"PUT":     "#127AA8",
	"DELETE":  "#EC3E5C",
	"PATCH":   "#8B5CF6",
	"HEAD":    "#139C8D",
	"OPTIONS": "#5F61E4",
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
		Foreground(lipgloss.Color(styles[prefix])).
		// Background(lipgloss.Color("#121212")).
		PaddingLeft(1).
		PaddingRight(1)

}
