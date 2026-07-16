package fileselection

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.FileCursor > 0 {
				m.FileCursor--
			}
		case "down", "j":
			if m.FileCursor < len(m.Files)-1 {
				m.FileCursor++
			}
		}
	}
	return m, nil
}
