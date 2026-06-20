package ui

import (
	"io"
	"log"
	"net/http"

	tea "charm.land/bubbletea/v2"
	"github.com/twodigitss/apio/internal/core/runner"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

			m.response = http.Response{}
			m.responseBody = ""
			m.currentRequest = m.requests[m.cursor]

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.requests)-1 {
				m.cursor++
			}

			m.response = http.Response{}
			m.responseBody = ""
			m.currentRequest = m.requests[m.cursor]

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		// case "space":
		// if m.selected == m.cursor {
		// 	m.selected = -1
		// } else {
		// 	m.selected = m.cursor
		// m.response = http.Response{}
		// m.responseBody = ""
		// m.currentRequest = m.requests[m.cursor]
		// }

		case "enter":
			// if m.selected == -1 {
			// 	return m, nil
			// }

			// var i int = m.selected
			// if configs.RunAtCursor {
			// 	i = m.cursor
			// }

			res, err := runner.Run(m.requests[m.cursor])

			if err != nil {
				// FIX: do something with this
				log.Fatal("Error running block:", err)
			}

			m.response = res

			if res.Body != nil {
				bodyBytes, _ := io.ReadAll(res.Body)
				m.responseBody = string(bodyBytes)
				res.Body.Close()
			} else {
				m.responseBody = ""
			}

		case "c":
			m.response.Body = nil
			m.response.StatusCode = 0
			m.responseBody = ""
		}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
