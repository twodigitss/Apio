package ui

import (
	"fmt"
	// "net/http"
	"strings"

	// tea "charm.land/bubbletea/v2"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	// "github.com/twodigitss/apio/configs"
)

func (m Model) View() tea.View {
	sidebar := m.renderSidebar()
	viewer := m.renderViewer()

	sidebarWidth := m.Width / 2
	viewerWidth := m.Width - sidebarWidth

	right := lipgloss.NewStyle().Width(viewerWidth).
		PaddingLeft(5).
		PaddingRight(5).
		PaddingTop(2).
		Render(viewer)

	left := lipgloss.NewStyle().
		Width(sidebarWidth).
		BorderRight(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#3C3C3C")).
		PaddingLeft(5).
		PaddingRight(5).
		PaddingTop(2).
		Render(sidebar)

	v := tea.NewView(lipgloss.JoinHorizontal(lipgloss.Top, left, right))
	v.AltScreen = true
	return v
}

func (m Model) renderSidebar() string {
	var s strings.Builder
	s.WriteString("Press q to quit.\n\n")
	// s.WriteString("Requests Found\n\n")
	for i, choice := range m.requests {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// checked := " "
		// if m.selected == i {
		// 	checked = "x"
		// }

		style := StyleHttpMethod(choice.Method)

		// var line string = "%s [%s] %s %s\n"
		// if configs.RunAtCursor {
		// 	line = "%s %s %s\n"
		// }

		s.WriteString(
			fmt.Sprintf("%s %s %s\n",
				cursor, style.Render(choice.Method), choice.URL,
			))
	}
	return s.String()
}

func (m Model) renderViewer() string {
	if m.response.StatusCode == 0 {
		// return "No response yet.\n"
		return m.currentRequest.Print()
	}

	// ponytail: render the cached response body instead of reading from stream in View
	return fmt.Sprintf("Status Code: %d \n\n%s\n", m.response.StatusCode, m.responseBody)
}
