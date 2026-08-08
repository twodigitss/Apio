package fileselection

import (
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"charm.land/lipgloss/v2/tree"
	"github.com/twodigitss/apio/internal/core/config"
)

var cfg = config.Default()

func (m Model) View(Height, Width int) string {
	var s strings.Builder
	var t = tree.New()
	content := `
┏┓  •  
┣┫┏┓┓┏┓
┛┗┣┛┗┗┛
  ┛    
	`

	s.WriteString(content)
	s.WriteString("Select HTTP/REST File\n\n")

	// calcula el ancho máximo del nombre de archivo para uniformar
	maxLen := 0
	for _, file := range m.Files {
		if l := len(file.Name()); l > maxLen {
			maxLen = l
		}
	}

	for i, file := range m.Files {
		cursorStyle := lipgloss.NewStyle().Width(1)
		nameStyle := lipgloss.NewStyle().Width(maxLen)

		cursor := cursorStyle.Render(" ")
		if m.FileCursor == i {
			cursor = cursorStyle.
				Background(compat.AdaptiveColor{
					Light: lipgloss.Color("#000000"),
					Dark:  lipgloss.Color("#f1f1f1"),
				}).
				Blink(true).
				Render(" ")
			nameStyle = nameStyle.Bold(true).Foreground(lipgloss.Color(config.Default().Colors.GET))
		}

		body := cursor + " " + nameStyle.Render(file.Name())
		t.Child(body)
	}

	s.WriteString(t.String() + "\n\n\n")
	s.WriteString(
		lipgloss.NewStyle().Foreground(lipgloss.Color(config.Default().Colors.INFRATEXT)).Render("[esc/f] cancel\n\n"),
	)

	return lipgloss.NewStyle().
		Padding(1, 3).
		Height(Height).
		Width(Width).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.NormalBorder(), cfg.UI.Borders, cfg.UI.Borders).
		BorderForeground(lipgloss.Color(config.Default().Colors.INFRABORDER)).
		Render(s.String())
}
