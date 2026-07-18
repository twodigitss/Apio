package fileselection

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
)

func (m Model) View() string {
	var s strings.Builder
	s.WriteString("  Select HTTP/REST File\n\n")
	for i, file := range m.Files {
		cursor := " "
		style := lipgloss.NewStyle()
		if m.FileCursor == i {
			cursor = lipgloss.NewStyle().
				Background(compat.AdaptiveColor{
					Light: lipgloss.Color("#000000"),
					Dark:  lipgloss.Color("#f1f1f1"),
				}).
				PaddingLeft(1).
				Blink(true).
				Render("")

			style = style.Bold(true)
		}
		s.WriteString(fmt.Sprintf("%s %s\n", cursor, style.Render(file.Name())))
	}
	s.WriteString("\n  [esc/f] cancel")

	return lipgloss.NewStyle().
		Padding(1, 3).
		Render(s.String())
}
