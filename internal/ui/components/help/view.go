package help

import (
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
)

var rows = [][]string{
	{"j / ↓", "Move down"},
	{"k / ↑", "Move up"},
	{"ctrl+j", "Scroll down"},
	{"ctrl+k", "Scroll up"},
	{"enter", "Run"},
	{"r", "Reload"},
	{"c", "Clear"},
	{"f", "Files"},
	{"h / ?", "Help"},
	{"q", "Quit"},
	{"y", "Copy body"},
}

var headerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#A8E6CF")).Bold(true).Align(lipgloss.Center)
var cellStyle = lipgloss.NewStyle().Padding(0, 1).Width(25).Align(lipgloss.Center, lipgloss.Center)
var oddRowStyle = cellStyle.Foreground(lipgloss.Color("#a0a0a0"))
var evenRowStyle = cellStyle.Foreground(lipgloss.Color("#efefef"))

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
	s.WriteString("Keybinds\n\n")

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#4e4e4e"))).
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
		lipgloss.NewStyle().Foreground(lipgloss.Color("#b0b0b0")).Render("[esc/h/q/?] cancel\n\n"),
	)

	return lipgloss.NewStyle().
		Height(Height).
		Width(Width).
		Border(lipgloss.NormalBorder(), true, true).
		Align(lipgloss.Center, lipgloss.Center).
		BorderForeground(lipgloss.Color("#2a2a2a")).
		Render(s.String())
}
