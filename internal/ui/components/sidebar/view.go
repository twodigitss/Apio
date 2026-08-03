package sidebar

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/twodigitss/apio/internal/ui/data"
)

func (m Model) View(width int) string {
	var s strings.Builder
	for i, choice := range m.Requests {
		cursor := " "
		style := data.StyleHttpMethod(choice.Method)

		if m.Cursor == i {
			cursor = lipgloss.NewStyle().
				Background(compat.AdaptiveColor{
					Light: style.GetBackground(),
					Dark:  style.GetBackground(),
				}).
				PaddingLeft(1).
				Blink(true).
				Render("")
		}

		url := strings.TrimPrefix(choice.URL, "https://")
		url = strings.TrimPrefix(url, "http://")
		urlTitle := data.Truncate(url, width-25) //magic number goes brrr

		s.WriteString(
			fmt.Sprintf("%s %s %s\n",
				cursor, style.Bold(true).Render(choice.Method),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#e0e0e0")).Render(urlTitle),
			))
	}
	return s.String()
}
