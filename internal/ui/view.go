package ui

import (
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/ui/components/help"
)

func (m Model) View() tea.View {

	sidebarWidth := m.Width / 3
	viewerWidth := m.Width - sidebarWidth

	if m.showHelp {
		help := help.View(m.Height, m.Width)
		v := tea.NewView(lipgloss.Place(m.Width, m.Height,
			lipgloss.Center, lipgloss.Center, help))
		v.AltScreen = true
		return v
	}

	if m.selectingFile {
		fileSel := m.fileSelection.View(m.Height, m.Width)
		v := tea.NewView(lipgloss.Place(m.Width, m.Height,
			lipgloss.Center, lipgloss.Center, fileSel))
		v.AltScreen = true
		return v
	}

	// color := data.GetColorByHttpMethod(m.currentRequest.Method)
	// square := lipgloss.NewStyle().Background(lipgloss.Color("#fff")).Render(" ")

	topBar := lipgloss.NewStyle().Width(sidebarWidth - 6).
		PaddingLeft(2).
		PaddingRight(2).
		MarginTop(1).
		MarginBottom(1).
		Bold(true).
		// Border(lipgloss.RoundedBorder(), true, true).
		// BorderForeground(lipgloss.Color("#4a4a4a")).
		// Foreground(lipgloss.Color(color)).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				" ",
				lipgloss.NewStyle().
					Render(m.fileSelection.Files[m.fileSelection.FileCursor].Name()),
				" - ",
				lipgloss.NewStyle().
					Foreground(lipgloss.Color("#aeaeae")).
					Render(strconv.Itoa(len(m.sidebar.Requests))+" Requests"),
				//
			),
		)

	left := lipgloss.NewStyle().Width(sidebarWidth).
		PaddingLeft(2).
		PaddingRight(2).
		Height(m.Height).
		Border(lipgloss.NormalBorder(), true, true).
		BorderForeground(lipgloss.Color("#afafaf")).
		// Background(lipgloss.Color("#121212")).
		Render(
			lipgloss.JoinVertical(lipgloss.Top, topBar, m.sidebar.View(sidebarWidth)),
			// m.sidebar.View(sidebarWidth),
		)

	right := lipgloss.NewStyle().Width(viewerWidth).
		PaddingLeft(4).
		PaddingRight(4).
		PaddingTop(1).
		Height(m.Height).
		Border(lipgloss.RoundedBorder(), true, true).
		BorderForeground(lipgloss.Color("#2a2a2a")).
		// Background(lipgloss.Color("#1f1f1f")).
		Render(m.viewer.View())

	v := tea.NewView(
		lipgloss.JoinHorizontal(lipgloss.Top, left, right),
	)
	v.AltScreen = true
	v.MouseMode = tea.MouseModeCellMotion
	return v
}
