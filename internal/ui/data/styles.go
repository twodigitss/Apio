package data

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/shared"
)

var styles = map[string]string{
	"GET":     "#A8E6CF",
	"POST":    "#FFD3B6",
	"PUT":     "#A9DEF9",
	"DELETE":  "#FFADAD",
	"PATCH":   "#D8B4F8",
	"HEAD":    "#A0E7E5",
	"OPTIONS": "#FFC6FF",
	"EXTRA":   "#b3e6a8",
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
		Width(9).
		Align(lipgloss.Center).
		PaddingLeft(1).
		PaddingRight(1)

}

func ColorResponse(httpCode int) string {
	switch true {
	case httpCode >= 100 && httpCode < 200:
		return styles["PUT"]
	case httpCode >= 200 && httpCode < 300:
		return styles["EXTRA"]
	case httpCode >= 300 && httpCode < 400:
		return styles["PATCH"]
	case httpCode >= 400 && httpCode < 500:
		return styles["POST"]
	case httpCode >= 500 && httpCode < 600:
		return styles["DELETE"]
	default:
		return styles["OPTIONS"]
	}
}

func GetColorByHttpMethod(line string) string {
	var prefix string
	for _, v := range shared.HttpMethods {
		if strings.HasPrefix(strings.TrimSpace(line), v) {
			prefix = v
		}
	}
	return styles[prefix]
}
