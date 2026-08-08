package ui

import (
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/core/shared"
	"github.com/twodigitss/apio/internal/ui/components/help"
	data "github.com/twodigitss/apio/internal/ui/data"
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

	color := data.GetColorByHttpMethod(m.currentRequest.Method)
	// square := lipgloss.NewStyle().Background(lipgloss.Color("#fff")).Render(" ")

	topBar := lipgloss.NewStyle().Width(sidebarWidth - 6).
		PaddingLeft(2).
		PaddingRight(2).
		MarginTop(1).
		MarginBottom(1).
		Bold(true).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Center,

				shared.Label(
					lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render(""),
					lipgloss.NewStyle().Render(m.fileSelection.Files[m.fileSelection.FileCursor].Name()),
					cfg.UI.Glyphs,
				),
				" - ",
				lipgloss.NewStyle().
					Foreground(lipgloss.Color(cfg.Colors.SUBTEXT)).
					Render(strconv.Itoa(len(m.sidebar.Requests))+" Requests"),
				//
			),
		)

	leftStyle := lipgloss.NewStyle().Width(sidebarWidth).
		PaddingLeft(2).
		PaddingRight(2).
		Height(m.Height)

	if cfg.UI.Borders {
		leftStyle = leftStyle.Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(cfg.Colors.BORDER))
	}

	left := leftStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, topBar, m.sidebar.View(sidebarWidth)),
		// m.sidebar.View(sidebarWidth),
	)

	rightStyle := lipgloss.NewStyle().Width(viewerWidth).
		PaddingLeft(4).
		PaddingRight(4).
		PaddingTop(1).
		Height(m.Height)

	if cfg.UI.Borders {
		rightStyle = rightStyle.Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(cfg.Colors.INFRABORDER))
	}

	right := rightStyle.Render(m.viewer.View())

	var res string
	if cfg.UI.Sidebar == "left" {
		res = lipgloss.JoinHorizontal(lipgloss.Top, left, right)
	} else {
		res = lipgloss.JoinHorizontal(lipgloss.Top, right, left)
	}

	v := tea.NewView(res)

	v.AltScreen = true
	v.MouseMode = tea.MouseModeCellMotion
	return v
}
