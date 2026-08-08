package viewer

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/twodigitss/apio/internal/core/config"
)

func (m Model) View() string {
	if m.Loading {
		return m.Spinner.View() + " Loading..."
	}

	vpView := m.Viewport.View()
	percent := m.Viewport.ScrollPercent()

	var percentStr string
	if percent >= 1.0 {
		percentStr = "100%"
	} else if percent <= 0.0 {
		percentStr = "0%"
	} else {
		percentStr = fmt.Sprintf("%2.f%%", percent*100)
	}

	footerText := fmt.Sprintf(" %s ", percentStr)
	footerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(config.Default().Colors.SUBTEXT)).Italic(true)
	footerVal := footerStyle.Render(footerText)

	width := m.Viewport.Width()
	dashCount := max(width-lipgloss.Width(footerVal)-2, 0)
	dashes := strings.Repeat(" ", dashCount)
	footer := dashes + footerVal

	return vpView + "\n" + footer
}
