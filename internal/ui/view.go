package ui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/ui/components/help"
)

func (m Model) View() tea.View {

	sidebarWidth := m.Width / 3
	viewerWidth := m.Width - sidebarWidth

	if m.showHelp {
		help := help.View()
		v := tea.NewView(lipgloss.Place(m.Width, m.Height,
			lipgloss.Center, lipgloss.Center, help))
		v.AltScreen = true
		return v
	}

	if m.selectingFile {
		fileSel := m.fileSelection.View()
		v := tea.NewView(lipgloss.Place(m.Width, m.Height,
			lipgloss.Center, lipgloss.Center, fileSel))
		v.AltScreen = true
		return v
	}

	right := lipgloss.NewStyle().Width(viewerWidth).
		PaddingLeft(5).
		PaddingRight(5).
		PaddingTop(2).
		Render(m.viewer.View())

	left := lipgloss.NewStyle().
		Width(sidebarWidth).
		BorderRight(true).
		PaddingLeft(5).
		PaddingRight(5).
		PaddingTop(2).
		Render(m.sidebar.View(sidebarWidth))

	v := tea.NewView(lipgloss.JoinHorizontal(lipgloss.Top, left, right))
	v.AltScreen = true
	v.MouseMode = tea.MouseModeCellMotion
	return v
}
