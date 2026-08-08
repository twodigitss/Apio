package data

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/core/config"
	"github.com/twodigitss/apio/internal/shared"
)

var styles = map[string]string{
	"GET":     config.Default().Colors.GET,
	"POST":    config.Default().Colors.POST,
	"PUT":     config.Default().Colors.PUT,
	"DELETE":  config.Default().Colors.DELETE,
	"PATCH":   config.Default().Colors.PATCH,
	"HEAD":    config.Default().Colors.HEAD,
	"OPTIONS": config.Default().Colors.OPTIONS,
	"EXTRA":   config.Default().Colors.EXTRA,
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
