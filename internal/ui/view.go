package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	// The header
	var s strings.Builder
	s.WriteString("What should we buy at the market?\n\n")

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice))
	}

	// The footer
	s.WriteString("\nPress q to quit.\n")

	// Send the UI for rendering
	return tea.NewView(s.String())
}
