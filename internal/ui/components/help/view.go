package help

import (
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"github.com/twodigitss/apio/internal/core/config"
)

var rows = [][]string{
	{"j / ↓", "Move down"},
	{"k / ↑", "Move up"},
	{"ctrl+j", "Scroll down"},
	{"ctrl+k", "Scroll up"},
	{"enter", "Run"},
	// {"r", "Reload"},
	{"c", "Clear"},
	{"f", "Files"},
	{"h / ?", "Help"},
	{"q", "Quit"},
	{"y", "Copy body"},
}

var headerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(config.Default().Colors.GET)).Bold(true).Align(lipgloss.Center)
var cellStyle = lipgloss.NewStyle().Padding(0, 1).Width(25).Align(lipgloss.Center, lipgloss.Center)
var oddRowStyle = cellStyle.Foreground(lipgloss.Color(config.Default().Colors.INFRATEXT))
var evenRowStyle = cellStyle.Foreground(lipgloss.Color(config.Default().Colors.TEXT))
var cfg = config.Default()

func View(Height, Width int) string {
	var s strings.Builder
	content := `
┏┓  •  
┣┫┏┓┓┏┓
┛┗┣┛┗┗┛
  ┛    
`

	s.WriteString(content)
	s.WriteString("v1.1\n\n")

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color(config.Default().Colors.SUBBORDER))).
		Wrap(true).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		Headers("Keybind", "Description").
		Rows(rows...)

	s.WriteString(t.Render() + "\n\n\n")
	s.WriteString(
		lipgloss.NewStyle().Foreground(lipgloss.Color(config.Default().Colors.INFRATEXT)).Render("[esc/h/q/?] cancel\n\n"),
	)

	return lipgloss.NewStyle().
		Height(Height).
		Width(Width).
		Border(lipgloss.NormalBorder(), cfg.UI.Borders, cfg.UI.Borders).
		Align(lipgloss.Center, lipgloss.Center).
		BorderForeground(lipgloss.Color(config.Default().Colors.INFRABORDER)).
		Render(s.String())
}
