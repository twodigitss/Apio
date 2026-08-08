package shared

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"charm.land/lipgloss/v2"
)

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") || path == "~" {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		if path == "~" {
			return home
		}
		return filepath.Join(home, path[2:])
	}
	return path
}

func Label(glyph, text string, enable bool) string {
	if enable {
		return glyph + " " + text
	}
	return text
}

func PrettyHeaders[V string | []string](h map[string]V, color string) string {
	keys := make([]string, 0, len(h))
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		switch v := any(h[k]).(type) {
		case string:
			fmt.Fprintf(&sb, " • %s: %s\n",
				lipgloss.NewStyle().Render(k),
				lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Italic(true).Render(v),
			)
		case []string:
			for _, val := range v {
				fmt.Fprintf(&sb, " • %s: %s\n",
					lipgloss.NewStyle().Render(k),
					lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Italic(true).Render(val),
				)
			}
		}
	}
	return sb.String()
}

// TODO: do a better truncation logic.
func Truncate(s string, max int) string {
	if max <= 0 || len(s) <= max {
		return s
	}
	return s[:max-1] + "…"
}
