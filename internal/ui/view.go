package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/ui/data"
)

func (m Model) View() tea.View {

	sidebarWidth := m.Width / 3
	viewerWidth := m.Width - sidebarWidth

	sidebar := m.renderSidebar(sidebarWidth)
	viewer := m.renderViewer()

	right := lipgloss.NewStyle().Width(viewerWidth).
		PaddingLeft(5).
		PaddingRight(5).
		PaddingTop(2).
		Render(viewer)

	left := lipgloss.NewStyle().
		Width(sidebarWidth).
		BorderRight(true).
		// BorderStyle(lipgloss.NormalBorder()).
		// BorderForeground(lipgloss.Color("#3C3C3C")).
		PaddingLeft(5).
		PaddingRight(5).
		PaddingTop(2).
		Render(sidebar)

	v := tea.NewView(lipgloss.JoinHorizontal(lipgloss.Top, left, right))
	v.AltScreen = true
	v.MouseMode = tea.MouseModeCellMotion
	return v
}

func (m Model) renderSidebar(width int) string {
	var s strings.Builder
	// s.WriteString("Press q to quit.\n\n")
	// s.WriteString("Requests Found\n\n")
	for i, choice := range m.requests {
		cursor := " "
		if m.cursor == i {
			cursor = lipgloss.NewStyle().
				Background(lipgloss.Color("#f5f5f5")).
				PaddingLeft(1).
				Blink(true).
				Render("")
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
				cursor, style.Render(choice.Method), data.Truncate(choice.URL, width-25),
			))
	}
	return s.String()
}

func (m Model) renderViewer() string {
	if m.loading {
		return m.spinner.View() + " loading..."
	}

	// ponytail: render the viewport view which handles vertical scrolling
	vpView := m.viewport.View()

	// ponytail: display scroll percentage at the bottom
	percent := m.viewport.ScrollPercent()
	var percentStr string
	if percent >= 1.0 {
		percentStr = "100%"
	} else if percent <= 0.0 {
		percentStr = "0%"
	} else {
		percentStr = fmt.Sprintf("%2.f%%", percent*100)
	}

	footerText := fmt.Sprintf(" %s ", percentStr)
	footerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))
	footerVal := footerStyle.Render(footerText)

	width := m.viewport.Width()
	dashCount := max(width-lipgloss.Width(footerVal), 0)
	dashes := strings.Repeat(" ", dashCount)
	footer := dashes + footerVal

	return vpView + "\n" + footer
}
