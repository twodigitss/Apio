package sidebar

import tea "charm.land/bubbletea/v2"

// ponytail: Init is a no-op since the sidebar component does not need to trigger initial commands
func (m Model) Init() tea.Cmd {
	return nil
}

// ponytail: Update handles navigation keypresses to update its own Cursor state
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Requests)-1 {
				m.Cursor++
			}
		}
	}
	return m, nil
}
